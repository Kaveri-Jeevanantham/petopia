#!/bin/bash

# Clean up unnecessary files
echo "Cleaning up unnecessary files..."

# Define color codes for output
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

# Files to keep
KEEP_FILES=(
  # Core files
  "main.go"
  "server/server.go"
  
  # Repository files
  "repository/interfaces.go"
  "repository/mongo_product_repository.go"
  "repository/mongo_brand_repository.go"
  
  # Service files
  "service/service.go"
  "service/service_impl.go"
  
  # Controller files
  "controller/product_controller_final.go"
  
  # DTO files
  "dto/dto_mapper.go"
  "dto/product_request_dto.go"
  "dto/product_response_dto.go"
  
  # Configuration files
  "config/config.go"
  
  # Database files
  "db/db.go"
  "db/mongodb.go"
  
  # Models
  "models/mongo/product_details.go"
  "models/mongo/product_brand.go"
)

# Files to remove
REMOVE_FILES=(
  "controller/product_controller.go"
  "controller/product_controller_new.go"
  "controller/product_controller_v2.go"
  "service/product_service.go"
  "service/product_service_new.go"
  "service/product_service_interface.go"
  "main_refactored.go"
  "repository/mongo_product_repository_additions.go"
)

# Function to check if file should be kept
should_keep() {
  local file=$1
  for keep_file in "${KEEP_FILES[@]}"; do
    if [[ "$file" == *"$keep_file" ]]; then
      return 0
    fi
  done
  return 1
}

# Function to check if file should be removed
should_remove() {
  local file=$1
  for remove_file in "${REMOVE_FILES[@]}"; do
    if [[ "$file" == *"$remove_file" ]]; then
      return 0
    fi
  done
  return 1
}

# Rename the controller file
echo -e "${YELLOW}Renaming controller file...${NC}"
mv -f controller/product_controller_final.go controller/product_controller.go && echo -e "${GREEN}✓ Renamed controller file${NC}" || echo -e "${RED}✗ Failed to rename controller file${NC}"

# Remove the specified files
for file in "${REMOVE_FILES[@]}"; do
  if [ -f "$file" ]; then
    rm "$file" && echo -e "${GREEN}✓ Removed $file${NC}" || echo -e "${RED}✗ Failed to remove $file${NC}"
  else
    echo -e "${YELLOW}! File $file not found${NC}"
  fi
done

echo -e "${GREEN}Cleanup complete!${NC}"
