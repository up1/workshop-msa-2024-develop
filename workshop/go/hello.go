package demo

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

func GetRandomData(w http.ResponseWriter, r *http.Request) {
	roll := 1 + rand.Intn(6)

	fmt.Printf("Rolled a %d\n", roll)

	resp := strconv.Itoa(roll)
	if _, err := io.WriteString(w, resp); err != nil {
		log.Printf("Write failed: %v\n", err)
	}
}
