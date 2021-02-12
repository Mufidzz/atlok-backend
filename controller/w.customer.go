package controller

import (
	"../structs"
	"net/url"
)

func GetAllCustomers(idb *InDB) (data []structs.Customer, err error) {
	fr := idb.DB.
		Find(&data)

	if fr.Error != nil {
		return data, fr.Error
	}

	return data, nil
}

func GetAllCustomersUsingParam(idb *InDB, query string, start int, count int, urlQuery url.Values) (data []structs.CustomerWSubstationPowerRating, err error) {
	fr := idb.DB.
		Debug().
		Preload("Substation").
		Preload("PowerRating").
		Where("id_pel LIKE ? OR full_name LIKE ?", "%"+query+"%", "%"+query+"%").
		Where("verified = ?", true)

	q := urlQuery["pr"]
	if len(q) > 0 {
		for i := 0; i < len(q); i++ {
			fr.Where("power_rating_id = ?", q[i])
		}
	}

	q = urlQuery["s"]
	if len(q) > 0 {
		for i := 0; i < len(q); i++ {
			fr.Where("substation_id = ?", q[i])
		}
	}

	fr.Order("LENGTH(id_pel), LENGTH(full_name)").
		Limit(count).
		Offset(start).
		Find(&data)

	if fr.Error != nil {
		return data, fr.Error
	}

	return data, nil
}

func GetAllCustomerCountUsingParam(idb *InDB, query string, urlQuery url.Values) (count int, err error) {
	var (
		data []structs.Customer
	)
	fr := idb.DB.
		Debug().
		Select("id").
		Where("id_pel LIKE ? OR full_name LIKE ?", "%"+query+"%", "%"+query+"%").
		Where("verified = ?", true)

	q := urlQuery["pr"]
	if len(q) > 0 {
		for i := 0; i < len(q); i++ {
			fr.Where("power_rating_id = ?", q[i])
		}
	}

	q = urlQuery["s"]
	if len(q) > 0 {
		for i := 0; i < len(q); i++ {
			fr.Where("substation_id = ?", q[i])
		}
	}

	fr.Find(&data)

	if fr.Error != nil {
		return len(data), fr.Error
	}

	return len(data), nil
}

func GetSingleCustomerUsingID(idb *InDB, id string) (data structs.Customer, err error) {
	fr := idb.DB.
		Debug().
		Where("id = ?", id).
		First(&data)

	if fr.Error != nil {
		return data, fr.Error
	}

	return data, nil
}
func GetSingleAssociatedCustomerUsingID(idb *InDB, id string) (data structs.CustomerWSubstationPowerRating, err error) {
	fr := idb.DB.
		Preload("Substation").
		Preload("PowerRating").
		Where("id = ?", id).
		First(&data)

	if fr.Error != nil {
		return data, fr.Error
	}

	return data, nil
}

func CreateCustomer(idb *InDB, Customer structs.Customer) (data structs.Customer, err error) {
	fr := idb.DB.
		Create(&Customer)
	if fr.Error != nil {
		return data, fr.Error
	}

	data = Customer
	return data, nil
}

func UpdateCustomer(idb *InDB, newCustomer structs.Customer) (data structs.Customer, err error) {
	fr := idb.DB.
		Save(&newCustomer)
	if fr.Error != nil {
		return data, fr.Error
	}

	data = newCustomer
	return data, nil
}

func DeleteCustomer(idb *InDB, id string) (data structs.Customer, err error) {
	data, err = GetSingleCustomerUsingID(idb, id)
	if err != nil {
		return data, err
	}
	fr := idb.DB.
		Delete(&data)
	if fr.Error != nil {
		return data, fr.Error
	}
	return data, nil
}
