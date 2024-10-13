## Orderbook Matching with Cross-Margin
A simple orderbook matching engine with cross-margin support web server developed in GoLang.
Using Tarantool as the database for the orderbook.

### Start
1. copy the `.env.sample` to `.env` and fill in the values
2. run `docker-compose up --build`
3. run `go run main.go`

### Endpoints
#### User
POST `/user` - Create a user

| Parameter | Type   | Description    |
|-----------|--------|----------------|
| id        | string | Unique User ID |
| amount    | float  | Initial amount |

GET `/user/:id` - Get user details

#### Order
POST `/order` - Create an order

| Parameter | Type   | Description       |
|-----------|--------|-------------------|
| userId    | string | Unique User ID    |
| market    | string | Market BTC/ETH    |
| side      | string | Side BUY/SELL     |
| size      | float  | Size of the order |

GET `/order/:id` - Get order details

### TODO
- [ ] Setup Mq for new order event
- [ ] Replace goroutines to Mq for order matching
- [ ] Optimising speed (order matching algorithm, database caching)
- [ ] Add test (unit test, integration test)
- [ ] Optimising the module structure