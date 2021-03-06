package handler

import (
	"bwastartup/campaign"
	"bwastartup/helper"
	"bwastartup/user"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// tangkap parameter di handler
// handler ke service
// service yang menentukan repository mana yang digunakan atau call
// repository :FindAll, FindByUserID
// db

type campaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service}
}

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.service.GetCampaigns(userID)
	if err != nil {
		response := helper.APIResponseError("Error to get camapigns", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := campaign.FormatCampaigns(campaigns)
	response := helper.APIResponseSuccess("List of campaigns", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) GetCampaign(c *gin.Context) {

	var input campaign.GetCampaignDetailInput

	err := c.ShouldBindUri(&input)

	if err != nil {
		response := helper.APIResponseError("Failed to get detail campaign", http.StatusBadRequest, err)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	campaignDetail, err := h.service.GetCampaignById(input)
	if err != nil {
		response := helper.APIResponseError("Failed to get detail campaign", http.StatusBadRequest, err)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := campaign.FormatCampaignDetail(campaignDetail)
	response := helper.APIResponseSuccess("Campaign Detail", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) CreateCampaign(c *gin.Context) {
	var input campaign.CreateCampaignInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponseError("Failed to create campaign", http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// Ambil current user
	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	newCampaign, err := h.service.CreateCampaign(input)
	if err != nil {
		response := helper.APIResponseError("Failed to create campaign 3", http.StatusBadRequest, err)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponseSuccess("Success to create campaign", campaign.FormatCampaign(newCampaign))
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) UpdateCampaign(c *gin.Context) {

	var inputId campaign.GetCampaignDetailInput

	err := c.ShouldBindUri(&inputId)

	if err != nil {
		response := helper.APIResponseError("Failed to update campaign", http.StatusBadRequest, err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData campaign.CreateCampaignInput
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponseError("Failed to update campaign", http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// Ambil current user
	currentUser := c.MustGet("currentUser").(user.User)
	inputData.User = currentUser

	updatedCampaign, err := h.service.UpdateCampaign(inputId, inputData)
	if err != nil {
		response := helper.APIResponseError("Failed to update campaign", http.StatusBadRequest, err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponseSuccess("Success to update campaign", campaign.FormatCampaign(updatedCampaign))
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) UploadImage(c *gin.Context) {
	// var input campaign.CreateImageInput
	// err := c.ShouldBindJSON(&input)
	// if err != nil {
	// 	errors := helper.FormatValidationError(err)
	// 	errorMessage := gin.H{"errors": errors}

	// 	response := helper.APIResponseError("Failed to create campaign image", http.StatusUnprocessableEntity, errorMessage)
	// 	c.JSON(http.StatusUnprocessableEntity, response)
	// 	return
	// }

	// var campaignID campaign.GetCampaignDetailInput

	// err = c.ShouldBindUri(&campaignID)

	// if err != nil {
	// 	response := helper.APIResponseError("Failed to upload campaign imag", http.StatusBadRequest, err)
	// 	c.JSON(http.StatusBadRequest, response)
	// 	return
	// }

	// file, err := c.FormFile("file")

	// if err != nil {
	// 	data := gin.H{"is_uploaded": false}
	// 	response := helper.APIResponseError("Failed to upload campaign image", http.StatusBadRequest, data)
	// 	c.JSON(http.StatusBadRequest, response)
	// 	return
	// }

	// path := fmt.Sprintf("images-campaign/%d-%s", campaignID.ID, file.Filename)
	// err = c.SaveUploadedFile(file, path)
	// if err != nil {
	// 	data := gin.H{"is_uploaded": false}
	// 	response := helper.APIResponseError("Failed to upload campaign image", http.StatusBadRequest, data)
	// 	c.JSON(http.StatusBadRequest, response)
	// 	return
	// }

	// newImage, err := h.service.SaveCampaignImage(campaignID.ID, input, path)
	// if err != nil {
	// 	data := gin.H{"is_uploaded": false}
	// 	response := helper.APIResponseError("Failed to upload campaign image", http.StatusBadRequest, data)
	// 	c.JSON(http.StatusBadRequest, response)
	// 	return
	// }

	// response := helper.APIResponseSuccess("Success to upload campaign image", newImage)
	// c.JSON(http.StatusBadRequest, response)
	var input campaign.CreateImageInput
	err := c.ShouldBind(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponseError("Failed to update campaign", http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponseError("Failed to upload campaign image", http.StatusBadRequest, data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Ambil current user
	currentUser := c.MustGet("currentUser").(user.User)
	// HArusnya dari JwT
	userID := currentUser.ID
	input.User = currentUser
	path := fmt.Sprintf("images/campaign/%d-%s", userID, file.Filename)
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponseError("Failed to upload campaign image", http.StatusBadRequest, data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newImage, err := h.service.SaveCampaignImage(input, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponseError("Failed to upload campaign image", http.StatusBadRequest, data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponseSuccess("Success to upload campaign image", newImage)
	c.JSON(http.StatusBadRequest, response)

}
