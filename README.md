# taxcalculator

## Running build

```
docker-compose up
```

## Notes

- Stored price as `int8`, to avoid [floating points math](http://0.30000000000000004.com/) problem
- API response is following [JSend](https://labs.omniti.com/labs/jsend) spec