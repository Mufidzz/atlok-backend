package controller

import (
	"../structs"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (idb *InDB) SearchSubstations(c *gin.Context) {
	var (
		data []structs.Substation
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

	data, err := GetAllSubstationsUsingParam(idb, p, iStart, iCount)
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

func (idb *InDB) GetSubstations(c *gin.Context) {
	var (
		data []structs.Substation
	)

	data, err := GetAllSubstations(idb)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
	return
}

func (idb *InDB) GetSubstationUsingID(c *gin.Context) {
	var (
		data structs.Substation
	)
	id := c.Param("q1")

	data, err := GetSingleSubstationUsingID(idb, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
	return
}

func (idb *InDB) CreateSubstation(c *gin.Context) {
	var (
		data structs.Substation
	)

	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	data, err = CreateSubstation(idb, data)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
	return
}

func (idb *InDB) UpdateSubstation(c *gin.Context) {
	var (
		data structs.Substation
	)

	id := c.Param("q1")
	data, err := GetSingleSubstationUsingID(idb, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	data, err = UpdateSubstation(idb, data)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
	return
}

func (idb *InDB) DeleteSubstation(c *gin.Context) {
	var (
		data structs.Substation
	)

	id := c.Param("q1")
	data, err := DeleteSubstation(idb, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
	return
}
