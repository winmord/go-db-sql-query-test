package main

import (
	"database/sql"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	_ "modernc.org/sqlite"
)

func TestSuite(t *testing.T) {
	suite.Run(t, new(SQLiteTestSuite))
}

type SQLiteTestSuite struct {
	suite.Suite
	db *sql.DB
}

const driverName = "sqlite"
const databaseTestName = "demo.db"

func (s *SQLiteTestSuite) SetupSuite() {
	var err error
	s.db, err = sql.Open(driverName, databaseTestName)
	require.NoError(s.T(), err)
}

func (s *SQLiteTestSuite) TearDownSuite() {
	err := s.db.Close()
	require.NoError(s.T(), err)
}

func (s *SQLiteTestSuite) Test_SelectClient_WhenOk() {
	clientID := 1

	s.T().Run("Test_SelectClient_WhenOk", func(t *testing.T) {
		client, err := selectClient(s.db, clientID)
		require.NoError(t, err)
		assert.Equal(t, clientID, client.ID)
		assert.NotEmpty(t, client.FIO)
		assert.NotEmpty(t, client.Birthday)
		assert.NotEmpty(t, client.Email)
		assert.NotEmpty(t, client.FIO)
	})
}

func (s *SQLiteTestSuite) Test_SelectClient_WhenNoClient() {
	clientID := -1
	s.T().Run("Test_SelectClient_WhenNoClient", func(t *testing.T) {
		client, err := selectClient(s.db, clientID)
		require.Error(t, err)

		assert.Equal(t, sql.ErrNoRows, err)
		assert.Empty(t, client.ID)
		assert.Empty(t, client.FIO)
		assert.Empty(t, client.Birthday)
		assert.Empty(t, client.Email)
		assert.Empty(t, client.FIO)
	})
}

func (s *SQLiteTestSuite) Test_InsertClient_ThenSelectAndCheck() {
	cl := Client{
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}

	s.T().Run("Test_InsertClient_ThenSelectAndCheck", func(t *testing.T) {
		id, err := insertClient(s.db, cl)
		cl.ID = id
		require.NotEmpty(t, id)
		require.NoError(t, err)

		client, err := selectClient(s.db, id)
		require.NoError(t, err)
		require.Equal(t, cl.ID, client.ID)
		require.Equal(t, cl.FIO, client.FIO)
		require.Equal(t, cl.Birthday, client.Birthday)
		require.Equal(t, cl.Email, client.Email)
		require.Equal(t, cl.FIO, client.FIO)
	})
}

func (s *SQLiteTestSuite) Test_InsertClient_DeleteClient_ThenCheck() {
	cl := Client{
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}

	s.T().Run("Test_InsertClient_ThenSelectAndCheck", func(t *testing.T) {
		id, err := insertClient(s.db, cl)
		require.NotEmpty(t, id)
		require.NoError(t, err)

		client, err := selectClient(s.db, id)
		require.NoError(t, err)
		require.Equal(t, id, client.ID)
		require.Equal(t, cl.FIO, client.FIO)
		require.Equal(t, cl.Birthday, client.Birthday)
		require.Equal(t, cl.Email, client.Email)
		require.Equal(t, cl.FIO, client.FIO)

		err = deleteClient(s.db, id)
		require.NoError(t, err)

		client, err = selectClient(s.db, id)
		require.Error(t, err)
		require.Equal(t, sql.ErrNoRows, err)
		require.Empty(t, client.ID)
		require.Empty(t, client.FIO)
		require.Empty(t, client.Birthday)
		require.Empty(t, client.Email)
		require.Empty(t, client.FIO)
	})
}
