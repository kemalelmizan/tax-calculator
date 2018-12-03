# tax-calculator

## Running build

```
docker-compose up
```

## API Documentation
- (Documentation)[http://htmlpreview.github.io/?https://raw.githubusercontent.com/kemalelmizan/tax-calculator/master/docs/tax-calculator.html]
- (API Blueprint)[docs/tax-calculator.apib]

## Notes

- Stored price as `int8`, to avoid [floating points math](http://0.30000000000000004.com/) problem
- API response is following [JSend](https://labs.omniti.com/labs/jsend) spec