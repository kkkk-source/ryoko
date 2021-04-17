#!/bin/bash
curl --header "Content-Type: application/json" \
  --request PUT \
  --data '{ "id":9, "name":"xyz"}' \
  http://localhost:8080/heroes
