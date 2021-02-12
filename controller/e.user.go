package controller

import (
	"../structs"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func (idb *InDB) CreateUser(c *gin.Context) {
	var (
		data structs.User
	)

	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	data, err = CreateUser(idb, data)
	if err != nil {
		c.JSON(http.StatusFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
	return
}

func (idb *InDB) GetUnverifiedUsers(c *gin.Context) {
	var (
		data []structs.User
		ns   int
	)
	start := c.Param("start")
	count := c.Param("count")

	iStart, _ := strconv.Atoi(start)
	iCount, _ := strconv.Atoi(count)

	data, err := GetAllUnverifiedUsers(idb, iStart, iCount)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if len(data) < iCount {
		ns = -1
	} else {
		ns = iStart + iCount
	}

	c.JSON(http.StatusOK, gin.H{
		"NextStart": ns,
		"Data":      data,
	})
	return
}

func (idb *InDB) VerifyUser(c *gin.Context) {
	var (
		data  structs.User
		datas []structs.User
	)

	datas, err := GetAllVerifiedUsers(idb)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if len(datas) >= 35 {
		c.JSON(http.StatusLocked, "quotas exceed")
		return
	}

	id := c.Param("id")
	data, err = GetSingleUserUsingID(idb, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	now := time.Now()
	data.VerifiedAt = &now

	data, err = UpdateUserWithoutVerification(idb, data)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
	return
}

func (idb *InDB) DeleteUser(c *gin.Context) {
	var (
		data structs.User
	)

	id := c.Param("id")
	data, err := DeleteUser(idb, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
	return
}
