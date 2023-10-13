package main

func main() {
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type QueueNode struct {
	val  *TreeNode
	next *QueueNode
}

type Queue struct {
	head *QueueNode
	tail *QueueNode
}

func (q *Queue) enqueue(node *TreeNode) {
	n := &QueueNode{val: node}
	if q.head == nil {
		q.head = n
		q.tail = n
		return
	}
	q.tail.next, q.tail = n, n
}

func (q *Queue) dequeue() *QueueNode {
	r := q.head
	if r == nil {
		return nil
	}
	q.head = q.head.next
	return r
}

/*
*
Iterative BFS solution
*/
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	q := Queue{}
	q.enqueue(root)

	level := 0
	nodesAtCurrentLevel := 1
	nodesAtNextLevel := 0
	for q.head != nil {
		qn := q.dequeue()
		nodesAtCurrentLevel--
		if qn.val.Left != nil {
			q.enqueue(qn.val.Left)
			nodesAtNextLevel++
		}
		if qn.val.Right != nil {
			q.enqueue(qn.val.Right)
			nodesAtNextLevel++
		}
		if nodesAtCurrentLevel == 0 {
			level++
			nodesAtCurrentLevel = nodesAtNextLevel
			nodesAtNextLevel = 0
		}
	}
	return level
}

/*
Recursive DFS solution
*/
func maxDepthRecursive(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := maxDepthRecursive(root.Left)
	r := maxDepthRecursive(root.Right)
	max := r
	if l > r {
		max = l
	}
	return 1 + max
}
