package controller

import (
	"../structs"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (idb *InDB) SearchPowerRates(c *gin.Context) {
	var (
		data []structs.PowerRating
		ns   int
	)
	p := c.Param("q1")

	if p == "-" {
		p = ""
	}

	start := c.Param("start")
	count := c.Param("count")

	iStart, _ := strconv.Atoi(start)
	iCount, _ := strconv.Atoi(count)

	data, err := GetAllPowerRatesUsingParam(idb, p, iStart, iCount)
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

func (idb *InDB) GetPowerRatings(c *gin.Context) {
	var (
		data []structs.PowerRating
	)

	data, err := GetAllPowerRatings(idb)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
	return
}

func (idb *InDB) GetPowerRatingUsingID(c *gin.Context) {
	var (
		data structs.PowerRating
	)
	id := c.Param("q1")

	data, err := GetSinglePowerRatingUsingID(idb, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
	return
}

func (idb *InDB) CreatePowerRating(c *gin.Context) {
	var (
		data structs.PowerRating
	)

	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	data, err = CreatePowerRating(idb, data)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
	return
}

func (idb *InDB) UpdatePowerRating(c *gin.Context) {
	var (
		data structs.PowerRating
	)

	id := c.Param("q1")
	data, err := GetSinglePowerRatingUsingID(idb, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	data, err = UpdatePowerRating(idb, data)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
	return
}

func (idb *InDB) DeletePowerRating(c *gin.Context) {
	var (
		data structs.PowerRating
	)

	id := c.Param("q1")
	data, err := DeletePowerRating(idb, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
	return
}
