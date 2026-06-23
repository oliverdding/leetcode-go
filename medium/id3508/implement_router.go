package id3508

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=3508 lang=golang
 *
 * [3508] Implement Router
 */

// @lc code=start
type Router struct {
	limit   int
	data    []packet
	destMap map[int][]int
}

type packet struct {
	source      int
	destination int
	timestamp   int
}

func Constructor(memoryLimit int) Router {
	return Router{
		memoryLimit,
		make([]packet, 0, memoryLimit),
		make(map[int][]int),
	}
}

func (this *Router) AddPacket(source int, destination int, timestamp int) bool {
	for i := len(this.data) - 1; i >= 0 && this.data[i].timestamp == timestamp; i-- {
		if this.data[i].source == source && this.data[i].destination == destination {
			return false
		}
	}
	if len(this.data) == this.limit {
		element := this.data[0]
		this.data = this.data[1:]
		this.destMap[element.destination] = this.destMap[element.destination][1:]
	}
	this.data = append(this.data, packet{source, destination, timestamp})
	this.destMap[destination] = append(this.destMap[destination], timestamp)
	return true
}

func (this *Router) ForwardPacket() []int {
	if len(this.data) == 0 {
		return nil
	}
	element := this.data[0]
	this.data = this.data[1:]
	this.destMap[element.destination] = this.destMap[element.destination][1:]
	return []int{element.source, element.destination, element.timestamp}
}

func (this *Router) GetCount(destination int, startTime int, endTime int) int {
	arr := this.destMap[destination]
	return sort.SearchInts(arr, endTime+1) - sort.SearchInts(arr, startTime)
}

/**
 * Your Router object will be instantiated and called as such:
 * obj := Constructor(memoryLimit);
 * param_1 := obj.AddPacket(source,destination,timestamp);
 * param_2 := obj.ForwardPacket();
 * param_3 := obj.GetCount(destination,startTime,endTime);
 */
// @lc code=end
