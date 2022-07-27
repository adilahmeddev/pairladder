package storages

import (
	"context"
	"github.com/google/uuid"
	"pairladder/models"
)

type Storage interface {
	GetPairCount(ctx context.Context, id uuid.UUID, id2 uuid.UUID) (int, error)
	SetPairCount(ctx context.Context, personA models.Person, personB models.Person, count int) error
}
