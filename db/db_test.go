package db

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Spin up an in memory db (since we're using sqlite3) for quick testing
//
// This function already includes a cleanup function where when the test completes, the database is closed
func InMemDb(t *testing.T) *DB {
	db, err := Init(":memory:")
	assert.NoError(t, err)
	t.Cleanup(func() {
		db.Close()
		db = nil
	})
	return db
}

// Actual test cases should are in seperate t.Run() instances, unless the test is reasonable concise enough to
// be put under it's own function.
//
// For the most part, if you see function calls before a t.Run(), it's more than likely a setup for those test cases.

func TestInit(t *testing.T) {
	tmp := t.TempDir()
	r, err := Init(tmp)
	assert.NoError(t, err)

	expectedDir := filepath.Join(tmp, "blackheaven")
	expectedDBFile := filepath.Join(expectedDir, "blackheaven.db")

	assert.NotNil(t, r)
	assert.NotEmpty(t, r.Dir)
	assert.NotEmpty(t, r.FilePath)
	assert.Equal(t, expectedDir, r.Dir)
	assert.Equal(t, expectedDBFile, r.FilePath)
}
