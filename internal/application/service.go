package application

import "library-music/internal/infrastructure"

type Service struct {
	IMusicService
}

func NewService(repos *infrastructure.Repository) *Service {
	return &Service{
		IMusicService: NewMusicService(repos.IMusicRepository),
	}
}
