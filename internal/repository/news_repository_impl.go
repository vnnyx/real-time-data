package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/rs/zerolog"
	"github.com/vnnyx/real-time-data/internal/config"
	"github.com/vnnyx/real-time-data/internal/domain"
)

type newsRepository struct {
	httpClient *http.Client
	cfg        *config.Config
	log        *zerolog.Logger
}

func NewNewsRepository(httpClient *http.Client, cfg *config.Config, log *zerolog.Logger) NewsRepository {
	return &newsRepository{
		httpClient: httpClient,
		cfg:        cfg,
		log:        log,
	}
}

func (nr *newsRepository) GetNews(ctx context.Context, query string) (*domain.NewsAPIResponse, error) {
	endDate := time.Now().Format("2006-01-02")
	startDate := time.Now().AddDate(0, 0, -7).Format("2006-01-02")
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/everything", nr.cfg.NewsAPIBaseURL), nil)
	if err != nil {
		nr.log.Error().Err(err).Msg("failed to create request")
		return nil, err
	}

	q := req.URL.Query()
	q.Add("q", query)
	q.Add("from", startDate)
	q.Add("to", endDate)
	q.Add("sortBy", "popularity")
	q.Add("apiKey", nr.cfg.NewsAPIKey)
	q.Add("language", "en")

	req.URL.RawQuery = q.Encode()

	resp, err := nr.httpClient.Do(req)
	if err != nil {
		nr.log.Error().Err(err).Msg("failed to do request")
		return nil, err
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		nr.log.Error().Err(err).Msg("failed to read response body")
		return nil, err
	}

	var newsAPIResponse domain.NewsAPIResponse
	err = json.Unmarshal(responseBody, &newsAPIResponse)
	if err != nil {
		nr.log.Error().Err(err).Msg("failed to unmarshal response body")
		return nil, err
	}

	return &newsAPIResponse, nil
}
