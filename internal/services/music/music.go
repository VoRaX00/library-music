package music

import (
	"errors"
	"fmt"
	"library-music/internal/services"
	"library-music/internal/storage/music"
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
	ErrMusicNotFound      = errors.New("music not found")
	ErrMusicAlreadyExists = errors.New("music already exists")
)

func New(log *slog.Logger, repo Repo) *Music {
	return &Music{
		log:    log,
		repo:   repo,
		mapper: mapper.MusicMapper{},
	}
}

func (s *Music) Add(music services.MusicToAdd) (int, error) {
	const op = "music.Add"
	log := s.log.With(
		slog.String("op", op),
	)

	data, err := s.mapper.AddToMusic(music)
	if err != nil {
		log.Info("error mapping", err.Error())
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	log.Info("adding a song")
	id, err := s.repo.Add(data)
	if err != nil {
		if errors.Is(err, musicrepo.ErrMusicAlreadyExists) {
			log.Warn("music already exists", err.Error())
			return 0, fmt.Errorf("%s: %w", op, ErrMusicAlreadyExists)
		}

		log.Error("failed to add a song", err.Error())
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	log.Info("successfully added a song")

	return id, err
}

func (s *Music) Delete(id int) error {
	const op = "music.Delete"
	log := s.log.With(
		slog.String("op", op),
	)

	log.Info("deleting a song")
	err := s.repo.Delete(id)
	if err != nil {
		if errors.Is(err, musicrepo.ErrMusicNotFound) {
			log.Warn("music not found", err.Error())
			return fmt.Errorf("%s: %w", op, ErrMusicNotFound)
		}
		log.Error("failed to delete a song", err.Error())
		return fmt.Errorf("%s: %w", op, err)
	}
	log.Info("successfully deleted a song")
	return nil
}

func (s *Music) Update(music services.MusicToUpdate, id int) error {
	const op = "music.Update"
	log := s.log.With(
		slog.String("op", op),
	)

	data, err := s.mapper.UpdateToMusic(music)
	if err != nil {
		log.Info("error mapping", err.Error())
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("updating a song")
	err = s.repo.Update(data, id)
	if err != nil {
		if errors.Is(err, musicrepo.ErrMusicNotFound) {
			log.Warn("music not found", err.Error())
			return fmt.Errorf("%s: %w", op, ErrMusicNotFound)
		}

		if errors.Is(err, musicrepo.ErrMusicAlreadyExists) {
			log.Warn("music already exists", err.Error())
			return fmt.Errorf("%s: %w", op, ErrMusicAlreadyExists)
		}

		log.Error("failed to update a song", err.Error())
		return fmt.Errorf("%s: %w", op, err)
	}
	log.Info("successfully updated a song")
	return nil
}

func (s *Music) GetAll(params services.MusicFilterParams, page int) ([]services.MusicToGet, error) {
	const op = "music.GetAll"
	log := s.log.With(
		slog.String("op", op),
	)

	log.Info("fetching all songs")
	res, err := s.repo.GetAll(s.mapper.FilterToMusic(params), page)
	if err != nil {
		if errors.Is(err, musicrepo.ErrMusicNotFound) {
			log.Error("failed to get all songs", err.Error())
			return nil, fmt.Errorf("%s: %w", op, ErrMusicNotFound)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	arr := make([]services.MusicToGet, len(res))
	for i, v := range res {
		arr[i] = s.mapper.MusicForGet(v)
	}
	log.Info("successfully fetched all songs")
	return arr, nil
}

func (s *Music) Get(song, group string) (services.MusicToGet, error) {
	const op = "music.Get"
	log := s.log.With(
		slog.String("op", op),
	)

	log.Info("fetching a song")
	music, err := s.repo.Get(song, group)
	if err != nil {
		if errors.Is(err, musicrepo.ErrMusicNotFound) {
			log.Warn("music not found", err.Error())
			return services.MusicToGet{}, fmt.Errorf("%s: %w", op, ErrMusicNotFound)
		}
		log.Error("failed to get a song", err.Error())
		return services.MusicToGet{}, fmt.Errorf("%s: %w", op, err)
	}
	log.Info("successfully fetched a song")
	return s.mapper.MusicForGet(music), nil
}

func (s *Music) GetText(song, group string, page int) (string, error) {
	const op = "music.GetText"
	log := s.log.With(
		slog.String("op", op),
	)

	log.Info("fetching a song")
	text, err := s.repo.GetText(song, group)
	if err != nil {
		if errors.Is(err, musicrepo.ErrMusicNotFound) {
			log.Warn("music not found", ErrMusicNotFound)
			return "", fmt.Errorf("%s: %w", op, ErrMusicNotFound)
		}

		log.Error("failed to fetch a song", err.Error())
		return "", fmt.Errorf("%s: %w", op, err)
	}

	verses := strings.Split(text, "\n\n")
	if len(verses) < page {
		log.Error("page is out of range")
		return "", fmt.Errorf("%s: %w", op, ErrMusicNotFound)
	}
	log.Info("successfully fetched a song")
	return verses[page-1], nil
}
