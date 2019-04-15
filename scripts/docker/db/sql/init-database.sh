#!/usr/bin/env bash
cd docker-entrypoint-initdb.d
rm -f out
for file in `\find . -type f -name '*.sql' | sort -n`; do
    cat $file >> out
done
mysql -u docker -pdocker mysql_db < "out"