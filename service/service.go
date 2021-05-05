package service

import "net/http"

func didcommHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		writer.WriteHeader(405)
	}
}

func Start() {
	http.HandleFunc("/didcomm", didcommHandler)
}
