package main

func (d *DB) CreateTable(value interface{}) *DB {
	return d.clone().NewScope(value).createTable().db
}
