#!/bin/bash

# Length of the secret key
KEY_LENGTH=32

# Generate a complex JWT secret key using openssl
JWT_SECRET_KEY=$(openssl rand -base64 $KEY_LENGTH)

# Optionally, you can save the generated key to an environment variable file
ENV_FILE=".env"
echo "JWT_SECRET_KEY=$JWT_SECRET_KEY" > $ENV_FILE

# Output the key to the terminal
echo "Generated JWT Secret Key: $JWT_SECRET_KEY"
echo "Key saved to $ENV_FILE"
