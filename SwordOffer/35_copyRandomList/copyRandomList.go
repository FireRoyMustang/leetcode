package main

// 请实现 copyRandomList 函数，复制一个复杂链表。在复杂链表中，每个节点除了有一个 next 指针指向下一个节点，
// 还有一个 random 指针指向链表中的任意节点或者 null。

func main() {

}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}
	mp := make(map[*Node]*Node)
	node := head
	for node != nil {
		mp[node] = &Node{node.Val, nil, nil}
		node = node.Next
	}
	node = head
	for node != nil {
		mp[node].Next = mp[node.Next]
		mp[node].Random = mp[node.Random]
		node = node.Next
	}
	return mp[head]
}
