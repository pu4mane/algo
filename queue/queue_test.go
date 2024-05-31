package queue

import "testing"

func TestQueue(t *testing.T) {
	queue := Queue{}
	t.Run("Test Enqueue", func(t *testing.T) {
		queue.Enqueue(10)
		queue.Enqueue(20)
		queue.Enqueue(30)

		if queue.Peek() != 10 {
			t.Errorf("Peek() returned incorrect value")
		}
	})
	t.Run("Test Dequeue", func(t *testing.T) {
		val := queue.Dequeue()

		if val != 10 {
			t.Errorf("Dequeue() returned incorrect value")
		}
	})
}
