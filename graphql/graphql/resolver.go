package graphql

import (
	"github.com/edgar-care/graphql/database"
)

type Resolver struct{}

var db *database.DB

func Init(newDb *database.DB) {
	db = newDb
}
