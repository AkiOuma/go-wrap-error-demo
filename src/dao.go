package src

import (
	"database/sql"
	"math/rand"
	"strconv"

	"github.com/pkg/errors"
)

type dao struct {
	db *db
}

// create dao object
func NewDAO(db *db) *dao {
	return &dao{
		db: db,
	}
}

// find a user by id
func (d *dao) FindUserByID(id int) (map[string]interface{}, error) {
	syntax := "SELECT id, name FROM user WHERE id = " + strconv.Itoa(rand.Intn(10))
	row, err := d.db.Query(syntax)
	if errors.Cause(err) == sql.ErrNoRows {
		err = errors.Wrapf(err, "dao: %v", syntax)
	} else {
		errors.Wrap(err, "dao: unknown error")
	}
	return row, err
}
