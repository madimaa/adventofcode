package priorityqueue

import (
	"container/heap"
	"fmt"
)

type Item[T any] struct {
	Value    T
	Priority int
}

func (item Item[T]) ToString() string {
	return fmt.Sprintf("%v", item.Value)
}

type PriorityQueue[T any] []Item[T]

func (pq PriorityQueue[_]) Len() int { return len(pq) }

func (pq PriorityQueue[_]) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue[T]) Push(x any) {
	*pq = append(*pq, x.(Item[T]))
}

func (pq *PriorityQueue[_]) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	//old[n-1] = nil // avoid memory leak
	*pq = old[:n-1]
	return item
}

func (pq *PriorityQueue[T]) GPush(value T, priority int) {
	heap.Push(pq, Item[T]{value, priority})
}

func (pq *PriorityQueue[T]) GPop() (T, int) {
	item := heap.Pop(pq).(Item[T])
	return item.Value, item.Priority
}
