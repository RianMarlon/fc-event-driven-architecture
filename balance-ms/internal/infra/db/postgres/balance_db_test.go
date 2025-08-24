package database

import (
	"balance-ms/internal/entity"
	"database/sql"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3" // driver sqlite
	"github.com/stretchr/testify/suite"
)

type BalanceDBTestSuite struct {
	suite.Suite
	DB        *sql.DB
	BalanceDB *BalanceDB
}

func (s *BalanceDBTestSuite) SetupTest() {
	var err error
	s.DB, err = sql.Open("sqlite3", ":memory:")
	s.Require().NoError(err)

	// Criar tabela
	_, err = s.DB.Exec(`
		CREATE TABLE IF NOT EXISTS balances (
			id VARCHAR(36) NOT NULL PRIMARY KEY,
			account_id VARCHAR(36) NOT NULL,
			amount DECIMAL(10, 2) NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);
	`)
	s.Require().NoError(err)

	s.BalanceDB = &BalanceDB{DB: s.DB}
}

func (s *BalanceDBTestSuite) TearDownTest() {
	s.DB.Close()
}

func (s *BalanceDBTestSuite) TestGetByAccountID() {
	balance, err := entity.NewBalance("123e4567-e89b-12d3-a456-426614174000", 100.00)
	s.Require().NoError(err)

	stmt, err := s.DB.Prepare("INSERT INTO balances (id, account_id, amount, created_at, updated_at) VALUES(?, ?, ?, ?, ?)")
	s.Require().NoError(err)

	_, err = stmt.Exec(balance.ID, balance.AccountID, balance.Amount, balance.CreatedAt, balance.UpdatedAt)
	s.Require().NoError(err)

	balanceCreated := s.BalanceDB.GetByAccountID(balance.AccountID)
	s.Require().Equal(balance.ID, balanceCreated.ID)
	s.Require().Equal(balance.AccountID, balanceCreated.AccountID)
	s.Require().Equal(balance.Amount, balanceCreated.Amount)
	s.Require().WithinDuration(balance.CreatedAt, balanceCreated.CreatedAt, time.Second)
	s.Require().WithinDuration(balance.UpdatedAt, balanceCreated.UpdatedAt, time.Second)
}

func (s *BalanceDBTestSuite) TestCreate() {
	balance, err := entity.NewBalance("123e4567-e89b-12d3-a456-426614174000", 100.00)
	s.Require().NoError(err)

	err = s.BalanceDB.Create(balance)
	s.Require().NoError(err)

	var balanceCreated entity.Balance
	row := s.DB.QueryRow("SELECT id, account_id, amount, created_at, updated_at FROM balances WHERE id = ?", balance.ID)
	err = row.Scan(&balanceCreated.ID, &balanceCreated.AccountID, &balanceCreated.Amount, &balanceCreated.CreatedAt, &balanceCreated.UpdatedAt)
	s.Require().NoError(err)
	s.Require().NotNil(balanceCreated)
}

func (s *BalanceDBTestSuite) TestUpdate() {
	balance, err := entity.NewBalance("123e4567-e89b-12d3-a456-426614174000", 100.00)
	s.Require().NoError(err)

	stmt, err := s.DB.Prepare("INSERT INTO balances (id, account_id, amount, created_at, updated_at) VALUES(?, ?, ?, ?, ?)")
	s.Require().NoError(err)

	_, err = stmt.Exec(balance.ID, balance.AccountID, balance.Amount, balance.CreatedAt, balance.UpdatedAt)
	s.Require().NoError(err)

	balance.UpdateBalance(300)

	err = s.BalanceDB.Update(balance)
	s.Require().NoError(err)

	var balanceUpdated entity.Balance
	row := s.DB.QueryRow("SELECT id, account_id, amount, created_at, updated_at FROM balances WHERE id = ?", balance.ID)
	err = row.Scan(&balanceUpdated.ID, &balanceUpdated.AccountID, &balanceUpdated.Amount, &balanceUpdated.CreatedAt, &balanceUpdated.UpdatedAt)
	s.Require().NoError(err)
	s.Require().NotNil(balanceUpdated)
	s.Require().Equal(balance.Amount, balanceUpdated.Amount)
}

func TestBalanceDBTestSuite(t *testing.T) {
	suite.Run(t, new(BalanceDBTestSuite))
}
