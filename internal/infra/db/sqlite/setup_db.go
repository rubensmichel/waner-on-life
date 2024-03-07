package sqlite

import (
	"testing"
)

func Setup(t *testing.T) *InMemoryDB {
	database, err := NewInMemoryDatabase(AllLimitsTables())
	if err != nil {
		panic(err)
	}
	t.Cleanup(func() {
		database.Close()
	})
	return database
}
