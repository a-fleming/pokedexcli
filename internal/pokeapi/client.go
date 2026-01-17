package pokeapi

import (
	"net/http"
	"pokedexcli/internal/pokecache"
	"time"
)

type Client struct {
	httpClient http.Client
	pokeCache  pokecache.Cache
}

func NewClient(timeout time.Duration, cacheInterval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		pokeCache: pokecache.NewCache(cacheInterval),
	}
}
