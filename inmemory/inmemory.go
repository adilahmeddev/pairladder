package inmemory

import (
	"context"
	"github.com/google/uuid"
	"pairladder/models"
)

type InMemory struct {
	pairs map[uuid.UUID]map[uuid.UUID]int
}

func NewInMemory() *InMemory {
	return &InMemory{pairs: map[uuid.UUID]map[uuid.UUID]int{}}
}

func (i *InMemory) GetPairCount(ctx context.Context, id uuid.UUID, id2 uuid.UUID) (int, error) {
	if id.ID() < id2.ID() {
		temp := id2
		id2 = id
		id = temp
	}

	a, ok := i.pairs[id]
	if !ok {
		return 0, nil
	}
	count, ok := a[id2]
	if !ok {
		return 0, nil
	}

	return count, nil
}

func (i *InMemory) SetPairCount(_ context.Context, personA models.Person, personB models.Person, count int) error {
	if personA.ID.ID() < personB.ID.ID() {
		temp := personB
		personB = personA
		personA = temp
	}

	_, ok := i.pairs[personA.ID]
	if !ok {
		i.pairs[personA.ID] = map[uuid.UUID]int{}
	}

	i.pairs[personA.ID][personB.ID] = count
	return nil
}
