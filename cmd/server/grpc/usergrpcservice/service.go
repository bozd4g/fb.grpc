package usergrpcservice

import (
	"context"
	"github.com/bozd4g/fb.grpc/cmd/server/internal/application/userservice"
	pb "github.com/bozd4g/fb.grpc/pkg/proto/userservice"
	"github.com/mitchellh/mapstructure"
	"net/http"
)

func (service UserGrpcService) Create(ctx context.Context, request *pb.CreateRequest) (*pb.CreateReply, error) {
	var requestDto userservice.UserCreateRequestDto
	err := mapstructure.Decode(request, &requestDto)
	if err != nil {
		return nil, err
	}

	if err = service.UserService.Create(requestDto); err != nil {
		return nil, err
	}

	return &pb.CreateReply{StatusCode: http.StatusCreated}, nil
}

func (service UserGrpcService) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllReply, error) {
	userDtos, err := service.UserService.GetAll()
	if err != nil {
		return &pb.GetAllReply{}, err
	}

	dtos := make([]*pb.UserDtoReply, 0)
	for _, v := range userDtos {
		dtos = append(dtos, &pb.UserDtoReply{
			Id:      v.Id,
			Name:    v.Name,
			Surname: v.Surname,
			Email:   v.Email,
		})
	}

	return &pb.GetAllReply{StatusCode: http.StatusOK, Users: dtos}, nil
}
