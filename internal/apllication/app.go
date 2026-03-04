package apllication

import "context"

type App struct {
}

func (a *App) CreateShortURL(ctx context.Context, fullURL string) (*url.URL, error) {

}

func (a *App) GetFullURL(ctx context.Context, shortUrl string) (*url.URL, error) {

}
