package controller

import (
	"../structs"
)

func GetAllSubstations(idb *InDB) (data []structs.Substation, err error) {
	fr := idb.DB.
		Find(&data)

	if fr.Error != nil {
		return data, fr.Error
	}

	return data, nil
}

func GetAllSubstationsUsingParam(idb *InDB, query string, start int, count int) (data []structs.Substation, err error) {
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

func GetSingleSubstationUsingID(idb *InDB, id string) (data structs.Substation, err error) {
	fr := idb.DB.
		Where("id = ?", id).
		First(&data)

	if fr.Error != nil {
		return data, fr.Error
	}

	return data, nil
}

func CreateSubstation(idb *InDB, Substation structs.Substation) (data structs.Substation, err error) {
	fr := idb.DB.
		Create(&Substation)
	if fr.Error != nil {
		return data, fr.Error
	}

	data = Substation
	return data, nil
}

func UpdateSubstation(idb *InDB, newSubstation structs.Substation) (data structs.Substation, err error) {
	fr := idb.DB.
		Save(&newSubstation)
	if fr.Error != nil {
		return data, fr.Error
	}

	data = newSubstation
	return data, nil
}

func DeleteSubstation(idb *InDB, id string) (data structs.Substation, err error) {
	data, err = GetSingleSubstationUsingID(idb, id)
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
