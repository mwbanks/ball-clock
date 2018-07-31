package main

import (
	"fmt"
	"sort"
)

type BallQueue struct {
	Size  int
	Array []int
	Name  string
}

func (q *BallQueue) Len() int {
	return q.Size
}

func (q *BallQueue) ValString() string {
	retStr := "["
	for i, val := range q.Array {
		if val == 0 {
			break
		} else if i != 0 {
			retStr += ","
		}

		retStr += fmt.Sprintf("%d", val)

	}
	return retStr + "]"
}

func (q BallQueue) String() string {
	return fmt.Sprintf("%s", q.ValString())
}

func (q *BallQueue) Less(i, j int) bool {
	return q.Array[i] < q.Array[j]
}

func (q *BallQueue) Swap(i, j int) {
	q.Array[i], q.Array[j] = q.Array[j], q.Array[i]
}

func (q *BallQueue) InOrder() bool {
	return sort.IsSorted(q) && q.Size == len(q.Array)
}

func (q *BallQueue) Append(i int) {
	// fmt.Printf("%s %d %v\n", q.Name, q.Size, q.Array)
	q.Array[q.Size] = i
	q.Size++
}

func (q *BallQueue) Empty(destQueue *BallQueue) int {
	for i := len(q.Array) - 2; i >= 0; i-- {
		destQueue.Append(q.Array[i])
		q.Array[i] = 0
	}
	retVal := q.Array[len(q.Array)-1]
	q.Array[len(q.Array)-1] = 0
	q.Size = 0
	return retVal
}

func (q *BallQueue) Pop(i int) int {
	retVal := q.Array[i]
	for i := 1; i < q.Size; i++ {
		q.Array[i-1] = q.Array[i]
	}
	q.Size--
	q.Array[q.Size] = 0
	return retVal
}

func (q *BallQueue) IsFull() bool {

	return q.Size == len(q.Array)
}
