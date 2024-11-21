# orderbook core backend

## Migration

https://github.com/golang-migrate/migrate/tree/master/cmd/migrate

### crate migrate script
```shell
  migrate create -ext sql -dir ./internal/adapter/database/postgres/migration -seq `nameOfMigration`
```

### run migrate script
```shell
  go run cmd/migration/main.go up 
```

```shell
  go run cmd/migration/main.go down 
```

## Start app

```shell
 go run cmd/server/main.go
```