package views

import (
	"encoding/json"
	"net/http"

	"github.com/AakashKath/PubSub/models"
)

type subReqBody struct {
	TopicID        string `json:"topic_id"`
	SubscriptionID string `json:"subscription_id"`
}

func AddSubscription(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var rBody subReqBody
	encoder := json.NewEncoder(w)
	if err := decoder.Decode(&rBody); err != nil {
		encoder.Encode(Response{Message: err.Error()})
		return
	}
	topic := models.Topic{Name: rBody.TopicID}
	topic.SelectByID()
	subs := models.Subscription{Name: rBody.SubscriptionID, Topic: topic}
	if err := subs.Insert(); err != nil {
		encoder.Encode(Response{Message: err.Error()})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	response := Response{Message: "Subscription Successfully inserted."}
	if err := encoder.Encode(&response); err != nil {
		encoder.Encode(Response{Message: err.Error()})
		return
	}
}

func DeleteSubscription(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var rBody subReqBody
	encoder := json.NewEncoder(w)
	if err := decoder.Decode(&rBody); err != nil {
		encoder.Encode(Response{Message: err.Error()})
		return
	}
	subs := models.Subscription{Name: rBody.SubscriptionID}
	if err := subs.Delete(); err != nil {
		encoder.Encode(Response{Message: err.Error()})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	response := Response{Message: "Subscription Successfully deleted."}
	if err := encoder.Encode(&response); err != nil {
		encoder.Encode(Response{Message: err.Error()})
		return
	}
}

type subscribeReqBody struct {
	SubscriptionID string `json:"subscription_id"`
	UserID         string `json:"user_id"`
}

func Subscribe(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var rBody subscribeReqBody
	encoder := json.NewEncoder(w)
	if err := decoder.Decode(&rBody); err != nil {
		encoder.Encode(Response{Message: err.Error()})
		return
	}
	user := models.User{Email: rBody.UserID}
	user.SelectByID()
	subs := models.Subscription{Name: rBody.SubscriptionID}
	subs.SelectByID()
	subscriber := models.Subscriber{UserSub: user, Subscription: subs}
	err := subscriber.SelectByID()
	if err != nil {
		if err.Error() != "record not found" {
			encoder.Encode(Response{Message: err.Error()})
			return
		}
		if err := subscriber.Insert(); err != nil {
			encoder.Encode(Response{Message: err.Error()})
			return
		}
		res := Response{Message: "Subscriber Successfully added."}
		encoder.Encode(&res)
		return
	}
	subscriber.SetSubscribe()
	w.Header().Set("Content-Type", "application/json")
	response := Response{Message: "You are already a subscriber."}
	if err := encoder.Encode(&response); err != nil {
		encoder.Encode(Response{Message: err.Error()})
		return
	}
}

func UnSubscribe(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var rBody subscribeReqBody
	encoder := json.NewEncoder(w)
	if err := decoder.Decode(&rBody); err != nil {
		encoder.Encode(Response{Message: err.Error()})
		return
	}
	user := models.User{Email: rBody.UserID}
	user.SelectByID()
	subs := models.Subscription{Name: rBody.SubscriptionID}
	subs.SelectByID()
	subscriber := models.Subscriber{UserSub: user, Subscription: subs}
	if err := subscriber.SelectByID(); err != nil {
		encoder.Encode(Response{Message: err.Error()})
		return
	}
	subscriber.SetUnsubscribe()
	w.Header().Set("Content-Type", "application/json")
	response := Response{Message: "Successfully unsubscribed."}
	if err := encoder.Encode(&response); err != nil {
		encoder.Encode(Response{Message: err.Error()})
		return
	}
}
