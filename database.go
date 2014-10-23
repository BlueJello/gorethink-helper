package main

import (
	"time"

	r "github.com/dancannon/gorethink"
)

// NowFunc returns current time, this function is exported in order to be able
// to give the flexiblity to the developer to costumize it accoring to their
// needs
//
//   e.g: return time.Now().UTC()
//
var NowFunc = func() time.Time {
	return time.Now()
}

type DB struct {
	session       *r.Session
	parent        *DB
	singularTable bool
	Value         interface{}
}

func Connect(address, database string) (DB, error) {
	return ConnectWithOpts(r.ConnectOpts{
		Address:     address,
		Database:    database,
		MaxIdle:     10,
		IdleTimeout: time.Second * 10,
	})
}

func ConnectWithOpts(opts r.ConnectOpts) (DB, error) {
	var db DB
	var err error

	db = DB{}

	db.session, err = r.Connect(opts)
	if err != nil {
		return db, err
	}

	db.parent = &db

	return db, nil
}
func (d *DB) Session() *r.Session {
	return d.session
}

func (d *DB) clone() *DB {
	db := DB{session: d.session, parent: d.parent, Value: d.Value}
	return &db
}
