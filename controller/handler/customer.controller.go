package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"superapp/controller"
	"superapp/middlewares"
	"superapp/models/dto"
	"superapp/usecase"
	"superapp/utils/httpparse"
	"superapp/utils/httpresponse"
	"superapp/utils/status"
)

type CustomerController struct {
	router         *mux.Router
	parseJson      *httpparse.JsonParse
	responder      httpresponse.IResponder
	service        usecase.ICustomerUsecase
}

func (c *CustomerController) InitRoute(mdw ...mux.MiddlewareFunc) {
	u := c.router.PathPrefix("/customers").Subrouter()
	u.Use(mdw...)
	u.HandleFunc("", middlewares.SetMiddlewareJSON(c.CreateCustomer)).Methods("POST")
	u.HandleFunc("/auth/login", middlewares.SetMiddlewareJSON(c.LoginCustomer)).Methods("POST")
	u.HandleFunc("/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(c.EditCustomer))).Methods("PUT")
	u.HandleFunc("/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(c.FindByIdCustomer))).Methods("GET")
	u.HandleFunc("/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(c.DeleteCustomer))).Methods("DELETE")
}

func NewCustomerController(router *mux.Router, parse *httpparse.JsonParse, responder httpresponse.IResponder, service usecase.ICustomerUsecase) controller.IDelivery {
	return &CustomerController{
		router,
		parse,
		responder,
		service,
	}
}

func (c *CustomerController) CreateCustomer(w http.ResponseWriter, r *http.Request)  {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		c.responder.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	var customer dto.CustomerRequest
	err = json.Unmarshal(body, &customer)
	if err != nil {
		c.responder.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	custCreated, errRegist := c.service.CreateCustomer(&customer)
	if errRegist != nil {
		c.responder.Error(w, http.StatusConflict, err.Error())
		return
	}

	c.responder.Data(w, http.StatusCreated, status.StatusText(status.CREATED), custCreated)
}

func (c *CustomerController) LoginCustomer(w http.ResponseWriter, r *http.Request)  {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		c.responder.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	var customer dto.LoginRequest
	err = json.Unmarshal(body, &customer)
	if err != nil {
		c.responder.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	login, errLogin := c.service.LoginCustomer(&customer)
	if errLogin != nil {
		c.responder.Error(w, http.StatusNotFound, errLogin.Error())
		return
	}
	c.responder.Data(w, http.StatusOK, status.StatusText(status.Success), login)
}

func (c *CustomerController) EditCustomer(w http.ResponseWriter, r *http.Request)  {
	param := mux.Vars(r)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		c.responder.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	var customer dto.CustomerRequest
	err = json.Unmarshal(body, &customer)
	if err != nil {
		c.responder.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	cust, errCust := c.service.UpdateCustomer(&customer, param["id"])
	if errCust != nil {
		c.responder.Error(w, http.StatusBadRequest, err.Error())
	}

	c.responder.Data(w, http.StatusCreated, status.StatusText(status.CREATED), cust)
}

func (c *CustomerController) DeleteCustomer(w http.ResponseWriter, r *http.Request)  {
	param := mux.Vars(r)
	customerId, err := c.service.DeleteCustomer(param["id"])
	if err != nil {
		c.responder.Error(w, http.StatusNotFound, err.Error())
		return
	}
	c.responder.Data(w, http.StatusOK, http.StatusText(status.Success), customerId)
}

func (c *CustomerController) FindByIdCustomer(w http.ResponseWriter, r *http.Request)  {
	param := mux.Vars(r)
	customerId, err := c.service.FindByIdCustomer(param["id"])
	if err != nil {
		c.responder.Error(w, http.StatusNotFound, err.Error())
		return
	}
	c.responder.Data(w, http.StatusOK, http.StatusText(status.Success), customerId)
}