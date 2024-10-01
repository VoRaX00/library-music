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

func (s *MusicService) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *MusicService) Update(music domain.MusicToUpdate, id int) error {
	return s.repo.Update(music, id)
}

func (s *MusicService) GetAll(page int) ([]domain.MusicToGet, error) {
	return s.repo.GetAll(page)
}

func (s *MusicService) Get(song, group string) (domain.MusicToGet, error) {
	return s.repo.Get(song, group)
}

func (s *MusicService) GetText(song, group string, page int) (string, error) {
	return s.repo.GetText(song, group)
}
