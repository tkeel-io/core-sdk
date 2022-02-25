package v1

import (
	"context"
	"fmt"

	proto "github.com/tkeel-io/core-sdk/proto/v1"
)

// CreateEntity create entity.
func (s *EntityService) AppendMapper(ctx context.Context, in *proto.AppendMapperRequest) (*proto.AppendMapperResponse, error) {
	// construct url.
	urlText := fmt.Sprintf(s.baseUrl+"/%s/mappers?type=%s&owner=%s&source=%s", in.EntityId, in.Type, in.Owner, in.Source)
	result := s.httpClient.Post(ctx, urlText, in.Mapper)

	// parse response.
	var res proto.AppendMapperResponse
	return &res, result.Parse(&res)
}

func (s *EntityService) GetMapper(ctx context.Context, in *proto.GetMapperRequest) (*proto.GetMapperResponse, error) {
	// construct url.
	urlText := fmt.Sprintf(s.baseUrl+"/%s/mappers/%s?type=%s&owner=%s&source=%s", in.EntityId, in.Id, in.Type, in.Owner, in.Source)
	result := s.httpClient.Put(ctx, urlText, nil)

	// parse response.
	var res proto.GetMapperResponse
	return &res, result.Parse(&res)
}

func (s *EntityService) ListMapper(ctx context.Context, in *proto.ListMapperRequest) (*proto.ListMapperResponse, error) {
	// construct url.
	urlText := fmt.Sprintf(s.baseUrl+"/%s/mappers?type=%s&owner=%s&source=%s", in.EntityId, in.Type, in.Owner, in.Source)
	result := s.httpClient.Get(ctx, urlText, nil)

	// parse response.
	var res proto.ListMapperResponse
	return &res, result.Parse(&res)
}

func (s *EntityService) RemoveMapper(ctx context.Context, in *proto.RemoveMapperRequest) (*proto.RemoveMapperResponse, error) {
	// construct url.
	urlText := fmt.Sprintf(s.baseUrl+"/%s/mappers/%s?type=%s&owner=%s&source=%s", in.EntityId, in.Id, in.Type, in.Owner, in.Source)
	result := s.httpClient.Delete(ctx, urlText, nil)

	// parse response.
	var res proto.RemoveMapperResponse
	return &res, result.Parse(&res)
}
