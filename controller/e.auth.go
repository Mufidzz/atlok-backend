package controller

import (
	"../security/jwt"
	"../structs"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func (idb *InDB) Login(c *gin.Context) {
	var (
		data     structs.User
		dataUser structs.User
	)

	err := c.BindJSON(&dataUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	data, err = GetSingleUserUsingUsername(idb, dataUser.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	if data.VerifiedAt == nil {
		c.JSON(http.StatusUpgradeRequired, fmt.Errorf("user not verified"))
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(dataUser.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	token, err := jwt.Generate(jwt.Body{Uid: data.ID, Acs: data.Access})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, token)

}
