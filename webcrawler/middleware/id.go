package middleware

import (
	"math"
	"sync"
)

// ID生成器接口类型
type IdGenerator interface {
	GetUint32() uint32 // 获取一个uint32类型的ID
}

// ID生成器实现类
type cyclicIdGenerator struct {
	sn    uint32     // 当前ID
	ended bool       // 前一个ID是否已经为其类型所能表示的最大值
	mutex sync.Mutex // 互斥锁
}

// 创建ID生成器
func NewIdGenerator() IdGenerator {
	return &cyclicIdGenerator{}
}

func (gen *cyclicIdGenerator) GetUint32() uint32 {
	gen.mutex.Lock()
	defer gen.mutex.Unlock()
	if gen.ended {
		defer func() { gen.ended = false }()
		gen.sn = 0
		return gen.sn
	}
	id := gen.sn
	if id < math.MaxUint32 {
		gen.sn++
	} else {
		gen.ended = true
	}
	return id
}

// ID生成器接口类型2
type IdGenerator2 interface {
	GetUint64() uint64 // 获得一个uint64类型的ID
}

// ID生成器实现类型2
type cyclicIdGenerator2 struct {
	base       cyclicIdGenerator // base ID 生成器
	cycleCount uint64            // 基于uint32类型的取值范围的周期计数
}

// 创建ID生成器2
func NewIdGenerator2() IdGenerator2 {
	return &cyclicIdGenerator2{}
}

func (gen *cyclicIdGenerator2) GetUint64() uint64 {
	var id64 uint64
	if gen.cycleCount%2 == 1 {
		id64 += math.MaxUint32
	}
	id32 := gen.base.GetUint32()
	if id32 == math.MaxUint32 {
		gen.cycleCount++
	}
	id64 += uint64(id32)
	return id64
}
