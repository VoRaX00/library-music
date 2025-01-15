package externalApi

import (
	"encoding/json"
	"fmt"
	"library-music/internal/services"
	"log/slog"
	"net/http"
	"net/url"
	"os"
)

type ExternalApi struct {
	log *slog.Logger
}

func New(log *slog.Logger) *ExternalApi {
	return &ExternalApi{
		log: log,
	}
}

func (s *ExternalApi) Info(song, group string) (services.SongDetail, error) {
	const op = "externalApi.Info"
	log := s.log.With(
		slog.String("op", op),
		slog.String("song", song),
		slog.String("group", group),
	)

	log.Debug(
		"remade song",
		slog.String("song", song),
		slog.String("group", group),
	)

	log.Info("getting info")
	var info services.SongDetail
	resp, err := s.FetchInfo(group, song)

	if err != nil {
		log.Error("error fetching info", slog.String("err", err.Error()))
		return services.SongDetail{}, fmt.Errorf("%s: %w", op, err)
	}

	if resp == nil {
		log.Error("The api request did not return anything")
		return services.SongDetail{}, fmt.Errorf("%s: %s", op, "empty request")
	}

	if resp.StatusCode != http.StatusOK {
		log.Error(fmt.Sprintf("The api request ended with the code %d", resp.StatusCode))
		return services.SongDetail{}, fmt.Errorf("%s: %s", op, "status not OK")
	}

	defer resp.Body.Close()
	log.Debug("response", resp.Body)

	err = json.NewDecoder(resp.Body).Decode(&info)
	if err != nil {
		log.Error("error decoding info", slog.String("err", err.Error()))
		return services.SongDetail{}, fmt.Errorf("%s: %w", op, err)
	}

	log.Info("info received")
	log.Debug(
		"info for song",
		slog.String("text", info.Text),
		slog.String("releaseDate", info.ReleaseDate),
		slog.String("link", info.Link),
	)
	return info, nil
}

func (s *ExternalApi) FetchInfo(group, song string) (*http.Response, error) {
	link := fmt.Sprintf("%s/info", os.Getenv("API"))

	params := url.Values{}
	params.Add("group", group)
	params.Add("song", song)
	requestURL := fmt.Sprintf("%s?%s", link, params.Encode())

	resp, err := http.Get(requestURL)
	return resp, err
}
