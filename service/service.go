package service

import (
	"context"
	"pairladder/models"
	"pairladder/storages"
)

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
