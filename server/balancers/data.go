package balancers

import (
	"database/sql"
)

type Balancer struct {
	Id                 int64 `json:"id"`
	TotalMachinesCount int64 `json:"totalMachinesCount"`
}

type Store struct {
	Db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{Db: db}
}

func (s *Store) ListBalancers() ([]*Balancer, error) {
	rows, err := s.Db.Query("SELECT id, count(machine_id) AS total_machines FROM balancers JOIN connections c on balancers.id = c.balancer_id GROUP BY id;")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var res []*Balancer
	for rows.Next() {
		var b Balancer
		if err := rows.Scan(&b.Id, &b.TotalMachinesCount); err != nil {
			return nil, err
		}
		res = append(res, &b)
	}
	if res == nil {
		res = make([]*Balancer, 0)
	}
	return res, nil
}
