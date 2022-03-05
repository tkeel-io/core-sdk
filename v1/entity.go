package v1

import (
	"context"
	"fmt"
	"strings"

	proto "github.com/tkeel-io/core-sdk/proto/v1"
)

const baseEntityUrl string = "http://localhost:3500/v1.0/invoke/keel/method/apis/core/v1/entities"

type EntityService struct {
	baseUrl    string
	httpClient *Client
}

func NewEntityService(client *Client) *EntityService {
	return &EntityService{
		baseUrl:    baseEntityUrl,
		httpClient: client,
	}
}

// CreateEntity create entity.
func (s *EntityService) CreateEntity(ctx context.Context, in *proto.CreateEntityRequest) (*proto.EntityResponse, error) {
	// construct url.
	urlText := fmt.Sprintf(s.baseUrl+"?type=%s&owner=%s&source=%s", in.Type, in.Owner, in.Source)
	result := s.httpClient.Post(ctx, urlText, in.Properties)

	// parse response.
	var res proto.EntityResponse
	return &res, result.Parse(&res)
}

func (s *EntityService) UpdateEntity(ctx context.Context, in *proto.UpdateEntityRequest) (*proto.EntityResponse, error) {
	// construct url.
	urlText := fmt.Sprintf(s.baseUrl+"/%s?type=%s&owner=%s&source=%s", in.Id, in.Type, in.Owner, in.Source)
	result := s.httpClient.Put(ctx, urlText, in.Properties)

	// parse response.
	var res proto.EntityResponse
	return &res, result.Parse(&res)
}

func (s *EntityService) GetEntity(ctx context.Context, in *proto.GetEntityRequest) (*proto.EntityResponse, error) {
	// construct url.
	urlText := fmt.Sprintf(s.baseUrl+"/%s?type=%s&owner=%s&source=%s", in.Id, in.Type, in.Owner, in.Source)
	result := s.httpClient.Get(ctx, urlText, nil)

	// parse response.
	var res proto.EntityResponse
	return &res, result.Parse(&res)
}

func (s *EntityService) ListEntity(ctx context.Context, in *proto.ListEntityRequest) (*proto.ListEntityResponse, error) {
	// construct url.
	urlText := fmt.Sprintf(s.baseUrl+"?type=%s&owner=%s&source=%s", in.Type, in.Owner, in.Source)
	result := s.httpClient.Get(ctx, urlText, nil)

	// parse response.
	var res proto.ListEntityResponse
	return &res, result.Parse(&res)
}

func (s *EntityService) DeleteEntity(ctx context.Context, in *proto.DeleteEntityRequest) (*proto.DeleteEntityResponse, error) {
	// construct url.
	urlText := fmt.Sprintf(s.baseUrl+"/%s?type=%s&owner=%s&source=%s", in.Id, in.Type, in.Owner, in.Source)
	result := s.httpClient.Delete(ctx, urlText, nil)

	// parse response.
	var res proto.DeleteEntityResponse
	return &res, result.Parse(&res)
}

func (s *EntityService) UpdateEntityProps(ctx context.Context, in *proto.UpdateEntityPropsRequest) (*proto.EntityResponse, error) {
	// construct url.
	urlText := fmt.Sprintf(s.baseUrl+"/%s/properties?type=%s&owner=%s&source=%s", in.Id, in.Type, in.Owner, in.Source)
	result := s.httpClient.Put(ctx, urlText, in.Properties)

	// parse response.
	var res proto.EntityResponse
	return &res, result.Parse(&res)
}

func (s *EntityService) PatchEntityProps(ctx context.Context, in *proto.PatchEntityPropsRequest) (*proto.EntityResponse, error) {
	// construct url.
	urlText := fmt.Sprintf(s.baseUrl+"/%s/properties/patch?type=%s&owner=%s&source=%s", in.Id, in.Type, in.Owner, in.Source)
	result := s.httpClient.Put(ctx, urlText, in.Properties)

	// parse response.
	var res proto.EntityResponse
	return &res, result.Parse(&res)
}

func (s *EntityService) GetEntityProps(ctx context.Context, in *proto.GetEntityPropsRequest) (*proto.EntityResponse, error) {
	// construct url.
	urlText := fmt.Sprintf(s.baseUrl+"/%s/properties?type=%s&owner=%s&source=%s&property_keys=%s",
		in.Id, in.Type, in.Owner, in.Source, strings.Join(in.PropertyKeys, ","))
	result := s.httpClient.Get(ctx, urlText, nil)

	// parse response.
	var res proto.EntityResponse
	return &res, result.Parse(&res)
}

func (s *EntityService) RemoveEntityProps(ctx context.Context, in *proto.RemoveEntityPropsRequest) (*proto.EntityResponse, error) {
	// construct url.
	urlText := fmt.Sprintf(s.baseUrl+"/%s/properties?type=%s&owner=%s&source=%s&property_keys=%s",
		in.Id, in.Type, in.Owner, in.Source, strings.Join(in.PropertyKeys, ","))
	result := s.httpClient.Delete(ctx, urlText, nil)

	// parse response.
	var res proto.EntityResponse
	return &res, result.Parse(&res)
}

func (s *EntityService) UpdateEntityConfigs(ctx context.Context, in *proto.UpdateEntityConfigsRequest) (*proto.EntityResponse, error) {
	// construct url.
	urlText := fmt.Sprintf(s.baseUrl+"/%s/configs?type=%s&owner=%s&source=%s", in.Id, in.Type, in.Owner, in.Source)
	result := s.httpClient.Put(ctx, urlText, in.Configs)

	// parse response.
	var res proto.EntityResponse
	return &res, result.Parse(&res)
}

func (s *EntityService) PatchEntityConfigs(ctx context.Context, in *proto.PatchEntityConfigsRequest) (*proto.EntityResponse, error) {
	// construct url.
	urlText := fmt.Sprintf(s.baseUrl+"/%s/configs/patch?type=%s&owner=%s&source=%s", in.Id, in.Type, in.Owner, in.Source)
	result := s.httpClient.Put(ctx, urlText, in.Configs)

	// parse response.
	var res proto.EntityResponse
	return &res, result.Parse(&res)
}

func (s *EntityService) GetEntityConfigs(ctx context.Context, in *proto.GetEntityConfigsRequest) (*proto.EntityResponse, error) {
	// construct url.
	urlText := fmt.Sprintf(s.baseUrl+"/%s/configs?type=%s&owner=%s&source=%s&property_keys=%s",
		in.Id, in.Type, in.Owner, in.Source, strings.Join(in.PropertyKeys, ","))
	result := s.httpClient.Get(ctx, urlText, nil)

	// parse response.
	var res proto.EntityResponse
	return &res, result.Parse(&res)
}

func (s *EntityService) RemoveEntityConfigs(ctx context.Context, in *proto.RemoveEntityConfigsRequest) (*proto.EntityResponse, error) {
	// construct url.
	urlText := fmt.Sprintf(s.baseUrl+"/%s/configs?type=%s&owner=%s&source=%s&property_keys=%s",
		in.Id, in.Type, in.Owner, in.Source, strings.Join(in.PropertyKeys, ","))
	result := s.httpClient.Delete(ctx, urlText, nil)

	// parse response.
	var res proto.EntityResponse
	return &res, result.Parse(&res)
}

func (s *EntityService) SetBaseUrl(baseUrl string) {
	s.baseUrl = baseUrl
}
