package service

import (
	"github.com/advor2102/socialnetwork/internal/contracts"
)

type Service struct {
	repository contracts.RepositoryI
}

func NewService(repository contracts.RepositoryI) *Service {
	return &Service{repository: repository}
}
