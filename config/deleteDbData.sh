#!/bin/bash
DB_PATH="./dbData"

for dir in $(ls -d ${DB_PATH}/*/); do
  /bin/chown -R 1000:1000 "${dir}data/"
  /bin/rm -rf "${dir}data/*"
done