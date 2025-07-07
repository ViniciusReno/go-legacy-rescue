package services

import (
    "context"
    "fmt"
    "myapp/factories"
    "myapp/handlers"
)

type UserService interface {
    ProcessUser(ctx context.Context, userID string) error
}

type userServiceImpl struct {
    handlerFactory factories.HandlerFactory
}

func NewUserService(factory factories.HandlerFactory) UserService {
    return &userServiceImpl{handlerFactory: factory}
}

func (s *userServiceImpl) ProcessUser(ctx context.Context, userID string) error {
    handler := s.handlerFactory.GetHandler("user")
    result := handler.Handle(ctx, userID)
    fmt.Println("User processed:", result)
    return nil
}
