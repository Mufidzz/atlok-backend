package controller

import (
	"../structs"
	"golang.org/x/crypto/bcrypt"
	"strconv"
)

func GetAllUsers(idb *InDB) (data []structs.User, err error) {
	fr := idb.DB.
		Find(&data)

	if fr.Error != nil {
		return data, fr.Error
	}

	return data, nil
}

func GetAllUnverifiedUsers(idb *InDB, start int, count int) (data []structs.User, err error) {
	fr := idb.DB.
		Debug().
		Where("verified_at IS NULL").
		Limit(count).
		Offset(start).
		Find(&data)

	if fr.Error != nil {
		return data, fr.Error
	}

	return data, nil
}
func GetAllVerifiedUsers(idb *InDB) (data []structs.User, err error) {
	fr := idb.DB.
		Debug().
		Where("verified_at IS NOT NULL").
		Find(&data)

	if fr.Error != nil {
		return data, fr.Error
	}

	return data, nil
}

func GetSingleUserUsingID(idb *InDB, id string) (data structs.User, err error) {
	fr := idb.DB.
		Where("id = ?", id).
		First(&data)

	if fr.Error != nil {
		return data, fr.Error
	}

	return data, nil
}

func GetSingleUserUsingUsername(idb *InDB, username string) (data structs.User, err error) {
	fr := idb.DB.
		Where("username = ?", username).
		First(&data)

	if fr.Error != nil {
		return data, fr.Error
	}

	return data, nil
}

func CreateUser(idb *InDB, user structs.User) (data structs.User, err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return data, err
	}

	user.Password = string(hash)

	fr := idb.DB.
		Create(&user)
	if fr.Error != nil {
		return data, fr.Error
	}

	data = user
	return data, nil
}

func UpdateUser(idb *InDB, newUser structs.User) (data structs.User, err error) {
	data, err = GetSingleUserUsingID(idb, strconv.Itoa(int(newUser.ID)))
	if err != nil {
		return data, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(newUser.Password)); err != nil {
		return data, err
	}

	fr := idb.DB.
		Save(&newUser)

	if fr.Error != nil {
		return data, fr.Error
	}

	data = newUser
	return data, nil
}

func UpdateUserWithoutVerification(idb *InDB, newUser structs.User) (data structs.User, err error) {
	data, err = GetSingleUserUsingID(idb, strconv.Itoa(int(newUser.ID)))
	if err != nil {
		return data, err
	}

	fr := idb.DB.
		Save(&newUser)

	if fr.Error != nil {
		return data, fr.Error
	}

	data = newUser
	return data, nil
}

func DeleteUser(idb *InDB, id string) (data structs.User, err error) {
	data, err = GetSingleUserUsingID(idb, id)
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
