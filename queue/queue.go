package queue

type Queue struct {
	collection []int
}

func (q *Queue) Enqueue(val int) {
	q.collection = append(q.collection, val)
}

func (q *Queue) Dequeue() int {
	if q.isEmpty() {
		return 0
	}

	val := q.collection[0]
	q.collection = q.collection[1:]

	return val
}

func (q *Queue) Peek() int {
	if q.isEmpty() {
		return 0
	}

	return q.collection[0]
}

func (q *Queue) isEmpty() bool {
	return len(q.collection) == 0
}
