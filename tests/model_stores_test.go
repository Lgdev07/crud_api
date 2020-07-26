package tests

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAllStores(t *testing.T) {
	err := refreshStoreTable()
	if err != nil {
		log.Fatal(err)
	}

	_, err = seedStores()
	if err != nil {
		log.Fatal(err)
	}

	users, err := storeInstance.GetStores(app.DB)
	if err != nil {
		t.Errorf("There was an error when getting all stores, err: %v\n", err)
		return
	}

	assert.Equal(t, len(*users), 2)

}
