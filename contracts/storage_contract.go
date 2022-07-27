package contracts

import (
	"context"
	"github.com/adamluzsi/testcase/assert"
	"pairladder/models"
	"pairladder/storages"
	"testing"
)

type TCStorage struct {
	Subject      func(testing.TB) storages.Storage
	MakeCTX      func(tb testing.TB) context.Context
	ResetStorage func(tb testing.TB, storage storages.Storage)
}

func (s TCStorage) Test(t *testing.T) {
	t.Run("when storage has neither of the pairs recorded, then the pair count will not be retrieved", func(t *testing.T) {
		storage := s.Subject(t)
		ctx := s.MakeCTX(t)

		personA := models.NewPerson("a")
		personB := models.NewPerson("b")

		pairCount, err := storage.GetPairCount(ctx, personA.ID, personB.ID)
		assert.NoError(t, err)
		assert.True(t, pairCount == 0, "expected 0 as nethier of them have paired before")
	})

	t.Run("when storage has both people recorded, and they have not paired the count will be 0", func(t *testing.T) {
		storage := s.Subject(t)
		ctx := s.MakeCTX(t)
		personA := models.NewPerson("a")
		personB := models.NewPerson("b")
		personC := models.NewPerson("c")

		assert.NoError(t, storage.SetPairCount(ctx, personA, personC, 1))
		assert.NoError(t, storage.SetPairCount(ctx, personB, personC, 1))

		count, err := storage.GetPairCount(ctx, personA.ID, personB.ID)
		assert.NoError(t, err)

		assert.Equal(t, 0, count)
		s.ResetStorage(t, storage)
	})

	t.Run("when storage has both people recorded, and they have paired the count will be returned", func(t *testing.T) {
		storage := s.Subject(t)
		ctx := s.MakeCTX(t)
		personA := models.NewPerson("a")
		personB := models.NewPerson("b")

		assert.NoError(t, storage.SetPairCount(ctx, personA, personB, 1))

		pairCount, err := storage.GetPairCount(ctx, personA.ID, personB.ID)
		assert.NoError(t, err)
		assert.NoError(t, storage.SetPairCount(ctx, personB, personA, pairCount+1))

		count, err := storage.GetPairCount(ctx, personA.ID, personB.ID)
		assert.NoError(t, err)

		assert.Equal(t, 2, count)
		s.ResetStorage(t, storage)
	})
}
