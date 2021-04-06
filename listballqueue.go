package main

import (
	"fmt"
)

type ballnum = uint8

type ABallQueue struct {
	Array                     []ballnum
	Name                      string
	Start, End, Size, MaxSize int
}

func NewABallQueue(arr []ballnum, name string, maxSize int) *ABallQueue {
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
		q.Array[q.End] = ballnum(i) + 1
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
		retStr += fmt.Sprintf("%d", int(val))

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
	for i, j := 0, q.Start; i < q.Size; i++ {
		val = q.Array[j]
		j = (j + 1) & 0x7F
		if val != ballnum(i+1) {
			return false
		}
	}
	return true
}

func (q *ABallQueue) Append(i ballnum) {
	q.Array[q.End] = i
	q.End = (q.End + 1) & 0x7F
	q.Size++
}

func (q *ABallQueue) Empty(destQueue *ABallQueue) ballnum {
	retVal := q.Array[q.MaxSize-1]
	for i := q.MaxSize - 2; i != -1; i-- {
		destQueue.Array[destQueue.End] = q.Array[i]
		// q.Array[i] = 0
		destQueue.End = (destQueue.End + 1) & 0x7F
	}
	destQueue.Size += (q.MaxSize - 1)
	q.End = 0
	q.Size = 0
	return retVal
}

func (q *ABallQueue) ReverseReturn() ballnum {
	// if q.Start+5 < len(q.Array) {
	retVal := q.Array[(q.Start+4)%128]

	q.Array[q.End],
		q.Array[(q.End+1)%128],
		q.Array[(q.End+2)%128],
		q.Array[(q.End+3)%128] =
		q.Array[(q.Start+3)%128],
		q.Array[(q.Start+2)%128],
		q.Array[(q.Start+1)%128],
		q.Array[q.Start]
	q.Start = (q.Start + 5) % 128
	q.End = (q.End + 4) % 128
	// if q.End >= len(q.Array) {
	// 	q.End = q.End - len(q.Array)
	// }
	q.Size -= 1
	return retVal
	// }
	// return 0
}

func (q *ABallQueue) FastReverseReturn() ballnum {
	retVal := q.Array[(q.Start+4)&0x7F]

	q.Array[q.End],
		q.Array[(q.End+1)&0x7F],
		q.Array[(q.End+2)&0x7F],
		q.Array[(q.End+3)&0x7F] =
		q.Array[(q.Start+3)&0x7F],
		q.Array[(q.Start+2)&0x7F],
		q.Array[(q.Start+1)&0x7F],
		q.Array[q.Start]
	q.Start = (q.Start + 5) & 0x7F
	q.End = (q.End + 4) & 0x7F
	q.Size -= 1
	return retVal
}

func (q *ABallQueue) Pop5(arr []ballnum) {
	if q.Start+len(arr) < len(q.Array) {
		copy(arr[:], q.Array[q.Start:q.Start+5])
		q.Start += 5
	} else {
		for i := 0; i < 5; i++ {
			arr[i] = q.Array[q.Start]
			q.Start++
			if q.Start == len(q.Array) {
				q.Start = 0
			}
		}
	}
	q.Size -= 5
}

func (q *ABallQueue) AppendInReverse(arr []ballnum) {
	arr[0], arr[1], arr[2], arr[3] = arr[3], arr[2], arr[1], arr[0]
	if q.End+4 < len(q.Array) {
		copy(q.Array, arr[:])
		q.End += 4
	} else {
		for i := 0; i < 4; i++ {
			q.Array[q.End] = arr[i]
			q.End++
			if q.End == len(q.Array) {
				q.End = 0
			}
		}
	}
	q.Size += 4
	// fmt.Printf("%#v\n", q.Array)
}

func (q *ABallQueue) PopFront() ballnum {
	retVal := q.Array[q.Start]
	q.Start += 1
	if q.Start == len(q.Array) {
		q.Start = 0
	}
	q.Size--
	return retVal
}

func (q *ABallQueue) IsFull() bool {
	return q.Size == q.MaxSize
}
