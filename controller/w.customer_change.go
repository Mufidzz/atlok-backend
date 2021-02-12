package controller

import (
	"../structs"
	"fmt"
)

func GetAllCustomerChanges(idb *InDB, start int, count int) (data []structs.CustomerChangeWCustomer, err error) {
	fr := idb.DB.
		Preload("CurrentCustomer").
		Limit(count).
		Offset(start).
		Find(&data)

	if fr.Error != nil {
		return data, fr.Error
	}

	return data, nil
}

func GetSingleCustomerChangeWCustomerUsingID(idb *InDB, id string) (data structs.CustomerChangeWCustomer, err error) {
	fr := idb.DB.
		Preload("CurrentCustomer").
		Where("id = ?", id).
		First(&data)

	if fr.Error != nil {
		return data, fr.Error
	}

	return data, nil
}

func GetSingleCustomerChangeUsingID(idb *InDB, id string) (data structs.CustomerChange, err error) {
	fr := idb.DB.
		Where("id = ?", id).
		Last(&data)

	if fr.Error != nil {
		return data, fr.Error
	}

	return data, nil
}

func CreateCustomerChange(idb *InDB, CustomerChange structs.CustomerChange) (data structs.CustomerChange, err error) {
	fr := idb.DB.
		Create(&CustomerChange)
	if fr.Error != nil {
		return data, fr.Error
	}

	data = CustomerChange
	return data, nil
}

func VerifyCustomerChangeData(idb *InDB, id int) (r bool, err error) {
	var (
		data structs.CustomerChange
	)

	fr := idb.DB.
		Where("current_customer_id = ?", id).
		Last(&data)

	if fr.Error != nil {
		return true, nil
	}

	return false, fmt.Errorf("data is in review")
}

func UpdateCustomerChange(idb *InDB, newCustomerChange structs.CustomerChange) (data structs.CustomerChange, err error) {
	fr := idb.DB.
		Save(&newCustomerChange)
	if fr.Error != nil {
		return data, fr.Error
	}

	data = newCustomerChange
	return data, nil
}

func DeleteCustomerChange(idb *InDB, id string) (data structs.CustomerChange, err error) {

	_, err = GetSingleCustomerUsingID(idb, "20")
	if err != nil {
		return data, err
	}

	data, err = GetSingleCustomerChangeUsingID(idb, id)
	if err != nil {
		return data, err
	}
	fr := idb.DB.
		Debug().
		Delete(&data)
	if fr.Error != nil {
		return data, fr.Error
	}
	return data, nil
}
