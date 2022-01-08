package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/nexusmed/nexusmed-go/consult"
)

func main() {
	// Create the client
	clt := consult.NewClient(http.DefaultClient)
	// Set the API key
	clt.SetApiKey(os.Getenv("NEXUSMED_SECRET_KEY"))

	// Create a new asynchronous consultation
	newConsult, err := clt.CreateConsultation(consult.CreateConsultationInput{
		Type: consult.ConsultationInputTypeAsynchronous,
		QuestionnaireAnswers: &consult.QuestionnaireAnswersInput{
			ID: "qans_23GTIkQARYHGfJgyD622wCq5LG9",
		},
		Products: []*consult.ProductInput{
			{ID: "prd_23Fvs4I8Y4MJJDnG2n8n2F2Fhg4"},
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	b, _ := json.Marshal(newConsult)
	fmt.Println(string(b))
}
