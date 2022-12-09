package graphql

import (
	"github.com/edgar-care/graphql/cmd/main/database"
)

type Resolver struct{}

var db *database.DB

func Init(newDb *database.DB) {
	db = newDb
}
