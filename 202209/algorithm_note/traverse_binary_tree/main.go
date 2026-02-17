package main

import "fmt"

var pre, in, post, level []any

type tree struct {
	val   any
	left  *tree
	right *tree
}

type queue struct {
	node []any
}

func con() queue {
	return queue{[]any{}}
}

func (q *queue) empty() bool {
	return len(q.node) == 0
}

func (q *queue) enQueue(tr *tree) {
	q.node = append(q.node, *tr)
}

func (q *queue) deQueue() (tr tree) {
	if !q.empty() {
		tr = q.node[0].(tree)
		q.node = q.node[1:]
		return
	}
	panic("Queue is empty!")
}

func preOrder(tr *tree) []any {
	if tr != nil {
		pre = append(pre, tr.val)
		preOrder(tr.left)
		preOrder(tr.right)
	}
	return pre
}

func inOrder(tr *tree) []any {
	if tr != nil {
		inOrder(tr.left)
		in = append(in, tr.val)
		inOrder(tr.right)
	}
	return in
}

func postOrder(tr *tree) []any {
	if tr != nil {
		postOrder(tr.left)
		postOrder(tr.right)
		post = append(post, tr.val)
	}
	return post
}

func levelOrder(tr *tree) []any {
	//初始化辅助队列
	myQueue := con()
	//创建辅助指针
	var t tree
	//根节点入队
	myQueue.enQueue(tr)
	for !myQueue.empty() {
		//出队，并访问该节点，如果左（右）子树非空，则将其根节点入队
		t = myQueue.deQueue()
		level = append(level, t.val)
		if t.left != nil {
			myQueue.enQueue(t.left)
		}
		if t.right != nil {
			myQueue.enQueue(t.right)
		}
	}
	return level
}

func main() {
	six := tree{6, nil, nil}
	five := tree{5, nil, nil}
	four := tree{4, &six, nil}
	three := tree{3, nil, &five}
	two := tree{2, nil, &four}
	one := tree{1, &two, &three}
	fmt.Println("preOrder =", preOrder(&one))
	fmt.Println("inOrder =", inOrder(&one))
	fmt.Println("postOrder =", postOrder(&one))
	fmt.Println("levelOrder =", levelOrder(&one))

	pre, in, post, level = []any{}, []any{}, []any{}, []any{}
	b := tree{"B", nil, nil}
	m := tree{"M", nil, nil}
	a := tree{"A", nil, nil}
	h := tree{"H", nil, nil}
	d := tree{"D", &b, nil}
	g := tree{"G", &m, nil}
	c := tree{"C", &a, &d}
	e := tree{"E", &h, &g}
	f := tree{"F", &c, &e}
	fmt.Println("----------------------------------------")
	fmt.Println("preOrder =", preOrder(&f))
	fmt.Println("inOrder =", inOrder(&f))
	fmt.Println("postOrder =", postOrder(&f))
	fmt.Println("levelOrder =", levelOrder(&f))
}
