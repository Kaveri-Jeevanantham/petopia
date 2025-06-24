#!/bin/bash

# Step 1: Build the UI Code
echo "Building the UI code..."
npm run build

# Step 2: Deploy the Built Code
echo "Deploying the built code..."
# Assuming the build output is in the 'dist' directory
# Replace '/path/to/nginx/html' with the actual path where Nginx serves the files
cp -r dist/* /path/to/nginx/html/

# Step 3: Reload Nginx
echo "Reloading Nginx..."
sudo nginx -s reload

echo "Deployment completed successfully."
