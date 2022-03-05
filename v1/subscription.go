package v1

import (
	"context"
	"fmt"

	proto "github.com/tkeel-io/core-sdk/proto/v1"
)

const subscriptionType = "SUBSCRIPTION"
const baseSubscriptionUrl string = "http://localhost:3500/v1.0/invoke/keel/method/apis/core/v1/subscriptions"

type SubscriptionService struct {
	baseUrl    string
	httpClient *Client
}

func NewSubscriptionService(client *Client) *SubscriptionService {
	return &SubscriptionService{
		baseUrl:    baseSubscriptionUrl,
		httpClient: client,
	}
}

// CreateEntity create entity.
func (s *EntityService) CreateSubscription(ctx context.Context, in *proto.CreateSubscriptionRequest) (*proto.SubscriptionResponse, error) {
	// construct url.
	urlText := fmt.Sprintf(s.baseUrl+"?type=%s&owner=%s&source=%s", subscriptionType, in.Owner, in.Source)
	result := s.httpClient.Post(ctx, urlText, in.Subscription)

	// parse response.
	var res proto.SubscriptionResponse
	return &res, result.Parse(&res)
}

func (s *EntityService) UpdateSubscription(ctx context.Context, in *proto.UpdateSubscriptionRequest) (*proto.SubscriptionResponse, error) {
	// construct url.
	urlText := fmt.Sprintf(s.baseUrl+"/%s?type=%s&owner=%s&source=%s", in.Id, subscriptionType, in.Owner, in.Source)
	result := s.httpClient.Put(ctx, urlText, in.Subscription)

	// parse response.
	var res proto.SubscriptionResponse
	return &res, result.Parse(&res)
}

func (s *EntityService) GetSubscription(ctx context.Context, in *proto.GetSubscriptionRequest) (*proto.SubscriptionResponse, error) {
	// construct url.
	urlText := fmt.Sprintf(s.baseUrl+"/%s?type=%s&owner=%s&source=%s", in.Id, subscriptionType, in.Owner, in.Source)
	result := s.httpClient.Get(ctx, urlText, nil)

	// parse response.
	var res proto.SubscriptionResponse
	return &res, result.Parse(&res)
}

func (s *EntityService) ListSubscription(ctx context.Context, in *proto.ListSubscriptionRequest) (*proto.ListSubscriptionResponse, error) {
	// construct url.
	urlText := fmt.Sprintf(s.baseUrl+"?type=%s&owner=%s&source=%s", subscriptionType, in.Owner, in.Source)
	result := s.httpClient.Get(ctx, urlText, nil)

	// parse response.
	var res proto.ListSubscriptionResponse
	return &res, result.Parse(&res)
}

func (s *EntityService) DeleteSubscription(ctx context.Context, in *proto.DeleteSubscriptionRequest) (*proto.DeleteSubscriptionResponse, error) {
	// construct url.
	urlText := fmt.Sprintf(s.baseUrl+"/%s?type=%s&owner=%s&source=%s", in.Id, subscriptionType, in.Owner, in.Source)
	result := s.httpClient.Delete(ctx, urlText, nil)

	// parse response.
	var res proto.DeleteSubscriptionResponse
	return &res, result.Parse(&res)
}

func (s *SubscriptionService) SetBaseUrl(baseUrl string) {
	s.baseUrl = baseSubscriptionUrl
}
