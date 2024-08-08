package service

import (
	"github.com/m1kr0b/message-processing/internal/kafka"
	"github.com/m1kr0b/message-processing/internal/repository"
	Intern "github.com/m1kr0b/message-processing/model/message"
)

type Message interface {
	CreateMessage(message Intern.Message) (int, error)
	ProcessMessage(id int)
	GetMessageById(id int) (Intern.Message, error)
	GetStats() (int, error)
}

type Service struct {
	Message
}

func NewService(repos *repository.Repository, producer *kafka.Producer) *Service {
	return &Service{
		Message: NewMessageService(repos, producer),
	}
}
