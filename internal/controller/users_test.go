package controller

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	mock_contracts "github.com/advor2102/socialnetwork/internal/contracts/mocks"
	"github.com/advor2102/socialnetwork/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHandler_CreateUser(t *testing.T) {
	type mockBehaviour func(s *mock_contracts.MockServiceI, u models.User)

	testTable := []struct {
		name                 string
		inputBody            string
		inputUser            models.User
		mockBehaviour        mockBehaviour
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "OK",
			inputBody: `{
				"user_name": "sasha",
				"email":"sasha@gmail.com",
				"age": 35
			}`,
			inputUser: models.User{
				UserName: "sasha",
				Email:    "sasha@gmail.com",
				Age:      35,
			},
			mockBehaviour: func(s *mock_contracts.MockServiceI, u models.User) {
				s.EXPECT().CreateUser(u).Return(nil)
			},
			expectedStatusCode:   http.StatusCreated,
			expectedResponseBody: `{"message":"User created successfully"}`,
		},
		{
			name:                 "Empty fields",
			inputBody:            `{}`,
			mockBehaviour:        func(s *mock_contracts.MockServiceI, u models.User) {},
			expectedStatusCode:   http.StatusUnprocessableEntity,
			expectedResponseBody: `{"error":"invalid field value"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			svc := mock_contracts.NewMockServiceI(ctrl)
			testCase.mockBehaviour(svc, testCase.inputUser)

			handler := NewController(svc)

			r := gin.New()
			r.POST("/users", handler.CreateUser)

			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponseBody, w.Body.String())
		})
	}
}
