package tests

import (
	"encoding/json"
	"errors"
	"github.com/decadevs/lunch-api/cmd/server"
	"github.com/decadevs/lunch-api/internal/adapters/api"
	"github.com/decadevs/lunch-api/internal/adapters/repository/mocks"
	"github.com/decadevs/lunch-api/internal/core/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestNotificationHandle(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDb := mocks.NewMockUserRepository(ctrl)
	mockAWS := mocks.NewMockAWSRepository(ctrl)
	mockMail := mocks.NewMockMailerRepository(ctrl)

	r := &api.HTTPHandler{
		UserService:   mockDb,
		MailerService: mockMail,
		AWSService:    mockAWS,
	}
	router := server.SetupRouter(r, mockDb)

	year, mon, day := time.Now().Date()
	month := int(mon)

	notification := models.Notification{
		Message: "serving",
		Year:    year,
		Month:   month,
		Day:     day,
	}

	notifications := []models.Notification{notification}

	bytes, _ := json.Marshal(notifications)

	t.Run("testing successful notification", func(t *testing.T) {
		mockDb.EXPECT().FindNotificationByDate(year, month, day).Return(notifications, nil)
		rw := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/user/notifications", strings.NewReader(string(bytes)))
		router.ServeHTTP(rw, req)
		assert.Equal(t, http.StatusOK, rw.Code)
		assert.Contains(t, rw.Body.String(), "notification successfully loaded")
	})

	t.Run("testing error", func(t *testing.T) {
		mockDb.EXPECT().FindNotificationByDate(year, month, day).Return(nil, errors.New("error in notification"))
		rw := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/user/notifications", strings.NewReader(string(bytes)))
		router.ServeHTTP(rw, req)
		assert.Equal(t, http.StatusInternalServerError, rw.Code)
		assert.Contains(t, rw.Body.String(), "Error in getting Notification")
	})
}
