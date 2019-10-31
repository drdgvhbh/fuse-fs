package pkg

type INodeSequentialGenerator struct {
	next uint64
}

func NewINodeSequentialGenerator(initialINode uint64) *INodeSequentialGenerator {
	return &INodeSequentialGenerator{
		next: initialINode,
	}
}

func (i *INodeSequentialGenerator) Next() uint64 {
	next := i.next
	i.next++
	return next
}
