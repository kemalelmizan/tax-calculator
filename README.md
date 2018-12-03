# tax-calculator

## Running build

```
docker-compose up
```

## API Documentation
- [Documentation](http://htmlpreview.github.io/?https://raw.githubusercontent.com/kemalelmizan/tax-calculator/master/docs/tax-calculator.html)
- [API Blueprint](docs/tax-calculator.apib)

### cURL Example

```
curl -X POST \
  http://localhost:3000/bill \
  -H 'Content-Type: application/json' \
  -d '{
    "data": [
        {
            "name": "Lucky Strike",
            "tax_code": 2,
            "price": 110000
        },
        {
            "name": "Big Mac",
            "tax_code": 1,
            "price": 102000
        },
        {
            "name": "Movie",
            "tax_code": 3,
            "price": 15050
        }
    ]
}'
```

## Notes

- API response is following [JSend](https://labs.omniti.com/labs/jsend) spec
- Stored price as `int8`, to tackle [floating points math](http://0.30000000000000004.com/) problem
- `ProductInput.Price` is still `float64`