#!/bin/bash
# Create mocks directories if they don't exist
mkdir -p services/mocks
mkdir -p controllers/mocks

# Generate mocks
mockgen -source=services/email_service.go > services/mocks/email_service.go
