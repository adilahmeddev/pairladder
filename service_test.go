package pairladder_test

import (
	"github.com/adamluzsi/testcase"
	"pairladder/models"
	"pairladder/storages"
	"testing"
)

func TestService(tt *testing.T) {
	s := testcase.NewSpec(tt)

	service := testcase.Var[Service]{
		ID: "idk",
		Init: testcase.VarInitFunc[Service](func(t *testcase.T) Service {
			return Service{}
		}),
		Before: nil,
		OnLet:  nil,
	}
	s.Describe("#RecordPair", func(s *testcase.Spec) {
		adil := models.Person{}
		chris := models.Person{}
		subject := func(t *testcase.T) error {
			return service.Get(t).RecordPair(adil, chris)
		}

		s.Test("can record pairs", func(t *testcase.T) {
			err := subject(t)
			t.Must.NoError(err)

			count, err := service.Get(t).GetPair(adil, chris)
			t.Must.NoError(err)

			t.Must.Equal(1, count)
		})
	})

}

type Service struct {
	storage storages.Storage
}

func (s Service) RecordPair(personA models.Person, personB models.Person) error {
	return nil
}

func (s Service) GetPair(personA models.Person, personB models.Person) (int, error) {
	return 1, nil
}
