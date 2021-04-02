package main

import (
	"fmt"
	"sort"
)

type ABallQueue struct {
	Array                     []int
	Name                      string
	Start, End, Size, MaxSize int
}

func NewABallQueue(arr []int, name string, maxSize int) *ABallQueue {
	return &ABallQueue{
		Array:   arr,
		Start:   0,
		End:     0,
		Size:    0,
		MaxSize: maxSize,
		Name:    name,
	}
}

func (q *ABallQueue) Init(balls int) {
	for i := 0; i < balls; i++ {
		q.Array[q.End] = i + 1
		q.End++
	}
	if q.End == len(q.Array) {
		q.End = 0
	}
	q.Size = balls
}

func (q *ABallQueue) Len() int {
	return q.Size
}

func (q *ABallQueue) ValString() string {
	retStr := "["
	for i, j := 0, q.Start; i < q.Size; i++ {
		val := q.Array[j]
		j++
		if j == len(q.Array) {
			j = 0
		}
		if i != 0 {
			retStr += ","
		}
		retStr += fmt.Sprintf("%d", val)

	}
	// for i := q.Start; i < q.Size+q.End; i++ {

	// 	retStr += fmt.Sprintf("%d", val)

	// }
	// for i, val := range q.Array {
	// 	if val == 0 {
	// 		break
	// 	} else if i != 0 {
	// 		retStr += ","
	// 	}

	// }
	return retStr + "]"
}

func (q ABallQueue) String() string {
	return fmt.Sprintf(`{
	"array": 
		%s, 
	"Start": %d, 
	"End": %d, 
	"Size": %d
	}`, q.ValString(), q.Start, q.End, q.Size)
}

func (q *ABallQueue) Less(i, j int) bool {
	di := i + q.Start
	if di >= len(q.Array) {
		di = di - len(q.Array)
	}
	dj := j + q.Start
	if dj >= len(q.Array) {
		dj = dj - len(q.Array)
	}
	return q.Array[di] < q.Array[dj]
}

func (q *ABallQueue) Swap(i, j int) {
	q.Array[i], q.Array[j] = q.Array[j], q.Array[i]
}

func (q *ABallQueue) InOrder() bool {
	return q.Size == q.MaxSize && sort.IsSorted(q)
}

func (q *ABallQueue) Append(i int) {
	// fmt.Printf("%s %d %v\n", q.Name, q.Size, q.Array)
	q.Array[q.End] = i
	q.End++
	if q.End == len(q.Array) {
		q.End = 0
	}
	q.Size++
}

func (q *ABallQueue) Empty(destQueue *ABallQueue) int {
	retVal := q.Array[q.MaxSize-1]
	for i := q.MaxSize - 2; i != -1; i-- {
		destQueue.Array[destQueue.End] = q.Array[i]
		// q.Array[i] = 0
		destQueue.End++
		if destQueue.End == len(destQueue.Array) {
			destQueue.End = 0
		}
	}
	destQueue.Size += (q.MaxSize - 1)
	q.End = 0
	q.Size = 0
	return retVal
}

// func (q *ABallQueue) RunFiveMinutes() int {

// }

func (q *ABallQueue) PopFront() int {
	retVal := q.Array[q.Start]
	q.Start++
	if q.Start == len(q.Array) {
		q.Start = 0
	}
	q.Size--
	return retVal
}

func (q *ABallQueue) IsFull() bool {
	return q.Size == q.MaxSize
}
