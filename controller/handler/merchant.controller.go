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
	"superapp/utils/formaterror"
	"superapp/utils/httpparse"
	"superapp/utils/httpresponse"
	"superapp/utils/status"
)

type merchantController struct {
	router    *mux.Router
	parseJson *httpparse.JsonParse
	responder httpresponse.IResponder
	service usecase.IMerchantUsecase
}

func (m *merchantController) InitRoute(mdw ...mux.MiddlewareFunc) {
	u := m.router.PathPrefix("/merchants").Subrouter()
	u.Use(mdw...)
	u.HandleFunc("", middlewares.SetMiddlewareJSON(m.CreateMerchant)).Methods("POST")
	u.HandleFunc("", middlewares.SetMiddlewareJSON(m.UpdatedMerchant)).Methods("PUT")
	u.HandleFunc("", middlewares.SetMiddlewareJSON(m.GetAllMerchant)).Methods("GET")
	u.HandleFunc("/{id}", middlewares.SetMiddlewareJSON(m.DeleteMerchant)).Methods("DELETE")
}

func NewMerchantController(router *mux.Router, parse *httpparse.JsonParse, responder httpresponse.IResponder, service usecase.IMerchantUsecase) controller.IDelivery {
	return &merchantController{
		router,
		parse,
		responder,
		service,
	}
}

func (m *merchantController) CreateMerchant(w http.ResponseWriter, r *http.Request)  {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		m.responder.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	var merchant models.Merchant

	err = json.Unmarshal(body, &merchant)
	if err != nil {
		m.responder.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	createMerchant, errMerchant := m.service.CreateMerchant(&merchant)
	if errMerchant != nil {
		formattedError := formaterror.FormatError(err.Error())
		m.responder.Error(w, http.StatusBadRequest, formattedError.Error())
		return
	}
	m.responder.Data(w, http.StatusCreated, status.StatusText(status.CREATED), createMerchant)
}

func (m *merchantController) UpdatedMerchant(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		m.responder.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	var merchant models.Merchant

	err = json.Unmarshal(body, &merchant)
	if err != nil {
		m.responder.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	editMerchant, errMerchant := m.service.CreateMerchant(&merchant)
	if errMerchant != nil {
		formattedError := formaterror.FormatError(err.Error())
		m.responder.Error(w, http.StatusBadRequest, formattedError.Error())
		return
	}
	m.responder.Data(w, http.StatusCreated, status.StatusText(status.CREATED), editMerchant)
}

func (m *merchantController) GetAllMerchant(w http.ResponseWriter, r *http.Request)  {
	param := r.URL.Query()
	merchants, err := m.service.GetAllMerchant(param.Get("page"),param.Get("size"))
	if err != nil {
		m.responder.Error(w, http.StatusNotFound, err.Error())
		return
	}
	m.responder.Data(w, status.Success, status.StatusText(http.StatusOK), merchants)
}

func (m merchantController) DeleteMerchant(w http.ResponseWriter, r *http.Request)  {
	param := mux.Vars(r)
	err := m.service.DeleteMerchant(param["id"])
	if err != nil {
		m.responder.Error(w, http.StatusNotFound, err.Error())
		return
	}
	m.responder.Data(w, status.Success, status.StatusText(http.StatusOK), "Data has been Deleted")

}
