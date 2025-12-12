#!/bin/bash

curl -X POST "http://localhost:8080/complex" \
	-H "Content-Type: application/json" \
	-d '{"data": {"key": "apple", "value": 123}}'
