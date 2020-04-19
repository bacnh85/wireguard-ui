package model

import (
	"time"
)

// Server model
type Server struct {
	KeyPair   *ServerKeypair
	Interface *ServerInterface
}

// ServerKeypair model
type ServerKeypair struct {
	PrivateKey string    `json:"private_key"`
	PublicKey  string    `json:"pulbic_key"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// ServerInterface model
type ServerInterface struct {
	Addresses  []string  `json:"addresses"`
	ListenPort int       `json:"listen_port,string"` // ,string to get listen_port string input as int
	UpdatedAt  time.Time `json:"updated_at"`
}