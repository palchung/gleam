#!/bin/bash

SQL_User_name=postgres
SQL_Password=password
SQL_Host=127.0.0.1
SQL_Port=5432
SQL_Database=thefreepress
SQL_SSL_mode=disable

Migration_file_path=db/migration
PostgreSQL_address=postgres://$SQL_User_name:$SQL_Password@$SQL_Host:$SQL_Port/$SQL_Database?sslmode=$SQL_SSL_mode

case "$1" in
    up)     migrate -path $Migration_file_path -database $PostgreSQL_address -verbose up;;
    down)   migrate -path $Migration_file_path -database $PostgreSQL_address -verbose down;;
    *)      echo "Invalid command";;
esac