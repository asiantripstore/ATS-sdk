package gwtools

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/asiantripstore/ATS-sdk/config"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

//RegisterToGateway send a request to register an api to the gateway
func RegisterToGateway(conf config.ApiConfig, route *gin.Engine) {
	routesInfos := route.Routes()
	var lstRoutes ListRoutes
	for _, route := range routesInfos {
		name := strings.Split(route.Handler, ".")
		lstRoutes.ListRoutes = append(lstRoutes.ListRoutes, RouteInfo{Path: route.Path, Method: route.Method, Name: name[len(name)-1]})
	}
	lstRoutes.AppName = conf.Name
	lstRoutes.Host = conf.Gateway.HostAPI
	strRoutes, err := json.Marshal(lstRoutes)
	if err != nil {
		log.Errorf("Error when creating list of routes", err.Error())
	}
	secret, err := bcrypt.GenerateFromPassword([]byte(conf.Gateway.RegisterKey), 8)
	if err != nil {
		log.Errorf("Error when creating secret", err.Error())
	}
	request, err := http.NewRequest("POST", "http://"+conf.Gateway.Host+":"+conf.Gateway.Port+"/v1/register", bytes.NewBuffer(strRoutes))
	request.Header.Set("Content-type", "application/json")
	request.Header.Set("Authorization", "Bearer "+string(secret))
	client := http.Client{Timeout: time.Duration(5 * time.Second)}
	resp, err := client.Do(request)
	if err != nil {
		log.Errorf("Error when register to GW, %s", err.Error())
		return
	}
	if resp.StatusCode == 200 {
		log.Info("registered")
	} else {
		log.Info("Not registered")

	}
}
