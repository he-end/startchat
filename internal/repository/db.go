package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/joho/godotenv"
)

var DB *sql.DB

type DBConf struct {
	DBUsername string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
}

func loadDBConfig() (conf DBConf) {
	path := "../configs/.dbconfig"
	if err := godotenv.Load(path); err != nil {
		log.Fatal("config not found")
	}
	conf.DBUsername = os.Getenv("DB_USER")
	conf.DBPassword = os.Getenv("DB_PASSWORD")
	conf.DBHost = os.Getenv("HOST")
	conf.DBPort = os.Getenv("PORT")
	conf.DBName = os.Getenv("DB_NAME")
	return
}

func (conf DBConf) loadDSN() string {
	// "postgres://postgres:yourpass@localhost:5432/startchat"
	dsn := fmt.Sprintf(
		"postgres://%v:%v@%v:%v/%v",
		conf.DBUsername, conf.DBPassword, conf.DBHost, conf.DBPort, conf.DBName)
	return dsn
}

func init() {
	conf := loadDBConfig()
	dsn := conf.loadDSN()
	fmt.Println(dsn)
	// sql.Register("pgx", stdlib.GetDefaultDriver())
	var err error
	DB, err = sql.Open("pgx", dsn)
	if err != nil {
		fmt.Println(err)
	}
	if err := DB.Ping(); err != nil {
		fmt.Println(err)
	}
}
