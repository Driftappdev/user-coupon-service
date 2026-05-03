package admincontrol

import (
	"os"
	"strings"
)

type Config struct {
	Subject string
	Durable string
}

func LoadConfig(defaultSubject, defaultDurable string) Config {
	subject := strings.TrimSpace(os.Getenv("ADMIN_CONTROL_SUBJECT"))
	if subject == "" {
		subject = defaultSubject
	}
	durable := strings.TrimSpace(os.Getenv("ADMIN_CONTROL_DURABLE"))
	if durable == "" {
		durable = defaultDurable
	}
	return Config{Subject: subject, Durable: durable}
}

