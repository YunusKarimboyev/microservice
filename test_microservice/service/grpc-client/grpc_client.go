package grpcclient

import (
  "fmt"

  config "github.com/double/test_microservice/config"
  cu "github.com/double/test_microservice/genproto/post"
  "google.golang.org/grpc"
  "google.golang.org/grpc/credentials/insecure"
)

type Clients interface {
  Post() cu.PostServiceClient
}

type ServiceManager struct {
  Config      config.Config
  postService cu.PostServiceClient
}

func New(cfg config.Config) (*ServiceManager, error) {
  connPost, err := grpc.Dial(
    fmt.Sprintf("%s:%s", cfg.PostServiceHost, cfg.PostServicePort),
    grpc.WithTransportCredentials(insecure.NewCredentials()))
  if err != nil {
    return nil, fmt.Errorf("user service dial host:%s port: %s", cfg.PostServiceHost, cfg.PostServicePort)
  }

  return &ServiceManager{
    Config:      cfg,
    postService: cu.NewPostServiceClient(connPost),
  }, nil

}

func (s *ServiceManager) User() cu.PostServiceClient {
  return s.postService
}