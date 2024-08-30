package employee

import (
	"context"
	"database/sql"
	"fmt"
)

type repository struct {
	db *sql.DB
}

// newEmpolyee implements repositoryContract.

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

func (r repository) newEmployee(ctx context.Context, req Employee) (err error) {
	query := `
	 INSERT INTO employees (name, address, nip, created_at, updated_at)
		VALUES (
			$1, $2, $3, now(), now()
		)`

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, req.Name, req.Address, req.NIP)
	return
}

func (r repository) deleteEmployee(ctx context.Context, id string) (err error) {
	query := `
	DELETE FROM employees WHERE id = $1;
	`

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("gagal menyiapkan statement: %w", err)
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, id)
	if err != nil {
		return fmt.Errorf("gagal mengeksekusi statement: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("gagal mendapatkan jumlah baris yang terpengaruh: %w", err)
	}

	if rowsAffected == 0 {
		return
	}
	return nil
}
