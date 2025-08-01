package gomail

import (
	"errors"
	"time"
)

type Config struct {
	Host               string
	Port               int
	Username           string
	Password           string
	From               string
	FromName           string
	InsecureSkipVerify bool
	Timeout            time.Duration
	KeepAlive          bool
	AuthMethod         string
}

func NewConfig(options ...func(*Config)) (*Config, error) {
	// Create a config struct with default values
	cfg := &Config{
		Host:       "smtp.gmail.com",
		Port:       587,
		Timeout:    30 * time.Second,
		AuthMethod: "PLAIN",
	}

	// Override default values / missing values with functional options
	for _, option := range options {
		option(cfg)
	}

	return cfg, nil
}

func WithHost(host string) func(*Config) {
	return func(cfg *Config) {
		cfg.Host = host
	}
}

func WithPort(port int) func(*Config) {
	return func(cfg *Config) {
		cfg.Port = port
	}
}

func WithUsername(username string) func(*Config) {
	return func(cfg *Config) {
		cfg.Username = username
	}
}

func WithPassword(password string) func(*Config) {
	return func(cfg *Config) {
		cfg.Password = password
	}
}

func WithFrom(from string) func(*Config) {
	return func(cfg *Config) {
		cfg.From = from
	}
}

func WithFromName(fromName string) func(*Config) {
	return func(cfg *Config) {
		cfg.FromName = fromName
	}
}

func WithInsecureSkipVerify(insecureSkipVerify bool) func(*Config) {
	return func(cfg *Config) {
		cfg.InsecureSkipVerify = insecureSkipVerify
	}
}

func WithTimeout(timeout time.Duration) func(*Config) {
	return func(cfg *Config) {
		cfg.Timeout = timeout
	}
}

func WithKeepAlive(keepAlive bool) func(*Config) {
	return func(cfg *Config) {
		cfg.KeepAlive = keepAlive
	}
}

func WithAuthMethod(authMethod string) func(*Config) {
	return func(cfg *Config) {
		cfg.AuthMethod = authMethod
	}
}

func (c *Config) Validate() error {
	if c.Username == "" {
		return errors.New("username is empty")
	}
	if c.Password == "" {
		return errors.New("password is empty")
	}
	return nil
}
