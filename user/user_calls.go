package user

import (
	"bytes"
	"encoding/json"
	"net/http"

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
	return &user
}
