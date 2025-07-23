package repository

import (
	"database/sql"
	"log"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// DBManager handles database connections
type DBManager struct {
	ProductDB *sql.DB
	PaymentDB *sql.DB
	mu        sync.Mutex
}

// NewDBManager creates a new instance of DBManager
func NewDBManager() *DBManager {
	manager := &DBManager{}
	return manager
}

// ConnectDatabases establishes connections to both databases
func (m *DBManager) ConnectDatabases() error {
	var err error

	// Connect to Product database
	m.ProductDB, err = sql.Open("mysql", "user:password@tcp(localhost:3306)/products")
	if err != nil {
		return err
	}

	// Set connection parameters
	m.ProductDB.SetMaxIdleConns(10)
	m.ProductDB.SetMaxOpenConns(100)
	m.ProductDB.SetConnMaxLifetime(time.Hour)

	// Test connection
	err = m.ProductDB.Ping()
	if err != nil {
		log.Println("Warning: Could not connect to product database:", err)
	} else {
		log.Println("Successfully connected to product database")
	}

	// Connect to Payment database
	m.PaymentDB, err = sql.Open("mysql", "user:password@tcp(localhost:3306)/payments")
	if err != nil {
		return err
	}

	// Set connection parameters
	m.PaymentDB.SetMaxIdleConns(10)
	m.PaymentDB.SetMaxOpenConns(100)
	m.PaymentDB.SetConnMaxLifetime(time.Hour)

	// Test connection
	err = m.PaymentDB.Ping()
	if err != nil {
		log.Println("Warning: Could not connect to payment database:", err)
	} else {
		log.Println("Successfully connected to payment database")
	}

	return nil
}

// Close closes all database connections
func (m *DBManager) Close() {
	if m.ProductDB != nil {
		m.ProductDB.Close()
	}
	if m.PaymentDB != nil {
		m.PaymentDB.Close()
	}
}
