package controller

import (
	"../structs"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (idb *InDB) GetCustomerUsingID(c *gin.Context) {
	var (
		data structs.CustomerWSubstationPowerRating
	)
	id := c.Param("q1")

	data, err := GetSingleAssociatedCustomerUsingID(idb, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
	return
}

func (idb *InDB) SearchCustomers(c *gin.Context) {
	var (
		data []structs.CustomerWSubstationPowerRating
		ns   int
		tdf  int
	)
	p := c.Param("q1")

	if p == "-" {
		p = ""
	}

	start := c.Param("start")
	count := c.Param("count")

	urlQuery := c.Request.URL.Query()

	fmt.Println("URLQ")
	fmt.Println(urlQuery)

	iStart, _ := strconv.Atoi(start)
	iCount, _ := strconv.Atoi(count)

	data, err := GetAllCustomersUsingParam(idb, p, iStart, iCount, urlQuery)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	tdf, err = GetAllCustomerCountUsingParam(idb, p, urlQuery)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if len(data) < iCount {
		ns = -1
	} else {
		ns = iStart + iCount
	}

	c.JSON(http.StatusOK, gin.H{
		"NextStart":      ns,
		"TotalDataFound": tdf,
		"Data":           data,
	})
	return
}

func (idb *InDB) CreateCustomer(c *gin.Context) {
	var (
		data structs.Customer
	)

	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	data, err = CreateCustomer(idb, data)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
	return
}

func (idb *InDB) UpdateCustomer(c *gin.Context) {
	var (
		data structs.Customer
	)

	id := c.Param("q1")
	data, err := GetSingleCustomerUsingID(idb, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	data, err = UpdateCustomer(idb, data)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
	return
}

func (idb *InDB) DeleteCustomer(c *gin.Context) {
	var (
		data structs.Customer
	)

	id := c.Param("q1")
	data, err := DeleteCustomer(idb, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
	return
}
