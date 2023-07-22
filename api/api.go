package api

import (
	"github.com/Grady-Saccullo/activity-pub-go/pkg/w3c/activity_streams/v2/impl"
	"log"
	"net/http"
)

var t activity_streams_v2_impl.Activity

func main() {
	http.HandleFunc("/.well-known/webfinger", webFingerHandler)
	http.HandleFunc("/v1/inbox", inboxHandler)
	http.HandleFunc("/v1/outbox", outboxHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func webFingerHandler(w http.ResponseWriter, r *http.Request) {
}

func inboxHandler(w http.ResponseWriter, r *http.Request) {

}

func outboxHandler(w http.ResponseWriter, r *http.Request) {

}
