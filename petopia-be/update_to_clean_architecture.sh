#!/bin/bash

# Script to update the main.go file to use clean architecture

# Set color codes
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

echo -e "${YELLOW}Updating main.go to use clean architecture...${NC}"

# Create backup of original main.go
cp main.go main.go.bak && echo -e "${GREEN}✓ Created backup of main.go${NC}" || echo -e "${RED}✗ Failed to create backup of main.go${NC}"

# Replace the server.Start line with server.StartV2 in main.go
sed -i '' 's/server.Start(cfg, database)/server.StartV2(cfg, database)/g' main.go && echo -e "${GREEN}✓ Updated main.go to use clean architecture${NC}" || echo -e "${RED}✗ Failed to update main.go${NC}"

# Rename files to use the clean versions
echo -e "${YELLOW}Renaming files to use clean architecture versions...${NC}"

# Rename routes_new.go to routes.go
cp routes/routes_new.go routes/routes.go && echo -e "${GREEN}✓ Updated routes.go${NC}" || echo -e "${RED}✗ Failed to update routes.go${NC}"

# Rename server_new.go to server.go
cp server/server_new.go server/server.go && echo -e "${GREEN}✓ Updated server.go${NC}" || echo -e "${RED}✗ Failed to update server.go${NC}"

# Note: product_controller.go is already the correct file

echo -e "${GREEN}Update complete! You can now run the application using clean architecture.${NC}"
echo -e "${YELLOW}If you need to revert changes, restore main.go from main.go.bak${NC}"
