package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	cache := NewCache(3, &Connection{})

	cache.items["1"] = Rankable{key: "1", rank: 1}
	cache.items["2"] = Rankable{key: "2", rank: 2}

	a := cache.Get("1")
	assert.Equal(t, a.key, "1")
	assert.Equal(t, a.rank, 1)

	a = cache.Get("2")
	assert.Equal(t, a.key, "2")
	assert.Equal(t, a.rank, 2)
}
