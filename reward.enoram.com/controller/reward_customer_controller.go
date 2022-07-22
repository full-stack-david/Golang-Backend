package controller

import (
	"log"
	"mauappa-go/service"
	"mauappa-go/util"
	"net/http"

	"github.com/gorilla/mux"
)

type RewardCustomerController struct {
	rewardCustomerService service.RewardCustomerService
}

func NewRewardCustomerController(service service.RewardCustomerService) *RewardCustomerController {
	return &RewardCustomerController{rewardCustomerService: service}
}

func (x RewardCustomerController) GetRewardCustomerList(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	dropshipId := pathParams["dropshipId"]
	log.Printf("GetRewardCustomerList dropshipId %v: ", dropshipId)

	var resp map[string]interface{}

	if dropshipId == "" {
		log.Print("Invalid dropship ID.")
		resp = util.Message(false, "Invalid dropship ID.")
		util.Respond(w, resp)
		return
	}
	log.Printf("Dropship Id : %v ", dropshipId)

	customerList, err := x.rewardCustomerService.GetRewardCustomerList(dropshipId)

	if err != nil {
		log.Print(err)
		resp = util.Message(false, "Unable to get customer reward points list - "+err.Error())
		util.Respond(w, resp)
		return
	}

	resp = util.Message(true, "Reward Customer List got successfully")
	resp["data"] = customerList
	util.Respond(w, resp)
}

func (x RewardCustomerController) GetRewardCustomer(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	dropshipId := pathParams["dropshipId"]
	customerId := pathParams["customerId"]
	log.Printf("GetRewardCustomer dropshipId: %v, customerId: %v ", dropshipId, customerId)

	var resp map[string]interface{}

	if dropshipId == "" {
		log.Print("Invalid dropship ID.")
		resp = util.Message(false, "Invalid dropship ID.")
		util.Respond(w, resp)
		return
	}

	if customerId == "" {
		log.Print("Invalid customer ID.")
		resp = util.Message(false, "Invalid customer ID.")
		util.Respond(w, resp)
		return
	}

	customer, err := x.rewardCustomerService.GetRewardCustomer(dropshipId, customerId)

	if err != nil {
		log.Print(err)
		resp = util.Message(false, "Unable to get customer reward points - "+err.Error())
		util.Respond(w, resp)
		return
	}

	resp = util.Message(true, "Reward Customer got successfully")
	resp["data"] = customer
	util.Respond(w, resp)
}
