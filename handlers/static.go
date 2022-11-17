package handlers

import (
	"fmt"
	"net/http"
	"os"
)

func Home(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		hostname = "0x000"
	}
	w.Write([]byte(fmt.Sprintf("Hello World from %s", hostname)))
}
