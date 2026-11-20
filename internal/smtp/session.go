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
	"github.com/emersion/go-smtp"
)

type session struct {
	ctx    context.Context
	app    *app.App
	domain string

	from     string
	sender   db.EmailSender
	repo     db.Repo
	user     *string
	verified bool
}

func (s *session) Mail(from string, _ *smtp.MailOptions) error {
	s.from = from

	return nil
}

func (s *session) Rcpt(to string, _ *smtp.RcptOptions) error {
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

	verified := true

	user, err := s.app.DB.GetUserByEmail(s.ctx, s.from)
	if errors.Is(err, errs.ErrInvalidEmail) {
		verified = false
	} else if err != nil {
		log.Printf("Error getting user by email: %v\n", err)
		return &smtp.SMTPError{
			Code:    451,
			Message: "Internal server error",
		}
	}

	if verified {
		s.user = &user
		s.verified = true
	}

	s.sender = db.EmailSender{
		Email:    s.from,
		Name:     s.user,
		Verified: s.verified,
	}

	return nil
}

func (s *session) Data(r io.Reader) error {
	data, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	if err := s.app.DB.AddEmail(
		s.ctx,
		s.sender,
		s.repo,
		data,
	); err != nil {
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
	s.sender = db.EmailSender{}
	s.repo = db.Repo{}
	s.user = nil
	s.verified = false
}

func (s *session) Logout() error {
	return nil
}
