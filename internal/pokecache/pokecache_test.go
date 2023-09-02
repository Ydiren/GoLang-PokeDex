package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestCache_AddGet(t *testing.T) {
	const interval = 1 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "test",
			val: []byte("test"),
		},
		{
			key: "test2",
			val: []byte("test2"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test Case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("Expected to get value for key '%s'", c.key)
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("Expected to get value '%s' for key '%s' but got '%s'", string(c.val), c.key, string(val))
				return
			}
		})
	}
}

func Test_ReapLoop(t *testing.T) {
	const interval = 10 * time.Millisecond
	const waitTime = interval + 10*time.Millisecond
	c := NewCache(interval)
	c.Add("test", []byte("test"))

	_, ok := c.Get("test")
	if !ok {
		t.Errorf("Expected to get value for key 'test'")
	}

	time.Sleep(waitTime)

	_, ok = c.Get("test")
	if ok {
		t.Errorf("Expected to not get value for key 'test'")
	}
}
