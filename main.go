package main

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
	"os"
)

var client *redis.Client

func main() {
	fmt.Println("BABYMETAL TOP")

	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		redisAddr = "localhost:6379"
	}

	client = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "",
		DB:       0,
	})

	http.HandleFunc("/", handleRequest)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	count, err := client.Incr(r.Context(), "counter").Result()

	if err != nil {
		http.Error(w, "Error fetching counter", http.StatusInternalServerError)
		return
	}

	var htmlText string = "<body>\n" +
		"<h1 style=\"color: red; align: center; \" >BABYMETAL TOP</h1>\n" +
		"<script>alert('BABYMETAL TOP')</script>" +
		"</body>"

	fmt.Fprintf(w, htmlText+"\nThis page has been viewed %d times!", count)
}
