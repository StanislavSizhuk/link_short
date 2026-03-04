package apllication

import (
	"context"
	"fmt"
	"linkshort/internal/domain/url"
	"linkshort/internal/util"
)

type App struct {
	urlRepo url.Repository
}

func NewApp(r url.Repository) *App {
	return &App{
		urlRepo: r,
	}
}

func (a *App) CreateShortURL(ctx context.Context, fullURL string) (*url.URL, error) {
	var (
		modelURL *url.URL
		err      error
	)
	modelURL = url.NewURL(fullURL)

	modelURL.SetShort(util.ShortenURL(modelURL.Full()))

	err = a.urlRepo.SaveURL(ctx, modelURL)
	if err != nil {
		return nil, fmt.Errorf("error save url: %w", err)
	}
	return modelURL, nil
}

func (a *App) GetFullURL(ctx context.Context, shortUrl string) (*url.URL, error) {
	var (
		modelURL *url.URL
		err      error
	)
	modelURL, err = a.urlRepo.GetURLByShort(ctx, shortUrl)
	if err != nil {
		return nil, fmt.Errorf("error get url by : %w", err)
	}
	return modelURL, nil
}
