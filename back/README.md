# Portforio backend by Go

## Migrate

### Create

```sh
migrate create -ext sql -dir migrations -seq create_xxxx_table
```

### Up

```sh
migrate -database 'mysql://root:password@tcp(sample-mariadb:3306)/sample' -path migrations up
```

### Down

```sh
migrate -database 'mysql://root:password@tcp(sample-mariadb:3306)/sample' -path migrations down
```
