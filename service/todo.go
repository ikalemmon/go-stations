package service

import (
	"context"
	"database/sql"

	"github.com/TechBowl-japan/go-stations/model"
)

// A TODOService implements CRUD of TODO entities.
type TODOService struct {
	db *sql.DB
}

// NewTODOService returns new TODOService.
func NewTODOService(db *sql.DB) *TODOService {
	return &TODOService{
		db: db,
	}
}

// CreateTODO creates a TODO on DB.
func (s *TODOService) CreateTODO(ctx context.Context, subject, description string) (*model.TODO, error) {
	const (
		insert  = `INSERT INTO todos(subject, description) VALUES(?, ?)`
		confirm = `SELECT subject, description, created_at, updated_at FROM todos WHERE id = ?`
	)
	var todo model.TODO

	res, err := s.db.ExecContext(ctx, insert, subject, description)
	if err != nil {
		return &todo, err
	}

	rows, err := res.LastInsertId()
	if err != nil {
		return &todo, err
	}
	todo.ID = rows

	stmt, err := s.db.PrepareContext(ctx, confirm)
	if err != nil {
		return &todo, err
	} else {
		defer stmt.Close()
		err = stmt.QueryRowContext(ctx, rows).Scan(&todo.Subject, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt)
		if err != nil {
			return &todo, err
		}
	}
	return &todo, err
}

// ReadTODO reads TODOs on DB.
func (s *TODOService) ReadTODO(ctx context.Context, prevID, size int64) ([]*model.TODO, error) {
	const (
		read       = `SELECT id, subject, description, created_at, updated_at FROM todos ORDER BY id DESC LIMIT ?`
		readWithID = `SELECT id, subject, description, created_at, updated_at FROM todos WHERE id < ? ORDER BY id DESC LIMIT ?`
	)
	todos := []*model.TODO{}
	var err error

	if prevID == 0 {
		stmt, err := s.db.PrepareContext(ctx, read)
		if err != nil {
			return todos, err
		} else {
			defer stmt.Close()
			rows, err := stmt.QueryContext(ctx, size)
			if err != nil {
				return todos, err
			}
			defer rows.Close()
			for rows.Next() {
				var todo model.TODO
				if err := rows.Scan(&todo.ID, &todo.Subject, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt); err != nil {
					return todos, err
				}
				todos = append(todos, &todo)
			}
		}

	} else {
		stmt, err := s.db.PrepareContext(ctx, readWithID)
		if err != nil {
			return todos, err
		} else {
			defer stmt.Close()
			rows, err := stmt.QueryContext(ctx, prevID, size)
			if err != nil {
				return todos, err
			}
			defer rows.Close()
			for rows.Next() {
				var todo model.TODO
				if err := rows.Scan(&todo.ID, &todo.Subject, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt); err != nil {
					return todos, err
				}
				todos = append(todos, &todo)
			}
		}
	}

	return todos, err
}

// UpdateTODO updates the TODO on DB.
func (s *TODOService) UpdateTODO(ctx context.Context, id int64, subject, description string) (*model.TODO, error) {
	const (
		update  = `UPDATE todos SET subject = ?, description = ? WHERE id = ?`
		confirm = `SELECT subject, description, created_at, updated_at FROM todos WHERE id = ?`
	)

	var todo model.TODO

	res, err := s.db.ExecContext(ctx, update, subject, description, id)
	if err != nil {
		return &todo, err
	}

	row, err := res.RowsAffected()
	if err != nil {
		return &todo, err
	}
	if row == 0 {
		return nil, model.Run()
	}

	stmt, err := s.db.PrepareContext(ctx, confirm)
	if err != nil {
		return &todo, err
	} else {
		defer stmt.Close()
		err = stmt.QueryRowContext(ctx, id).Scan(&todo.Subject, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt)
		if err != nil {
			return &todo, err
		}
	}
	todo.ID = id
	return &todo, err
}

// DeleteTODO deletes TODOs on DB by ids.
func (s *TODOService) DeleteTODO(ctx context.Context, ids []int64) error {
	const deleteFmt = `DELETE FROM todos WHERE id IN (?%s)`

	return nil
}
