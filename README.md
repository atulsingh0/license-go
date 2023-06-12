# license-go

# To run locally

```
go run /src/api/main.go
```

## To generate a license

```
curl --location 'localhost:9200/generate' \
--header 'Content-Type: application/json' \
--data '{
    "version": "v1",
    "customer": "atul",
    "valid-from": "2021-01-01T00:00:00Z",
    "expiry-date": "2022-01-01T00:00:00Z",
    "hard-expiry-date": "2022-01-01T00:00:00Z",
    "seats": 1,
    "hard-seats": 1,
    "type": "test"
}'
```
