package seed

import (
	"context"
	"log"

	"petopia-be/dao"
	"petopia-be/db"
	mongo_models "petopia-be/models/mongo"
)

// SeedMongoDatabase seeds MongoDB with initial data
func SeedMongoDatabase() error {
	ctx := context.Background()

	// Get MongoDB collections
	brandCollection := db.GetMongoCollection("brands")
	productCollection := db.GetMongoCollection("products")
	reviewCollection := db.GetMongoCollection("reviews")

	// Create DAOs
	brandDAO := dao.NewProductBrandDAO(brandCollection)
	productDAO := dao.NewProductDetailsDAO(productCollection)
	reviewDAO := dao.NewCustomerReviewDAO(reviewCollection)

	// Seed brands first
	if err := seedProductBrands(ctx, brandDAO); err != nil {
		return err
	}

	// Seed products
	if err := seedProductDetails(ctx, productDAO); err != nil {
		return err
	}

	// Seed reviews
	if err := seedCustomerReviews(ctx, reviewDAO); err != nil {
		return err
	}

	log.Println("MongoDB seeding completed successfully!")
	return nil
}

// seedProductBrands seeds the product brands collection
func seedProductBrands(ctx context.Context, brandDAO *dao.ProductBrandDAO) error {
	brands := []*mongo_models.ProductBrand{
		{
			BrandID:         1,
			BrandName:       "PetNutrition Plus",
			Description:     "Premium pet food and nutrition products for healthy pets",
			LogoURL:         "https://example.com/logos/petnutrition.png",
			Website:         "https://petnutritionplus.com",
			Country:         "USA",
			EstablishedYear: 2010,
			Categories:      []string{"dog-food", "cat-food", "supplements"},
			IsActive:        true,
		},
		{
			BrandID:         2,
			BrandName:       "PlayTime Pets",
			Description:     "Fun and engaging toys for dogs and cats",
			LogoURL:         "https://example.com/logos/playtime.png",
			Website:         "https://playtimepets.com",
			Country:         "Canada",
			EstablishedYear: 2015,
			Categories:      []string{"dog-toys", "cat-toys", "interactive-toys"},
			IsActive:        true,
		},
		{
			BrandID:         3,
			BrandName:       "HealthyPaws",
			Description:     "Natural and organic pet care products",
			LogoURL:         "https://example.com/logos/healthypaws.png",
			Website:         "https://healthypaws.com",
			Country:         "Germany",
			EstablishedYear: 2008,
			Categories:      []string{"grooming", "supplements", "natural-treats"},
			IsActive:        true,
		},
		{
			BrandID:         4,
			BrandName:       "AquaLife",
			Description:     "Complete aquarium and fish care solutions",
			LogoURL:         "https://example.com/logos/aqualife.png",
			Website:         "https://aqualife.com",
			Country:         "Japan",
			EstablishedYear: 2005,
			Categories:      []string{"fish-food", "aquarium-supplies", "water-treatment"},
			IsActive:        true,
		},
		{
			BrandID:         5,
			BrandName:       "FeatherFriends",
			Description:     "Specialized products for birds and small animals",
			LogoURL:         "https://example.com/logos/featherfriends.png",
			Website:         "https://featherfriends.com",
			Country:         "Australia",
			EstablishedYear: 2012,
			Categories:      []string{"bird-food", "bird-toys", "small-animal-care"},
			IsActive:        true,
		},
		{
			BrandID:         6,
			BrandName:       "WildNature",
			Description:     "Raw and natural pet food inspired by nature",
			LogoURL:         "https://example.com/logos/wildnature.png",
			Website:         "https://wildnature.com",
			Country:         "New Zealand",
			EstablishedYear: 2018,
			Categories:      []string{"raw-food", "natural-treats", "freeze-dried"},
			IsActive:        true,
		},
		{
			BrandID:         7,
			BrandName:       "ComfortCare",
			Description:     "Comfortable beds and accessories for pets",
			LogoURL:         "https://example.com/logos/comfortcare.png",
			Website:         "https://comfortcare.com",
			Country:         "UK",
			EstablishedYear: 2014,
			Categories:      []string{"beds", "blankets", "comfort-accessories"},
			IsActive:        true,
		},
		{
			BrandID:         8,
			BrandName:       "TechPet",
			Description:     "Smart technology products for modern pet care",
			LogoURL:         "https://example.com/logos/techpet.png",
			Website:         "https://techpet.com",
			Country:         "South Korea",
			EstablishedYear: 2020,
			Categories:      []string{"smart-feeders", "gps-trackers", "health-monitors"},
			IsActive:        true,
		},
		{
			BrandID:         9,
			BrandName:       "EcoTails",
			Description:     "Sustainable and eco-friendly pet products",
			LogoURL:         "https://example.com/logos/ecotails.png",
			Website:         "https://ecotails.com",
			Country:         "Netherlands",
			EstablishedYear: 2019,
			Categories:      []string{"eco-toys", "biodegradable-waste", "sustainable-food"},
			IsActive:        true,
		},
		{
			BrandID:         10,
			BrandName:       "VetApproved",
			Description:     "Veterinarian recommended health and wellness products",
			LogoURL:         "https://example.com/logos/vetapproved.png",
			Website:         "https://vetapproved.com",
			Country:         "USA",
			EstablishedYear: 2016,
			Categories:      []string{"medication", "supplements", "health-monitoring"},
			IsActive:        true,
		},
	}

	for _, brand := range brands {
		if err := brandDAO.CreateProductBrand(ctx, brand); err != nil {
			log.Printf("Error seeding brand %s: %v", brand.BrandName, err)
		}
	}

	log.Println("Product brands seeded successfully!")
	return nil
}

