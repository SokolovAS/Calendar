package database

import (
	"fmt"
	"testing"
)

func TestInitDB(t *testing.T) {
	err := InitDatabase()
	if err != nil {
		fmt.Println("Error with SQLite!")
	}
}
