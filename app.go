package main

import (
	"fmt"
	"log"
	"net/http"
    "gopkg.in/redis.v3"
)

func loggerMiddleware(nextHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Request recieved: %v\n", r)
		nextHandler.ServeHTTP(w, r)
		fmt.Println("Request handled successfully")
	})
}

func indexPageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("indexPageHandler called")
	client := redis.NewClient(&redis.Options{
		Addr:     "db:6379",
		Password: "",
		DB:       0,
	})

	err := client.Set("key", "foo", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("key").Result()
	if err != nil {
		panic(err)
	} else {
        fmt.Fprintln(w, "value is: %q", val)
    }
  }

func main() {
	fmt.Println("It's running and what not.")

	mux := http.NewServeMux()


	mux.Handle("/", loggerMiddleware(http.HandlerFunc(indexPageHandler)))
	mux.Handle("/about", loggerMiddleware(http.HandlerFunc(indexPageHandler)))

    log.Fatal(http.ListenAndServe(":5000", mux))
}
