package skip

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

//SkipNode - an element in SkipList
type SkipNode struct {
	data   int
	next   *SkipNode
	prev   *SkipNode
	lower  *SkipNode
	higher *SkipNode
}

//SkipList - as the name suggests
type SkipList struct {
	head     *SkipNode
	maxLevel int
}

//Print - as the name suggests
func (sl *SkipList) Print() {
	n := sl.head
	l := sl.maxLevel
	for n != nil {
		fmt.Printf("Printing Level %d:\t", l)
		sn := n
		for sn != nil {
			fmt.Printf("%d,", sn.data)
			sn = sn.next
		}
		fmt.Println()
		n = n.higher
		l--
	}
}

//NewSkipList - creates a new SkipList ad returns it
func NewSkipList(ml int) *SkipList {
	h := createHeadNode()
	t := createTailNode()
	h.next = t
	t.prev = h

	th := h
	tt := t
	for i := 1; i < ml; i++ {
		nh := SkipNode{th.data, nil, nil, th, nil}
		th.higher = &nh
		th = &nh

		nt := SkipNode{tt.data, nil, nil, tt, nil}
		tt.higher = &nt
		tt = &nt

		th.next = tt
		tt.prev = th
	}

	sl := SkipList{h, ml}
	return &sl
}

func createHeadNode() *SkipNode {
	return &SkipNode{math.MinInt32, nil, nil, nil, nil}
}

func createTailNode() *SkipNode {
	return &SkipNode{math.MaxInt32, nil, nil, nil, nil}
}

//Insert - as the name suggests
func (sl *SkipList) Insert(i int) *SkipNode {
	sn := &SkipNode{i, nil, nil, nil, nil}
	h := sl.head

	if h.insert(sn) == nil {
		return nil
	}

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	for {
		n := r.Intn(100)

		if n%2 == 0 && h.higher != nil {
			hsn := &SkipNode{sn.data, nil, nil, sn, nil}
			sn.higher = hsn
			sn = hsn
			h = h.higher
			h.insert(sn)
		} else {
			break
		}
	}

	return sn.getBottom()
}

func (sn *SkipNode) insert(nsn *SkipNode) *SkipNode {
	c := sn
	for c != nil {
		if c.data == nsn.data {
			return nil
		}

		if c.data < nsn.data {
			c = c.next
			continue
		}

		c.prev.next = nsn
		nsn.prev = c.prev
		nsn.next = c
		c.prev = nsn
		break
	}
	return nsn
}

func (sn *SkipNode) getBottom() *SkipNode {
	b := sn
	for b.lower != nil {
		b = b.lower
	}
	return b
}

func (sn *SkipNode) getTop() *SkipNode {
	t := sn
	for t.higher != nil {
		t = t.higher
	}
	return t
}

//Find - as the name suggests
func (sl *SkipList) Find(f int) *SkipNode {
	th := sl.head.getTop()
	return th.find(f)
}

func (sn *SkipNode) find(f int) *SkipNode {
	c := sn
	for c.data < f {
		c = c.next
	}

	if c.data == f {
		return c.getBottom()
	}

	if c.prev.lower == nil {
		return nil
	}

	return c.prev.lower.find(f)
}

//Delete - as the name suggests
func (sl *SkipList) Delete(d int) *SkipNode {
	f := sl.head.getTop().find(d)

	n := f
	for n != nil {
		n.prev.next = n.next
		n.next.prev = n.prev
		n = n.higher
	}

	return f
}
