package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"project_dwi/auth"
	"project_dwi/helper"
	"project_dwi/users"
)

type userHandler struct {
	userService users.Service
	authService auth.Service
}

func NewUserHandler(userService users.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

// RegisterUser godoc
// @Summary      Register User
// @Description  Register an account
// @Tags         User
// @Accept		 json
// @Produce      json
// @Param        payload   body      users.RegisterUserInput  true  "Request Body"
// @Success      200  {object}  helper.Response{data=users.UserFormatResponse} "Success to register an account"
// @Failure      400  {object}  helper.Response{data=interface{}} "Failed to register an account"
// @Router       /register [post]
func (h *userHandler) RegisterUser(c *gin.Context) {
	var input users.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Failed to register an account", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	user, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("Failed to register an account", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	token, err := h.authService.GenerateToken(user.ID)
	if err != nil {
		response := helper.APIResponse("Failed to register an account", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := users.FormatUserResponse(user, token)
	response := helper.APIResponse("Account has been registered", http.StatusOK, "Success", formatter)
	c.JSON(http.StatusOK, response)
}

// LoginUser godoc
// @Summary      Login User
// @Description  Login an account
// @Tags         User
// @Accept		 json
// @Produce      json
// @Param        payload   body      users.LoginUserInput  true  "Request Body"
// @Success      200  {object}  helper.Response{data=users.UserFormatResponse} "Success to login"
// @Failure      400  {object}  helper.Response{data=interface{}} "Failed to login"
// @Router       /sessions [post]
func (h *userHandler) LoginUser(c *gin.Context) {
	var input users.LoginUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	loggedInUser, err := h.userService.Login(input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	token, err := h.authService.GenerateToken(loggedInUser.ID)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	formatter := users.FormatUserResponse(loggedInUser, token)
	response := helper.APIResponse("Login Success", http.StatusOK, "Success", formatter)
	c.JSON(http.StatusOK, response)
}

// CheckEmailAvailability godoc
// @Summary      Check Email Availability
// @Description  Check an email whether it's available or not
// @Tags         User
// @Accept		 json
// @Produce      json
// @Param        payload   body      users.CheckEmail  true  "Request Body"
// @Success      200  {object}  helper.Response{data=users.UserFormatResponse} "Success to check email"
// @Failure      400  {object}  helper.Response{data=interface{}} "Failed to check email"
// @Router       /checkEmail [post]
func (h *userHandler) CheckEmailAvailability(c *gin.Context) {
	var input users.CheckEmail

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Failed to check email", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	isEmailAvailable, err := h.userService.IsEmailAvailable(input)

	if err != nil {
		errorMessage := gin.H{"errors": "Server Error"}
		response := helper.APIResponse("Failed to check email", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	data := gin.H{
		"is_available": isEmailAvailable,
	}

	metaMessage := "Email has been registered"

	if isEmailAvailable {
		metaMessage = "Email is available"
	}
	response := helper.APIResponse(metaMessage, http.StatusOK, "Success", data)
	c.JSON(http.StatusOK, response)
}

// UploadAvatar godoc
// @Summary      Upload Avatar
// @Description  Upload account avatar
// @Tags         User
// @Accept		 mpfd
// @Produce      json
// @Param        AvatarFileName   formData      string  true  "Request"
// @Security	 ApiAuthKey
// @Success      200  {object}  helper.Response{data=users.UserFormatResponse} "Success to upload avatar"
// @Failure      400  {object}  helper.Response{data=interface{}} "Failed to upload avatar"
// @Router       /avatars [post]
func (h *userHandler) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar")

	if err != nil {
		data := gin.H{
			"is_uploaded": false,
		}
		response := helper.APIResponse("Failed to upload avatar", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	currentUser := c.MustGet("currentUser").(users.User)
	userID := currentUser.ID

	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)

	if err != nil {
		data := gin.H{
			"is_uploaded": false,
		}
		response := helper.APIResponse("Failed to upload avatar", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.userService.SaveAvatar(userID, path)
	if err != nil {
		data := gin.H{
			"is_uploaded": false,
		}
		response := helper.APIResponse("Failed to upload avatar", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	data := gin.H{
		"is_uploaded": true,
	}
	response := helper.APIResponse("Avatar successfully uploaded", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}

// FetchUser godoc
// @Summary      Fetch User
// @Description  Fetch User
// @Tags         User
// @Accept		 json
// @Produce      json
// @Param        payload   body      users.User  true  "Request Body"
// @Security	 ApiAuthKey
// @Success      200  {object}  helper.Response{data=users.UserFormatResponse} "Success to fetch user"
// @Failure      400  {object}  helper.Response{data=interface{}} "Failed to fetch user"
// @Router       /users/fetch [post]
func (h *userHandler) FetchUser(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(users.User)

	formatter := users.FormatUserResponse(currentUser, "")
	response := helper.APIResponse("Success to fetch user data", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
