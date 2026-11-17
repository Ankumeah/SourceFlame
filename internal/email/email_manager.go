package email

import "context"

type EmailManager interface {
	SendEmail(ctx context.Context, email string, data []byte) error
}
