package tests

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/Lgdev07/crud_api/controllers"
	"github.com/Lgdev07/crud_api/models"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var app = controllers.App{}
var storeInstance = models.Store{}

func TestMain(m *testing.M) {
	var err error
	err = godotenv.Load(os.ExpandEnv("./../.env"))
	if err != nil {
		log.Fatalf("Error getting env %v\n", err)
	}

	Database()

	os.Exit(m.Run())

}

func Database() {

	var err error

	DBURL := fmt.Sprintf(`host=%s port=%s user=%s dbname=%s 
		sslmode=disable password=%s`, os.Getenv("TEST_DB_HOST"),
		os.Getenv("DB_PORT"), os.Getenv("DB_USER"),
		os.Getenv("TEST_DB_NAME"), os.Getenv("DB_PASSWORD"))

	app.DB, err = gorm.Open("postgres", DBURL)
	if err != nil {
		fmt.Printf("Cannot connect to %s database\n", "postgres")
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database\n", "postgres")
	}

}

func refreshStoreTable() error {
	err := app.DB.DropTableIfExists(&models.Store{}).Error
	if err != nil {
		return err
	}
	err = app.DB.AutoMigrate(&models.Store{}).Error
	if err != nil {
		return err
	}
	log.Printf("Successfully refreshed table")
	return nil
}

func seedOneStore() (models.Store, error) {
	store := models.Store{
		Name:   "Store 1",
		Type:   "Type 1",
		Active: true,
	}

	err := app.DB.Model(&models.Store{}).Create(&store).Error
	if err != nil {
		return models.Store{}, err
	}
	return store, nil
}

func seedStores() ([]models.Store, error) {

	var err error
	if err != nil {
		return nil, err
	}

	stores := []models.Store{
		{
			Name:   "Store 2",
			Type:   "Type 2",
			Active: true,
		},
		{
			Name:   "Store 3",
			Type:   "Type 3",
			Active: true,
		},
	}

	for i := range stores {
		err := app.DB.Model(&models.Store{}).Create(&stores[i]).Error
		if err != nil {
			return []models.Store{}, err
		}
	}
	return stores, nil

}
