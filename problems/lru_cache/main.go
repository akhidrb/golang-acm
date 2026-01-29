package main

type LRUCache struct {
	capacity    int
	cache       map[int]*Node
	left, right *Node
}

type Node struct {
	key, value int
	prev, next *Node
}

func Constructor(capacity int) LRUCache {
	left := &Node{key: 0}
	right := &Node{key: 0}
	left.next, right.prev = right, left
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
		left:     left,
		right:    right,
	}
}

func (this *LRUCache) remove(node *Node) {
	prev, next := node.prev, node.next
	prev.next, next.prev = next, prev
}

func (this *LRUCache) insert(node *Node) {
	prev := this.right.prev
	prev.next, this.right.prev = node, node
	node.prev, node.next = prev, this.right
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.remove(node)
		this.insert(node)
		return node.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		node.value = value
		this.remove(node)
		this.insert(node)
		return
	}
	temp := &Node{key: key, value: value}
	this.insert(temp)
	this.cache[key] = temp
	if len(this.cache) > this.capacity {
		lru := this.left.next
		leftKey := lru.key
		this.remove(lru)
		delete(this.cache, leftKey)
	}
}
