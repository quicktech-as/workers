package connection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	t.Log("Open connection")
	db, err := Get()
	assert.Nil(t, err)

	t.Log("Ping Pong")
	err = db.Ping()
	assert.Nil(t, err)
}
