package main

// ! Linked list example
type Node struct {
	Val  int
	Next *Node
}

type MyLinkedList struct {
	Head *Node
	Size int
}

func Constructor() MyLinkedList {
	return MyLinkedList{&Node{}, 0}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.Size {
		return -1
	}
	curr := this.Head
	for i := 0; i <= index; i++ {
		curr = curr.Next
	}
	return curr.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	node := &Node{val, this.Head.Next}
	this.Head.Next = node
	this.Size++
}

func (this *MyLinkedList) AddAtTail(val int) {
	node := &Node{val, nil}
	curr := this.Head
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = node
	this.Size++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > this.Size {
		return
	}
	node := &Node{val, nil}
	curr := this.Head
	for i := 0; i < index; i++ {
		curr = curr.Next
	}
	node.Next = curr.Next
	curr.Next = node
	this.Size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.Size {
		return
	}
	curr := this.Head
	for i := 0; i < index; i++ {
		curr = curr.Next
	}
	curr.Next = curr.Next.Next
	this.Size--
}

// ! HashSet bucket implementation
type MyHashSet struct {
	buckets []*Bucket
}

type Bucket struct {
	list *List
}

func NewBucket() *Bucket {
	return &Bucket{
		list: nil,
	}
}

type List struct {
	Val  int
	Next *List
}

func Constructor() MyHashSet {
	buckets := make([]*Bucket, 16)
	for i := range buckets {
		buckets[i] = NewBucket()
	}
	return MyHashSet{
		buckets: buckets,
	}
}

func (h *MyHashSet) Len() int {
	return len(h.buckets)
}

func (h *MyHashSet) Add(key int) {
	hash := key % h.Len()
	cur := h.buckets[hash].list

	if cur == nil {
		h.buckets[hash].list = &List{Val: key}
		return
	}

	for cur.Next != nil {
		if cur.Val == key {
			return
		}
		cur = cur.Next
	}

	if cur.Val != key {
		cur.Next = &List{Val: key}
	}
}

func (h *MyHashSet) Remove(key int) {
	hash := key % h.Len()
	cur := h.buckets[hash].list

	if cur != nil && cur.Val == key {
		h.buckets[hash].list = h.buckets[hash].list.Next
		return
	}

	for cur != nil && cur.Next != nil {
		if cur.Next.Val == key {
			cur.Next = cur.Next.Next
			return
		}

		cur = cur.Next
	}
}

func (h *MyHashSet) Contains(key int) bool {
	hash := key % h.Len()
	cur := h.buckets[hash].list

	for cur != nil {
		if cur.Val == key {
			return true
		}

		cur = cur.Next
	}

	return false
}

// ! HashMap low level
type node struct {
	key  int
	val  int
	next *node
}

type MyHashMap struct {
	size int
	arr  []*node
}

func Constructor() MyHashMap {
	return MyHashMap{
		size: 1000,
		arr:  make([]*node, 1000),
	}
}

func (this *MyHashMap) Put(key int, value int) {
	//Use mod operation to find the position in the array
	position := key % this.size

	//Create a new node with key and value
	newNode := &node{
		key:  key,
		val:  value,
		next: nil,
	}

	//If current position is empty, assign it to newNode then return
	if this.arr[position] == nil {
		this.arr[position] = newNode
		return
	}

	//Check whether the new key already exists or not
	//Iterate through the chaining to, if the key matches, replace the old value by newValue, then return
	head := this.arr[position]

	for head != nil {
		if head.key == key {
			head.val = value
			return
		}

		head = head.next
	}

	//When we reach here, it means that the new key does not exist in the map
	//We need to insert the new key to the head of the linked list
	//Why not insert tail? Because insert head or tail does not make any sense
	//Moreover, insert head is O(1) while insert tail is O(n)
	newNode.next = this.arr[position]
	this.arr[position] = newNode
}

func (this *MyHashMap) Get(key int) int {
	position := key % this.size

	head := this.arr[position]

	//Iterate through the whole linked list at the position
	//If the key matches, return the value
	//Otherwise, return -1
	for head != nil {
		if head.key == key {
			return head.val
		}

		head = head.next
	}

	return -1
}

func (this *MyHashMap) Remove(key int) {
	//Find the position to remove
	position := key % this.size
	head := this.arr[position]

	//If the position is empty, it means that the key does not exist in the map, just return
	if head == nil {
		return
	}

	//Special case, when the key is the head of the linked list, just update the arr[position] to the next value of head
	if head.key == key {
		head = head.next
		this.arr[position] = head
		return
	}

	//Otherwise, we iterate to find the key in the linked list, then remove it from the linked list (Actually, we use prev.next to skip the node)
	var prev *node

	for head != nil {
		if head.key == key {
			prev.next = head.next
			return
		}

		prev = head
		head = head.next
	}
}
