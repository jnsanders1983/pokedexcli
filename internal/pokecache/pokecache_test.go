package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://pokeapi.co",
			val: []byte("testdata_1"),
		},
		{
			key: "https://pokeapi.co",
			val: []byte("testdata_2"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %d", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			actual, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key %s", c.key)
				return
			}
			if string(actual) != string(c.val) {
				t.Errorf("expected value %s, got %s", string(c.val), string(actual))
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	// A slightly more generous base time and buffer to prevent flakiness
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + (5 * time.Millisecond) // 10ms total ensures the ticker fires at least once

	cache := NewCache(baseTime)
	cache.Add("https://pokeapi.co", []byte("testdata"))

	_, ok := cache.Get("https://pokeapi.co")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://pokeapi.co")
	if ok {
		t.Errorf("expected to not find key (it should have been reaped due to scheduling buffer)")
		return
	}
}
