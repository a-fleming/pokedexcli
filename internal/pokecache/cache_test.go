package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const internal = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(internal)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Error("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Error("expected to find value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewCache(baseTime)
	url := "https://example.com"
	cache.Add(url, []byte("testdata"))

	_, ok := cache.Get(url)
	if !ok {
		t.Error("expected to find key")
		return
	}
	time.Sleep(waitTime)

	_, ok = cache.Get(url)
	if ok {
		t.Error("expected to not find key")
		return
	}
}
