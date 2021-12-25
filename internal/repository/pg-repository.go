package repository

import (
	"Calendar/entity"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type RepoPG struct {
	conn *sql.DB
}

func NewRepoPG() *RepoPG {
	db, err := InitPG()
	if err != nil {
		panic(err)
	}
	return &RepoPG{
		conn: db,
	}
}

func (r *RepoPG) GetAllEvents() ([]entity.Event, error) {
	rows, err := r.conn.Query("select * from event")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var events []entity.Event

	for rows.Next() {
		e := entity.Event{}
		err := rows.Scan(&e.Id, &e.IdUser, &e.Title, &e.Description)
		if err != nil {
			fmt.Println(err)
			continue
		}
		events = append(events, e)
	}
	return events, err
}

func (r *RepoPG) GetOne(id int) (entity.Event, error) {
	row := r.conn.QueryRow("select * from event where id = $1", id)
	e := entity.Event{}
	err := row.Scan(&e.Id, &e.IdUser, &e.Title, &e.Description)
	if err != nil {
		panic(err)
	}
	return e, err
}

func (r *RepoPG) Add(e entity.Event) error {
	_, err := r.conn.Exec("insert into event (iduser, title, description, id) values ($1, $2, $3, $4)", e.IdUser, e.Title, e.Description, e.Id)
	if err != nil {
		return RepoError{500, "error getting data"}
	}
	return err
}

func (r *RepoPG) Update(e entity.Event) error {
	_, err := r.conn.Exec("update event set title = $1 where description = $2", e.Title, e.Title)
	if err != nil {
		return RepoError{500, "error updating data"}
	}
	return err
}

func (r *RepoPG) Delete(id int) error {
	_, err := r.conn.Exec("delete from event where id = $1", id)
	if err != nil {
		return RepoError{500, "Error deleting an event"}
	}
	return err
}

type RepoError struct {
	Code    int
	Message string
}

func (e RepoError) Error() string {
	return fmt.Sprintf("Code %d, message: %v", e.Code, e.Message)
}

func InitPG() (*sql.DB, error) {
	connStr := "user=postgres password=somePassword dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return db, RepoError{500, "error PG connection"}
	}
	defer db.Close()
	return db, nil
}
