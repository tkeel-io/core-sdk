package v1

import (
	"context"
	"fmt"
	"testing"

	proto "github.com/tkeel-io/core-sdk/proto/v1"
)

func TestEntityService_CreateEntity(t *testing.T) {
	client := NewClient("tomas", "manager", "tkeel")
	entitySrv := NewEntityService(&client)
	entitySrv.SetBaseUrl("http://localhost:6789/v1/entities")
	// create entity.
	res, err := entitySrv.CreateEntity(context.Background(), &proto.CreateEntityRequest{
		Id:     "device123",
		Type:   "DEVICE",
		Owner:  "tomas",
		Source: "CORE-SDK",
		Properties: map[string]interface{}{
			"temp": 20,
		},
	})

	fmt.Println("error: ", err)
	fmt.Println("result: ", res)
}
