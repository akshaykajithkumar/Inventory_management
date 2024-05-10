package handler

import (
	services "main/pkg/usecase/interface"
	models "main/pkg/utils/models"
	"main/pkg/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	adminUseCase services.AdminUseCase
}

func NewAdminHandler(usecase services.AdminUseCase) *AdminHandler {
	return &AdminHandler{
		adminUseCase: usecase,
	}
}

// @Summary		Admin Login
// @Description	Login handler for admins
// @Tags			Admin
// @Accept			json
// @Produce		json
// @Param			admin	body		models.AdminLogin	true	"Admin login details"
// @Success		200		{object}	response.Response{}
// @Failure		500		{object}	response.Response{}
// @Router			/admin/adminlogin [post]
func (ad *AdminHandler) LoginHandler(c *gin.Context) { // login handler for the admin

	// var adminDetails models.AdminLogin
	var adminDetails models.AdminLogin
	if err := c.BindJSON(&adminDetails); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	admin, err := ad.adminUseCase.LoginHandler(adminDetails)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "cannot authenticate user", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	c.SetSameSite(http.SameSiteLaxMode)
	//c.SetCookie("Authorization", admin.Token, 3600, "/", "teeverse.online", true, false)
	c.SetCookie("Authorization", admin.AccessToken, 0, "/", "", false, true)
	c.SetCookie("Refreshtoken", admin.RefreshToken, 0, "/", "", false, true)

	successRes := response.ClientResponse(http.StatusOK, "Admin authenticated successfully", admin, nil)
	c.JSON(http.StatusOK, successRes)

}

// Logout is a handler for admin logout
// @Summary		admin Logout
// @Description	admin can logout
// @Tags			Admin
// @Accept			json
// @Produce		    json
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/logout [post]
func (ad *AdminHandler) Logout(c *gin.Context) {

	successRes := response.ClientResponse(http.StatusOK, "admin successfully logged out", nil, nil)
	c.SetCookie("Authorization", "", -1, "", "", true, false)
	c.JSON(http.StatusOK, successRes)
}
