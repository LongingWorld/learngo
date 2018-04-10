package queue

//An FIFO queue
type Queue []int

//Pops element from head
func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

//Pushes the element into the queue
func (q *Queue) Push(i int) {
	*q = append(*q, i)
}

//Returns if the queue is empty or not
func (q Queue) IsEmpty() bool {
	return len(q) == 0
}
