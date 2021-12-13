package handler

import (
	"fmt"
	"go-startup/campaign"
	"go-startup/helper"
	"go-startup/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// tangkap parameter di handler
// handler ke service
// service yang menentukan repo mana yang di call
// repo GetAll, GetByUserID
// db

type campaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service}
}

// api/v1/campaign

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	// mengambil id dari params
	userID, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.service.GetCampaigns(userID)

	if err != nil {
		response := helper.APIResponse("Error to get campaigns", http.StatusBadRequest, "error", nil)

		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("LIst Of Campaigns campaigns", http.StatusOK, "success", campaign.FormatCampaigns(campaigns))
	c.JSON(http.StatusOK, response)

}

func (h *campaignHandler) GetCampaign(c *gin.Context) {
	var input campaign.GetCampaignDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Error to get campaigns", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	campaignDetail, err := h.service.GetCampaignsByID(input)
	if err != nil {
		response := helper.APIResponse("Error to get campaigns", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return

	}

	response := helper.APIResponse("campaign detail", http.StatusOK, "success", campaign.FormatCampaignDetail(campaignDetail))
	c.JSON(http.StatusOK, response)

	// handler : mapping id dari url ke struct input => service, call formatter
	// service : inputnya struct untuk menangkap id di url
	// repository : get campaign by id
}

// tangkap parameter dari user ke struct input
// ambil current user dair jwt/handler
// panggil service , parameternya input struct (Dan juga slug)
// panggil repo untuk simpan data campaign baru
func (h *campaignHandler) CreateCampaign(c *gin.Context) {
	var input campaign.CreateCampaignInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}

		response := helper.APIResponse("failed to create campaign ", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return

	}
	currentUser := c.MustGet("currentUser").(user.User)

	input.User = currentUser

	newCampaign, err := h.service.CreateCampaign(input)
	if err != nil {
		response := helper.APIResponse("failed to create campaign ", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return

	}
	response := helper.APIResponse("success to created campaign ", http.StatusOK, "success", campaign.FormatCampaign(newCampaign))
	c.JSON(http.StatusOK, response)

}

// user input
// handler
// mapping input ke input struct ada 2 dari user dan ur
// repository update
// service
// repository update data campaign
//

func (h *campaignHandler) UpdateCampaign(c *gin.Context) {
	var inputID campaign.GetCampaignDetailInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to update campaign", http.StatusBadRequest, "error", nil)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	// sudah dapat id nya

	var inputData campaign.CreateCampaignInput

	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := helper.APIResponse("Failed to update Campaign", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, errorMessage)
		return

	}

	currentUser := c.MustGet("currentUser").(user.User)

	inputData.User = currentUser
	updatedCampaign, err := h.service.UpdateCampaign(inputID, inputData)

	if err != nil {
		response := helper.APIResponse("Failed to update Campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to update Campaign", http.StatusOK, "Success", updatedCampaign)
	c.JSON(http.StatusOK, response)
}

// membuat upload campaign images
// tangka input dan ubah ke struct input
// save images campaign ke folder
// service kondisi manggil point 1 dan 2 di repo
// repository
// 1. create images/save data ke database campaign_images
// 2. ubah is_primary true ke false

func (h *campaignHandler) UploadImage(c *gin.Context) {
	var input campaign.CreateCampaignImageInput

	err := c.ShouldBind(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := helper.APIResponse("Failed to upload Campaign image", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, errorMessage)
		return

	}

	file, err := c.FormFile("file")

	if err != nil {
		data := gin.H{
			"is_uploaded": false,
		}
		response := helper.APIResponse("failed upload avatar campaign image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return

	}

	currentUser := c.MustGet("currentUser").(user.User)
	userId := currentUser.ID

	//path := "images/" + file.Filename
	path := fmt.Sprintf("images/%d-%s", userId, file.Filename)

	err = c.SaveUploadedFile(file, path)

	if err != nil {
		data := gin.H{
			"is_uploaded": false,
		}
		response := helper.APIResponse("failed upload campaign image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return

	}

	_, err = h.service.SaveCampaignImage(input, path)
	if err != nil {
		data := gin.H{
			"is_uploaded": false,
		}
		response := helper.APIResponse("failed upload campaign image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{
		"is_uploaded": true,
	}
	response := helper.APIResponse("campaign image success uploaded", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
	return

}
