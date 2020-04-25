package main

import (
	"fmt"
	"log" // Import the redigo/redis package.
	"net/http"

	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/mux"
)

var (
	c     redis.Conn
	err   error
	reply interface{}
)

func init() {
	c, err = redis.Dial("tcp", "redis:6379")
	if err != nil {
		log.Fatal(err)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "try : /add_word/word=foo, /autocomplete/query=fo")
}

func addWord(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	_, err := c.Do("ZADD", "word", 1, vars["keyword"])
	if err != nil {
		fmt.Fprintf(w, "Fail")
		log.Fatal("Unable to add to redis :", err)
	}
	fmt.Fprintf(w, "Success")
	log.Printf("Added %s\n", vars["keyword"])
}

func queryWord(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rank, err := redis.Int(c.Do("ZRANK", "word", vars["keyword"]))
	if err != nil {
		fmt.Fprintf(w, "Faild")
		log.Fatal("Unable to query ZRANK redis : ", err)
	}
	results, err := redis.Strings(c.Do("ZRANGE", "word", rank, -1))
	if err != nil {
		fmt.Fprintf(w, "Failed")
		log.Fatal("Unable to query ZRANGE redis : ", err)
	}
	log.Printf("result returned")
	fmt.Fprintf(w, "%s", results)
}

func main() {
	defer c.Close()
	c.Do("FLUSHALL")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", home)
	router.HandleFunc("/add_word/word={keyword}", addWord)
	router.HandleFunc("/autocomplete/query={keyword}", queryWord)
	log.Fatal(http.ListenAndServe(":8080", router))
}
