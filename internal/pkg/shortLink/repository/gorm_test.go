package repository

import (
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"regexp"
	"testing"
)

type Suite struct {
	suite.Suite
	DB *gorm.DB
	mock sqlmock.Sqlmock
	testRep ShortLinkRep
}

func (s *Suite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	require.NoError(s.T(), err)


	s.testRep = ShortLinkRep{
		s.DB,
	}
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) TestCreate() {
	originalLink := "originalLink"
	shortenLink := "shortenLink"
	s.mock.MatchExpectationsInOrder(false)
	s.mock.ExpectBegin()
	s.mock.ExpectQuery("INSERT INTO").
		WithArgs(
			originalLink,
			shortenLink,
		).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).
			AddRow(1))
	s.mock.ExpectCommit()

	err :=s.testRep.Create(originalLink, shortenLink)
	require.NoError(s.T(), err)

	s.mock.MatchExpectationsInOrder(false)
	s.mock.ExpectBegin()
	s.mock.ExpectQuery("INSERT INTO").
		WithArgs(
			originalLink,
			shortenLink,
		).
		WillReturnError(errors.New("Some db error"))
	s.mock.ExpectRollback()

	err =s.testRep.Create(originalLink, shortenLink)
	require.Error(s.T(), err)
}

func (s *Suite) TestGet() {
	originalLink := "originalLink"
	shortenLink := "shortenLink"
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT "original_link" FROM "links" WHERE shorten_link=$1 LIMIT 1`)).
		WithArgs(shortenLink).
		WillReturnRows(sqlmock.NewRows([]string{
			"original_link",
		}).AddRow(
			originalLink,
		))
	ol, err := s.testRep.Get(shortenLink)
	require.NoError(s.T(), err)
	assert.Equal(s.T(), originalLink, ol)


	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT "original_link" FROM "links" WHERE shorten_link=$1 LIMIT 1`)).
		WithArgs(shortenLink).
		WillReturnError(errors.New("Some db error"))
	_, err = s.testRep.Get(shortenLink)
	require.Error(s.T(), err)
}

func (s *Suite) TestExists() {
	originalLink := "originalLink"
	shortenLink := "shortenLink"
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT "shorten_link" FROM "links" WHERE original_link=$1 LIMIT 1`)).
		WithArgs(originalLink).
		WillReturnRows(sqlmock.NewRows([]string{
			"shorten_link",
		}).AddRow(
			shortenLink,
		))
	sl, err := s.testRep.Exists(originalLink)
	require.NoError(s.T(), err)
	assert.Equal(s.T(), shortenLink, sl)


	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT "shorten_link" FROM "links" WHERE original_link=$1 LIMIT 1`)).
		WithArgs(originalLink).
		WillReturnError(errors.New("Some db error"))
	_, err = s.testRep.Exists(originalLink)
	require.Error(s.T(), err)
}

