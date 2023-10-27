package server

import (
	"errors"

	"github.com/google/uuid"
)

// Email entity
//   - ID : string
//   - Name : string
//   - Address : string
//   - Port : string
//   - User : string
//   - Password : string
//   - Cryptography : string
type Server struct {
	ID           string `json:"id" gorm:"primaryKey"`
	Name         string `json:"name"`
	Address      string `json:"address"`
	Port         string `json:"port"`
	User         string `json:"user"`
	Password     string `json:"password"`
	Cryptography string `json:"cryptography" gorm:"default: TLS"`
}

// NewServer creates a new instance of the Server struct.
//
// Parameters:
//   - id: The ID of the server.
//   - name: The name of the server.
//   - address: The address of the server.
//   - port: The port of the server.
//   - user: The user for authentication.
//   - password: The password for authentication.
//   - cryptography: The cryptography method used.
//
// Returns:
// - A pointer to the new Server instance.
func NewServer(name, address, port, user, password, cryptography string) *Server {
	return &Server{
		ID:           uuid.New().String(),
		Name:         name,
		Address:      address,
		Port:         port,
		User:         user,
		Password:     password,
		Cryptography: cryptography,
	}
}

// Validate checks if the Server object is valid.
//
// It validates the name, address, port, user, password, and cryptography
// fields of the Server object. If any of these fields are empty, an error
// is returned.
//
// Returns:
//   - error: If any of the required fields are empty.
func (s *Server) Validate() error {
	if s.Name == "" {
		return errors.New("name is required")
	}
	if s.Address == "" {
		return errors.New("address is required")
	}
	if s.Port == "" {
		return errors.New("port is required")
	}
	if s.User == "" {
		return errors.New("user is required")
	}
	if s.Password == "" {
		return errors.New("password is required")
	}

	return nil
}
