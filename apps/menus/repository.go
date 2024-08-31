package menus

import (
	"context"
	"database/sql"
)

type repository struct {
	db *sql.DB
}

func newRepository(db *sql.DB) repository {
	return repository{
		db: db,
	}
}

func (r repository) insertMenu(ctx context.Context, menu Menu) (err error) {
	query := `
	INSERT INTO menus (name, category, description, price, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return
	}

	defer stmt.Close()

	_, err = stmt.Exec(menu.Name, menu.Category, menu.Desc, menu.Price, menu.CreatedAt, menu.UpdatedAt)

	return
}
func (r repository) findAll(ctx context.Context) (model []Menu, err error) {
	query := `
		SELECT 
			id, name, category, description, price
			, created_at
			, updated_at
		FROM menus
	`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var menu = Menu{}
		err = rows.Scan(
			&menu.Id,
			&menu.Name,
			&menu.Category,
			&menu.Desc,
			&menu.Price,
			&menu.CreatedAt,
			&menu.UpdatedAt,
		)

		model = append(model, menu)
	}

	return
}
func (r repository) findByID(ctx context.Context, id int)(model Menu, err error) {
	query := `
		SELECT 
			id, name, category, description, price
			, created_at
			, updated_at
		FROM menus
		WHERE id = $1
	`

	row := r.db.QueryRowContext(ctx, query)
	
	err = row.Scan(
		&model.Id, &model.Name, &model.Category, &model.Desc, &model.Price, &model.CreatedAt, &model.UpdatedAt,
	)

	if err != nil {
		return
	}
	return
}
