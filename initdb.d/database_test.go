package initdb_d

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitDB(t *testing.T) {
	err := InitDatabase()
	assert.NoError(t, err)
}
