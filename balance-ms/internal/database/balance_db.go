package database

import (
	"balance-ms/internal/entity"
	"database/sql"
)

type BalanceDB struct {
	DB *sql.DB
}

func NewBalanceDB(db *sql.DB) *BalanceDB {
	return &BalanceDB{DB: db}
}

func (b *BalanceDB) GetByAccountID(accountID string) *entity.Balance {
	row := b.DB.QueryRow("SELECT id, account_id, amount, created_at, updated_at FROM balances WHERE account_id = $1", accountID)
	var balance entity.Balance
	err := row.Scan(&balance.ID, &balance.AccountID, &balance.Amount, &balance.CreatedAt, &balance.UpdatedAt)
	if err != nil {
		return nil
	}
	return &balance
}

func (b *BalanceDB) Create(balance *entity.Balance) error {
	stmt, err := b.DB.Prepare("INSERT INTO balances (id, account_id, amount, created_at, updated_at) VALUES(?, ?, ?, ?, ?)")
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

func (b *BalanceDB) Update(balance *entity.Balance) error {
	stmt, err := b.DB.Prepare("UPDATE balances SET amount = ?, updated_at = ? WHERE id = ?")
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
