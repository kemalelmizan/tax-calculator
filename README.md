# tax-calculator
[![Coverage Status](https://img.shields.io/badge/Coverage-90.15%25-brightgreen.svg)](http://htmlpreview.github.io/?https://raw.githubusercontent.com/kemalelmizan/tax-calculator/master/docs/coverage.html)

## Running build

```
docker-compose up
```

## Docs
- [API Documentation](http://htmlpreview.github.io/?https://raw.githubusercontent.com/kemalelmizan/tax-calculator/master/docs/tax-calculator.html)
- [API Blueprint](docs/tax-calculator.apib)
- [Test Coverage](http://htmlpreview.github.io/?https://raw.githubusercontent.com/kemalelmizan/tax-calculator/master/docs/coverage.html)
- [Database Design](docs/db-schema.md)

## Notes

- API response is following [JSend](https://labs.omniti.com/labs/jsend) spec
- Stored price as `int8`, to tackle [floating points math](http://0.30000000000000004.com/) problem
- `ProductInput.Price` is still `float64`
- using [mockery](https://github.com/vektra/mockery) to generate mocks