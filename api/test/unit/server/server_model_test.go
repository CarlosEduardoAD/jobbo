package unit

import (
	"errors"

	"testing"

	"github.com/CarlosEduardoAD/jobbo-api/internal/api/domain/model/server"
	"github.com/google/uuid"
)

func TestNewServer(t *testing.T) {
	name := "Test Server"
	address := "127.0.0.1"
	port := "8080"
	user := "testuser"
	password := "testpassword"
	cryptography := "TLS"

	s := server.NewServer(name, address, port, user, password, cryptography)

	if s.ID == "" {
		t.Error("NewServer() did not set ID")
	}

	if s.Name != name {
		t.Error("NewServer() did not set Name")
	}

	if s.Address != address {
		t.Error("NewServer() did not set Address")
	}

	if s.Port != port {
		t.Error("NewServer() did not set Port")
	}

	if s.User != user {
		t.Error("NewServer() did not set User")
	}

	if s.Password != password {
		t.Error("NewServer() did not set Password")
	}

	if s.Cryptography != cryptography {
		t.Error("NewServer() did not set Cryptography")
	}
}

func TestServer_Validate(t *testing.T) {
	tests := []struct {
		name     string
		server   *server.Server
		expected error
	}{
		{
			name: "valid server",
			server: &server.Server{
				ID:           uuid.New().String(),
				Name:         "Test Server",
				Address:      "127.0.0.1",
				Port:         "8080",
				User:         "testuser",
				Password:     "testpassword",
				Cryptography: "TLS",
			},
			expected: nil,
		},
		{
			name: "missing name",
			server: &server.Server{
				ID:           uuid.New().String(),
				Address:      "127.0.0.1",
				Port:         "8080",
				User:         "testuser",
				Password:     "testpassword",
				Cryptography: "TLS",
			},
			expected: errors.New("name is required"),
		},
		{
			name: "missing address",
			server: &server.Server{
				ID:           uuid.New().String(),
				Name:         "Test Server",
				Port:         "8080",
				User:         "testuser",
				Password:     "testpassword",
				Cryptography: "TLS",
			},
			expected: errors.New("address is required"),
		},
		{
			name: "missing port",
			server: &server.Server{
				ID:           uuid.New().String(),
				Name:         "Test Server",
				Address:      "127.0.0.1",
				User:         "testuser",
				Password:     "testpassword",
				Cryptography: "TLS",
			},
			expected: errors.New("port is required"),
		},
		{
			name: "missing user",
			server: &server.Server{
				ID:           uuid.New().String(),
				Name:         "Test Server",
				Address:      "127.0.0.1",
				Port:         "8080",
				Password:     "testpassword",
				Cryptography: "TLS",
			},
			expected: errors.New("user is required"),
		},
		{
			name: "missing password",
			server: &server.Server{
				ID:           uuid.New().String(),
				Name:         "Test Server",
				Address:      "127.0.0.1",
				Port:         "8080",
				User:         "testuser",
				Cryptography: "TLS",
			},
			expected: errors.New("password is required"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.server.Validate()

			if err == nil && tt.expected != nil {
				t.Errorf("Validate() returned nil, expected %v", tt.expected)
			} else if err != nil && tt.expected == nil {
				t.Errorf("Validate() returned %v, expected nil", err)
			} else if err != nil && tt.expected != nil && err.Error() != tt.expected.Error() {
				t.Errorf("Validate() returned %v, expected %v", err, tt.expected)
			}
		})
	}
}
