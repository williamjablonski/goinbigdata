package mongo

import (
	"math/rand"
	"errors"
)

type Service interface {
	Health() error
	// Business methods go here
}

func New() Service {
	return &service{}
}

type service struct {
	// Some fields
}

func (s *service) Health() error {
	if rand.Intn(2) > 0 {
		return errors.New("Service unavailable")
	}
	return nil
}
