package pairladder_test

import (
	"context"
	"github.com/adamluzsi/testcase"
	"github.com/google/uuid"
	"pairladder/inmemory"
	"pairladder/models"
	"pairladder/storages"
	"testing"
)

func TestService(tt *testing.T) {
	s := testcase.NewSpec(tt)
	ctx := context.Background()
	storage := testcase.Var[*inmemory.InMemory]{
		ID: uuid.New().String(),
		Init: testcase.VarInitFunc[*inmemory.InMemory](func(t *testcase.T) *inmemory.InMemory {
			return inmemory.NewInMemory()
		}),
	}

	service := testcase.Var[*Service]{
		ID: "idk",
		Init: testcase.VarInitFunc[*Service](func(t *testcase.T) *Service {
			return NewService(storage.Get(t))
		}),
		Before: nil,
		OnLet:  nil,
	}
	s.Describe("#RecordPair", func(s *testcase.Spec) {
		adil := models.Person{}
		chris := models.Person{}
		subject := func(t *testcase.T) error {
			return service.Get(t).RecordPair(ctx, adil, chris)
		}

		s.Test("can record pairs", func(t *testcase.T) {
			err := subject(t)
			t.Must.NoError(err)

			count, err := service.Get(t).GetPair(ctx, adil, chris)
			t.Must.NoError(err)

			t.Must.Equal(1, count)

			err = subject(t)
			t.Must.NoError(err)

			count, err = service.Get(t).GetPair(ctx, adil, chris)
			t.Must.NoError(err)

			t.Must.Equal(2, count)

		})
	})

}

type Service struct {
	storage storages.Storage
}

func NewService(storage storages.Storage) *Service {
	return &Service{storage: storage}
}

func (s Service) RecordPair(ctx context.Context, personA models.Person, personB models.Person) error {
	count, err := s.storage.GetPairCount(ctx, personA.ID, personB.ID)
	if err != nil {
		return err
	}

	return s.storage.SetPairCount(ctx, personA, personB, count+1)
}

func (s Service) GetPair(ctx context.Context, personA models.Person, personB models.Person) (int, error) {
	return s.storage.GetPairCount(ctx, personA.ID, personB.ID)
}
