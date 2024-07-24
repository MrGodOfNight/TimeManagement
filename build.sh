#!/bin/bash
# Build the server
echo "Building server..."
BUILD_CMD="go build -o ./build/server/ ./src/server"
$BUILD_CMD

# Build the client
echo "Building client..."

echo "Building updater..."
# Path to updater project
PROJECT_PATH="./src/client/Updater/Updater.csproj"

# Restore dependencies
echo "Restoring dependencies..."
dotnet restore $PROJECT_PATH

# Build the project
echo "Building project..."
dotnet build $PROJECT_PATH --configuration Release

echo "Building time management..."
# Path to time management project
PROJECT_PATH="./src/client/TimeManagement/TimeManagement.csproj"

# Restore dependencies
echo "Restoring dependencies..."
dotnet restore $PROJECT_PATH

# Build the project
echo "Building project..."
dotnet build $PROJECT_PATH --configuration Release

echo "Build completed successfully."

# Run the server
# RUN_CMD="./build/server/server.exe"
# $RUN_CMD