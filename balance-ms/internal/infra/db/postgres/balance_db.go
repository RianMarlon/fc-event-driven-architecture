package database

import (
	"balance-ms/internal/domain/entity"
	"database/sql"
)

type BalanceRepositoryDB struct {
	DB *sql.DB
}

func NewBalanceRepositoryDB(db *sql.DB) *BalanceRepositoryDB {
	return &BalanceRepositoryDB{DB: db}
}

func (b *BalanceRepositoryDB) GetByAccountID(accountID string) *entity.Balance {
	row := b.DB.QueryRow("SELECT id, account_id, amount, created_at, updated_at FROM balances WHERE account_id = $1", accountID)
	var balance entity.Balance
	err := row.Scan(&balance.ID, &balance.AccountID, &balance.Amount, &balance.CreatedAt, &balance.UpdatedAt)
	if err != nil {
		return nil
	}
	return &balance
}

func (b *BalanceRepositoryDB) Create(balance *entity.Balance) error {
	stmt, err := b.DB.Prepare("INSERT INTO balances (id, account_id, amount, created_at, updated_at) VALUES($1, $2, $3, $4, $5)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(balance.ID, balance.AccountID, balance.Amount, balance.CreatedAt, balance.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (b *BalanceRepositoryDB) Update(balance *entity.Balance) error {
	stmt, err := b.DB.Prepare("UPDATE balances SET amount = $1, updated_at = $2 WHERE id = $3")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(balance.Amount, balance.UpdatedAt, balance.ID)
	if err != nil {
		return err
	}
	return nil
}
