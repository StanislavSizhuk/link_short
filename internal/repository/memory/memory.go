package memory

import (
	"context"
	"linkshort/internal/domain/url"
)

type Storage struct {
	URLs map[string]*url.URL
}

func (s *Storage) SaveURl(_ context.Context, url *url.URL) error {
	s.URLs[(url.Short())] = url
	return nil
}
func (s *Storage) GetURLByFull(_ context.Context, fullUrl string) (*url.URL, error) {
	for _, url := range s.URLs {
		if url.Full() == fullUrl {
			return url, nil
		}

	}
	return nil, url.ErrURLNotFound
}
func (s *Storage) GetURLByShort(_ context.Context, short string) (*url.URL, error) {
	if el, ok := s.URLs[short]; ok {
		return el, nil
	}
	return nil, url.ErrURLNotFound
}
