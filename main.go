package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
      "practice/sqll"
	"github.com/julienschmidt/httprouter"
)

func main() {
 
	var err error
	db, err = sql.Open("mysql", "root:mayank99@tcp(127.0.0.1:3306)/recordings")
	if err != nil {
		panic(err)
	}
	fmt.Print("Successfully Created")
	defer db.Close()

	albID, err := AddAlbum(Album{
		Title:  "The Modern Sound of Betty Carter",
		Artist: "Betty Carter",
		Price:  49.99,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID of added album: %v\n", albID)

	router := httprouter.New()
	router.GET("/", Index)
	log.Fatal(http.ListenAndServe(":8080", router))

}
