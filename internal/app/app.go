package app

import (
	"github.com/Ankumeah/SourceFlame/internal/db"
	"github.com/Ankumeah/SourceFlame/internal/email"
	"github.com/Ankumeah/SourceFlame/internal/git"
	"github.com/Ankumeah/SourceFlame/internal/hash"
	"github.com/Ankumeah/SourceFlame/internal/otp"
)

type Config struct {
	Host        string
	DBURL       string
	ReposPath   string
	HashConfig  string
	DefaultUser string
	LockUsers   bool

	SMTPHost string
	Domain   string

	EmailURL  string
	EmailUser string
	EmailPass string
}

type App struct {
	Config        Config
	Hasher        hash.Hasher
	GitController git.GitController
	DB            *db.DBController
	OTPManager    *otp.OTPManager
	EmailManager  email.EmailManager
}
