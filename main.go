package main

// Custom deque implementation
type Deque struct {
	data [][]int
}

func (d *Deque) PushBack(val []int) {
	d.data = append(d.data, val)
}

func (d *Deque) PushFront(val []int) {
	d.data = append([][]int{val}, d.data...)
}

func (d *Deque) PopFront() []int {
	if len(d.data) == 0 {
		return nil
	}
	val := d.data[0]
	d.data = d.data[1:]
	return val
}

func (d *Deque) Len() int {
	return len(d.data)
}
