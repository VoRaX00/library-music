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

func (s *MusicService) Add(music domain.MusicToAdd) (int, error) {
	return s.repo.Add(music)
}

func (s *MusicService) Delete(music domain.MusicToDelete) error {
	return s.repo.Delete(music)
}

func (s *MusicService) Update(music domain.MusicToUpdate) error {
	return s.repo.Update(music)
}

func (s *MusicService) GetAll(page int) ([]domain.Music, error) {
	return s.repo.GetAll()
}

func (s *MusicService) Get(song domain.MusicToGet, page int) (domain.Music, error) {
	return s.repo.Get(song)
}

func (s *MusicService) GetText(music domain.MusicToGet, page int) (string, error) {
	foundMusic, err := s.repo.Get(music)
	if err != nil {
		return "", err
	}
	return foundMusic.Text, nil
}
