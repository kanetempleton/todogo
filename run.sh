#!/bin/bash

# Compile the application
go build

# Run the tests
go test

# Check the test exit status
if [ $? -eq 0 ]; then
    # If tests passed, run the application
    ./todogo
else
    echo "Tests failed. Not running the application."
fi

