package controller

import (
	"../structs"
)

func GetAllPowerRatesUsingParam(idb *InDB, query string, start int, count int) (data []structs.PowerRating, err error) {
	fr := idb.DB.
		Where("code LIKE ?", "%"+query+"%").
		Or("name LIKE ?", "%"+query+"%").
		Limit(count).
		Offset(start).
		Find(&data)

	if fr.Error != nil {
		return data, fr.Error
	}

	return data, nil
}

func GetAllPowerRatings(idb *InDB) (data []structs.PowerRating, err error) {
	fr := idb.DB.
		Find(&data)

	if fr.Error != nil {
		return data, fr.Error
	}

	return data, nil
}

func GetSinglePowerRatingUsingID(idb *InDB, id string) (data structs.PowerRating, err error) {
	fr := idb.DB.
		Where("id = ?", id).
		First(&data)

	if fr.Error != nil {
		return data, fr.Error
	}

	return data, nil
}

func CreatePowerRating(idb *InDB, PowerRating structs.PowerRating) (data structs.PowerRating, err error) {
	fr := idb.DB.
		Create(&PowerRating)
	if fr.Error != nil {
		return data, fr.Error
	}

	data = PowerRating
	return data, nil
}

func UpdatePowerRating(idb *InDB, newPowerRating structs.PowerRating) (data structs.PowerRating, err error) {
	fr := idb.DB.
		Save(&newPowerRating)
	if fr.Error != nil {
		return data, fr.Error
	}

	data = newPowerRating
	return data, nil
}

func DeletePowerRating(idb *InDB, id string) (data structs.PowerRating, err error) {
	data, err = GetSinglePowerRatingUsingID(idb, id)
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
