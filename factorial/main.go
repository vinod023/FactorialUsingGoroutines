package main

import (
	"encoding/json"
	"errors"
	"factorial/models"
	"factorial/util"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.POST("/calculate", calculate)

	log.Fatal(http.ListenAndServe(":8989", router))
}

func calculate(responseWriter http.ResponseWriter, request *http.Request, ps httprouter.Params) {
	var requestBody models.InputData
	factorialOfA := make(chan uint64)
	factorialOfB := make(chan uint64)

	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil || isIncorrectValues(requestBody) {
		util.RespondError(responseWriter, http.StatusBadRequest, errors.New("Incorrect input").Error())
		return
	}

	//To get factorial of A value
	go getFactorial(requestBody.A, factorialOfA)

	//To get factorial of B value
	go getFactorial(requestBody.B, factorialOfB)

	response := models.Product{
		A: <-factorialOfA,
		B: <-factorialOfB,
	}
	util.RespondJSON(responseWriter, http.StatusOK, response)
}

func getFactorial(value int, factorial chan uint64) {

	var computation uint64 = 1

	for i := 1; i <= value; i++ {
		computation *= uint64(i)

	}
	factorial <- computation
}

func isIncorrectValues(val models.InputData) bool {
	//In future we can do more validations here
	if val.A <= 0 || val.B <= 0 {
		return true
	}
	return false
}
