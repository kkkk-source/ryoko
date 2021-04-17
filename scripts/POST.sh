#!/bin/bash
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{ "name":"xyz" }' \
  http://localhost:8080/heroes
