package main

import "fmt"

type A12Queue struct {
	Array []ballnum
	// Name  string
	End int
}

func (q *A12Queue) ValString() string {
	retStr := "["
	for i := 0; i < len(q.Array); i++ {
		retStr += fmt.Sprintf("%d, ", int(q.Array[i]))

	}
	return retStr + "]"
}

func (q A12Queue) String() string {
	return fmt.Sprintf(`{
	"array": 
		%#v, 
	"End": %d, 
	"Size": %d
	}`, q.ValString(), q.End, q.End)
}

func New12Queue(arr []ballnum, name string) *A12Queue {
	return &A12Queue{
		Array: arr,
		End:   0,
		// Name:  name,
	}
}

func (q *A12Queue) Append(i ballnum) {
	q.Array[q.End] = i
	q.End++
}

func (q *A12Queue) Empty2(destQueue *ABallQueue) (retVal ballnum) {
	retVal = q.Array[0]
	end := destQueue.Start + destQueue.Size
	for i := sizenum(1); i < 12; i++ {
		destQueue.Array[end+i] = q.Array[i]
	}
	destQueue.Size += 11
	q.End = 11
	return
}

func (q *A12Queue) Empty(destQueue *ABallQueue) (retVal ballnum) {
	retVal = q.Array[len(q.Array)-1]
	end := destQueue.Start + destQueue.Size
	for i := int(10); i != -1; i, end = i-1, end+1 {
		destQueue.Array[end] = q.Array[i]
	}
	destQueue.Size += 11
	q.End = 0
	return
}

func (q *A12Queue) IsFull() bool {
	return q.End == len(q.Array)
}
