package main

import (
	"log"

	lrucache "github.com/handytantyo/memory-cache/lru"
)

func main() {
	cache := lrucache.NewLRU(3)
	cache.Set("test1", 2)
	cache.Set("test2", "handy")
	cache.Set("test3", "handy===========")
	cache.Set("test4", "berhasil")

	log.Println(cache.Get("test1"))
	log.Println(cache.Get("test3"))
	log.Println(cache.Count())
}
