package repository

import (
	"github.com/jmoiron/sqlx"
	Intern "github.com/m1kr0b/message-processing/model/message"
)

type MessageRepository interface {
	CreateMessage(message Intern.Message) (int, error)
	ProcessMessage(id int)
	GetMessageById(id int) (Intern.Message, error)
	GetStats() (int, error)
}

type Repository struct {
	MessageRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		MessageRepository: NewMessagePostgresDB(db),
	}
}
