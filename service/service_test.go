package service_test

import (
	"context"
	"github.com/adamluzsi/testcase"
	"github.com/google/uuid"
	"pairladder/inmemory"
	"pairladder/models"
	service2 "pairladder/service"
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

	service := testcase.Var[*service2.Service]{
		ID: "idk",
		Init: testcase.VarInitFunc[*service2.Service](func(t *testcase.T) *service2.Service {
			return service2.NewService(storage.Get(t))
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
