package minheap

//IntHeap - a heap of type int
type MinHeap []int

//IsEmpty - returns if the heap is empty
func (mh MinHeap) IsEmpty() bool {
	return len(mh) == 0
}

//Add - to add an element to IntHeap
func (mh MinHeap) Add(n int) MinHeap {
	l := len(mh)

	h := append(mh, n)

	h.balance(l)

	return h
}

//Delete - to delete an element from heap
func (mh MinHeap) Delete(i int) MinHeap {
	l := len(mh) - 1

	mh[i], mh[l] = mh[l], mh[i]
	mh = mh[:l]
	mh.balance(i)

	return mh
}

//Balance - balance the heap from given element
func (mh MinHeap) Balance(i int, l int) {
	if mh.IsEmpty() {
		return
	}

	//parent check
	pI := (i - 1) / 2
	if mh[i] < mh[pI] {
		mh[i], mh[pI] = mh[pI], mh[i]
		mh.Balance(pI, l)
		return
	}

	//left child check
	lcI := 2*i + 1
	if lcI > l {
		return
	}

	if mh[i] > mh[lcI] {
		mh[i], mh[lcI] = mh[lcI], mh[i]
		mh.Balance(lcI, l)
		mh.Balance(i, l)
		return
	}

	//right child check
	rcI := lcI + 1
	if rcI > l {
		return
	}

	if mh[i] > mh[rcI] {
		mh[i], mh[rcI] = mh[rcI], mh[i]
		mh.Balance(rcI, l)
		return
	}
}

func (mh MinHeap) balance(i int) {
	if mh.IsEmpty() {
		return
	}

	//parent check
	pI := (i - 1) / 2
	if mh[i] < mh[pI] {
		mh[i], mh[pI] = mh[pI], mh[i]
		mh.balance(pI)
		return
	}

	//left child check
	lcI := 2*i + 1
	if lcI >= len(mh) {
		return
	}

	if mh[i] > mh[lcI] {
		mh[i], mh[lcI] = mh[lcI], mh[i]
		mh.balance(lcI)
		mh.balance(i)
		return
	}

	//right child check
	rcI := lcI + 1
	if rcI >= len(mh) {
		return
	}

	if mh[i] > mh[rcI] {
		mh[i], mh[rcI] = mh[rcI], mh[i]
		mh.balance(rcI)
		return
	}
}
