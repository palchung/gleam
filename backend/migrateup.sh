#!/bin/bash

migrate -path db/migration -database postgres://postgres:password@127.0.0.1:5432/thefreepress?sslmode=disable -verbose up