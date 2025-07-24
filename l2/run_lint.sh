#!/bin/bash

LINT=~/go/bin/golangci-lint

if [ -z "$1" ]; then
  echo "Usage: ./run_lint.sh <directory_number>"
  echo "Example: ./run_lint.sh 1"
  exit 1
fi

DIR="./$1"

if [ ! -d "$DIR" ]; then
  echo "Error: Directory $DIR does not exist"
  exit 1
fi

if [ ! -f "$LINT" ]; then
  echo "Error: golangci-lint not found at ~/go/bin/golangci-lint"
  exit 1
fi

echo "Running golangci-lint in $DIR..."
cd $DIR
$LINT run

if [ $? -eq 0 ]; then
  echo "✅ Linting passed for $DIR"
else
  echo "❌ Linting failed for $DIR"
  exit 1
fi
