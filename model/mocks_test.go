package model

import "github.com/stretchr/testify/mock"

type mockTracker struct {
	mock.Mock
	entities map[string]EntityType
}

func (m *mockTracker) Add(entity EntityType) error {
	m.Called(entity)
	return nil
}

func (m *mockTracker) Get(id string) (EntityType, bool) {
	args := m.Called(id)
	return args.Get(0).(EntityType), args.Bool(1)
}