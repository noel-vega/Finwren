package email

import (
	"context"

	"github.com/resend/resend-go/v3"
)

type Service struct {
	client *resend.Client
}

func NewService(apiKey string) *Service {
	client := resend.NewClient(apiKey)
	return &Service{
		client,
	}
}

type SendParams struct {
	From    string
	To      []string
	HTML    string
	Subject string
}

func (s *Service) send(ctx context.Context, params SendParams) error {
	_, err := s.client.Emails.SendWithContext(ctx, &resend.SendEmailRequest{
		From:    params.From,
		To:      params.To,
		Html:    params.HTML,
		Subject: params.Subject,
	})
	return err
}

type MagicLinkData struct {
	Name             string
	Link             string
	ExpiresInMinutes int
}

func (s *Service) SendMagicLink(ctx context.Context, to string, data MagicLinkData) error {
	html, err := render("magic_link.html", data)
	if err != nil {
		return err
	}

	return s.send(ctx, SendParams{
		From:    "Moolatta <noreplay@auth.moolatta.com>",
		To:      []string{to},
		HTML:    html,
		Subject: "Sign in to you account",
	})
}
