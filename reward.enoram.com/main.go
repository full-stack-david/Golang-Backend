package main

import (
	"log"
	"mauappa-go/config"
	"mauappa-go/controller"
	"mauappa-go/repository/firestore"
	"mauappa-go/service"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	configBasePath := ""
	if len(os.Args) > 1 {
		configBasePath = os.Args[1]
	}

	log.Printf("## Config Base Path = %s", configBasePath)

	appConfig, err := config.InitConfig(configBasePath)
	if err != nil {
		log.Panicln("## Error while reading config file", err)
		return
	}

	router := mux.NewRouter()

	rewardOrderRepo, err := firestore.NewRewardOrderFirestoreRepository()
	if err != nil {
		log.Panicln("## Error while creating firestore repository", err)
	}

	invoiceRepo, err := firestore.NewInvoiceFirestoreRepository()
	if err != nil {
		log.Panicln("## Error while creating firestore repository", err)
	}

	custRepo, err := firestore.NewCustomerFirestoreRepository()
	if err != nil {
		log.Panicln("## Error while creating firestore repository", err)
	}

	storeProfileRepo, err := firestore.NewStoreProfileFirestoreRepository()
	if err != nil {
		log.Panicln("## Error while creating firestore repository", err)
	}

	rewardOrderService := service.NewRewardOrderService(rewardOrderRepo, invoiceRepo, custRepo, storeProfileRepo)

	rewardCustomerRepo, err := firestore.NewRewardCustomerFirestoreRepository()
	if err != nil {
		log.Panicln("## Error while creating firestore repository", err)
	}

	rewardCustomerService := service.NewRewardCustomerService(rewardCustomerRepo)

	rewardController := controller.NewRewardController(rewardOrderService, rewardCustomerService)

	router.HandleFunc("/_api/reward/calc", rewardController.CalculateReward).Methods("POST")

	rewardCustomerController := controller.NewRewardCustomerController(rewardCustomerService)

	router.HandleFunc("/_api/reward/{dropshipId}/customers", rewardCustomerController.GetRewardCustomerList).Methods("GET")
	router.HandleFunc("/_api/reward/{dropshipId}/customers/{customerId}", rewardCustomerController.GetRewardCustomer).Methods("GET")

	healthController := controller.NewHealthController()
	router.HandleFunc("/_api/reward/health", healthController.HealthCheck).Methods("GET")

	originsOk := handlers.AllowedOrigins([]string{os.Getenv("*")})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "OPTIONS"})

	err = http.ListenAndServe(":"+appConfig.PORT, handlers.CORS(originsOk, methodsOk)(router))
	if err != nil {
		log.Fatal("Error while initializing server", err)
	}

}
