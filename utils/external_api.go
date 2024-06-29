package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type UserInfo struct {
	Surname    string `json:"surname"`
	Name       string `json:"name"`
	Patronymic string `json:"patronymic"`
	Address    string `json:"address"`
}

// TODO:utils->FetchUserInfo
func FetchUserInfo(passportNumber string) (*UserInfo, error) {
	parts := strings.Split(passportNumber, " ")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid passport number format")
	}

	passportSerie, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, fmt.Errorf("invalid passport serie format")
	}

	passportNum, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, fmt.Errorf("invalid passport number format")
	}

	// Update URL to point to local mock server
	url := fmt.Sprintf("http://localhost:8081/info?passportSerie=%d&passportNumber=%d", passportSerie, passportNum)
	log.Printf("Fetching user info from URL: %s", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch user info, status code: %d", resp.StatusCode)
	}

	var userInfo UserInfo
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return nil, err
	}

	log.Printf("Fetched user info: %+v", userInfo)
	return &userInfo, nil
}
