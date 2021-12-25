package repository

import "testing"

func TestInitPG(t *testing.T) {
	_, err := InitPG()
	if err != nil {
		t.Errorf("error %w", err)
	}
}
