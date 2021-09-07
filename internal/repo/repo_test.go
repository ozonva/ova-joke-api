//go:build test_unit
// +build test_unit

package repo

import (
	"database/sql"
	"errors"
	"github.com/stretchr/testify/require"
	sqlxmock "github.com/zhashkevych/go-sqlxmock"
	"testing"

	"github.com/ozonva/ova-joke-api/internal/models"
)

var testSomeDbError = errors.New("some-db-error")

func TestJokePgRepo_AddJokes(t *testing.T) {
	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	s := NewJokePgRepo(db)

	tests := []struct {
		name    string
		s       *JokePgRepo
		jokes   []models.Joke
		mock    func()
		want    error
		wantErr bool
	}{
		{
			// When everything works as expected
			name: "OK",
			s:    s,
			jokes: []models.Joke{
				{
					ID:       2,
					Text:     "joke #2",
					AuthorID: 24,
				},
			},
			mock: func() {
				mock.ExpectExec("^INSERT INTO joke \\(id, text, author_id\\) VALUES \\(\\?, \\?, \\?\\)$").
					WithArgs(2, "joke #2", 24).
					WillReturnResult(sqlxmock.NewResult(1, 1))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			err := tt.s.AddJokes(tt.jokes)

			if tt.wantErr {
				require.ErrorIs(t, err, tt.want)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestJokePgRepo_ListJokes(t *testing.T) {
	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	s := NewJokePgRepo(db)

	tests := []struct {
		name    string
		s       *JokePgRepo
		limit   uint64
		offset  uint64
		mock    func()
		want    []models.Joke
		wantErr bool
		err     error
	}{
		{
			name:   "OK",
			s:      s,
			limit:  3,
			offset: 5,
			want: []models.Joke{
				{3, "joke #3", 3},
				{4, "joke #4", 4},
				{5, "joke #5", 5},
			},
			mock: func() {
				rows := sqlxmock.NewRows([]string{"id", "text", "author_id"}).
					AddRow(3, "joke #3", 3).
					AddRow(4, "joke #4", 4).
					AddRow(5, "joke #5", 5)

				mock.ExpectQuery("^SELECT id, text, author_id FROM joke ORDER BY id LIMIT \\d OFFSET \\d$").
					WillReturnRows(rows)
			},
		},
		{
			name:   "empty result",
			s:      s,
			limit:  3,
			offset: 5,
			mock: func() {
				mock.ExpectQuery("^SELECT id, text, author_id FROM joke ORDER BY id LIMIT \\d OFFSET \\d$").
					WillReturnError(sql.ErrNoRows)
			},
		},
		{
			name:   "returns error",
			s:      s,
			limit:  3,
			offset: 5,
			mock: func() {
				mock.ExpectQuery("^SELECT id, text, author_id FROM joke ORDER BY id LIMIT \\d OFFSET \\d$").
					WillReturnError(testSomeDbError)
			},
			wantErr: true,
			err:     testSomeDbError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := tt.s.ListJokes(tt.limit, tt.offset)

			if tt.wantErr {
				require.Error(t, err)
				require.ErrorIs(t, err, tt.err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.want, got)
			}
		})
	}
}

func TestJokePgRepo_DescribeJoke(t *testing.T) {
	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	s := NewJokePgRepo(db)

	tests := []struct {
		name    string
		s       *JokePgRepo
		id      uint64
		mock    func()
		want    *models.Joke
		wantErr bool
		err     error
	}{
		{
			name: "OK",
			s:    s,
			id:   3,
			want: &models.Joke{ID: 3, Text: "joke #3", AuthorID: 3},
			mock: func() {
				rows := sqlxmock.NewRows([]string{"id", "text", "author_id"}).
					AddRow(3, "joke #3", 3)

				mock.ExpectQuery("^SELECT id, text, author_id FROM joke WHERE id=\\$1 LIMIT 1$").
					WithArgs(3).
					WillReturnRows(rows)
			},
		},
		{
			name: "empty",
			s:    s,
			id:   3,
			mock: func() {
				mock.ExpectQuery("^SELECT id, text, author_id FROM joke WHERE id=\\$1 LIMIT 1$").
					WithArgs(3).
					WillReturnError(sql.ErrNoRows)
			},
			wantErr: true,
			err:     sql.ErrNoRows,
		},
		{
			name: "returns error",
			s:    s,
			id:   3,
			mock: func() {
				mock.ExpectQuery("^SELECT id, text, author_id FROM joke WHERE id=\\$1 LIMIT 1$").
					WithArgs(3).
					WillReturnError(testSomeDbError)
			},
			wantErr: true,
			err:     testSomeDbError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := tt.s.DescribeJoke(tt.id)

			if tt.wantErr {
				require.Error(t, err)
				require.ErrorIs(t, err, tt.err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.want, got)
			}
		})
	}
}

func TestJokePgRepo_UpdateJoke(t *testing.T) {
	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	s := NewJokePgRepo(db)

	tests := []struct {
		name    string
		s       *JokePgRepo
		joke    models.Joke
		mock    func()
		wantErr bool
		err     error
	}{
		{
			name: "OK",
			s:    s,
			joke: models.Joke{
				ID:       3,
				Text:     "joke#3",
				AuthorID: 33,
			},
			mock: func() {
				mock.ExpectExec("^UPDATE joke SET text=\\?, author_id=\\? WHERE id=\\?$").
					WithArgs("joke#3", 33, 3).
					WillReturnResult(sqlxmock.NewResult(0, 1))
			},
		},
		{
			name: "returns error",
			s:    s,
			joke: models.Joke{
				ID:       42,
				Text:     "joke#42",
				AuthorID: 4242,
			},
			mock: func() {
				mock.ExpectExec("^UPDATE joke SET text=\\?, author_id=\\? WHERE id=\\?$").
					WithArgs("joke#42", 4242, 42).
					WillReturnError(testSomeDbError)
			},
			wantErr: true,
			err:     testSomeDbError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			err := tt.s.UpdateJoke(tt.joke)

			if tt.wantErr {
				require.Error(t, err)
				require.ErrorIs(t, err, tt.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestJokePgRepo_RemoveJoke(t *testing.T) {
	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	s := NewJokePgRepo(db)

	tests := []struct {
		name    string
		s       *JokePgRepo
		id      uint64
		mock    func()
		wantErr bool
		err     error
	}{
		{
			name: "OK",
			s:    s,
			id:   3,
			mock: func() {
				mock.ExpectExec("^DELETE FROM joke WHERE id=\\$1$").
					WithArgs(3).
					WillReturnResult(sqlxmock.NewResult(0, 1))
			},
		},
		{
			name: "empty",
			s:    s,
			id:   42,
			mock: func() {
				mock.ExpectExec("^DELETE FROM joke WHERE id=\\$1$").
					WithArgs(42).
					WillReturnResult(sqlxmock.NewResult(0, 0))
			},
		},
		{
			name: "returns error",
			s:    s,
			id:   42,
			mock: func() {
				mock.ExpectExec("^DELETE FROM joke WHERE id=\\$1$").
					WithArgs(42).
					WillReturnError(testSomeDbError)
			},
			wantErr: true,
			err:     testSomeDbError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			err := tt.s.RemoveJoke(tt.id)

			if tt.wantErr {
				require.Error(t, err)
				require.ErrorIs(t, err, tt.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
