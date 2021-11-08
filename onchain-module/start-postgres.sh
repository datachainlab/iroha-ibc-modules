#!/bin/bash

docker run --name some-postgres \
	--rm \
	-e POSTGRES_USER=postgres \
	-e POSTGRES_PASSWORD=mysecretpassword \
	-p 5432:5432 \
	-d postgres:9.5 \
	-c 'max_prepared_transactions=100'
