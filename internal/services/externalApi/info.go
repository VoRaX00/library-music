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

	log.Info("getting info")
	var info services.SongDetail
	resp, err := s.FetchInfo(group, song)
	if err != nil {
		return services.SongDetail{}, fmt.Errorf("%s: %w", op, err)
	}

	if resp == nil {
		log.Warn("The api request did not return anything")
	} else if resp.StatusCode != http.StatusOK {
		log.Warn(fmt.Sprintf("The api request ended with the code %d", resp.StatusCode))
	} else {
		defer resp.Body.Close()
		err = json.NewDecoder(resp.Body).Decode(&info)
		if err != nil {
			return services.SongDetail{}, fmt.Errorf("%s: %w", op, err)
		}
	}
	log.Info("info received")
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
