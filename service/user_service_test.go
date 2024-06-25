package service_test

import (
	"assignment_1/entity"
	"assignment_1/service"
	mock_service "assignment_1/test/mock/services"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRepo := mock_service.NewMockIUserRepository(ctrl)
	userService := service.NewUserService(mockRepo)

	ctx := context.Background()

	user := &entity.User{
		Name:      "pepeg",
		Email:     "pepeg@handsome.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	t.Run("successfully create user", func(t *testing.T) {
		mockRepo.EXPECT().CreateUser(ctx, user).Return(*user, nil)

		createdUser, err := userService.CreateUser(ctx, user)
		assert.NoError(t, err)
		assert.Equal(t, *user, createdUser)
	})
}
