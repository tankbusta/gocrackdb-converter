# GoCrack DB Converter

This convert's [GoCrack's](https://github.com/mandiant/gocrack) database from it's flatfile BoltDB/Storm backend to Postgres. You probably don't need this (yet)

## Running

    go run cmd\convert\convert.go -bolt-db="<path to gocracks .db file> -pg-conn-str="<destination postgres connection string>"
