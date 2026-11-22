package smtp

import (
	"context"
	"errors"
	"io"
	"log"
	"strings"

	"github.com/Ankumeah/SourceFlame/internal/app"
	"github.com/Ankumeah/SourceFlame/internal/db"
	"github.com/Ankumeah/SourceFlame/internal/errs"
	"github.com/Ankumeah/SourceFlame/internal/utils"
	"github.com/emersion/go-smtp"
)

type session struct {
	ctx    context.Context
	app    *app.App
	domain string

	from     string
	subject  string
	email    db.Email
	repo     db.Repo
	user     *string
	verified bool
}

func (s *session) Mail(from string, _ *smtp.MailOptions) error {
	s.from = from

	return nil
}

func (s *session) Rcpt(to string, _ *smtp.RcptOptions) error {
	if s.repo != (db.Repo{}) {
		return nil
	}

	parts := strings.Split(to, "@")
	if len(parts) != 2 {
		return &smtp.SMTPError{
			Code:    501,
			Message: "Invalid email format",
		}
	}

	if parts[1] != s.domain {
		return &smtp.SMTPError{
			Code:    550,
			Message: "Relaying not allowed",
		}
	}

	repoParts := strings.Split(parts[0], "/")
	if len(repoParts) != 2 {
		if s.app.Config.DefaultUser == "" {
			return &smtp.SMTPError{
				Code:    550,
				Message: "Invalid repo",
			}
		} else {
			s.repo = db.Repo{
				Owner: s.app.Config.DefaultUser,
				Name:  parts[0],
			}
		}
	} else {
		s.repo = db.Repo{
			Owner: repoParts[0],
			Name:  repoParts[1],
		}
	}

	return nil
}

func (s *session) Data(r io.Reader) error {
	data, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	user, key, err := s.app.DB.GetUserAndPubKey(s.ctx, s.from)
	if errors.Is(err, errs.ErrInvalidEmail) {
		s.verified = false
	} else if err != nil {
		log.Printf("Error getting user public key: %v\n", err)
		return &smtp.SMTPError{
			Code:    451,
			Message: "Internal server error",
		}
	} else {
		ok, err := utils.VerifySign(s.ctx, key, data)
		if err != nil {
			log.Printf("Error verifing email sign: %v\n", err)
			return &smtp.SMTPError{
				Code:    451,
				Message: "Internal server error",
			}
		}

		if ok {
			log.Println("Verified")
			s.verified = true
			s.user = &user
		} else {
			log.Println("Not Verified")
			s.verified = false
			s.user = nil
		}
	}

	s.email.Sender = db.EmailSender{
		Email:    s.from,
		Name:     s.user,
		Verified: s.verified,
	}

	s.email.Data = data

	if err := s.app.DB.AddEmail(
		s.ctx,
		s.repo,
		s.email,
	); errors.Is(err, errs.ErrInvalidRepo) {
		return &smtp.SMTPError{
			Code:    550,
			Message: "Invalid repo",
		}
	} else if err != nil {
		log.Printf("Error adding email: %v\n", err)
		return &smtp.SMTPError{
			Code:    451,
			Message: "Internal server error",
		}
	}

	return nil
}

func (s *session) Reset() {
	s.from = ""
	s.email = db.Email{}
	s.repo = db.Repo{}
	s.user = nil
	s.verified = false
}

func (s *session) Logout() error {
	return nil
}
