package analyze

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cdipaolo/sentiment"
)

// Input is the http post input
type Input struct {
	Sentence string `json:"sentence"`
}

// Analyze analyze the sentiment of a sentence
func Analyze(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" && r.URL.Path != "Analyze" {
		http.NotFound(w, r)
		return
	}

	var input Input
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	analysis := sentimentModel.SentimentAnalysis(input.Sentence, sentiment.English)
	fmt.Fprint(w, analysis.Score)
}
