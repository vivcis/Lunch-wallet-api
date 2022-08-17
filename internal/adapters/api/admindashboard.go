package api

import (
	"github.com/decadevs/lunch-api/internal/adapters/repository"
	"github.com/decadevs/lunch-api/internal/core/helpers"
	"github.com/gin-gonic/gin"
	"log"
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

	food, err := u.UserService.FindFoodByDate(year, int(month), day)
	if err != nil {
		helpers.JSON(c, "internal server error", 500, nil, []string{"internal server error"})
		return
	}

	helpers.JSON(c, "Timetable found", 200, food, nil)

}

func (u *HTTPHandler) GetScannedUsersByDate(c *gin.Context) {
	_, err := u.GetAdminFromContext(c)
	if err != nil {
		helpers.JSON(c, "internal server error", 500, nil, []string{"internal server error"})
		return
	}
	date := time.Now().Format("2006-01-02")
	log.Println(date)
	helpers.JSON(c, "Successful", 200, nil, nil)
}

func (u *HTTPHandler) GetGraphData(c *gin.Context) {
	_, err := u.GetAdminFromContext(c)
	if err != nil {
		helpers.JSON(c, "internal server error", 500, nil, []string{"internal server error"})
		return
	}
	graphData, err := u.UserService.FindActiveUsersByMonth()
	if err != nil {
		helpers.JSON(c, "internal server error", 500, nil, []string{"internal server error"})
		return
	}
	helpers.JSON(c, "Successful", 200, graphData, nil)
}

func (u *HTTPHandler) GetNumberOfScannedUsers(c *gin.Context) {
	_, err := u.GetAdminFromContext(c)
	if err != nil {
		helpers.JSON(c, "internal server error", 500, nil, []string{"internal server error"})
		return
	}

	date := time.Now().Format("2006-01-02")

	scanned, err := u.UserService.FindNumbersOfScannedUsers(date)
	if err != nil {
		helpers.JSON(c, "internal server error", 500, nil, []string{"internal server error"})
		return
	}
	total, err := u.UserService.GetTotalUsers()
	if err != nil {
		helpers.JSON(c, "internal server error", 500, nil, []string{"internal server error"})
		return
	}
	res := struct {
		Scanned    int64
		NotScanned int64
	}{}

	res.Scanned = scanned
	res.NotScanned = int64(total) - scanned
	helpers.JSON(c, "Successful", 200, res, nil)
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
