package heap

import "sync"

type Element interface {
	LessThan(e Element) bool
}

type Heap struct {
	sync.Mutex
	heap  []Element
	len   int
	isMin bool
}

func MinHeapConstructor() *Heap {
	return &Heap{heap: make([]Element, 0), isMin: true}
}

func MaxHeapConstructor() *Heap {
	return &Heap{heap: make([]Element, 0), isMin: false}
}

func (h *Heap) IsEmpty() bool {
	h.Lock()
	defer h.Unlock()

	return len(h.heap) == 0
}

func (h *Heap) Len() int {
	h.Lock()
	defer h.Unlock()

	return h.len
}

func (h *Heap) Insert(e Element) {
	h.Lock()
	defer h.Unlock()

	h.heap = append(h.heap, e)
	h.len++
	h.precolateUp(h.len - 1)
}

func (h *Heap) Extract() (e Element) {
	h.Lock()
	defer h.Unlock()

	if h.len == 0 {
		panic("Empty heap, cannot Extract.")
	}

	e, h.heap[0] = h.heap[0], h.heap[h.len-1]
	h.heap = h.heap[:h.len-1]
	h.len--
	h.precolateDown(0)
	return
}

func (h *Heap) Peek() (e Element) {
	h.Lock()
	defer h.Unlock()

	if h.len == 0 {
		panic("Empty heap, cannot Peek.")
	}
	return h.heap[0]
}

func (h *Heap) precolateUp(index int) {
	// 上滤，新元素在堆中上滤直到找出正确位置
	needUp, parent := index, index>>1
	for needUp > 0 && h.less(h.heap[needUp], h.heap[parent]) {
		h.heap[parent], h.heap[needUp] = h.heap[needUp], h.heap[parent]
		needUp, parent = parent, parent>>1
	}
}

func (h *Heap) precolateDown(index int) {
	// 下滤
	needDown, child := index, index<<1+1
	for needDown < h.len && child < h.len {
		// find min(leftChild, rightChild)
		if child+1 < h.len && h.less(h.heap[child+1], h.heap[child]) {
			child++
		}

		if h.less(h.heap[needDown], h.heap[child]) {
			break
		}

		h.heap[needDown], h.heap[child] = h.heap[child], h.heap[needDown]
		needDown, child = child, needDown<<1+1
	}
}

func (h *Heap) less(a, b Element) bool {
	if h.isMin {
		return a.LessThan(b)
	} else {
		return b.LessThan(a)
	}
}
