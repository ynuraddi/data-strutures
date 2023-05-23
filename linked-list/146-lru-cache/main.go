package main

import "container/list"

func main() {
}

type kv struct {
	k int
	v int
}

type LRUCache struct {
	cache    map[int]*list.Element
	linklist *list.List
	capacity int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cache:    make(map[int]*list.Element, capacity),
		linklist: list.New(),
		capacity: capacity,
	}
}

// 0 is key and 1 is value
func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; !ok {
		return -1
	}
	elem := this.cache[key]
	this.linklist.MoveToFront(elem)
	return elem.Value.([]int)[1]
}

func (this *LRUCache) Put(key int, value int) {
	// if capacity reached
	if elem, ok := this.cache[key]; ok {
		this.linklist.Remove(elem)
		newelem := this.linklist.PushFront([]int{key, value})
		this.cache[key] = newelem
		return
	}
	if len(this.cache) == this.capacity {
		elem := this.linklist.Back()
		v := this.linklist.Remove(elem)
		delete(this.cache, v.([]int)[0])
	}
	newelem := this.linklist.PushFront([]int{key, value})
	this.cache[key] = newelem
}
