package controller

import (
	"../structs"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (idb *InDB) GetCustomerChanges(c *gin.Context) {
	var (
		data []structs.CustomerChangeWCustomer
		ns   int
	)

	start := c.Param("q1")
	count := c.Param("q2")

	iStart, _ := strconv.Atoi(start)
	iCount, _ := strconv.Atoi(count)

	data, err := GetAllCustomerChanges(idb, iStart, iCount)
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

func (idb *InDB) GetCustomerChangeUsingID(c *gin.Context) {
	var (
		data structs.CustomerChangeWCustomer
	)
	id := c.Param("q1")

	data, err := GetSingleCustomerChangeWCustomerUsingID(idb, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
	return
}

func (idb *InDB) CreateCustomerChange(c *gin.Context) {
	var (
		newData structs.Customer
		cData   structs.CustomerChange
	)

	oldId, _ := strconv.Atoi(c.Param("q1"))

	_, err := VerifyCustomerChangeData(idb, oldId)

	if err != nil {
		c.JSON(http.StatusConflict, err.Error())
		return
	}

	err = c.BindJSON(&newData)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	newData.ID = 0
	newData.Verified = false
	unverifiedData, err := CreateCustomer(idb, newData)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	cData = structs.CustomerChange{
		CurrentCustomerID: uint(oldId),
		NewCustomerID:     unverifiedData.ID,
	}

	data, err := CreateCustomerChange(idb, cData)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
	return
}

func (idb *InDB) AcceptCustomerChange(c *gin.Context) {
	var (
		newData structs.Customer
		cData   structs.CustomerChange
	)

	cID := c.Param("q1")
	cData, err := GetSingleCustomerChangeUsingID(idb, cID)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	_, err = DeleteCustomer(idb, strconv.Itoa(int(cData.CurrentCustomerID)))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	newData, err = GetSingleCustomerUsingID(idb, strconv.Itoa(int(cData.NewCustomerID)))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	newData.Verified = true
	newData, err = UpdateCustomer(idb, newData)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	_, err = DeleteCustomerChange(idb, cID)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Success")
	return
}

func (idb *InDB) DenyCustomerChange(c *gin.Context) {
	cID := c.Param("q1")
	_, err := DeleteCustomerChange(idb, cID)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Success")
	return
}
