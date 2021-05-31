package datastruct

//双向列表，官方go已经实现：list.List。我这里就是用来练练手,学习一下 src/container
type Node struct {
	data interface{}

	pre  *Node
	next *Node
}

type List struct {
	head *Node
	tail *Node
	size int
}

func New() *List {
	head := &Node{}
	tail := &Node{pre: head}
	head.next = tail
	return &List{head, tail, 0}
}

func (l *List) Len() int {
	return l.size
}

func (l *List) Shift() *Node {
	node := l.head.next
	l.head.next = node.next
	node.next.pre = node.pre
	l.size--
	return node
}

func (l *List) Unshift(node *Node) {
	node.pre, node.next = l.head, l.head.next
	l.head.next, node.next.pre = node, node
	l.size++
}

func (l *List) Push(node *Node) {
	node.pre, node.next = l.tail.pre, l.tail
	node.pre.next, l.tail.pre = node, node
	l.size++
}

func (l *List) Pop() *Node {
	node := l.tail.pre
	l.tail.pre = node.pre
	node.pre.next = node.next
	l.size--
	return node
}

func (l *List) Erase(node *Node) {
	node.pre.next = node.next
	node.next.pre = node.pre
	l.size--
}
