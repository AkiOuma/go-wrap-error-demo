package src

import (
	"database/sql"
	"math/rand"

	"github.com/pkg/errors"
)

// define unknown error
var ErrUnknow = errors.New("sql: unkonwn error")

// a fake db object
type db struct{}

// create a fake db object
func NewDB() *db {
	return &db{}
}

// simulate a query method in fake db object, may cause 3 kinds of result
//
// return a map with id and name
//
// cause sql.ErrNoRow
//
// cause unknown error
func (d *db) Query(syntax string) (map[string]interface{}, error) {
	id := rand.Intn(10)
	status := rand.Intn(10)
	switch {
	case status < 3:
		return nil, sql.ErrNoRows
	case status >= 6:
		return map[string]interface{}{
			"id":   id,
			"name": "user",
		}, nil
	default:
		return nil, ErrUnknow
	}
}
