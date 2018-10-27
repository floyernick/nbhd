package api

import (
	"nbhd/models/request"
	"net/http"
)

func (api API) TasksCreate(w http.ResponseWriter, r *http.Request) {

	var req request.TasksCreate

	err := ProcessRequest(r, &req, api.checkHash, api.hashSalt)

	if err != nil {
		ReturnErrorResponse(w, err.Error())
		return
	}

	res, err := api.controller.TasksCreate(req)

	if err != nil {
		ReturnErrorResponse(w, err.Error())
		return
	}

	ReturnSuccessResponse(w, res)

}

func (api API) TasksGet(w http.ResponseWriter, r *http.Request) {

	var req request.TasksGet

	err := ProcessRequest(r, &req, api.checkHash, api.hashSalt)

	if err != nil {
		ReturnErrorResponse(w, err.Error())
		return
	}

	res, err := api.controller.TasksGet(req)

	if err != nil {
		ReturnErrorResponse(w, err.Error())
		return
	}

	ReturnSuccessResponse(w, res)

}

func (api API) TasksList(w http.ResponseWriter, r *http.Request) {

	var req request.TasksList

	err := ProcessRequest(r, &req, api.checkHash, api.hashSalt)

	if err != nil {
		ReturnErrorResponse(w, err.Error())
		return
	}

	res, err := api.controller.TasksList(req)

	if err != nil {
		ReturnErrorResponse(w, err.Error())
		return
	}

	ReturnSuccessResponse(w, res)

}

func (api API) TasksDelete(w http.ResponseWriter, r *http.Request) {

	var req request.TasksDelete

	err := ProcessRequest(r, &req, api.checkHash, api.hashSalt)

	if err != nil {
		ReturnErrorResponse(w, err.Error())
		return
	}

	res, err := api.controller.TasksDelete(req)

	if err != nil {
		ReturnErrorResponse(w, err.Error())
		return
	}

	ReturnSuccessResponse(w, res)

}
