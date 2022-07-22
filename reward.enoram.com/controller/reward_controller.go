package controller

import (
	"log"
	"mauappa-go/service"
	"mauappa-go/util"
	"net/http"
)

type RewardController struct {
	rewardOrderService    service.RewardOrderService
	rewardCustomerService service.RewardCustomerService
}

func NewRewardController(rewardOrderService service.RewardOrderService, rewardCustomerService service.RewardCustomerService) *RewardController {
	return &RewardController{rewardOrderService: rewardOrderService, rewardCustomerService: rewardCustomerService}
}

func (x RewardController) CalculateReward(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	storeId := queryParams.Get("storeId")
	dropshipId := queryParams.Get("dropshipId")
	orderId := queryParams.Get("orderId")

	var resp map[string]interface{}

	if storeId == "" {
		log.Print("Invalid store ID.")
		resp = util.Message(false, "Invalid store ID.")
		util.Respond(w, resp)
		return
	}
	log.Printf("Store Id : %v ", storeId)

	if dropshipId == "" {
		log.Print("Invalid dropship ID.")
		resp = util.Message(false, "Invalid dropship ID.")
		util.Respond(w, resp)
		return
	}
	log.Printf("Dropship Id : %v ", dropshipId)

	if orderId == "" {
		log.Print("Invalid order ID.")
		resp = util.Message(false, "Invalid order ID.")
		util.Respond(w, resp)
		return
	}
	log.Printf("Order Id : %v ", orderId)

	_, err := x.rewardCustomerService.DeleteRewardCustomerContent(dropshipId, orderId)

	if err != nil {
		log.Print(err)
		resp = util.Message(false, "Unable to calculate reward points - "+err.Error())
		util.Respond(w, resp)
		return
	}

	_, err = x.rewardOrderService.CalculateRewardOrder(storeId, dropshipId, orderId)

	if err != nil {
		log.Print(err)
		resp = util.Message(false, "Unable to calculate reward points - "+err.Error())
		util.Respond(w, resp)
		return
	}

	_, err = x.rewardCustomerService.CalculateRewardCustomer(dropshipId, orderId)

	if err != nil {
		log.Print(err)
		resp = util.Message(false, "Unable to calculate reward points - "+err.Error())
		util.Respond(w, resp)
		return
	}

	resp = util.Message(true, "Reward points calculated successfully")
	util.Respond(w, resp)
}
