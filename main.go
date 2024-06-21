package main

import (
	"lgwt/di"
	"lgwt/mocking"
	"net/http"
	"os"
	"time"
)

func MyGreetHandler(w http.ResponseWriter, r *http.Request) {
	di.Greet(w, "world")
}

func main() {
	// log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreetHandler)))
	sleeper := &mocking.ConfigurableSleeper{Duration: 1 * time.Second, SleepFn: time.Sleep}
	mocking.Countdown(os.Stdout, sleeper)
}
