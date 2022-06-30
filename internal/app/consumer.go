package app

import (
	"fmt"
	"net/http"
)

func StartConsumer() {
	http.HandleFunc("/v1/consume", consumeData)
	fmt.Println("server started at localhost:8080")
	if err := http.ListenAndServe(":9090", nil); err != nil {
		panic(err)
	}
}

func consumeData(w http.ResponseWriter, r *http.Request) {

}
