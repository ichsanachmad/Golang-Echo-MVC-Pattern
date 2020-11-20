package utils

import (
	"Golang-Echo-MVC-Pattern/responsegraph"
	"crypto/rand"
	"fmt"
	"github.com/jackc/pgx/v4"
	"regexp"
	"strings"
)

func IsStringEmpty(variable ...string) bool {
	for _, isi := range variable {
		if isi == "" {
			return true
		}
	}
	return false
}

func IsPhoneNumber(s string) (bool, error) {
	matched, err := regexp.Match(`[+][0-9]{2}[0-9]{10,}`, []byte(s))
	if err != nil {
		LogError(err, DetailFunction())
		return matched, err
	}
	return matched, nil
}

func IsEmailFormat(s string) (bool, error) {
	matched, err := regexp.Match(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`, []byte(s))
	if err != nil {
		LogError(err, DetailFunction())
		return matched, err
	}
	return matched, nil
}

func IsStringStrip(variable ...string) bool {
	for _, isi := range variable {
		if isi == "-" {
			return true
		}
	}
	return false
}

func GetResData(status string, message string, data interface{}) responsegraph.ResponseGenericGet {
	resp := responsegraph.ResponseGenericGet{
		Status:  status,
		Message: message,
		Data:    data,
	}
	return resp
}

func GetResNoData(status string, message string) responsegraph.ResponseGenericIn {
	resp := responsegraph.ResponseGenericIn{
		Status:  status,
		Message: message,
	}
	return resp
}

func GetUID() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		LogError(err, DetailFunction())
		return "", err
	}

	uuid := fmt.Sprintf("%x%x%x%x%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid, nil
}

func TrimString(message string, trim string) string {
	message = strings.Trim(message, trim)
	return message
}

func ResClose(res pgx.Rows) {
	res.Close()
}
