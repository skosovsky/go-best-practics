package main

type DB interface {
	GetUsers() []int64
}

type DBImpl struct {
}

func (db *DBImpl) GetUsers() []int64 {
	return []int64{42}
}
