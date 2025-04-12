package main

import (
	"context"
	"fmt"
	"grpc/internal/demoapi"
)

// DemoServiceServer ...
type DemoServiceServer struct {
	demoapi.UnimplementedDemoServiceServer
}

// NewDemoServiceServer ...
func NewDemoServiceServer() *DemoServiceServer {
	return &DemoServiceServer{}
}

// Echo ...
func (s *DemoServiceServer) Echo(ctx context.Context, request *demoapi.EchoData) (*demoapi.EchoData, error) {
	fmt.Println(request.Message)

	response := &demoapi.EchoData{
		Message: "Hello from Server!",
	}

	return response, nil
}

// Login ...
func (s *DemoServiceServer) Login(ctx context.Context, request *demoapi.LoginRequest) (*demoapi.LoginResponse, error) {
	response := &demoapi.LoginResponse{}

	if request.Username != "lendy" || request.Password != "qwerty" {
		response.Connected = false
	} else {
		response.Connected = true
	}

	return response, nil
}

// Get_user_id
func (s *DemoServiceServer) Get_user_id(ctx context.Context, request *demoapi.GetUserRequest) (*demoapi.UserResponse, error) {
	response := &demoapi.UserResponse{
		Id: reques.id,
		Username: lendy,
	}
	
	return response, nil
}