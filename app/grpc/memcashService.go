package grpc

import (
	"context"
	"time"

	v1 "github.com/speix/memcash/app/pb/v1"
)

func (s *Server) SaveDataToCache(ctx context.Context, req *v1.SetDataRequest) (*v1.Response, error) {

	err := s.MemcashService.Set(req.Key, req.Value, time.Duration(req.Expiration)*time.Second)
	if err != nil {
		return &v1.Response{Status: &v1.Status{Error: true, Message: err.Error()}}, nil
	}

	return &v1.Response{Status: &v1.Status{Error: false, Message: "Value cached"}}, nil
}

func (s *Server) GetDataFromCache(ctx context.Context, req *v1.GetDataRequest) (*v1.DataResponse, error) {

	item, err := s.MemcashService.Get(req.Key)
	if err != nil {
		return &v1.DataResponse{Status: &v1.Status{Error: true, Message: err.Error()}}, nil
	}

	return &v1.DataResponse{
		Status: &v1.Status{
			Error:   false,
			Message: "Item received",
		},
		Key:        req.Key,
		Value:      item.Value,
		Expiration: item.Expiration,
	}, nil
}
