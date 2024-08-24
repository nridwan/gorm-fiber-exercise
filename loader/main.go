package main

import (
	"fmt"
	"gofiber-boilerplate/modules/transactions/transactionsmodel"
	"gofiber-boilerplate/modules/user/usermodel"
	"io"
	"os"

	"ariga.io/atlas-provider-gorm/gormschema"
)

func main() {
	stmts, err := gormschema.New("postgres").Load(
		// user module
		&usermodel.UserModel{},
		// transaction module
		&transactionsmodel.TransactionModel{},
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load gorm schema: %v\n", err)
		os.Exit(1)
	}
	io.WriteString(os.Stdout, `CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)
	io.WriteString(os.Stdout, stmts)
}
