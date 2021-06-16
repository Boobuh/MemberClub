package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStorage(t *testing.T) {
	s := NewStorage()
	assert.NotNil(t, s)
}

func TestStorage_GetAll(t *testing.T) {
	s := NewStorage()
	assert.NotNil(t, s)
	m := s.GetAll()
	assert.True(t, len(m) == 0)
}
func TestStorage_SaveSuccess(t *testing.T) {
	s := NewStorage()
	assert.NotNil(t, s)
	err := s.Save(Member{Name: "test", Email: "test@gmail.com"})
	assert.NoError(t, err)
	m := s.GetAll()
	assert.True(t, len(m) > 0)
}
func TestStorage_SaveSaveError(t *testing.T) {
	s := NewStorage()
	assert.NotNil(t, s)
	t.Run("error bad name", func(t *testing.T) {
		err := s.Save(Member{Name: "3214", Email: "test@gmail.com"})
		assert.Error(t, err)
	})
	t.Run("error bad email", func(t *testing.T) {
		err := s.Save(Member{Name: "fdsgdfg", Email: "t4214m"})
		assert.Error(t, err)
	})

}
