package views

import (
	"encoding/json"
	"net/http"

	"github.com/AakashKath/PubSub/models"
)

type Response struct {
	Message string
}

type topicreqBody struct {
	TopicID string `json:"topic_id"`
}

func CreateTopic(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var rBody topicreqBody
	encoder := json.NewEncoder(w)
	if err := decoder.Decode(&rBody); err != nil {
		encoder.Encode(Response{Message: err.Error()})
		return
	}
	topic := models.Topic{Name: rBody.TopicID}
	if err := topic.Insert(); err != nil {
		encoder.Encode(Response{Message: err.Error()})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	response := Response{Message: "Topic Successfully inserted."}
	if err := encoder.Encode(&response); err != nil {
		encoder.Encode(Response{Message: err.Error()})
		return
	}
}

func DeleteTopic(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var rBody topicreqBody
	encoder := json.NewEncoder(w)
	if err := decoder.Decode(&rBody); err != nil {
		encoder.Encode(Response{Message: err.Error()})
		return
	}
	topic := models.Topic{Name: rBody.TopicID}
	if err := topic.Delete(); err != nil {
		encoder.Encode(Response{Message: err.Error()})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	response := Response{Message: "Topic Successfully deleted."}
	if err := encoder.Encode(&response); err != nil {
		encoder.Encode(Response{Message: err.Error()})
		return
	}
}
