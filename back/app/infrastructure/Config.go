package infrastructure

import "os"

type Config struct {
	DB struct {
		Main struct {
			Host     string
			Username string
			Password string
			DBName   string
		}
		Test struct {
			Host     string
			Username string
			Password string
			DBName   string
		}
	}
	Routing struct {
		Port string
	}
}

func NewConfig() *Config {

	c := new(Config)

	c.DB.Main.Host = os.Getenv("DB_HOST")
	c.DB.Main.Username = os.Getenv("DB_USERNAME")
	c.DB.Main.Password = os.Getenv("DB_PASSWORD")
	c.DB.Main.DBName = os.Getenv("DB_DBNAME")

	c.DB.Test.Host = os.Getenv("DB_TEST_HOST")
	c.DB.Test.Username = os.Getenv("DB_TEST_USERNAME")
	c.DB.Test.Password = os.Getenv("DB_TEST_PASSWORD")
	c.DB.Test.DBName = os.Getenv("DB_TEST_DBNAME")

	c.Routing.Port = ":" + os.Getenv("PORT")

	return c
}
