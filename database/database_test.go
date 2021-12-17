package database

import (
	"testing"
)

func TestNewGormDB(t *testing.T) {
	_, err := NewGormDB()
	if err != nil {
		t.Errorf("got err: %s", err)
	}
}
