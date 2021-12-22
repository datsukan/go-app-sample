package main

import (
	"sample/app/infrastructure"
)

func main() {
	infrastructure.LoadEnv()
	db := infrastructure.NewDB()
	r := infrastructure.NewRouting(db)
	r.Run()
}
