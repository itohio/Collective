package graph

import (
	"github.com/itohio/collective/backend/pkg/db"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Orm db.DBType
}
