package db

import "testing"

func TestDbConnection_ConnectionURL(t *testing.T) {
	conn := &Connection{
		DbName:     "test_db",
		User:       "test_user",
		Password:   "wef24fwef",
		Host:       "localhost",
		DisableSSL: true,
	}
	if conn.ConnectionURL() != "postgres://test_user:wef24fwef@localhost/test_db?sslmode=disable" {
		t.Error("Unexpected connection string")
	}
}
