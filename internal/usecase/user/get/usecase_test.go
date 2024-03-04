package get_test

import (
	"context"
	"testing"

	"github.com/rubensmichel/waner-on-life/internal/infra/validators"
	getuser "github.com/rubensmichel/waner-on-life/internal/usecase/user/get"
	"github.com/stretchr/testify/assert"
)

func TestFindUsers(t *testing.T) {
	c := context.Background()
	usc := getuser.NewUseCase(
		validators.NewInput(),
	)

	t.Run("Should return a list of users", func(t *testing.T) {
		result, err := usc.Get(c)

		assert.Nil(t, err)
		assert.Equal(t, result[0].ID, 1)
	})
}
