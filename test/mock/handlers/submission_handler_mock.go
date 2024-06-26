// Code generated by MockGen. DO NOT EDIT.
// Source: ./handler/submission_handler.go
//
// Generated by this command:
//
//	mockgen -source=./handler/submission_handler.go -destination=./test/mock/handlers/submission_handler_mock.go
//

// Package mock_handler is a generated GoMock package.
package mock_handler

import (
	reflect "reflect"

	gin "github.com/gin-gonic/gin"
	gomock "go.uber.org/mock/gomock"
)

// MockISubmissionHandler is a mock of ISubmissionHandler interface.
type MockISubmissionHandler struct {
	ctrl     *gomock.Controller
	recorder *MockISubmissionHandlerMockRecorder
}

// MockISubmissionHandlerMockRecorder is the mock recorder for MockISubmissionHandler.
type MockISubmissionHandlerMockRecorder struct {
	mock *MockISubmissionHandler
}

// NewMockISubmissionHandler creates a new mock instance.
func NewMockISubmissionHandler(ctrl *gomock.Controller) *MockISubmissionHandler {
	mock := &MockISubmissionHandler{ctrl: ctrl}
	mock.recorder = &MockISubmissionHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockISubmissionHandler) EXPECT() *MockISubmissionHandlerMockRecorder {
	return m.recorder
}

// CreateSubmission mocks base method.
func (m *MockISubmissionHandler) CreateSubmission(c *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CreateSubmission", c)
}

// CreateSubmission indicates an expected call of CreateSubmission.
func (mr *MockISubmissionHandlerMockRecorder) CreateSubmission(c any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSubmission", reflect.TypeOf((*MockISubmissionHandler)(nil).CreateSubmission), c)
}

// DeleteSubmission mocks base method.
func (m *MockISubmissionHandler) DeleteSubmission(c *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteSubmission", c)
}

// DeleteSubmission indicates an expected call of DeleteSubmission.
func (mr *MockISubmissionHandlerMockRecorder) DeleteSubmission(c any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSubmission", reflect.TypeOf((*MockISubmissionHandler)(nil).DeleteSubmission), c)
}

// GetAllSubmissions mocks base method.
func (m *MockISubmissionHandler) GetAllSubmissions(c *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetAllSubmissions", c)
}

// GetAllSubmissions indicates an expected call of GetAllSubmissions.
func (mr *MockISubmissionHandlerMockRecorder) GetAllSubmissions(c any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllSubmissions", reflect.TypeOf((*MockISubmissionHandler)(nil).GetAllSubmissions), c)
}

// GetSubmission mocks base method.
func (m *MockISubmissionHandler) GetSubmission(c *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetSubmission", c)
}

// GetSubmission indicates an expected call of GetSubmission.
func (mr *MockISubmissionHandlerMockRecorder) GetSubmission(c any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubmission", reflect.TypeOf((*MockISubmissionHandler)(nil).GetSubmission), c)
}
