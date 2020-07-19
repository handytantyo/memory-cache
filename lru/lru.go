package lrucache

import (
	"log"

	memorycache "github.com/handytantyo/memory-cache"
	"github.com/handytantyo/memory-cache/models"
)

// lru means least recently used
type lru struct {
	member    map[string]models.ModelCache
	newest    *node
	last      *node
	maxNumber int
	counter   int
}

type node struct {
	key  string
	next *node
	prev *node
}

func NewLRU(maxNumber int) memorycache.Cache {
	if maxNumber <= 0 {
		log.Fatal("Max Number must more than 0")
	}
	return &lru{member: make(map[string]models.ModelCache), maxNumber: maxNumber}
}

func (l *lru) Get(key string) (result interface{}) {
	if _, ok := l.member[key]; !ok {
		return nil
	}
	return l.member[key].Value
}

func (l *lru) Set(key string, value interface{}) (err error) {
	newNode := &node{key: key, prev: l.newest}
	if l.last == nil {
		l.last = newNode
		l.newest = newNode
	} else {
		l.newest.next = newNode
		l.newest = newNode
	}

	// log.Println("new member:", l.newest, "prev member:", l.newest.prev, "oldest member:", l.last)

	if l.counter >= l.maxNumber {
		delete(l.member, l.last.key)

		newLast := l.last.next
		newLast.prev = nil
		l.last = nil
		l.last = newLast
		l.counter--
	}

	l.member[key] = models.ModelCache{
		Value: value,
	}
	l.counter++

	return
}

func (l *lru) Count() int {
	return l.counter
}
