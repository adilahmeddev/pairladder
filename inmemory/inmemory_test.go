package inmemory

import (
	"context"
	"pairladder/contracts"
	"pairladder/storages"
	"testing"
)

func TestInMemory(t *testing.T) {

	storage := contracts.TCStorage{
		Subject: func(tb testing.TB) storages.Storage {
			return NewInMemory()
		},
		MakeCTX: func(tb testing.TB) context.Context {
			return context.Background()
		},
		ResetStorage: func(tb testing.TB, storage storages.Storage) {
			storage = NewInMemory()
		},
	}

	storage.Test(t)
}
