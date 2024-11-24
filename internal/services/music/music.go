package music

import (
	"errors"
	"library-music/internal/domain"
	"library-music/pkg/mapper"
	"log/slog"
	"strings"
)

type Music struct {
	log    *slog.Logger
	repo   Repo
	mapper mapper.MusicMapper
}

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrMusicExists        = errors.New("music already exists")
	ErrMusicNotFound      = errors.New("music not found")
)

func New(log *slog.Logger, repo Repo) *Music {
	return &Music{
		log:    log,
		repo:   repo,
		mapper: mapper.MusicMapper{},
	}
}

func (s *Music) Add(music ToAdd) (int, error) {
	data, err := s.mapper.AddToMusic(music)
	if err != nil {
		return 0, err
	}
	return s.repo.Add(data)
}

func (s *Music) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *Music) Update(music ToUpdate, id int) (domain.Music, error) {
	data, err := s.mapper.UpdateToMusic(music)
	if err != nil {
		return domain.Music{}, err
	}
	return s.repo.Update(data, id)
}

func (s *Music) GetAll(params FilterParams, page int) ([]ToGet, error) {
	res, err := s.repo.GetAll(s.mapper.FilterToMusic(params), page)
	if err != nil {
		return nil, err
	}

	arr := make([]ToGet, len(res))
	for i, v := range res {
		arr[i] = s.mapper.MusicForGet(v)
	}
	return arr, nil
}

func (s *Music) Get(song, group string) (ToGet, error) {
	music, err := s.repo.Get(song, group)
	if err != nil {
		return ToGet{}, err
	}
	return s.mapper.MusicForGet(music), nil
}

func (s *Music) GetText(song, group string, page int) (string, error) {
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
