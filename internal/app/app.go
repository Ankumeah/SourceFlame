package app

import (
	"github.com/Ankumeah/SourceFlame/internal/db"
	"github.com/Ankumeah/SourceFlame/internal/git"
	"github.com/Ankumeah/SourceFlame/internal/hash"
)

type Config struct {
	ApiVersion string
	Host       string
	DBURL      string
	ReposPath  string
	HashConfig string
}

type App struct {
	Config        Config
	Hasher        hash.Hasher
	GitController git.GitController
	DB            *db.DBController
}
