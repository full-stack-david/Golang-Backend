package controller

import (
	"log"
	"mauappa-go/util"
	"net/http"
)

type HealthController struct {
}

func NewHealthController() *HealthController {
	return &HealthController{}
}

func (j HealthController) HealthCheck(w http.ResponseWriter, r *http.Request) {
	log.Println("##  Health Check Request ##")
	util.Respond(w, util.Message(true, "Service is Up and Running"))
}
