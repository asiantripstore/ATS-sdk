package user

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/asiantripstore/ATS-sdk/models"
	log "github.com/sirupsen/logrus"
)

func CreateUser(gatewayHost string, gatewayPort string, user models.User) *models.User {
	body, err := json.Marshal(user)
	if err != nil {
		log.Errorf("Error in Marshal CreateUser")
		return nil
	}
	responseBody := bytes.NewBuffer(body)
	resp, err := http.Post("http://"+gatewayHost+":"+gatewayPort+"/v1/ats-user/v1/user", "application/json", responseBody)
	if err != nil {
		log.Errorf("Error when posting new user, %s", err.Error())
		return nil
	}
	if resp.StatusCode == 200 {
		log.Info("New user created")
	}
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("Error when reading response, %s", err.Error())
		return nil
	}
	defer resp.Body.Close()
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Errorf("Error when unmarshalling, %s", err.Error())
		return nil
	}
	return &user
}

func GetUser(gatewayHost string, gatewayPort string, id int) *models.User {
	var user models.User
	resp, err := http.Get("http://" + gatewayHost + ":" + gatewayPort + "/v1/ats-user/v1/user/" + strconv.Itoa(id))
	if err != nil {
		log.Errorf("Error when getting user, %s", err.Error())
		return nil
	}
	if resp.StatusCode != 200 {
		log.Info("Get user failed")
		return nil
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("Error when reading response, %s", err.Error())
		return nil
	}
	defer resp.Body.Close()
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Errorf("Error when unmarshalling, %s", err.Error())
		return nil
	}
	return &user
}
