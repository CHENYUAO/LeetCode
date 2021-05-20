/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU 缓存机制
 */

// @lc code=start
import "container/list"

type LRUCache struct {
	Capacity   int
	DoubleList *list.List
	Hashtable  map[int]*list.Element
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Capacity:   capacity,
		DoubleList: list.New(),
		Hashtable:  make(map[int]*list.Element),
	}
}

func (this *LRUCache) Get(key int) int {
	v, ok := this.Hashtable[key]
	if ok {
		this.DoubleList.MoveToFront(v)
		return key
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	//查找key，如果已经存在，则把该节点放到链表头
	v, ok := this.Hashtable[key]
	if ok {
		this.DoubleList.MoveToFront(v)
		return
	}
	//如果不存在，判断当前map大小
	currLen := len(this.Hashtable)
	//map已满，删除链表尾节点，将新节点插入表头
	if currLen >= this.Capacity {
		this.DoubleList.Remove(this.DoubleList.Back())
		delete(this.Hashtable, key)
	}
	//map未满，直接插入表头
	ins := this.DoubleList.PushFront(value)
	this.Hashtable[key] = ins
	return
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end

