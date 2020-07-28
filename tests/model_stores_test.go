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

	stores, err := storeInstance.GetStores(app.DB)
	if err != nil {
		t.Errorf("There was an error when getting all stores, err: %v\n", err)
		return
	}

	assert.Equal(t, len(*stores), 2)
	assert.Equal(t, (*stores)[0].Name, "Store 2")
	assert.Equal(t, (*stores)[0].Type, "Type 2")
	assert.Equal(t, (*stores)[0].Active, true)

}

func TestOneStore(t *testing.T) {
	err := refreshStoreTable()
	if err != nil {
		log.Fatal(err)
	}

	_, err = seedOneStore()
	if err != nil {
		log.Fatal(err)
	}

	store, err := storeInstance.GetStores(app.DB)
	if err != nil {
		t.Errorf("There was an error when getting a store, err: %v\n", err)
		return
	}

	assert.Equal(t, len(*store), 1)
	assert.Equal(t, (*store)[0].Name, "Store 1")
	assert.Equal(t, (*store)[0].Type, "Type 1")
	assert.Equal(t, (*store)[0].Active, true)
}
