package main

func main() {
}

type LRUCache struct {
	capacity int
	hash     map[int]int
}

type Node struct {
	val  int
	next *Node
	prev *Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{make(map[int]int, capacity)}
}

func (this *LRUCache) Get(key int) int {
	return this.hash[key]
}

func (this *LRUCache) Put(key int, value int) {
	this.hash[key] = value
}
