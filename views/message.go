package views

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"

	"github.com/AakashKath/PubSub/lib"
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
	if err := sendToAllClients(topic, rBody.Content); err != nil {
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

func sendToAllClients(topic models.Topic, content string) error {
	db := lib.DB.Connection
	fmt.Println("topic", topic, "id", topic.ID)
	var subscriptions []int64
	if err := db.Table("subscription").Where("topic_id = ?", topic.ID).Select("id").Find(&subscriptions).Error; err != nil {
		fmt.Println("Error while finding records", "table", "subscription", "error", err.Error())
		return err
	}
	var subscribers []models.Subscriber
	if err := db.Where("subscription_id in ? and is_subscribed='t'", subscriptions).Find(&subscribers).Error; err != nil {
		fmt.Println("Error while finding records", "table", "subscription", "error", err.Error())
		return err
	}
	if err := sendMessage(content); err != nil {
		fmt.Println("Error while sending message.")
		return err
	}
	return nil
}

func sendMessage(content string) error {
	conn, _ := net.Dial("tcp", "127.0.0.1:8000")
	fmt.Fprintf(conn, content+"\n")
	return nil
}
