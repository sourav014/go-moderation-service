package notification

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// SendEmail sends an email using SendGrid with retries.
func SendEmail(toEmail, subject, body string) error {
	from := mail.NewEmail("Moderation Service", "tryit.sourav@gmail.com")
	to := mail.NewEmail("User", toEmail)

	message := mail.NewSingleEmail(from, subject, to, body, body)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))

	operation := func() error {
		response, err := client.Send(message)
		if err != nil {
			log.Printf("SendGrid API request failed: %v", err)
			return err
		}

		if response.StatusCode >= 400 {
			errMsg := fmt.Errorf("SendGrid error: %s", response.Body)
			log.Printf("SendGrid returned error: %s", errMsg)
			return errMsg
		}

		log.Println("Email sent successfully!")
		return nil
	}

	// Retry with exponential backoff (max 3 attempts)
	notify := func(err error, duration time.Duration) {
		log.Printf("Retrying due to error: %v (waiting %s)", err, duration)
	}

	expBackoff := backoff.NewExponentialBackOff()
	expBackoff.MaxElapsedTime = 30 * time.Second // Stop retrying after 30s

	err := backoff.RetryNotify(operation, expBackoff, notify)
	if err != nil {
		log.Printf("Failed to send email after retries: %v", err)
		return err
	}

	return nil
}
