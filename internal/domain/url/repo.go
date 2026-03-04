package url

import "context"

type Repository interface {
	SaveURL(ctx context.Context, url *URL) error
	GetURLByFull(ctx context.Context, fullUrl string) (*URL, error)
	GetURLByShort(ctx context.Context, short string) (*URL, error)
}
