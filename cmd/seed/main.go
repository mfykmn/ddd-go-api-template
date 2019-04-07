package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/mafuyuk/ddd-go-api-template/config"
	"github.com/mafuyuk/ddd-go-api-template/infrastructure/db"
)

func main() {
	var dbConf = &config.DB{}
	flag.StringVar(&dbConf.Host, "h", "0.0.0.0", "DB host.")
	flag.StringVar(&dbConf.Port, "P", "3306", "DB port.")
	flag.StringVar(&dbConf.User, "u", "user", "DB user.")
	flag.StringVar(&dbConf.Password, "p", "pass", "DB password.")
	flag.StringVar(&dbConf.DBName, "n", "demo", "DB name.")

	// DBコネクションの初期化
	dbClient, err := db.NewMySQL(dbConf)
	if err != nil {
		os.Exit(1)
	}
	defer dbClient.Close()

	// Seedディレクトリの取得
	seedsDir, err := filepath.Abs("./_seeds")
	if err != nil {
		os.Exit(2)
	}

	seed := Seed{
		db:      dbClient.DB,
		dirPath: seedsDir,
	}
	fmt.Println("Start seeding")
	if err := seed.Execute(); err != nil {
		fmt.Printf("Failed seeding. err: %#v\n", err)
		os.Exit(3)
	}
	fmt.Println("Success seeding!!")
}
