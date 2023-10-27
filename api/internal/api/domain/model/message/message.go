package message

import (
	"time"

	"github.com/CarlosEduardoAD/jobbo-api/internal/api/domain/server"
)

type Message struct {
	ID      string `gorm:"primaryKey"`
	UserID  string `json:"userId" query:"userId"`
	From    string `json:"from" query:"from"`
	To      string `json:"to" query:"to"`
	Subject string `json:"subject" query:"subject"`
	Body    string `json:"body" query:"body"`
	SentAt time.Time
	Server server.Server `json:"serverId" query:"serverId"`
	
}
