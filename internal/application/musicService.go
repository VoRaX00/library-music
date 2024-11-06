package application

import (
	"errors"
	"library-music/internal/infrastructure"
	"library-music/pkg/mapper"
	"strings"
)

type MusicService struct {
	repo   infrastructure.IMusicRepository
	mapper mapper.MusicMapper
}

func NewMusicService(repo infrastructure.IMusicRepository) *MusicService {
	return &MusicService{
		repo:   repo,
		mapper: mapper.MusicMapper{},
	}
}

func (s *MusicService) Add(music MusicToAdd) (int, error) {
	data, err := s.mapper.AddToMusic(music)
	if err != nil {
		return 0, err
	}
	return s.repo.Add(data)
}

func (s *MusicService) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *MusicService) Update(music MusicToUpdate, id int) error {
	data, err := s.mapper.UpdateToMusic(music)
	if err != nil {
		return err
	}
	return s.repo.Update(data, id)
}

func (s *MusicService) GetAll(params MusicFilterParams, page int) ([]MusicToGet, error) {
	res, err := s.repo.GetAll(s.mapper.FilterToMusic(params), page)
	if err != nil {
		return nil, err
	}

	arr := make([]MusicToGet, len(res))
	for i, v := range res {
		arr[i] = s.mapper.MusicForGet(v)
	}
	return arr, nil
}

func (s *MusicService) Get(song, group string) (MusicToGet, error) {
	music, err := s.repo.Get(song, group)
	if err != nil {
		return MusicToGet{}, err
	}
	return s.mapper.MusicForGet(music), nil
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
