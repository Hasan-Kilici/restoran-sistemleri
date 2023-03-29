package dataset

import (
	"os"
	"encoding/csv"
	"strconv"
	"errors"
	"Kawethra/utils"
	"bufio"
	"fmt"
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

func readUserCSV(filePath string) ([]Users, error) {
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
			Token:        record[1],
			Username:     record[2],
			Email:        record[3],
			Password:     record[4],
			Perms:        record[5],
		}

		users = append(users, user)
	}
	return users, nil
}

func readPermsCSV(filePath string) ([]Perms, error) {
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
			Token:        		record[1],
			AllowedHours:     	record[2],
			Name:				record[3],
		}

		perms = append(perms, perm)
	}
	return perms, nil
}

func readTableCSV(filePath string) ([]Tables, error) {
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
			Token:				record[1],
		}

		tables = append(tables, table)
	}
	return tables, nil
}

func readFoodCSV(filePath string) ([]Foods, error) {
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
			Token:			record[1],
			Name:			record[2],
			Price: 			Price,
			ImagePath:		record[4],
		}

		foods = append(foods, food)
	}
	return foods, nil
}
func readOrdersCSV(filePath string) ([]Orders, error) {
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

func UserCSV() ([]Users, error){
	users , _ := readUserCSV("./data/users.csv");
	return users, nil
}

func PermCSV() ([]Perms, error){
	perms, _ := readPermsCSV("./data/perms.csv");
	return perms, nil
}

func TableCSV() ([]Tables, error){
	tables, _ := readTableCSV("./data/tables.csv");
	return tables, nil	
}
func FoodCSV() ([]Foods, error){
	foods, _ := readFoodCSV("./data/foods.csv");
	return foods, nil	
}
 
func loadUsersFromCSV() ([]Users, error) {
	file, err := os.Open("./data/users.csv")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	csvData, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	var users []Users
	for _, row := range csvData {
		userID, err := strconv.Atoi(row[0])
		if err != nil {
			return nil, err
		}
		user := Users{
			ID:           userID,
			Token:        row[1],
			Username:     row[2],
			Email:        row[3],
			Password:     row[4],
			Perms:        row[5],
		}
		users = append(users, user)
	}

	return users, nil
}

func FindUserByToken(token string) (*Users, error) {
	users, err := loadUsersFromCSV()
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		if user.Token == token {	
			return &user, nil
		}
	}

	return nil, errors.New("Kullanıcı Bulunamadı")
}

func AddFoodToCSV(filePath string, name string, price int, image string) error {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	defer file.Close()

	token := utils.GenerateToken(16)

	lines := 0
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines++
	}

	writer := csv.NewWriter(file)
	err = writer.Write([]string{strconv.Itoa(lines + 1), token, name, strconv.Itoa(price), image})
	if err != nil {
		return err
	}

	writer.Flush()

	return nil
}

func AddUserToCSV(filePath string, name string, password string, email string,perms string) error {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	defer file.Close()

	token := utils.GenerateToken(16)

	lines := 0
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines++
	}

	writer := csv.NewWriter(file)
	err = writer.Write([]string{strconv.Itoa(lines + 1),name, token, email,utils.HashPassword(password), perms})
	if err != nil {
		return err
	}

	writer.Flush()

	return nil
}


func AddTableToCSV(filePath string) error {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	defer file.Close()

	token := utils.GenerateToken(16)

	lines := 0
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines++
	}

	writer := csv.NewWriter(file)
	err = writer.Write([]string{strconv.Itoa(lines + 1), token})
	if err != nil {
		return err
	}

	writer.Flush()

	return nil
}

func AddPermToCSV(filePath string, AllowedHours string, Name string) error {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	defer file.Close()

	token := utils.GenerateToken(16)

	lines := 0
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines++
	}

	writer := csv.NewWriter(file)
	err = writer.Write([]string{strconv.Itoa(lines + 1), token, AllowedHours, Name})
	if err != nil {
		return err
	}

	writer.Flush()

	return nil
}

func ResetTable(Tableid string) error{
	file, err := os.OpenFile("./data/orders.csv", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Dosya açılırken hata oluştu")
	}
	defer file.Close()

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Dosya okunurken hata oluştu")
	}
	
	found := false
	for i, line := range lines {
		if line[1] == Tableid {
			lines = append(lines[:i], lines[i+1:]...)
			found = true
			break;
		}
	}

	if !found {
		fmt.Println("ID bulunamadı")
	}

	file.Truncate(0)
	file.Seek(0, 0)

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.WriteAll(lines)
	if err != nil {
		fmt.Println("Dosya yazılırken hata oluştu")
	}
	return nil
}

