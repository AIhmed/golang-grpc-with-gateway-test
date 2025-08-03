package server

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"grpc-test/proto/testdata"
)

type Server struct {
	testdata.UnimplementedTestDataServiceServer
}

func (s *Server) GetTestData(ctx context.Context, req *testdata.GetTestDataRequest) (*testdata.TestDataResponse, error) {
	// Create a list of persons
	persons := []*testdata.Person{
		{
			Name:       "John Doe",
			Age:        30,
			Profession: "Software Engineer",
		},
		{
			Name:       "Jane Smith",
			Age:        28,
			Profession: "Data Scientist",
		},
		{
			Name:       "Bob Johnson",
			Age:        45,
			Profession: "Product Manager",
		},
	}

	return &testdata.TestDataResponse{Persons: persons}, nil
}

func StartGRPCServer() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	testdata.RegisterTestDataServiceServer(s, &Server{})

	log.Println("gRPC server started on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
