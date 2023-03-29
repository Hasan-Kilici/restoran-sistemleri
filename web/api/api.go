package api

import(
	"os"
	"encoding/csv"
	"strconv"
)

type Users struct {
	ID           int
	Token        string
	Username     string
	Email        string
	Password     string
	Perms        string
}

type Foods struct{
	ID			int
	Token		string
	Name		string
	Price		int
	ImagePath	string
}

type Tables struct{
	ID			int
	Token		string
}

type Orders struct{
	ID				int
	TableID			int
	FoodName		string
	FoodPrice		int
}

type Perms struct{
	ID				int
	Token			string
	AllowedHours	string
	Name 			string
}

func ReadUsers(filePath string) ([]Users, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var users []Users
	header := true
	for {
		record, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return nil, err
		}

		if header {
			header = false
			continue
		}

		id, err := strconv.Atoi(record[0])
		if err != nil {
			return nil, err
		}

		user := Users{
			ID:           id,
			Token:        "Secret",
			Username:     record[2],
			Email:        "Secret",
			Password:     "Secret",
			Perms:        "Secret",
		}

		users = append(users, user)
	}
	return users, nil
}

func ReadTables(filePath string) ([]Tables, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var tables []Tables
	header := true
	for {
		record, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return nil, err
		}

		if header {
			header = false
			continue
		}

		id, err := strconv.Atoi(record[0])
		if err != nil {
			return nil, err
		}

		table := Tables{
			ID:           		id,
			Token:				"Secret",
		}

		tables = append(tables, table)
	}
	return tables, nil
}

func ReadFoods(filePath string) ([]Foods, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var foods []Foods
	header := true
	for {
		record, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return nil, err
		}

		if header {
			header = false
			continue
		}

		id, err := strconv.Atoi(record[0])
		if err != nil {
			return nil, err
		}

		Price , err := strconv.Atoi(record[3])
		if err != nil {
			return nil, err
		}

		food := Foods{
			ID:           	id,
			Token:			"Secret",
			Name:			record[2],
			Price: 			Price,
			ImagePath:		record[4],
		}

		foods = append(foods, food)
	}
	return foods, nil
}

func ReadOrders(filePath string) ([]Orders, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var orders []Orders
	header := true
	for {
		record, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return nil, err
		}

		if header {
			header = false
			continue
		}

		id, err := strconv.Atoi(record[0])
		if err != nil {
			return nil, err
		}

		TableID , err := strconv.Atoi(record[1])
		if err != nil {
			return nil, err
		}

		Price , err := strconv.Atoi(record[3])
		if err != nil {
			return nil, err
		}

		order := Orders{
			ID:           			id,
			TableID:				TableID,
			FoodName:				record[2],
			FoodPrice: 				Price,
		}

		orders = append(orders, order)
	}
	return orders, nil
}

func ReadPerms(filePath string) ([]Perms, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var perms []Perms
	header := true
	for {
		record, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return nil, err
		}

		if header {
			header = false
			continue
		}

		id, err := strconv.Atoi(record[0])
		if err != nil {
			return nil, err
		}

		perm := Perms{
			ID:           		id,
			Token:        		"Secret",
			AllowedHours:     	record[2],
			Name:				record[3],
		}

		perms = append(perms, perm)
	}
	return perms, nil
}