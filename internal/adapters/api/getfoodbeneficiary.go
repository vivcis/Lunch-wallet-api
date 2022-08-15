package api

import (
	"github.com/decadevs/lunch-api/internal/adapters/repository"
	"github.com/decadevs/lunch-api/internal/core/helpers"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (u *HTTPHandler) GetAllBeneficiaryHandle(c *gin.Context) {

	_, err := u.GetAdminFromContext(c)
	if err != nil {
		helpers.JSON(c, "internal server error", 500, nil, []string{"This is an internal server error"})
		return
	}
	pagination := repository.GeneratePaginationFromRequest(c)
	userLists, err := u.UserService.FindAllFoodBeneficiary(&pagination)

	if err != nil {
		helpers.JSON(c, "internal server error", 500, nil, []string{"An internal server error"})
		return
	}
	helpers.JSON(c, "food beneficiaries found successfully", http.StatusOK, userLists, nil)

}

func (u *HTTPHandler) GetMealTimetableHandle(c *gin.Context) {
	_, err := u.GetAdminFromContext(c)
	if err != nil {
		helpers.JSON(c, "internal server error", 500, nil, []string{"internal server error"})
		return
	}

	year, month, day := time.Now().Date()

	food, err := u.UserService.FindFoodByDate(year, month, day)
	if err != nil {
		helpers.JSON(c, "internal server error", 500, nil, []string{"internal server error"})
		return
	}

	helpers.JSON(c, "Dinner found", 200, food, nil)

}

func (u *HTTPHandler) GetTickets(c *gin.Context) {

}

func (u *HTTPHandler) AdminSearchFoodBeneficiaries(c *gin.Context) {
	_, err := u.GetAdminFromContext(c)
	if err != nil {
		helpers.JSON(c, "An internal server error", 500, nil, []string{"internal server error"})
		return
	}
	pagination := repository.GeneratePaginationFromRequest(c)
	query := c.Param("text")

	beneficiaries, err := u.UserService.SearchFoodBeneficiary(query, &pagination)
	if err != nil {
		helpers.JSON(c, "An internal server error", 500, nil, []string{"internal server error"})
		return
	}
	if len(beneficiaries) == 0 {
		helpers.JSON(c, "Record Not Found", 404, nil, []string{"Record Not Found"})
		return
	}
	helpers.JSON(c, "information gotten", http.StatusOK, beneficiaries, nil)
}

func (u *HTTPHandler) AdminGetTotalNumberOfUsers(c *gin.Context) {
	_, err := u.GetAdminFromContext(c)
	if err != nil {
		helpers.JSON(c, "An internal server error", 500, nil, []string{"internal server error"})
		return
	}

	totalNumber, err := u.UserService.GetTotalUsers()
	if err != nil {
		helpers.JSON(c, "internal server error", 500, nil, []string{"Unable to get total number of users"})
		return
	}
	helpers.JSON(c, "Total number of users", http.StatusOK, totalNumber, nil)

}
