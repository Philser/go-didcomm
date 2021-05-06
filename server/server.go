package server

import (
	"encoding/json"
	"fmt"
	"go-didcomm/db"
	"log"
	"net/http"
	"strings"
)

const PORT = "4444"
const DEFAULT_SECRET_KEY = "4inCoTsmorZdJTkmA8KEm1b41Pnhjj5VmRydDRhEHuPz1ewCb3kpq8nSd2HEHdPAGPEaekoJHeWe21Yc1ZMpu6C4" // throwaway
const DEFAULT_PUBLIC_KEY = "CKoTcMQw5jGCPBeSALWsSFrJdpddwusXhFdYUbWFea6x"                                             // throwaway

func createDidDocument(did string) string {
	return fmt.Sprintf(`
	{
		"@context": "https://w3id.org/did/v1",
        "id": "%s",
        "publicKey": [
            {
                "id": "%s#key-1",
                "type": "Secp256k1VerificationKey2018",
                "controller": "%s",
                "publicKeyBase58": "%s"
            }
        ],
        "authentication": [
            "%s#key-1"
        ],
        "created": "2011-11-11T11:11:11.111Z",
        "updated": "2011-11-11T11:11:11.111Z"
	}
	`, did, did, did, did, DEFAULT_PUBLIC_KEY)

}

func rejectDisallowedMethods(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		writer.WriteHeader(405)
	}
}

func didcommHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		writer.WriteHeader(405)
	} else {
		var message db.Message
		err := json.NewDecoder(request.Body).Decode(&message)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}

		db.LogMessage(message)

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusCreated)
		json.NewEncoder(writer).Encode(message)

		db.WriteMessage(message)
	}
}

// Mocking service endpoint that resolves a dynamically generated did doc until this application
// can resolve actual DIDs
func didResolver(writer http.ResponseWriter, request *http.Request) {
	rejectDisallowedMethods(writer, request)
	segments := strings.Split(request.URL.Path, "/")
	requestedDid := segments[len(segments)-1]
	writer.Write([]byte(createDidDocument(requestedDid)))
}

func Start() {
	fmt.Println("Starting didcomm endpoint on port " + PORT + ". Listening...")
	http.HandleFunc("/didcomm", didcommHandler)
	http.HandleFunc("/did/", didResolver)
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}
