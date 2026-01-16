package pokeapi

import (
	"net/http"
	"pokedexcli/internal/pokecache"
	"time"
)

type Client struct {
	httpClient http.Client
	pokeCache  *pokecache.Cache
}

func NewClient(timeout time.Duration) Client {
	var newCache *pokecache.Cache
	newCache = pokecache.NewCache(5 * time.Second)
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		pokeCache: newCache,
	}
}
