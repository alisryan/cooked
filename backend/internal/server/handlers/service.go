package handlers

import "github.com/GenerateNU/cooked/backend/internal/storage"

type Service struct {
	store storage.Storage
}

func NewService(storage storage.Storage) *Service {
	return &Service{storage}
}
