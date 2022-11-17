package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	sync.RWMutex
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	l.Lock()
	defer l.Unlock()

	i, ok := l.items[key]
	if ok {
		c, _ := i.Value.(cacheItem)
		c.value = value
		i.Value = c
		l.queue.MoveToFront(i)
	} else {
		l.items[key] = l.queue.PushFront(cacheItem{key: key, value: value})
		if l.queue.Len() > l.capacity {
			b := l.queue.Back()
			l.queue.Remove(b)
			c, _ := b.Value.(cacheItem)
			delete(l.items, c.key)
		}
	}

	return ok
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	l.RLock()
	defer l.RUnlock()
	i, ok := l.items[key]
	if ok {
		l.queue.MoveToFront(i)
		c, _ := i.Value.(cacheItem)
		return c.value, true
	}
	return nil, false
}

func (l *lruCache) Clear() {
	l.queue.Clear()
}

type cacheItem struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
