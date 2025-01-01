package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/Stupnikjs/musicsite/util"
)

type PostgresRepo struct {
	DB *sql.DB
}

func (m *PostgresRepo) InitTables() error {

	ctx, timeout := context.WithTimeout(context.Background(), time.Second*5)
	defer timeout()
	createTableSongSTMT := `
	CREATE TABLE IF NOT EXISTS song (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		visited INT, 
	);
`
	createTableCommentSTMT := `
	CREATE TABLE IF NOT EXISTS comment (
		id SERIAL PRIMARY KEY,
		content TEXT NOT NULL,
		songId INT, 
	);
`

	_, err := m.DB.ExecContext(ctx, createTableSongSTMT)
	if err != nil {
		return err
	}

	_, err = m.DB.ExecContext(ctx, createTableCommentSTMT)
	if err != nil {
		return err
	}
	songs := util.ListAllSong()

	for _, song := range songs {
		err := m.InsertSong(song)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *PostgresRepo) IncrementListenCount(id int) error {
	ctx, timeout := context.WithTimeout(context.Background(), time.Second*5)
	defer timeout()
	updateSTMT := ` 
		UPDATE song
		SET visited = visited + 1
		WHERE id = $1;
		`
	_, err := m.DB.ExecContext(ctx, updateSTMT, id)
	return err

}

func (m *PostgresRepo) InsertSong(name string) error {
	ctx, timeout := context.WithTimeout(context.Background(), time.Second*5)
	defer timeout()
	insertSTMT := ` 
		INSERT INTO song (name) values ($1) 
		ON CONFLICT (name) DO NOTHING
		`
	_, err := m.DB.ExecContext(ctx, insertSTMT, name)
	return err

}

func (m *PostgresRepo) LeaveComment(id int, comment string) error {
	return nil
}
