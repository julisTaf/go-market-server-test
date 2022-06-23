package database

import (
	"testing"
)

func TestInitDB(t *testing.T) {
	err := InitDatabase()
	if err != nil {
		return
	}
}
