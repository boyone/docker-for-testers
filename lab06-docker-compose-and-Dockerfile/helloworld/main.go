package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	port := ":8080"
	timeServer := "localhost:8000"
	dbUser := "root"
	dbPassword := "password"
	dbServer := "db:3306"

	if os.Getenv("PORT") != "" {
		port = fmt.Sprintf(":%s", os.Getenv("PORT"))
	}

	if os.Getenv("TIME_SERVER") != "" {
		timeServer = os.Getenv("TIME_SERVER")
	}

	if os.Getenv("DB_USER") != "" {
		dbUser = os.Getenv("DB_USER")
	}

	if os.Getenv("DB_PASSWORD") != "" {
		dbPassword = os.Getenv("DB_PASSWORD")
	}

	if os.Getenv("DB_SERVER") != "" {
		dbServer = os.Getenv("DB_SERVER")
	}

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/helloworld", dbUser, dbPassword, dbServer))

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		accessTime := getTime(timeServer)
		message := fmt.Sprintf("Hello, World! %s", accessTime)
		//fmt.Fprintf(w, message)
		w.Write([]byte(message))
		accessAt(db, accessTime)
	})

	log.Fatal(http.ListenAndServe(port, nil))
}

func getTime(timeServer string) string {
	resp, err := http.Get(fmt.Sprintf("http://%s/time", timeServer))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

func accessAt(db *sql.DB, mytime string) {
	datetime, _ := time.Parse(time.RFC3339, mytime)
	insert, err := db.Query(fmt.Sprintf("INSERT INTO access_log (time) VALUES ( '%s' )", datetime.Format("2006-01-02 15:04:05")))

	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
}
