package postgres

import (
	entity "github.com/rubensmichel/waner-on-life/internal/domain"
)

func AllLimitsTables() []interface{} {
	return []interface{}{
		&entity.User{},
	}
}
