# Tanker-Go

Tanker is an API developed for being a resource for load testing [restQL](http://restql.b2w.io). The raison d'être of this project is creating a fast API that returns a simple JSON object.

## Endpoints

Each endpoint in `Tanker-Go` corresponds to a different type of respose:

- `/product` - returns a JSON object with a name and an id
- `/multiProduct` - returns a list of products as JSON objects
- `/offer` - returns a JSON object with a price and an id
- `/installment` - returns a JSON object with a number of installments
- `/resource-status` - a simple health check which returns 204 status code

## Changing the response time

Query parameters can be used to set the repsonse time of each request.

- `responseTime` - sets a custom response time (in milliseconds)
- `minTime` - sets a minimum response time (in milliseconds)
- `maxTime` - sets a maximum response time (in milliseconds)

**Note:** the `responseTime` parameter has precedence above `minTime` and `maxTime`

If there are no query parameters, the response time will be random time between 100ms and 3s