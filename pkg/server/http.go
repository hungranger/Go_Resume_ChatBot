package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/hungranger/Go_Resume_ChatBot/pkg/config"
)

type HTTPServer struct {
	config config.IConfig
}

func ProvideServer(config config.IConfig) HTTPServer {
	return HTTPServer{
		config: config,
	}
}

func (s HTTPServer) Run() error {
	httpServer := &http.Server{
		Addr:    ":" + s.config.GetPort(),
		Handler: s.Handler(),
	}

	return httpServer.ListenAndServe()
}

func (s HTTPServer) Handler() http.Handler {
	mux := mux.NewRouter()

	mux.HandleFunc("/", s.home).Methods("GET")
	mux.HandleFunc("/webhook", s.messages).Methods("POST")
	// mux.HandleFunc("/webhook", s.verify).Methods("GET")

	return mux
}

func (s HTTPServer) home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello from home")
}

func (s HTTPServer) messages(w http.ResponseWriter, r *http.Request) {
	var callback Callback
	json.NewDecoder(r.Body).Decode(&callback)
	if callback.Object == "page" {
		for _, entry := range callback.Entry {
			for _, event := range entry.Messaging {
				processMessage(s, event)
			}
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Got your message"))
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Message not supported"))
	}
}

func processMessage(s HTTPServer, event Messaging) {
	client := &http.Client{}
	response := Response{
		Recipient: User{
			ID: event.Sender.ID,
		},
		Message: Message{
			Attachment: &Attachment{
				Type: "image",
				Payload: Payload{
					URL: s.config.GetImage(),
				},
			},
		},
	}

	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(&response)

	url := fmt.Sprintf(s.config.GetFbApi(), os.Getenv("PAGE_ACCESS_TOKEN"))
	req, err := http.NewRequest("POST", url, body)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
}
