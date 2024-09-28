package application

import (
	"library-music/internal/domain"
	"library-music/internal/infrastructure"
)

type MusicService struct {
	repo *infrastructure.Repository
}

func NewMusicService(repo *infrastructure.Repository) *MusicService {
	return &MusicService{repo: repo}
}

func (s *MusicService) Add(music domain.Music) error {
	return s.repo.Add(music)
}

func (s *MusicService) Delete(song string) error {
	return s.repo.Delete(song)
}

func (s *MusicService) Update(music domain.Music) error {
	return s.repo.Update(music)
}

func (s *MusicService) GetAll() ([]domain.Music, error) {
	return s.repo.GetAll()
}

func (s *MusicService) Get(song string) (domain.Music, error) {
	return s.repo.Get(song)
}

func (s *MusicService) GetText(song string) (string, error) {
	music, err := s.repo.Get(song)
	if err != nil {
		return "", err
	}
	return music.Text, nil
}
