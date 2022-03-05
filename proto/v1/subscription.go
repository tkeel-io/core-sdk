package v1

type Subscription struct {
	Mode       string `json:"mode,omitempty"`
	Source     string `json:"source,omitempty"`
	Filter     string `json:"filter,omitempty"`
	Target     string `json:"target,omitempty"`
	Topic      string `json:"topic,omitempty"`
	PubsubName string `json:"pubsub_name,omitempty"`
}

type CreateSubscriptionRequest struct {
	Id           string        `json:"id,omitempty"`
	Owner        string        `json:"owner,omitempty"`
	Source       string        `json:"source,omitempty"`
	Subscription *Subscription `json:"subscription,omitempty"`
}

type UpdateSubscriptionRequest struct {
	Id           string        `json:"id,omitempty"`
	Owner        string        `json:"owner,omitempty"`
	Source       string        `json:"source,omitempty"`
	Subscription *Subscription `json:"subscription,omitempty"`
}

type GetSubscriptionRequest struct {
	Id     string `json:"id,omitempty"`
	Owner  string `json:"owner,omitempty"`
	Source string `json:"source,omitempty"`
}

type DeleteSubscriptionRequest struct {
	Id     string `json:"id,omitempty"`
	Owner  string `json:"owner,omitempty"`
	Source string `json:"source,omitempty"`
}

type ListSubscriptionRequest struct {
	Source string `json:"source,omitempty"`
	Owner  string `json:"owner,omitempty"`
}

type DeleteSubscriptionResponse struct {
	Id     string `json:"id,omitempty"`
	Status string `json:"status,omitempty"`
}

type SubscriptionResponse struct {
	Id           string        `json:"id,omitempty"`
	Owner        string        `json:"owner,omitempty"`
	Source       string        `json:"source,omitempty"`
	Subscription *Subscription `json:"subscription,omitempty"`
}

type ListSubscriptionResponse struct {
	Count int32                   `json:"count,omitempty"`
	Items []*SubscriptionResponse `json:"items,omitempty"`
}
