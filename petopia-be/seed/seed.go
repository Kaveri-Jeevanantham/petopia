package seed

import (
	"log"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Product struct {
	ID          uint    `gorm:"primaryKey"`
	Name        string
	Description string
	Price       float64
}

func SeedDatabase(database *gorm.DB) error {
	products := []Product{
		{ID: 1, Name: "Sample Product", Description: "This is a sample product description.", Price: 19.99},
		{ID: 2, Name: "Dog Collar", Description: "A durable and stylish collar for dogs.", Price: 12.99},
		{ID: 3, Name: "Cat Toy", Description: "A fun toy for cats to play with.", Price: 5.49},
		{ID: 4, Name: "Bird Feeder", Description: "A feeder to attract and feed wild birds.", Price: 15.99},
		{ID: 5, Name: "Fish Tank", Description: "A 20-gallon tank for freshwater fish.", Price: 89.99},
		{ID: 6, Name: "Hamster Wheel", Description: "A wheel for hamsters to exercise.", Price: 9.99},
		{ID: 7, Name: "Rabbit Hutch", Description: "A spacious hutch for rabbits.", Price: 120.00},
		{ID: 8, Name: "Dog Bed", Description: "A comfortable bed for dogs of all sizes.", Price: 45.00},
		{ID: 9, Name: "Cat Scratching Post", Description: "A post for cats to scratch and climb.", Price: 25.00},
		{ID: 10, Name: "Pet Shampoo", Description: "A gentle shampoo for pets.", Price: 8.99},
		{ID: 11, Name: "Reptile Heat Lamp", Description: "A heat lamp for reptile enclosures.", Price: 22.50},
	}

	result := database.Clauses(clause.OnConflict{DoNothing: true}).Create(&products)
	if result.Error != nil {
		log.Printf("Error inserting sample product data: %v\n", result.Error)
		return result.Error
	}

	fmt.Println("Sample product data inserted successfully: %v", result.RowsAffected)
	return nil
}
