package bloom

// BloomFilter 布隆过滤器，用于高效去重
type BloomFilter struct {
	bits []uint8
	size uint32
}

// New 创建布隆过滤器
func New(size uint32) *BloomFilter {
	bytes := (size + 7) / 8
	return &BloomFilter{
		bits: make([]uint8, bytes),
		size: size,
	}
}

// hash 简单哈希函数
func (bf *BloomFilter) hash(data string) []uint32 {
	hashes := make([]uint32, 3)
	h := uint32(0)
	for _, ch := range data {
		h = h*31 + uint32(ch)
	}
	hashes[0] = h % bf.size
	hashes[1] = (h * 31) % bf.size
	hashes[2] = (h * 31 * 31) % bf.size
	return hashes
}

// Add 添加元素
func (bf *BloomFilter) Add(data string) {
	for _, idx := range bf.hash(data) {
		byteIdx := idx / 8
		bitIdx := idx % 8
		bf.bits[byteIdx] |= (1 << bitIdx)
	}
}

// Contains 检查元素是否存在
func (bf *BloomFilter) Contains(data string) bool {
	for _, idx := range bf.hash(data) {
		byteIdx := idx / 8
		bitIdx := idx % 8
		if bf.bits[byteIdx]&(1<<bitIdx) == 0 {
			return false
		}
	}
	return true
}
