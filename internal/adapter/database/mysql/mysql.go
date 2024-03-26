package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func NewMySQLConnection() (*sql.DB, error) {
	// viper.SetConfigFile(".env")
	// viper.SetConfigType("env")
	// viper.AutomaticEnv()

	// // devo adicionar um fallback aqui
	// if err := viper.ReadInConfig(); err != nil {
	// 	log.Fatalf("Error reading .env file: %s", err)
	// }
	// user := viper.GetString("DB_USER")
	// password := viper.GetString("DB_PASSWORD")
	// host := viper.GetString("DB_HOST")
	// port := viper.GetString("DB_PORT")
	// databaseName := viper.GetString("DB_NAME")
	// connString := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + databaseName
	connString := "root:secret@tcp(localhost:3306)/goapi"

	db, err := sql.Open("mysql", connString)
	if err != nil {
		return nil, err
	}
	return db, nil
}
