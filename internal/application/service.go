package application

import "library-music/internal/infrastructure"

type Service struct {
	repos *infrastructure.Repository
}

func NewService(repos *infrastructure.Repository) *Service {
	return &Service{
		repos: repos,
	}
}
