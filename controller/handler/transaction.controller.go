package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"superapp/controller"
	"superapp/middlewares"
	"superapp/models"
	"superapp/usecase"
	"superapp/utils/httpparse"
	"superapp/utils/httpresponse"
	"superapp/utils/status"
)

type trxController struct {
	router    	*mux.Router
	parseJson 	*httpparse.JsonParse
	responder 	httpresponse.IResponder
	service 	usecase.ITransactionUsecase
}

func (t *trxController) InitRoute(mdw ...mux.MiddlewareFunc) {
	u := t.router.PathPrefix("/transaction").Subrouter()
	u.Use(mdw...)
	u.HandleFunc("/transfer", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(t.Transfer))).Methods("POST")
}

func NewTrxController(router *mux.Router, parse *httpparse.JsonParse, responder httpresponse.IResponder, service usecase.ITransactionUsecase) controller.IDelivery {
	return &trxController{router, parse, responder, service}
}

func (t trxController) Transfer(w http.ResponseWriter, r *http.Request)  {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.responder.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	var transfer models.Transaction

	err = json.Unmarshal(body, &transfer)
	if err != nil {
		t.responder.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	trx,errTrx := t.service.Transfer(&transfer)
	if errTrx != nil {
		t.responder.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	t.responder.Data(w, http.StatusCreated, status.StatusText(status.CREATED), trx)

}
