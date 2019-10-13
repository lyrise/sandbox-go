package service

import (
	"context"
	"errors"

	"github.com/lyrise/sandbox-go/2019/09/25/grpc/gens/api"
)

type AnimalService struct {
}

func (c *AnimalService) GetName(ctx context.Context, in *api.GetNameArgument) (*api.GetNameResponse, error) {
	switch in.Id {
	case 1:
		return &api.GetNameResponse{
			Name: "Cat",
		}, nil
	case 2:
		return &api.GetNameResponse{
			Name: "Dog",
		}, nil
	}
	return nil, errors.New("Not Found")
}
