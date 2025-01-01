package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type PostgresRepo struct {
	DB *sql.DB
}

func (m *PostgresRepo) InitTables() error {

	ctx, timeout := context.WithTimeout(context.Background(), time.Second*5)
	defer timeout()
	createTableSTMT := `
	CREATE TABLE IF NOT EXISTS test (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
`
	_, err := m.DB.ExecContext(ctx, createTableSTMT)
	if err != nil {
		return err
	}

	return nil
}

func (m *PostgresRepo) IncrementListenCount(id int) {
	fmt.Println("incrementing")
}

func (m *PostgresRepo) LeaveComment(id int, comment string) {
	fmt.Println("incrementing")
}
