package id138

/*
 * @lc app=leetcode.cn id=138 lang=golang
 *
 * [138] Copy List with Random Pointer
 */

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// @lc code=start

func copyRandomList(head *Node) *Node {
	// insert new node in even point, and leave the original node in old point
	for cur := head; cur != nil; cur = cur.Next.Next {
		node := new(Node)
		node.Val = cur.Val
		node.Next = cur.Next
		cur.Next = node
	}
	// construct random pointer
	for cur := head; cur != nil; cur = cur.Next.Next {
		if cur.Random == nil { // it's importance that the cur.Random could be nil
			continue
		}
		cur.Next.Random = cur.Random.Next
	}
	// decontruct target list from original list
	var (
		// create an empty node as start point to reduce complexicity
		targetHead = new(Node)
		targetTail = targetHead
	)
	for cur := head; cur != nil; cur = cur.Next { // since we cur out the target node, so the step is 1
		targetTail.Next = cur.Next
		cur.Next = cur.Next.Next
		targetTail = targetTail.Next
	}
	targetTail.Next = nil
	return targetHead.Next
}

// @lc code=end
