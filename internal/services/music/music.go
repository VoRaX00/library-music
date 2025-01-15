package music

import (
	"errors"
	"fmt"
	"library-music/internal/domain/models"
	"library-music/internal/services"
	"library-music/internal/storage/music"
	"library-music/pkg/mapper"
	"log/slog"
	"strconv"
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

func (s *Music) Add(music models.Music) (int, error) {
	const op = "music.Add"
	log := s.log.With(
		slog.String("op", op),
	)

	log.Debug(
		"adding song",
		slog.String("song", music.Song),
		slog.String("group", music.Group.Name),
		slog.String("Text", music.Text),
		slog.String("Link", music.Link),
		slog.String("ReleaseDate", music.ReleaseDate.String()),
	)

	log.Info("start adding song")
	id, err := s.repo.Add(music)
	if err != nil {
		if errors.Is(err, musicrepo.ErrMusicAlreadyExists) {
			log.Warn("music already exists", slog.String("err", err.Error()))
			return 0, fmt.Errorf("%s: %w", op, ErrMusicAlreadyExists)
		}
		log.Error("failed to add a song", slog.String("err", err.Error()))
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	log.Info("successfully added a song")
	log.Debug(
		"id added song",
		slog.String("id", strconv.FormatInt(int64(id), 10)),
	)

	return id, err
}

func (s *Music) Delete(id int) error {
	const op = "music.Delete"
	log := s.log.With(
		slog.String("op", op),
	)

	log.Debug(
		"deleting song",
		slog.String("id", strconv.FormatInt(int64(id), 10)),
	)
	log.Info("start deleting a song")
	err := s.repo.Delete(id)
	if err != nil {
		if errors.Is(err, musicrepo.ErrMusicNotFound) {
			log.Warn("music not found", slog.String("err", err.Error()))
			return fmt.Errorf("%s: %w", op, ErrMusicNotFound)
		}
		log.Error("failed to delete a song", slog.String("err", err.Error()))
		return fmt.Errorf("%s: %w", op, err)
	}
	log.Info("successfully deleted a song")
	log.Debug(
		"delete song",
		slog.String("id", strconv.FormatInt(int64(id), 10)),
	)

	return nil
}

func (s *Music) Update(music services.MusicToUpdate, id int) error {
	const op = "music.Update"
	log := s.log.With(
		slog.String("op", op),
	)

	log.Debug(
		"updating song",
		slog.String("id", strconv.FormatInt(int64(id), 10)),
		slog.String("song", music.Song),
		slog.String("group", music.Group),
		slog.String("Text", music.Text),
		slog.String("Link", music.Link),
		slog.String("ReleaseDate", music.ReleaseDate),
	)

	data, err := s.mapper.UpdateToMusic(music)
	if err != nil {
		log.Warn("error mapping", slog.String("err", err.Error()))
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("start updating a song")
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
	log.Debug(
		"updated a song",
		slog.String("id", strconv.FormatInt(int64(id), 10)),
	)
	return nil
}

func (s *Music) GetAll(params services.MusicFilterParams, countSongs, page int) ([]services.MusicToGet, error) {
	const op = "music.GetAll"
	log := s.log.With(
		slog.String("op", op),
	)

	log.Debug(
		"parameters",
		slog.String("song", params.Song),
		slog.String("group", params.Group),
		slog.String("text", params.Text),
		slog.String("link", params.Link),
		slog.String("releaseData", params.ReleaseDate.String()),
		slog.String("countSongs", strconv.FormatInt(int64(countSongs), 10)),
		slog.String("page", strconv.FormatInt(int64(page), 10)),
	)

	log.Info("start fetching all songs")
	res, err := s.repo.GetAll(s.mapper.FilterToMusic(params), countSongs, page)
	if err != nil {
		if errors.Is(err, musicrepo.ErrMusicNotFound) {
			log.Warn("failed to get all songs", slog.String("err", err.Error()))
			return nil, fmt.Errorf("%s: %w", op, ErrMusicNotFound)
		}
		log.Error("failed to fetch all songs", slog.String("err", err.Error()))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	arr := make([]services.MusicToGet, len(res))
	for i, v := range res {
		arr[i] = s.mapper.MusicForGet(v)
	}
	log.Info("successfully fetched all songs")
	log.Debug(fmt.Sprintf("%d songs returned", len(res)))
	return arr, nil
}

func (s *Music) Get(song, group string) (services.MusicToGet, error) {
	const op = "music.Get"
	log := s.log.With(
		slog.String("op", op),
	)

	log.Debug(
		"fetching song",
		slog.String("song", song),
		slog.String("group", group),
	)

	log.Info("start fetching a song")
	music, err := s.repo.Get(song, group)
	if err != nil {
		if errors.Is(err, musicrepo.ErrMusicNotFound) {
			log.Warn("music not found", slog.String("err", err.Error()))
			return services.MusicToGet{}, fmt.Errorf("%s: %w", op, ErrMusicNotFound)
		}
		log.Error("failed to get a song", slog.String("err", err.Error()))
		return services.MusicToGet{}, fmt.Errorf("%s: %w", op, err)
	}
	log.Info("successfully fetched a song")

	log.Debug(
		"returned song",
		slog.String("id", strconv.FormatInt(int64(music.Id), 10)),
		slog.String("song", music.Song),
		slog.String("group", music.Group.Name),
		slog.String("text", music.Text),
		slog.String("link", music.Link),
		slog.String("releaseData", music.ReleaseDate.String()),
	)
	return s.mapper.MusicForGet(music), nil
}

func (s *Music) GetText(song, group string, countVerse, page int) (string, error) {
	const op = "music.GetText"
	log := s.log.With(
		slog.String("op", op),
	)

	log.Debug(
		"getting song",
		slog.String("song", song),
		slog.String("group", group),
		slog.String("countVerse", strconv.FormatInt(int64(countVerse), 10)),
		slog.String("page", strconv.FormatInt(int64(countVerse), 10)),
	)

	log.Info("fetching a song")
	text, err := s.repo.GetText(song, group)
	if err != nil {
		if errors.Is(err, musicrepo.ErrMusicNotFound) {
			log.Warn("music not found", slog.String("err", ErrMusicNotFound.Error()))
			return "", fmt.Errorf("%s: %w", op, ErrMusicNotFound)
		}

		log.Error("failed to fetch a song", slog.String("err", err.Error()))
		return "", fmt.Errorf("%s: %w", op, err)
	}

	verses := strings.Split(text, "\n\n")
	if len(verses)/countVerse < page {
		log.Warn("page is out of range")
		return "", fmt.Errorf("%s: %w", op, ErrMusicNotFound)
	}
	log.Info("successfully fetched a song")

	start := countVerse * (page - 1)
	end := countVerse*(page-1) + countVerse
	result := strings.Join(verses[start:end], "\n\n")

	log.Debug("text", result)
	return result, nil
}
