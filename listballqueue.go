package main

import (
	"fmt"
)

type ballnum = uint8
type sizenum = uint8

const sizemax int16 = 0xFF

type ABallQueue struct {
	Array                []ballnum
	Name                 string
	Start, Size, MaxSize sizenum
}

func NewABallQueue(arr []ballnum, name string, maxSize sizenum) *ABallQueue {
	return &ABallQueue{
		Array:   arr,
		Start:   0,
		Size:    0,
		MaxSize: maxSize,
		Name:    name,
	}
}

func (q *ABallQueue) Init(balls uint8) {
	for i := sizenum(0); i < balls; i++ {
		q.Array[q.Size] = ballnum(i) + 1
		// q.End = (q.End + 1) & 0x7F
		q.Size++
	}
}

func (q *ABallQueue) ValString() string {
	retStr := "["
	for i := 0; i < len(q.Array); i++ {
		retStr += fmt.Sprintf("%d, ", int(q.Array[i]))

	}
	return retStr + "]"
}

func (q ABallQueue) String() string {
	return fmt.Sprintf(`{
	"array": 
		%#v, 
	"Start": %d, 
	"Size": %d
	}`, q.ValString(), q.Start, q.Size)
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
	q.Array[q.Start+q.Size] = i
	q.Size++
}

func (q *ABallQueue) FastReverseReturn() (retVal ballnum) {
	retVal = q.Array[q.Start+4]
	end := q.Start + q.Size

	q.Array[end],
		q.Array[end+1],
		q.Array[end+2],
		q.Array[end+3] =
		q.Array[q.Start+3],
		q.Array[q.Start+2],
		q.Array[q.Start+1],
		q.Array[q.Start]
	q.Start += 5
	q.Size -= 1
	return
}

func (q *ABallQueue) IsFull() bool {
	return q.Size == q.MaxSize
}
