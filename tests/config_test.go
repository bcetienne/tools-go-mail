package tests

import (
	"testing"
	"time"

	gomail "github.com/bcetienne/tools-go-mail"
)

// TestNewConfig_DefaultValues checks if NewConfig correctly sets default values
// when no overriding options are provided.
func TestNewConfig_DefaultValues(t *testing.T) {
	// To pass validation, we must provide a username and password.
	cfg, err := gomail.NewConfig(
		gomail.WithUsername("testuser"),
		gomail.WithPassword("testpass"),
	)

	if err != nil {
		t.Fatalf("NewConfig() with required fields returned an unexpected error: %v", err)
	}

	if cfg.Host != "smtp.gmail.com" {
		t.Errorf("Expected Host to be 'smtp.gmail.com', got '%s'", cfg.Host)
	}
	if cfg.Port != 587 {
		t.Errorf("Expected Port to be 587, got %d", cfg.Port)
	}
	if cfg.Timeout != 30*time.Second {
		t.Errorf("Expected Timeout to be 30s, got %v", cfg.Timeout)
	}
	if cfg.AuthMethod != "PLAIN" {
		t.Errorf("Expected AuthMethod to be 'PLAIN', got '%s'", cfg.AuthMethod)
	}
	if cfg.Username != "testuser" {
		t.Errorf("Expected Username to be 'testuser', got '%s'", cfg.Username)
	}
	if cfg.Password != "testpass" {
		t.Errorf("Expected Password to be 'testpass', got '%s'", cfg.Password)
	}
}

// TestNewConfig_WithOptions verifies that all functional options correctly
// override the default configuration values.
func TestNewConfig_WithOptions(t *testing.T) {
	host := "mail.example.com"
	port := 465
	username := "user"
	password := "pass"
	from := "sender@example.com"
	fromName := "Sender Name"
	insecureSkipVerify := true
	timeout := 15 * time.Second
	keepAlive := true
	authMethod := "CRAM-MD5"

	cfg, err := gomail.NewConfig(
		gomail.WithHost(host),
		gomail.WithPort(port),
		gomail.WithUsername(username),
		gomail.WithPassword(password),
		gomail.WithFrom(from),
		gomail.WithFromName(fromName),
		gomail.WithInsecureSkipVerify(insecureSkipVerify),
		gomail.WithTimeout(timeout),
		gomail.WithKeepAlive(keepAlive),
		gomail.WithAuthMethod(authMethod),
	)

	if err != nil {
		t.Fatalf("NewConfig() with all options returned an unexpected error: %v", err)
	}

	if cfg.Host != host {
		t.Errorf("Expected Host to be '%s', got '%s'", host, cfg.Host)
	}
	if cfg.Port != port {
		t.Errorf("Expected Port to be %d, got %d", port, cfg.Port)
	}
	if cfg.Username != username {
		t.Errorf("Expected Username to be '%s', got '%s'", username, cfg.Username)
	}
	if cfg.Password != password {
		t.Errorf("Expected Password to be '%s', got '%s'", password, cfg.Password)
	}
	if cfg.From != from {
		t.Errorf("Expected From to be '%s', got '%s'", from, cfg.From)
	}
	if cfg.FromName != fromName {
		t.Errorf("Expected FromName to be '%s', got '%s'", fromName, cfg.FromName)
	}
	if cfg.InsecureSkipVerify != insecureSkipVerify {
		t.Errorf("Expected InsecureSkipVerify to be %v, got %v", insecureSkipVerify, cfg.InsecureSkipVerify)
	}
	if cfg.Timeout != timeout {
		t.Errorf("Expected Timeout to be %v, got %v", timeout, cfg.Timeout)
	}
	if cfg.KeepAlive != keepAlive {
		t.Errorf("Expected KeepAlive to be %v, got %v", keepAlive, cfg.KeepAlive)
	}
	if cfg.AuthMethod != authMethod {
		t.Errorf("Expected AuthMethod to be '%s', got '%s'", authMethod, cfg.AuthMethod)
	}
}

// TestNewConfig_Validation checks the validation logic within the NewConfig constructor.
func TestNewConfig_Validation(t *testing.T) {
	testCases := []struct {
		name          string
		options       []func(*gomail.Config)
		expectedError string
	}{
		{
			name: "Missing username",
			options: []func(*gomail.Config){
				gomail.WithPassword("password"),
			},
			expectedError: "username is empty",
		},
		{
			name: "Missing password",
			options: []func(*gomail.Config){
				gomail.WithUsername("username"),
			},
			expectedError: "password is empty",
		},
		{
			name:          "Missing both username and password",
			options:       []func(*gomail.Config){},
			expectedError: "username is empty",
		},
		{
			name: "Valid config with required fields",
			options: []func(*gomail.Config){
				gomail.WithUsername("username"),
				gomail.WithPassword("password"),
			},
			expectedError: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := gomail.NewConfig(tc.options...)
			if tc.expectedError != "" {
				if err == nil {
					t.Fatal("Expected an error but got nil")
				}
				if err.Error() != tc.expectedError {
					t.Errorf("Expected error message '%s', got '%s'", tc.expectedError, err.Error())
				}
			} else {
				if err != nil {
					t.Fatalf("Expected no error, but got: %v", err)
				}
			}
		})
	}
}

// TestConfig_Validate tests the Validate method directly.
func TestConfig_Validate(t *testing.T) {
	testCases := []struct {
		name          string
		config        *gomail.Config
		expectedError string
	}{
		{
			name: "Valid config",
			config: &gomail.Config{
				Username: "user",
				Password: "password",
			},
			expectedError: "",
		},
		{
			name: "Missing username",
			config: &gomail.Config{
				Password: "password",
			},
			expectedError: "username is empty",
		},
		{
			name: "Missing password",
			config: &gomail.Config{
				Username: "user",
			},
			expectedError: "password is empty",
		},
		{
			name:          "Missing both username and password",
			config:        &gomail.Config{},
			expectedError: "username is empty",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.config.Validate()
			if tc.expectedError != "" {
				if err == nil {
					t.Fatal("Expected an error but got nil")
				}
				if err.Error() != tc.expectedError {
					t.Errorf("Expected error message '%s', got '%s'", tc.expectedError, err.Error())
				}
			} else {
				if err != nil {
					t.Fatalf("Expected no error, but got: %v", err)
				}
			}
		})
	}
}
