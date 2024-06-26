package handler_test

import (
	"assignment_1/entity"
	"assignment_1/handler"
	mock_service "assignment_1/test/mock/services"
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func setupTestSubs(t *testing.T) (context.Context, *gomock.Controller, *mock_service.MockISubmissionService, handler.ISubmissionHandler) {
	ctrl := gomock.NewController(t)
	mockService := mock_service.NewMockISubmissionService(ctrl)
	subsHandler := handler.NewSubmissionHandler(mockService)
	ctx := context.Background()

	return ctx, ctrl, mockService, subsHandler
}

func TestSubHandler_CreateSubmission(t *testing.T) {
	_, ctrl, mockService, subsHandler := setupTestSubs(t)
	defer ctrl.Finish()

	r := gin.Default()
	gin.SetMode(gin.TestMode)
	r.POST("/submissions", subsHandler.CreateSubmission)

	t.Run("succesfully create submission", func(t *testing.T) {
		answerJson, _ := json.Marshal([]entity.Answer{
			{
				QuestionID: 2,
				Answer:     "≥ 10 tahun",
			},
		})
		sub := entity.Submission{
			UserID:  1,
			Answers: answerJson,
		}

		mockService.EXPECT().CreateSubmission(gomock.Any(), sub).Return(nil)

		reqBody, _ := json.Marshal(map[string]string{
			"user_id": "1",
			"answers": `[
			{
				"question_id": 2,
				"answer": "≥ 10 tahun"
			}`,
		})

		req, _ := http.NewRequest(http.MethodPost, "/submissions", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "created successfully")
	})
}
