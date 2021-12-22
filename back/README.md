# Portforio backend by Go

## Migrate

### Create

```sh
migrate create -ext sql -dir migrations -seq create_xxxx_table
```

### Up

```sh
migrate -database 'mysql://root:password@tcp(portfolio-mariadb:3306)/portfolio' -path migrations up
```

### Down

```sh
migrate -database 'mysql://root:password@tcp(portfolio-mariadb:3306)/portfolio' -path migrations down
```
