package views

import (
	"encoding/json"
	"net/http"

	"github.com/AakashKath/PubSub/models"
)

type publishReqBody struct {
	TopicID string `json:"topic_id"`
	Content string `json:"content"`
}

func Publish(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var rBody publishReqBody
	encoder := json.NewEncoder(w)
	if err := decoder.Decode(&rBody); err != nil {
		encoder.Encode(Response{Message: err.Error()})
		return
	}
	topic := models.Topic{Name: rBody.TopicID}
	topic.SelectByID()
	message := models.Message{Topic: topic, Content: rBody.Content}
	if err := message.Insert(); err != nil {
		encoder.Encode(Response{Message: err.Error()})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	response := Response{Message: "Message Successfully published."}
	if err := encoder.Encode(&response); err != nil {
		encoder.Encode(Response{Message: err.Error()})
		return
	}
}
