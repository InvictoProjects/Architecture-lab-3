package balancers

import (
	"database/sql"
)

type Balancer struct {
	Id                 int64   `json:"id"`
	UsedMachines       []int64 `json:"usedMachines"`
	TotalMachinesCount int64   `json:"totalMachinesCount"`
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
		machines, err := s.getBalancerMachines(b.Id)
		if err != nil {
			return nil, err
		}
		b.UsedMachines = machines
		res = append(res, &b)
	}
	if res == nil {
		res = make([]*Balancer, 0)
	}
	return res, nil
}

func (s *Store) getBalancerMachines(id int64) ([]int64, error) {
	rows, err := s.Db.Query("SELECT machine_id FROM connections c JOIN machines m on c.machine_id = m.id where balancer_id = $1 AND m.is_working = true;", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var machines []int64
	for rows.Next() {
		var m int64
		if err := rows.Scan(&m); err != nil {
			return nil, err
		}

		machines = append(machines, m)
	}
	if machines == nil {
		machines = make([]int64, 0)
	}

	return machines, nil
}
