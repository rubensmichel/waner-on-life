package get_test

import (
	"context"
	"testing"

	"github.com/rubensmichel/waner-on-life/internal/infra/db"
	"github.com/rubensmichel/waner-on-life/internal/infra/db/sqlite"
	"github.com/rubensmichel/waner-on-life/internal/infra/validators"
	getuser "github.com/rubensmichel/waner-on-life/internal/usecase/user/get"
	"github.com/stretchr/testify/assert"
)

var repoUser db.Users
var database *sqlite.InMemoryDB

func TestFindUsers(t *testing.T) {
	c := context.Background()
	database, _ = sqlite.NewInMemoryDatabase(sqlite.AllLimitsTables())
	defer database.Close()

	repoUser = db.NewUserDB(database.DB)

	usc := getuser.NewUseCase(
		repoUser,
		validators.NewInput(),
	)

	t.Run("Should return a list of users", func(t *testing.T) {
		result, err := usc.Get(c)
		assert.Nil(t, err)
		assert.NotNil(t, result)
	})
}
