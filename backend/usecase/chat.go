package usecase

import (
	"backend/models/chat"
	"backend/repository"
)

type IChatUsecase interface {
}

type chatUsecase struct {
	cr repository.IChatRepository
}

func NewChatUsecase(cr repository.IChatRepository) IChatUsecase {
	return &chatUsecase{cr}
}

func (cu *chatUsecase) SaveMessage(*chat.Message) error {
	return nil
}
