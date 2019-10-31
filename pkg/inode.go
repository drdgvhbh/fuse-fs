package pkg

type INodeGenerator struct {
	next uint64
}

func NewINodeGenerator(initialINode uint64) *INodeGenerator {
	return &INodeGenerator{
		next: initialINode,
	}
}

func (i INodeGenerator) Next() uint64 {
	next := i.next
	i.next++
	return next
}
