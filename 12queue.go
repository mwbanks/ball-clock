package main

import "fmt"

type A12Queue struct {
	Array []ballnum
	Name  string
	End   sizenum
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
		Name:  name,
	}
}

func (q *A12Queue) Append(i ballnum) {
	q.Array[q.End] = i
	q.End++
}

func (q *A12Queue) Empty(destQueue *ABallQueue) (retVal ballnum) {
	retVal = uint8(q.Array[len(q.Array)-1])
	// destQueue.Array[destQueue.End] = q.Array[10]
	// destQueue.Array[destQueue.End+1] = q.Array[9]
	// destQueue.Array[destQueue.End+2] = q.Array[8]
	// destQueue.Array[destQueue.End+3] = q.Array[7]
	// destQueue.Array[destQueue.End+4] = q.Array[6]
	// destQueue.Array[destQueue.End+5] = q.Array[5]
	// destQueue.Array[destQueue.End+6] = q.Array[4]
	// destQueue.Array[destQueue.End+7] = q.Array[3]
	// destQueue.Array[destQueue.End+8] = q.Array[2]
	// destQueue.Array[destQueue.End+9] = q.Array[1]
	// destQueue.Array[destQueue.End+10] = q.Array[0]
	// destQueue.End += 11
	for i := int(10); i != -1; i, destQueue.End = i-1, destQueue.End+1 {
		destQueue.Array[destQueue.End] = q.Array[i]
	}
	destQueue.Size += 11
	q.End = 0
	return
}

func (q *A12Queue) IsFull() bool {
	return int(q.End) == len(q.Array)
}
