package handler_test

import (
	"assignment_1/entity"
	"assignment_1/handler"
	mock_service "assignment_1/test/mock/services"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func setupTest(t *testing.T) (context.Context, *gomock.Controller, *mock_service.MockIUserService, handler.IUserHandler) {
	ctrl := gomock.NewController(t)
	mockService := mock_service.NewMockIUserService(ctrl)
	userHandler := handler.NewUserHandler(mockService)
	ctx := context.Background()

	return ctx, ctrl, mockService, userHandler
}

func TestUserHandler_CreateUser(t *testing.T) {
	_, ctrl, mockService, userHandler := setupTest(t)
	defer ctrl.Finish()

	r := gin.Default()
	gin.SetMode(gin.TestMode)
	r.POST("/users", userHandler.CreateUser)

	t.Run("success create user", func(t *testing.T) {
		userReq := &entity.User{
			Name:  "pepeg",
			Email: "pepeg@handsome.com",
		}

		mockService.EXPECT().CreateUser(gomock.Any(), userReq).Return(*userReq, nil)

		reqBody, _ := json.Marshal(map[string]string{
			"name":  "pepeg",
			"email": "pepeg@handsome.com",
		})

		req, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
		assert.Contains(t, w.Body.String(), fmt.Sprintf("user id %v created successfully", userReq.ID))
	})

	t.Run("gagal create user", func(t *testing.T) {
		mockService.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(entity.User{}, errors.New("error creating user"))

		reqBody, _ := json.Marshal(map[string]string{
			"name":  "pepeg",
			"email": "pepeg@handsome.com",
		})

		req, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "error creating user")
	})

	t.Run("error creating user, karena invalid json", func(t *testing.T) {
		reqBody := []byte(`{"name": "pepeg", email: "invalid-email"}`)
		req, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(reqBody))

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "error creating user")
	})

	t.Run("error validasi, nama dan email tidak boleh kosong", func(t *testing.T) {
		reqBody := []byte(`{}`)
		req, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(reqBody))

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "Name Can not be empty")
	})

	t.Run("error validasi, nama dan email harus valid", func(t *testing.T) {
		reqBody := []byte(`{"name": "pepeg", "email": "invalid-email"}`)
		req, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(reqBody))

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "Email Must be a valid email address")
	})
}

func TestUserHandler_GetUser(t *testing.T) {
	_, ctrl, mockService, userHandler := setupTest(t)
	defer ctrl.Finish()

	r := gin.Default()
	gin.SetMode(gin.TestMode)
	r.GET("/users/1", userHandler.GetUser)

	t.Run("successfully get user", func(t *testing.T) {
		user := entity.User{
			ID:   1,
			Name: "pepeg",
		}
		mockService.EXPECT().GetUserByID(gomock.Any(), gomock.Any()).Return(user, nil)

		req, _ := http.NewRequest(http.MethodGet, "/users/1", nil)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), user.Name)
	})

	t.Run("error get user", func(t *testing.T) {
		mockService.EXPECT().GetUserByID(gomock.Any(), gomock.Any()).Return(entity.User{}, errors.New("error get user"))

		req, _ := http.NewRequest(http.MethodGet, "/users/1", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "error get user")
	})
}

func TestUserHandler_UpdateUser(t *testing.T) {
	_, ctrl, mockService, userHandler := setupTest(t)
	defer ctrl.Finish()

	r := gin.Default()
	gin.SetMode(gin.TestMode)
	r.PUT("/users/1", userHandler.UpdateUser)

	updateUser := entity.User{
		Name:  "pepeg_update",
		Email: "pepeg_update@handsome.com",
	}

	t.Run("successfully update user", func(t *testing.T) {
		mockService.EXPECT().UpdateUser(gomock.Any(), gomock.Any(), updateUser).Return(updateUser, nil)
		reqJson, _ := json.Marshal(map[string]string{
			"name":  "pepeg_update",
			"email": "pepeg_update@handsome.com",
		})
		req, _ := http.NewRequest(http.MethodPut, "/users/1", bytes.NewBuffer(reqJson))

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "updated successfully")
	})

	t.Run("error update user", func(t *testing.T) {
		mockService.EXPECT().UpdateUser(gomock.Any(), gomock.Any(), updateUser).Return(entity.User{}, errors.New("error updating user"))
		reqJson, _ := json.Marshal(map[string]string{
			"name":  "pepeg_update",
			"email": "pepeg_update@handsome.com",
		})
		req, _ := http.NewRequest(http.MethodPut, "/users/1", bytes.NewBuffer(reqJson))

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "error updating user")
	})

	t.Run("error validasi, nama dan email kosong", func(t *testing.T) {
		reqBody := []byte(`{}`)
		req, _ := http.NewRequest(http.MethodPut, "/users/1", bytes.NewBuffer(reqBody))

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "Name Can not be empty")
	})

	t.Run("error validasi, nama dan email harus valid", func(t *testing.T) {
		reqBody := []byte(`{"name": "pepeg", "email": "invalid-email"}`)
		req, _ := http.NewRequest(http.MethodPut, "/users/1", bytes.NewBuffer(reqBody))

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "Email Must be a valid email address")
	})

	t.Run("gagal update user, invalid json request", func(t *testing.T) {
		reqBody := []byte(`{"name": "pepeg", email: "invalid-request"}`)
		req, _ := http.NewRequest(http.MethodPut, "/users/1", bytes.NewBuffer(reqBody))

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "error updating user")
	})
}
