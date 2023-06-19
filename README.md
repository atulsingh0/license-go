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
    "customer": "atul",
    "valid-from": "2021-01-01T00:00:00Z",
    "expiry-date": "2022-01-01T00:00:00Z",
    "hard-expiry-date": "2022-01-01T00:00:00Z",
    "seats": 1,
    "hard-seats": 1,
    "type": "test"
}'
```

## To generate a license

```
curl --location 'localhost:9200/validate' \
--header 'Content-Type: application/json' \
--data '{
    "id":"25912eab-aa97-4fc5-85c1-6782e7f7ed8b",
    "customer":"atul",
    "valid-from":"2021-01-01T00:00:00Z",
    "expiry-date":"2022-01-01T00:00:00Z",
    "hard-expiry-date":"2022-01-01T00:00:00Z",
    "seats":1,
    "hard-seats":1,
    "type":"test",
    "signature":"T4OWqj51P2bT80Tnyiu47nMbkruLvKeBJ+3hQT8kgnG5kBeMyytdjQTczP59ik5qvMXlWoWQq8BaR//4wwy7Bg=="
}'
```
