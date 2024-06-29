package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type UserInfo struct {
	Surname    string `json:"surname"`
	Name       string `json:"name"`
	Patronymic string `json:"patronymic"`
	Address    string `json:"address"`
}

func FetchUserInfo(passportNumber string) (*UserInfo, error) {
	parts := strings.Split(passportNumber, " ")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid passport number format")
	}

	passportSerie, passportNum := parts[0], parts[1]

	url := fmt.Sprintf("http://external.api/info?passportSerie=%s&passportNumber=%s", passportSerie, passportNum)
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

	return &userInfo, nil
}
