package main

import (
	"fmt"
	"idm/inner/database"
)

func main() {
	db := database.ConnectDb()

	exec := db.MustExec("select 1")
	fmt.Println(exec.RowsAffected())
}
