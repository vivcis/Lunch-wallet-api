package api

import (
	"fmt"
	"github.com/decadevs/lunch-api/internal/core/helpers"
	"github.com/decadevs/lunch-api/internal/core/middleware"
	"github.com/decadevs/lunch-api/internal/core/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// CreateMeal godoc
// @Summary      Admin creates meal
// @Description  creates meal by collecting fields in models.Food in a form data. Note: "images" is a file to be uploaded in jpeg or png. "name" is the name of the meal, "type" is either brunch or dinner, "weekday" can be ignored but it is either monday - sunday, "kitchen" is either uno, edo-tech park, etc. "year", "month" and "day" are numbers. It is an authorized route to only ADMIN
// @Tags         Food
// @Accept       json
// @Produce      json
// @Param Food in form data body models.Food true "images, type, name, kitchen, year, month, day"
// @Success      200  {string} string "Successfully Created"
// @Failure      500  {string}  string "internal server error"
// @Failure      400  {string}  string "bad request"
// @Router       /admin/createtimetable [post]
func (u *HTTPHandler) CreateFoodTimetableHandle(c *gin.Context) {
	admin, err := u.GetAdminFromContext(c)
	if err != nil {
		helpers.JSON(c, "internal server error", 500, nil, []string{"internal server error"})
		return
	}
	var food models.Food

	form, err := c.MultipartForm()

	if err != nil {
		log.Printf("error parsing multipart form: %v", err)
		helpers.JSON(c, "error parsing multipart form", 400, nil, []string{"bad request"})
		return
	}

	formImages := form.File["images"]
	var images []models.Image
	log.Println(formImages)
	log.Println(images)

	// upload the images to aws.
	for _, f := range formImages {
		file, err := f.Open()
		if err != nil {

		}
		fileExtension, ok := middleware.CheckSupportedFile(strings.ToLower(f.Filename))
		log.Printf(filepath.Ext(strings.ToLower(f.Filename)))
		fmt.Println(fileExtension)
		if ok {
			log.Println(fileExtension)
			helpers.JSON(c, "Bad Request", 400, nil, []string{fileExtension + " image file type is not supported"})
			return
		}

		session, tempFileName, err := middleware.PreAWS(fileExtension, "product")
		if err != nil {
			log.Println("could not upload file", err)
		}

		url, err := u.AWSService.UploadFileToS3(session, file, tempFileName, f.Size)
		if err != nil {
			helpers.JSON(c, "internal server error", http.StatusInternalServerError, nil, []string{"an error occurred while uploading the image"})
			return
		}

		img := models.Image{
			Url: url,
		}
		images = append(images, img)
	}

	mealType := c.PostForm("type")
	foodType := strings.ToUpper(mealType)

	foodName := c.PostForm("name")

	weekDay := c.PostForm("weekday")

	kitchen := c.PostForm("kitchen")

	year, err := strconv.Atoi(c.PostForm("year"))
	if err != nil {
		log.Println(err)

		helpers.JSON(c, "bad request", http.StatusBadRequest, nil, []string{"an error occur in converting year"})
		return

	}

	month, err := strconv.Atoi(c.PostForm("month"))
	if err != nil {
		log.Println(err)

		helpers.JSON(c, "bad request", http.StatusBadRequest, nil, []string{"an error occur in converting month"})
		return
	}

	date, err := strconv.Atoi(c.PostForm("day"))
	if err != nil {
		log.Println(err)

		helpers.JSON(c, "bad request", http.StatusBadRequest, nil, []string{"an error occur in converting date"})
		return
	}

	food.CreatedAt = time.Now()
	food.Name = foodName
	food.Type = foodType
	food.AdminName = admin.FullName
	food.Year = year
	food.Month = month
	food.Day = date
	food.Weekday = weekDay
	food.Images = images
	food.Kitchen = kitchen
	food.Status = "Not serving"

	err = u.UserService.CreateFoodTimetable(food)
	if err != nil {
		c.JSON(400, gin.H{"message": "This is a bad request"})
		return
	}

	notification := models.Notification{
		Message: admin.FullName + " added " + foodName + " to timetable",
		Year:    year,
		Month:   month,
		Day:     date,
	}

	err = u.UserService.CreateNotification(notification)
	if err != nil {
		helpers.JSON(c, "Error in getting Notification", http.StatusInternalServerError, nil, []string{"Error in getting Notification"})
		return
	}

	helpers.JSON(c, "Successfully Created", 201, nil, nil)
}
