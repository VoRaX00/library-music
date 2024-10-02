package application

import (
	"errors"
	"library-music/internal/domain"
	"library-music/internal/infrastructure"
	"strings"
)

type MusicService struct {
	repo infrastructure.IMusicRepository
}

func NewMusicService(repo infrastructure.IMusicRepository) *MusicService {
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

func (s *MusicService) GetAll(params domain.MusicFilterParams, page int) ([]domain.MusicToGet, error) {
	return s.repo.GetAll(params, page)
}

func (s *MusicService) Get(song, group string) (domain.MusicToGet, error) {
	return s.repo.Get(song, group)
}

func (s *MusicService) GetText(song, group string, page int) (string, error) {
	text, err := s.repo.GetText(song, group)
	if err != nil {
		return "", err
	} else if page < 1 {
		return "", errors.New("page is less than 1")
	}

	verses := strings.Split(text, "\n\n")
	if len(verses) <= page {
		return "", errors.New("page is out of range")
	}
	return verses[page-1], nil
}
