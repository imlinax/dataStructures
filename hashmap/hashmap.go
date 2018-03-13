package hashmap

import (
	"../dlist"
)

/*
Q: hashmap是啥？
A: map[key] = value
Q: 一个hashmap由哪些东西组成 ?
A:
	1. hash算法，或者叫hash函数
	2. 一个数组，叫桶数组
	3. 每个桶下挂一个链表
	4. 数组的长度,数组可能会扩容，长度也随之改变
	5. 总的元素个数， 也就是所有链表长度的和
*/

type HashMap struct {
	elementSize uint64
	bucketSize  uint32
	buckets     []dlist.DList
}

func Hash(key string) int {
	return 0
}

func MinPower(size int) (pow uint) {
	for size > 0 {
		size = size >> 1
		pow++
	}
	return
}
func New(size int) *HashMap {

	pow := MinPower(size)
	hmap := &HashMap{
		elementSize: 0,
		bucketSize:  0,
		buckets:     make([]dlist.DList, 2<<pow)}
	return hmap
}

func Insert(key, value string) {

}
