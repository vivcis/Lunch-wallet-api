package tests

//
//import (
//	"github.com/decadevs/lunch-api/internal/adapters/repository/mocks"
//	"github.com/decadevs/lunch-api/internal/core/middleware"
//	"github.com/gin-gonic/gin"
//	"github.com/golang/mock/gomock"
//	"github.com/stretchr/testify/assert"
//	"testing"
//)
//
//func TestAuthorizeKitchenStaff(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	mockDb := mocks.NewMockUserRepository(ctrl)
//
//	var c gin.Context
//
//	secret := "jakhjb67273hsbv"
//	accToken := middleware.GetTokenFromHeader(&c)
//	accessToken, _, _ := middleware.AuthorizeToken(&accToken, &secret)
//
//	kitchenStaff, _ := mockDb.FindKitchenStaffByEmail("okoasuquo@yahoo.com")
//	// set the user and token as context parameters.
//	c.Set("user", kitchenStaff)
//	c.Set("access_token", accessToken.Raw)
//
//	got := middleware.AuthorizeKitchenStaff(mockDb.FindKitchenStaffByEmail, mockDb.TokenInBlacklist)
//	want := c.Handler()
//
//	assert.Equal(t, want, got)
//
//}
