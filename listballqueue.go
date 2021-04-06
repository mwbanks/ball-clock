package main

import (
	"fmt"
)

type ballnum = uint8
type sizenum = uint8

const sizemax int16 = 0xFF

type ABallQueue struct {
	Array                     []ballnum
	Name                      string
	Start, End, Size, MaxSize sizenum
}

func NewABallQueue(arr []ballnum, name string, maxSize sizenum) *ABallQueue {
	return &ABallQueue{
		Array:   arr,
		Start:   0,
		End:     0,
		Size:    0,
		MaxSize: maxSize,
		Name:    name,
	}
}

func (q *ABallQueue) Init(balls uint8) {
	for i := sizenum(0); i < balls; i++ {
		q.Array[q.End] = ballnum(i) + 1
		// q.End = (q.End + 1) & 0x7F
		q.End++
	}
	q.Size = balls
}

func (q *ABallQueue) ValString() string {
	retStr := "["
	// l := int8(len(q.Array))
	// for i, j := int8(0), q.Start; i < q.Size; i++ {
	for i := 0; i < len(q.Array); i++ {
		// val := q.Array[j]
		// j++
		// if j == l {
		// 	j = 0
		// }
		// if i != 0 {
		// 	retStr += ","
		// }
		// retStr += fmt.Sprintf("%d", int(val))
		retStr += fmt.Sprintf("%d, ", int(q.Array[i]))

	}
	return retStr + "]"
}

func (q ABallQueue) String() string {
	return fmt.Sprintf(`{
	"array": 
		%#v, 
	"Start": %d, 
	"End": %d, 
	"Size": %d
	}`, q.ValString(), q.Start, q.End, q.Size)
}

func (q *ABallQueue) InOrder() bool {
	if q.Size != q.MaxSize {
		return false
	}
	// prevNum := -1
	var val ballnum
	for i, j := sizenum(0), q.Start; i < q.Size; i++ {
		val = q.Array[j]
		j++
		if val != ballnum(i+1) {
			return false
		}
	}
	return true
}

func (q *ABallQueue) Append(i ballnum) {
	q.Array[q.End] = i
	q.End++
	q.Size++
}

func (q *ABallQueue) Empty(destQueue *ABallQueue) ballnum {
	retVal := q.Array[q.MaxSize-1]
	for i := int(q.MaxSize - 2); i != -1; i-- {
		destQueue.Array[destQueue.End] = q.Array[i]
		// q.Array[i] = 0
		// destQueue.End = (destQueue.End + 1) & 0x7F
		destQueue.End++
	}
	destQueue.Size += (q.MaxSize - 1)
	q.End = 0
	q.Size = 0
	return retVal
}

func (q *ABallQueue) FastReverseReturn(val *ballnum) {
	*val = q.Array[q.Start+4]

	q.Array[q.End],
		q.Array[q.End+1],
		q.Array[q.End+2],
		q.Array[q.End+3] =
		q.Array[q.Start+3],
		q.Array[q.Start+2],
		q.Array[q.Start+1],
		q.Array[q.Start]
	q.Start += 5
	q.End += 4
	q.Size -= 1
}

func (q *ABallQueue) IsFull() bool {
	return q.Size == q.MaxSize
}
