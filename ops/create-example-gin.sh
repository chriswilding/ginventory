#!/usr/bin/env sh

aws dynamodb put-item \
    --table-name ginventory \
    --item '{
    "PK": { "S": "GIN#MONKEY_47" },
    "SK": { "S": "GIN#MONKEY_47" },
    "Name": { "S": "Tanqueray" },
    "CountryOfOrigin": { "S": "Germany" }
}'
