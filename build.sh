#!/bin/bash

# Check if the build folder exists
if [ -d "build" ]; then
  # Remove the build folder and its contents
  rm -rf "build"
  echo "The build folder was removed."
fi

echo "Create the build folder."

# Build the server
echo "Building server..."
BUILD_CMD="go build -o ./build/server/ ./src/server"
$BUILD_CMD
cp "./.env" "build/server"

# Build the client
echo "Building client..."
PROJECT_PATH="./src/client/TimeManagement/TimeManagement.csproj"

# Restore dependencies
echo "Restoring dependencies..."
dotnet restore $PROJECT_PATH

# Build the project
echo "Building project..."
dotnet build $PROJECT_PATH -c Release

echo "Build completed successfully."

read -p "Press any key to continue..." -n1 -s