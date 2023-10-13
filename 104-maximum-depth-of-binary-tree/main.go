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

type Stack struct {
	node  *TreeNode
	depth int
}

/*
*
Iterative DFS solution with slice
*/
func maxDepth(root *TreeNode) int {
	s := []Stack{
		{root, 1},
	}
	levels := 0
	for len(s) > 0 {
		n := s[0]
		if levels < n.depth {
			levels = n.depth
		}
		s = s[1:]
		if n.node.Left != nil {
			s = append([]Stack{{node: n.node.Left, depth: n.depth + 1}}, s...)
		}
		if n.node.Right != nil {
			s = append([]Stack{{node: n.node.Right, depth: n.depth + 1}}, s...)
		}
	}
	return levels
}

/*
*
Iterative BFS solution with slice
*/
func maxDepthWithSlice(root *TreeNode) int {
	q := []*TreeNode{root}
	levels := 0
	for len(q) > 0 {
		levels++
		elc := len(q)
		for i := 0; i < elc; i++ {
			node := q[0]
			q = q[1:]

			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}
	return levels
}

/*
*
Iterative BFS solution with custom struct
*/
func maxDepthCustomStruct(root *TreeNode) int {
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
	m := r
	if l > r {
		m = l
	}
	return 1 + m
}
