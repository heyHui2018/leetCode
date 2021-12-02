package main

import (
	"container/heap"
)

/*
要求：


解题思路：
取中值，保存一大顶堆一小顶堆来实现


关键点：


*/

type minHeap []int

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *minHeap) Pop() interface{} {
	min := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return min
}

type maxHeap []int

func (h maxHeap) Len() int {
	return len(h)
}

func (h maxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *maxHeap) Pop() interface{} {
	max := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return max
}

type MedianFinder struct {
	minHeap minHeap
	maxHeap maxHeap
}

func Constructor() MedianFinder {
	return MedianFinder{
		minHeap: minHeap{},
		maxHeap: maxHeap{},
	}
}

func (this *MedianFinder) AddNum(num int) {
	// 大顶堆堆顶为此堆最小值，小顶堆堆顶为此堆最大值
	// 大顶堆长度为0或小于大顶堆的最小值，则将num放入小顶堆，否则放入大顶堆
	if this.maxHeap.Len() == 0 || num <= this.maxHeap[0] {
		heap.Push(&this.maxHeap, num)
	} else {
		heap.Push(&this.minHeap, num)
	}
	// 保持大顶堆长度最多比小顶堆长度大1，否则从大顶堆移数至小顶堆
	if this.maxHeap.Len() > this.minHeap.Len()+1 {
		maxLeft := heap.Pop(&this.maxHeap)
		heap.Push(&this.minHeap, maxLeft)
	}
	// 保持小顶堆长度小于大顶堆长度，否则移数至大顶堆
	if this.minHeap.Len() > this.maxHeap.Len() {
		minRight := heap.Pop(&this.minHeap)
		heap.Push(&this.maxHeap, minRight)
	}
}

func (this *MedianFinder) FindMedian() float64 { // faster 77.94% less 48.53%
	// 大小顶堆长度一致，返回两堆顶堆顶数的平均值，否则返回大顶堆的堆顶值
	if this.minHeap.Len() == this.maxHeap.Len() {
		return float64(this.minHeap[0]+this.maxHeap[0]) / 2
	}
	return float64(this.maxHeap[0])
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
