package smtp

import (
	"context"

	"github.com/Ankumeah/SourceFlame/internal/app"
	"github.com/emersion/go-smtp"
)

type backend struct {
	ctx    context.Context
	domain string
	app    *app.App
}

func (b backend) NewSession(*smtp.Conn) (smtp.Session, error) {
	return &session{
		ctx:    b.ctx,
		domain: b.domain,
		app:    b.app,
	}, nil
}

func Listen(ctx context.Context, app *app.App) error {
	srv := smtp.NewServer(
		backend{
			ctx:    ctx,
			domain: app.Config.Domain,
			app:    app,
		},
	)

	srv.Addr = app.Config.SMTPHost
	srv.Domain = app.Config.Domain

	return srv.ListenAndServe()
}
