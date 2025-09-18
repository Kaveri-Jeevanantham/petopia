#!/bin/bash

# MongoDB Connection Test Script

echo "Testing MongoDB Connection and API Endpoints"
echo "============================================"

# Step 1: Check if the service is running
echo "Step 1: Checking if the service is running..."
curl -s -o /dev/null -w "%{http_code}" http://localhost:8080/api/health
if [ $? -eq 0 ]; then
  echo "✅ Service is running"
else
  echo "❌ Service is not running. Please start the service first."
  exit 1
fi

# Step 2: Test creating a product
echo -e "\nStep 2: Testing product creation..."
PRODUCT_RESPONSE=$(curl -s -X POST -H "Content-Type: application/json" -d '{
  "product_name": "Test Dog Collar",
  "description": "A comfortable collar for dogs",
  "brand_id": 1,
  "brand_name": "PetNutrition Plus",
  "seller_id": 101,
  "category": "dog-accessories",
  "price": 19.99,
  "discount": 0,
  "availability": true,
  "item_dimensions": {
    "weight": "0.2kg",
    "materials": ["Leather", "Metal"]
  }
}' http://localhost:8080/api/products)

echo "$PRODUCT_RESPONSE" | grep -q "product_name"
if [ $? -eq 0 ]; then
  echo "✅ Product created successfully"
  # Extract the product ID for future tests
  PRODUCT_ID=$(echo "$PRODUCT_RESPONSE" | grep -o '"id":"[^"]*"' | sed 's/"id":"//;s/"//')
  echo "   Product ID: $PRODUCT_ID"
else
  echo "❌ Failed to create product"
  echo "$PRODUCT_RESPONSE"
  exit 1
fi

# Step 3: Test listing products
echo -e "\nStep 3: Testing product listing..."
LIST_RESPONSE=$(curl -s http://localhost:8080/api/products)

echo "$LIST_RESPONSE" | grep -q "product_name"
if [ $? -eq 0 ]; then
  echo "✅ Products listed successfully"
else
  echo "❌ Failed to list products"
  echo "$LIST_RESPONSE"
  exit 1
fi

# Step 4: Test updating a product
if [ ! -z "$PRODUCT_ID" ]; then
  echo -e "\nStep 4: Testing product update..."
  UPDATE_RESPONSE=$(curl -s -X PUT -H "Content-Type: application/json" -d '{
    "product_name": "Updated Dog Collar",
    "description": "An improved comfortable collar for dogs",
    "brand_id": 1,
    "brand_name": "PetNutrition Plus",
    "seller_id": 101,
    "category": "dog-accessories",
    "price": 24.99,
    "discount": 2.00,
    "availability": true,
    "item_dimensions": {
      "weight": "0.2kg",
      "materials": ["Premium Leather", "Stainless Steel"]
    }
  }' http://localhost:8080/api/products/$PRODUCT_ID)

  echo "$UPDATE_RESPONSE" | grep -q "Updated Dog Collar"
  if [ $? -eq 0 ]; then
    echo "✅ Product updated successfully"
  else
    echo "❌ Failed to update product"
    echo "$UPDATE_RESPONSE"
    exit 1
  fi

  # Step 5: Test deleting a product
  echo -e "\nStep 5: Testing product deletion..."
  DELETE_RESPONSE=$(curl -s -X DELETE -w "%{http_code}" -o /dev/null http://localhost:8080/api/products/$PRODUCT_ID)

  if [ "$DELETE_RESPONSE" -eq 204 ]; then
    echo "✅ Product deleted successfully"
  else
    echo "❌ Failed to delete product (HTTP status: $DELETE_RESPONSE)"
    exit 1
  fi
else
  echo -e "\nSkipping update and delete tests as product ID is not available"
fi

echo -e "\n✅ All tests completed successfully!"
echo "MongoDB integration with repository pattern is working as expected."