// seedProductDetails seeds the product details collection
func seedProductDetails(ctx context.Context, productDAO *dao.ProductDetailsDAO) error {
	products := []*mongo_models.ProductDetails{
		{
			ProductID:   101,
			ProductName: "Premium Adult Dog Food - Chicken & Rice",
			Description: "High-quality protein-rich dog food made with real chicken and brown rice. Perfect for adult dogs of all sizes.",
			BrandID:     1,
			BrandName:   "PetNutrition Plus",
			SellerID:    1,
			Category:    "dog-food",
			ItemDimensions: map[string]interface{}{
				"weight": "5kg",
				"length": "30cm",
				"width":  "20cm",
				"height": "35cm",
			},
			Price:        29.99,
			Discount:     5.0,
			Availability: true,
		},
		{
			ProductID:   102,
			ProductName: "Interactive Puzzle Toy for Dogs",
			Description: "Mental stimulation toy that challenges your dog while providing treats. Great for reducing boredom and anxiety.",
			BrandID:     2,
			BrandName:   "PlayTime Pets",
			SellerID:    2,
			Category:    "dog-toys",
			ItemDimensions: map[string]interface{}{
				"diameter": "25cm",
				"height":   "8cm",
				"weight":   "500g",
			},
			Price:        19.99,
			Discount:     0.0,
			Availability: true,
		},
		{
			ProductID:   103,
			ProductName: "Natural Cat Shampoo - Lavender Scented",
			Description: "Gentle, natural shampoo for cats with sensitive skin. Infused with lavender for a calming effect.",
			BrandID:     3,
			BrandName:   "HealthyPaws",
			SellerID:    1,
			Category:    "grooming",
			ItemDimensions: map[string]interface{}{
				"volume": "250ml",
				"height": "18cm",
				"width":  "6cm",
			},
			Price:        12.99,
			Discount:     10.0,
			Availability: true,
		},
		{
			ProductID:   104,
			ProductName: "Tropical Fish Food Flakes",
			Description: "Nutritious fish food flakes suitable for all tropical fish species. Enhances color and promotes healthy growth.",
			BrandID:     4,
			BrandName:   "AquaLife",
			SellerID:    3,
			Category:    "fish-food",
			ItemDimensions: map[string]interface{}{
				"weight": "100g",
				"height": "12cm",
				"width":  "8cm",
			},
			Price:        8.99,
			Discount:     0.0,
			Availability: true,
		},
		{
			ProductID:   105,
			ProductName: "Bird Seed Mix - Premium Blend",
			Description: "Premium mixed seeds for parrots and other large birds. Contains sunflower seeds, nuts, and dried fruits.",
			BrandID:     5,
			BrandName:   "FeatherFriends",
			SellerID:    4,
			Category:    "bird-food",
			ItemDimensions: map[string]interface{}{
				"weight": "2kg",
				"height": "25cm",
				"width":  "15cm",
			},
			Price:        15.99,
			Discount:     15.0,
			Availability: true,
		},
		{
			ProductID:   106,
			ProductName: "Raw Freeze-Dried Dog Treats",
			Description: "Single-ingredient freeze-dried beef liver treats. Perfect for training and as a healthy snack.",
			BrandID:     6,
			BrandName:   "WildNature",
			SellerID:    2,
			Category:    "natural-treats",
			ItemDimensions: map[string]interface{}{
				"weight": "150g",
				"height": "20cm",
				"width":  "10cm",
			},
			Price:        22.99,
			Discount:     0.0,
			Availability: true,
		},
		{
			ProductID:   107,
			ProductName: "Orthopedic Memory Foam Dog Bed",
			Description: "Supportive memory foam bed for senior dogs or dogs with joint issues. Removable, washable cover.",
			BrandID:     7,
			BrandName:   "ComfortCare",
			SellerID:    5,
			Category:    "beds",
			ItemDimensions: map[string]interface{}{
				"length": "90cm",
				"width":  "60cm",
				"height": "15cm",
			},
			Price:        79.99,
			Discount:     20.0,
			Availability: true,
		},
		{
			ProductID:   108,
			ProductName: "Smart Pet GPS Tracker",
			Description: "Real-time GPS tracking collar for dogs and cats. Waterproof with long battery life and mobile app.",
			BrandID:     8,
			BrandName:   "TechPet",
			SellerID:    3,
			Category:    "gps-trackers",
			ItemDimensions: map[string]interface{}{
				"length": "5cm",
				"width":  "3cm",
				"height": "1.5cm",
				"weight": "35g",
			},
			Price:        49.99,
			Discount:     10.0,
			Availability: true,
		},
		{
			ProductID:   109,
			ProductName: "Biodegradable Waste Bags - 300 Count",
			Description: "Eco-friendly dog waste bags that decompose naturally. Strong and leak-proof for responsible pet ownership.",
			BrandID:     9,
			BrandName:   "EcoTails",
			SellerID:    4,
			Category:    "biodegradable-waste",
			ItemDimensions: map[string]interface{}{
				"count":  300,
				"length": "23cm",
				"width":  "15cm",
			},
			Price:        11.99,
			Discount:     5.0,
			Availability: true,
		},
		{
			ProductID:   110,
			ProductName: "Joint Health Supplement for Dogs",
			Description: "Veterinarian-formulated glucosamine and chondroitin supplement for joint health and mobility support.",
			BrandID:     10,
			BrandName:   "VetApproved",
			SellerID:    5,
			Category:    "supplements",
			ItemDimensions: map[string]interface{}{
				"tablets": 120,
				"height":  "10cm",
				"width":   "6cm",
			},
			Price:        34.99,
			Discount:     0.0,
			Availability: true,
		},
	}

	for _, product := range products {
		if err := productDAO.CreateProductDetails(ctx, product); err != nil {
			log.Printf("Error seeding product %s: %v", product.ProductName, err)
		}
	}

	log.Println("Product details seeded successfully!")
	return nil
}

