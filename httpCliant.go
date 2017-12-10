package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"log"
	"io/ioutil"
)

type Movie struct {
	Id int
	Name string
	Director string
	Rating float32
}

func main() {
	// json-server,db.jsonの定義に応じて変える
	url := "http://localhost:3000/movies"

	cliant := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	
	res, err := cliant.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
	
	var movies []Movie
	err = json.Unmarshal(body, &movies)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i< len(movies); i++ {
		fmt.Println(movies[i])
	}
}