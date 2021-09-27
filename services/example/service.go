package example

import (
	"github.com/rs/zerolog"
)

// Service represents a use case service.
type Service struct {
	repository Repository
	logger     *zerolog.Logger
}

// NewService creates a new example service.
func NewService(repository Repository, logger *zerolog.Logger) *Service {
	return &Service{
		repository: repository,
		logger:     logger,
	}
}

func (s *Service) GetExample(id string) ()
