package machines

import (
	"database/sql"
)

type Machine struct {
	Id        int64 `json:"id"`
	IsWorking bool  `json:"isWorking"`
}

type Store struct {
	Db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{Db: db}
}

func (s *Store) UpdateMachine(id int64, isWorking bool) error {
	_, err := s.Db.Exec("UPDATE machines SET is_working = $1 WHERE id = $2", isWorking, id)
	return err
}
