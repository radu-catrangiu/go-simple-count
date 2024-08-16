# Simple Counter Server

This is a simple counter server that writes the current count to a file.

## Read the current count
```sh
curl -X GET localhost:8080/count
``` 

## Increment the count
```sh
curl -X POST localhost:8080/count
```
