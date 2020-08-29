package heap

/**
	最大堆：任意非叶子节点值大于其左右孩子节点值
	常见操作：
		1. 构建二叉堆
		2. 入堆
		3. 出堆
 */

type MaxHeap struct {
	maxSize int
	size int
	integers []int
}

func NewMaxHeap(maxSize int) *MaxHeap {
	heap := &MaxHeap{
		maxSize: maxSize,
		size:     0,
		integers: []int{},
	}

	return heap
}

//
func (heap *MaxHeap) isLeaf(index int) bool {
	if index > (heap.size / 2) && index <= heap.size {
		return true
	}

	return false
}

// 父节点下标计算公式：leafIndex = 2 * parentIndex + 1
func (heap *MaxHeap) parent(index int) int {
	return (index - 1) / 2
}

func (head *MaxHeap)  {
	
}
