package repo

import (
	"database/sql"
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"

	"github.com/ozonva/ova-joke-api/internal/models"
)

var jokeTblName = "joke"

type JokePgRepo struct {
	db *sqlx.DB
}

// AddJokes from entities into database in single query.
func (j JokePgRepo) AddJokes(entities []models.Joke) error {
	for _, ee := range entities {
		_, err := j.db.NamedExec(
			fmt.Sprintf("INSERT INTO %s (id, text, author_id) VALUES (:id, :text, :author_id)", jokeTblName),
			map[string]interface{}{
				"id":        ee.ID,
				"text":      ee.Text,
				"author_id": ee.AuthorID,
			},
		)
		if err != nil {
			return err
		}
	}
	return nil
}

// ListJokes returns jokes from database by limit/offset ordered by id asc.
func (j JokePgRepo) ListJokes(limit, offset uint64) ([]models.Joke, error) {
	query, _, err := sq.Select("id", "text", "author_id").
		From(jokeTblName).
		Limit(limit).
		Offset(offset).
		OrderBy("id").
		ToSql()
	if err != nil {
		panic(fmt.Sprintf("invalid query: %q", query))
	}

	var jokes []models.Joke
	err = j.db.Select(&jokes, query)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return jokes, nil
}

// DescribeJoke from database where id=jokeID.
func (j JokePgRepo) DescribeJoke(jokeID models.JokeID) (*models.Joke, error) {
	joke := &models.Joke{}
	err := j.db.Get(
		joke,
		fmt.Sprintf("SELECT id, text, author_id FROM %s WHERE id=$1 LIMIT 1", jokeTblName),
		jokeID,
	)
	if err != nil {
		return nil, err
	}

	return joke, nil
}

// RemoveJoke from database where id=jokeID.
func (j JokePgRepo) RemoveJoke(jokeID models.JokeID) error {
	_, err := j.db.Exec(
		fmt.Sprintf("DELETE FROM %s WHERE id=$1", jokeTblName),
		jokeID,
	)
	return err
}

// UpdateJoke with from database where id=jokeID.
func (j JokePgRepo) UpdateJoke(joke models.Joke) error {
	_, err := j.db.NamedExec(
		fmt.Sprintf("UPDATE %s SET text=:text, author_id=:author_id WHERE id=:id", jokeTblName),
		map[string]interface{}{
			"id":        joke.ID,
			"text":      joke.Text,
			"author_id": joke.AuthorID,
		})
	return err
}

func NewJokePgRepo(db *sqlx.DB) *JokePgRepo {
	return &JokePgRepo{
		db: db,
	}
}
