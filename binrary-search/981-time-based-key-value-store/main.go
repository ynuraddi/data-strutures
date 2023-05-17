package main

import (
	"fmt"
	"sort"
)

func main() {
	obj := Constructor()
	obj.Set("foo", "bar", 1)
	fmt.Println(obj.Get("foo", 1))
	fmt.Println(obj.Get("foo", 3))
	obj.Set("foo", "bar2", 4)
	fmt.Println(obj.Get("foo", 4))
	fmt.Println(obj.Get("foo", 5))
}

type Value struct {
	value     string
	timestamp int
}

type TimeMap struct {
	mp map[string][]Value
}

func Constructor() TimeMap {
	return TimeMap{make(map[string][]Value)}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	if _, ok := this.mp[key]; !ok {
		this.mp[key] = make([]Value, 0)
	}
	this.mp[key] = append(this.mp[key], Value{value, timestamp})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	i := sort.Search(len(this.mp[key]), func(j int) bool {
		return timestamp < this.mp[key][j].timestamp
	})
	if i == 0 {
		return ""
	}
	return this.mp[key][i-1].value
}
