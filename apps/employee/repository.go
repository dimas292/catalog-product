package employee

import (
	"context"
	"database/sql"
)

type repository struct {
	db *sql.DB
}

func newRepository(db *sql.DB) repository {
	return repository{db: db}
}

func (r repository) findAllEmployees(ctx context.Context) (res []Employee, err error) {

	query := `
	SELECT id, name, address, nip, created_at
	FROM employees
	`

	row, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return
	}

	defer row.Close()

	for row.Next() {

		var emp = Employee{}

		if err := row.Scan(&emp.Id, &emp.Name, &emp.Address, &emp.NIP, &emp.CreatedAt); err != nil {
			return []Employee{}, nil
		}

		res = append(res, emp)
	}

	return res, nil
}
