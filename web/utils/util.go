package utils

import (
	"math/rand"
	"time"
    "encoding/csv"
	"log"
	"os"
	"strconv"
	"strings"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
    "errors"
)

func GetRandomNum(num int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(num)
}

func GenerateToken(length int) string {
	rand.Seed(time.Now().UnixNano())

	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}


func HashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return ""
	}
	return string(hash)
}

func HasRequiredPerms(ctx *gin.Context, requiredPerms []int) bool {
    token, err := ctx.Cookie("token")
    if err != nil {
        return false
    }
    
    file, err := os.Open("./data/users.csv")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    reader := csv.NewReader(file)
    lines, err := reader.ReadAll()
    if err != nil {
        log.Fatal(err)
    }

    var perms []string
    for _, line := range lines {
        if line[1] == token {
            perms = strings.Split(line[5], ">")
            break
        }
    }

    for _, requiredPerm := range requiredPerms {
        found := false
        for _, perm := range perms {
            permInt, err := strconv.Atoi(strings.TrimSpace(perm))
            if err != nil {
                log.Fatal(err)
            }
            if permInt == requiredPerm {
                found = true
                break
            }
        }
        if !found {
            return false
        }
    }
    return true
}

func LoginUser(email, password string) (string, error) {
    file, err := os.Open("./data/users.csv")
    if err != nil {
        return "", err
    }
    defer file.Close()

    reader := csv.NewReader(file)
    records, err := reader.ReadAll()
    if err != nil {
        return "", err
    }

    email = strings.ToLower(email)
    for _, record := range records {
        if len(record) != 6 {
            continue
        }
        if strings.ToLower(record[3]) == email {
            if CheckPassword(record[4], password) == true {
                token := record[1]

                file, err := os.OpenFile("./data/users.csv", os.O_WRONLY, 0755)
                if err != nil {
                    return "", err
                }
                defer file.Close()

                writer := csv.NewWriter(file)
                if err := writer.WriteAll(records); err != nil {
                    return "", err
                }
                return token, nil
            } else {
                return "", errors.New("Geçersiz şifre")
            }
        }
    }

    return "", errors.New("Kullanıcı Bulunamadı")
}

func CheckPassword(hashedPassword, password string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
    return err == nil
}

func ReplaceSpacesWithDash(text string) string {
    return strings.ReplaceAll(text, " ", "-")
}
