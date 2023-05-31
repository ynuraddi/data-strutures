package main

import (
	"container/heap"
	"fmt"
)

func main() {
	t := Constructor()
	t.PostTweet(1, 5)
	fmt.Println(t.GetNewsFeed(1))
	// t.GetNewsFeed(1)
	t.Follow(1, 2)
	t.PostTweet(2, 6)
	fmt.Println(t.GetNewsFeed(1))
	// t.GetNewsFeed(1)
	t.Unfollow(1, 2)
	fmt.Println(t.GetNewsFeed(1))
	// t.GetNewsFeed(1)
}

type Tweet struct {
	createdAt  int
	tweetId    int
	followeeId int
}

type TweetHeap []*Tweet

func (h TweetHeap) Len() int            { return len(h) }
func (h TweetHeap) Less(i, j int) bool  { return h[i].createdAt > h[j].createdAt }
func (h TweetHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *TweetHeap) Push(x interface{}) { *h = append(*h, x.(*Tweet)) }
func (h *TweetHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type Twitter struct {
	timer     int
	tweetMap  map[int][]*Tweet
	followMap map[int]map[int]bool
}

func Constructor() Twitter {
	return Twitter{tweetMap: make(map[int][]*Tweet), followMap: make(map[int]map[int]bool)}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	if _, ok := this.tweetMap[userId]; !ok {
		this.tweetMap[userId] = make([]*Tweet, 0)
	}
	this.tweetMap[userId] = append(this.tweetMap[userId], &Tweet{
		createdAt:  this.timer,
		tweetId:    tweetId,
		followeeId: userId,
	})
	this.timer += 1
}

func (this *Twitter) GetNewsFeed(userId int) (res []int) {
	h := &TweetHeap{}
	heap.Init(h)

	if _, ok := this.followMap[userId]; !ok {
		this.followMap[userId] = make(map[int]bool)
	}
	this.followMap[userId][userId] = true
	for followeeId := range this.followMap[userId] {
		if this.followMap[userId][followeeId] {
			length := len(this.tweetMap[followeeId]) - 1
			for i := length; i >= 0; i -= 1 {
				heap.Push(h, this.tweetMap[followeeId][i])
			}
		}
	}

	for len(res) < 10 && len(*h) > 0 {
		tweet := heap.Pop(h).(*Tweet)
		res = append(res, tweet.tweetId)
	}

	return res
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	if _, ok := this.followMap[followerId]; !ok {
		this.followMap[followerId] = make(map[int]bool)
	}
	this.followMap[followerId][followeeId] = true
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if _, ok := this.followMap[followerId]; !ok {
		this.followMap[followerId] = make(map[int]bool)
	}
	delete(this.followMap[followerId], followeeId)
}