func DeleteFood(id string) error {
	file, err := os.OpenFile("./data/foods.csv", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Dosya açılırken hata oluştu")
	}
	defer file.Close()

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Dosya okunurken hata oluştu")
	}
	
	found := false
	for i, line := range lines {
		if line[0] == id {
			lines = append(lines[:i], lines[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		fmt.Println("ID bulunamadı")
	}

	file.Truncate(0)
	file.Seek(0, 0)

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.WriteAll(lines)
	if err != nil {
		fmt.Println("Dosya yazılırken hata oluştu")
	}
	return nil
}

func DeleteTable(id string) error {
	file, err := os.OpenFile("./data/tables.csv", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Dosya açılırken hata oluştu")
	}
	defer file.Close()

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Dosya okunurken hata oluştu")
	}
	
	found := false
	for i, line := range lines {
		if line[0] == id {
			lines = append(lines[:i], lines[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		fmt.Println("ID bulunamadı")
	}

	file.Truncate(0)
	file.Seek(0, 0)

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.WriteAll(lines)
	if err != nil {
		fmt.Println("Dosya yazılırken hata oluştu")
	}
	return nil
}

func DeletePerm(id string) error {
	file, err := os.OpenFile("./data/foods.csv", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Dosya açılırken hata oluştu")
	}
	defer file.Close()

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Dosya okunurken hata oluştu")
	}
	
	found := false
	for i, line := range lines {
		if line[0] == id {
			lines = append(lines[:i], lines[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		fmt.Println("ID bulunamadı")
	}

	file.Truncate(0)
	file.Seek(0, 0)

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.WriteAll(lines)
	if err != nil {
		fmt.Println("Dosya yazılırken hata oluştu")
	}
	return nil
}

func DeleteUser(id string) error {
	file, err := os.OpenFile("./data/users.csv", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Dosya açılırken hata oluştu")
	}
	defer file.Close()

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Dosya okunurken hata oluştu")
	}
	
	found := false
	for i, line := range lines {
		if line[0] == id {
			lines = append(lines[:i], lines[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		fmt.Println("ID bulunamadı")
	}

	file.Truncate(0)
	file.Seek(0, 0)

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.WriteAll(lines)
	if err != nil {
		fmt.Println("Dosya yazılırken hata oluştu")
	}
	return nil
}

func AddFoodToTable(tableid int,foodname string, price int) error {
	file, err := os.OpenFile("./data/orders.csv", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	defer file.Close()

	lines := 0
	f, err := os.Open("./data/orders.csv")
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines++
	}

	writer := csv.NewWriter(file)
	err = writer.Write([]string{strconv.Itoa(lines + 1), strconv.Itoa(tableid), foodname, strconv.Itoa(price)})
	if err != nil {
		return err
	}

	writer.Flush()
	return nil
}

func AddFoodToOrder(tableid int,foodname string, price int) error {
	file, err := os.OpenFile("./data/foodlist.csv", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	defer file.Close()

	lines := 0
	f, err := os.Open("./data/foodlist.csv")
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines++
	}

	writer := csv.NewWriter(file)
	err = writer.Write([]string{strconv.Itoa(lines + 1), strconv.Itoa(tableid), foodname, strconv.Itoa(price)})
	if err != nil {
		return err
	}

	writer.Flush()
	return nil
}

func DeleteOrder(Orderid string) error{
	file, err := os.OpenFile("./data/orders.csv", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Dosya açılırken hata oluştu")
	}
	defer file.Close()

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Dosya okunurken hata oluştu")
	}
	
	found := false
	for i, line := range lines {
		if line[0] == Orderid {
			lines = append(lines[:i], lines[i+1:]...)
			found = true
			break;
		}
	}

	if !found {
		fmt.Println("ID bulunamadı")
	}

	file.Truncate(0)
	file.Seek(0, 0)

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.WriteAll(lines)
	if err != nil {
		fmt.Println("Dosya yazılırken hata oluştu")
	}
	return nil
}

func DeleteFoodOrder(Orderid string) error{
	file, err := os.OpenFile("./data/foodlist.csv", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Dosya açılırken hata oluştu")
	}
	defer file.Close()

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Dosya okunurken hata oluştu")
	}
	
	found := false
	for i, line := range lines {
		if line[0] == Orderid {
			lines = append(lines[:i], lines[i+1:]...)
			found = true
			break;
		}
	}

	if !found {
		fmt.Println("ID bulunamadı")
	}

	file.Truncate(0)
	file.Seek(0, 0)

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.WriteAll(lines)
	if err != nil {
		fmt.Println("Dosya yazılırken hata oluştu")
	}
	return nil
}