// seedCustomerReviews seeds the customer reviews collection
func seedCustomerReviews(ctx context.Context, reviewDAO *dao.CustomerReviewDAO) error {
	reviews := []*mongo_models.CustomerReview{
		{
			ProductID:        101,
			CustomerID:       1,
			CustomerName:     "Sarah Johnson",
			Rating:           4.5,
			Title:            "Great food for my Golden Retriever",
			Comment:          "My dog loves this food and his coat has never looked better. Great quality ingredients and good value for money.",
			Images:           []string{"review1_img1.jpg"},
			VerifiedPurchase: true,
			HelpfulVotes:     23,
			Filters:          map[string]string{"dog_size": "large", "age": "adult"},
			PetInfo: mongo_models.PetInfo{
				PetType:  "dog",
				PetBreed: "Golden Retriever",
				PetAge:   4,
				PetSize:  "large",
			},
		},
		{
			ProductID:        101,
			CustomerID:       2,
			CustomerName:     "Mike Chen",
			Rating:           5.0,
			Title:            "Excellent quality dog food",
			Comment:          "Been feeding this to my German Shepherd for 6 months. Excellent quality, no fillers, and great protein content.",
			Images:           []string{},
			VerifiedPurchase: true,
			HelpfulVotes:     15,
			Filters:          map[string]string{"dog_size": "large", "activity": "high"},
			PetInfo: mongo_models.PetInfo{
				PetType:  "dog",
				PetBreed: "German Shepherd",
				PetAge:   3,
				PetSize:  "large",
			},
		},
		{
			ProductID:        102,
			CustomerID:       3,
			CustomerName:     "Emily Rodriguez",
			Rating:           4.0,
			Title:            "Keeps my dog entertained",
			Comment:          "This puzzle toy is great for keeping my Border Collie busy when I'm at work. It's challenging but not too difficult.",
			Images:           []string{"review3_img1.jpg", "review3_img2.jpg"},
			VerifiedPurchase: true,
			HelpfulVotes:     8,
			Filters:          map[string]string{"intelligence": "high", "energy": "high"},
			PetInfo: mongo_models.PetInfo{
				PetType:  "dog",
				PetBreed: "Border Collie",
				PetAge:   2,
				PetSize:  "medium",
			},
		},
		{
			ProductID:        103,
			CustomerID:       4,
			CustomerName:     "Lisa Park",
			Rating:           4.8,
			Title:            "Perfect for sensitive skin",
			Comment:          "My cat has very sensitive skin and this shampoo works wonderfully. No irritation and she smells great afterward.",
			Images:           []string{"review4_img1.jpg"},
			VerifiedPurchase: true,
			HelpfulVotes:     12,
			Filters:          map[string]string{"skin_type": "sensitive", "scent": "lavender"},
			PetInfo: mongo_models.PetInfo{
				PetType:  "cat",
				PetBreed: "Persian",
				PetAge:   5,
				PetSize:  "medium",
			},
		},
		{
			ProductID:        104,
			CustomerID:       5,
			CustomerName:     "David Thompson",
			Rating:           4.2,
			Title:            "Good quality fish food",
			Comment:          "My tropical fish seem to love this food. Good color enhancement and they're growing well.",
			Images:           []string{},
			VerifiedPurchase: true,
			HelpfulVotes:     6,
			Filters:          map[string]string{"fish_type": "tropical", "tank_size": "medium"},
			PetInfo: mongo_models.PetInfo{
				PetType:  "fish",
				PetBreed: "Mixed Tropical",
				PetAge:   1,
				PetSize:  "small",
			},
		},
		{
			ProductID:        105,
			CustomerID:       6,
			CustomerName:     "Amanda White",
			Rating:           5.0,
			Title:            "My parrot absolutely loves this mix",
			Comment:          "High quality seed mix with great variety. My African Grey parrot picks through it happily and his feathers look amazing.",
			Images:           []string{"review6_img1.jpg"},
			VerifiedPurchase: true,
			HelpfulVotes:     18,
			Filters:          map[string]string{"bird_type": "large", "seed_variety": "mixed"},
			PetInfo: mongo_models.PetInfo{
				PetType:  "bird",
				PetBreed: "African Grey",
				PetAge:   7,
				PetSize:  "large",
			},
		},
		{
			ProductID:        106,
			CustomerID:       7,
			CustomerName:     "Robert Kim",
			Rating:           4.7,
			Title:            "Excellent training treats",
			Comment:          "Single ingredient treats that my dog goes crazy for. Perfect size for training sessions and very healthy.",
			Images:           []string{},
			VerifiedPurchase: true,
			HelpfulVotes:     11,
			Filters:          map[string]string{"use": "training", "ingredient": "single"},
			PetInfo: mongo_models.PetInfo{
				PetType:  "dog",
				PetBreed: "Labrador",
				PetAge:   1,
				PetSize:  "large",
			},
		},
		{
			ProductID:        107,
			CustomerID:       8,
			CustomerName:     "Jennifer Adams",
			Rating:           4.9,
			Title:            "Perfect for my senior dog",
			Comment:          "This orthopedic bed has made such a difference for my 12-year-old Lab. She sleeps much more comfortably now.",
			Images:           []string{"review8_img1.jpg", "review8_img2.jpg"},
			VerifiedPurchase: true,
			HelpfulVotes:     25,
			Filters:          map[string]string{"age": "senior", "condition": "arthritis"},
			PetInfo: mongo_models.PetInfo{
				PetType:  "dog",
				PetBreed: "Labrador",
				PetAge:   12,
				PetSize:  "large",
			},
		},
		{
			ProductID:        108,
			CustomerID:       9,
			CustomerName:     "Carlos Martinez",
			Rating:           4.3,
			Title:            "Good GPS tracker with long battery",
			Comment:          "The GPS accuracy is good and the battery lasts about a week. The app could be better but overall satisfied.",
			Images:           []string{"review9_img1.jpg"},
			VerifiedPurchase: true,
			HelpfulVotes:     9,
			Filters:          map[string]string{"feature": "gps", "battery": "long-lasting"},
			PetInfo: mongo_models.PetInfo{
				PetType:  "dog",
				PetBreed: "Beagle",
				PetAge:   3,
				PetSize:  "medium",
			},
		},
		{
			ProductID:        109,
			CustomerID:       10,
			CustomerName:     "Rachel Green",
			Rating:           4.6,
			Title:            "Great eco-friendly option",
			Comment:          "Love that these bags are biodegradable. They're strong and don't leak. Feel good about being environmentally responsible.",
			Images:           []string{},
			VerifiedPurchase: true,
			HelpfulVotes:     14,
			Filters:          map[string]string{"eco": "friendly", "strength": "strong"},
			PetInfo: mongo_models.PetInfo{
				PetType:  "dog",
				PetBreed: "Mixed Breed",
				PetAge:   5,
				PetSize:  "medium",
			},
		},
	}

	for _, review := range reviews {
		if err := reviewDAO.CreateCustomerReview(ctx, review); err != nil {
			log.Printf("Error seeding review for product %d: %v", review.ProductID, err)
		}
	}

	log.Println("Customer reviews seeded successfully!")
	return nil
}
