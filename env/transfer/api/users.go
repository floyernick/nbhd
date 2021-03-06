package api

import (
	"nbhd/models/request"
	"net/http"
)

func (api API) UsersSignUp(w http.ResponseWriter, r *http.Request) {

	var req request.UsersSignUp

	err := ProcessRequest(r, &req, api.checkHash, api.hashSalt)

	if err != nil {
		ReturnErrorResponse(w, err.Error())
		return
	}

	res, err := api.controller.UsersSignUp(req)

	if err != nil {
		ReturnErrorResponse(w, err.Error())
		return
	}

	ReturnSuccessResponse(w, res)

}

func (api API) UsersSignIn(w http.ResponseWriter, r *http.Request) {

	var req request.UsersSignIn

	err := ProcessRequest(r, &req, api.checkHash, api.hashSalt)

	if err != nil {
		ReturnErrorResponse(w, err.Error())
		return
	}

	res, err := api.controller.UsersSignIn(req)

	if err != nil {
		ReturnErrorResponse(w, err.Error())
		return
	}

	ReturnSuccessResponse(w, res)

}

func (api API) UsersSignOut(w http.ResponseWriter, r *http.Request) {

	var req request.UsersSignOut

	err := ProcessRequest(r, &req, api.checkHash, api.hashSalt)

	if err != nil {
		ReturnErrorResponse(w, err.Error())
		return
	}

	res, err := api.controller.UsersSignOut(req)

	if err != nil {
		ReturnErrorResponse(w, err.Error())
		return
	}

	ReturnSuccessResponse(w, res)

}

func (api API) UsersGet(w http.ResponseWriter, r *http.Request) {

	var req request.UsersGet

	err := ProcessRequest(r, &req, api.checkHash, api.hashSalt)

	if err != nil {
		ReturnErrorResponse(w, err.Error())
		return
	}

	res, err := api.controller.UsersGet(req)

	if err != nil {
		ReturnErrorResponse(w, err.Error())
		return
	}

	ReturnSuccessResponse(w, res)

}

func (api API) UsersLocationUpdate(w http.ResponseWriter, r *http.Request) {

	var req request.UsersLocationUpdate

	err := ProcessRequest(r, &req, api.checkHash, api.hashSalt)

	if err != nil {
		ReturnErrorResponse(w, err.Error())
		return
	}

	res, err := api.controller.UsersLocationUpdate(req)

	if err != nil {
		ReturnErrorResponse(w, err.Error())
		return
	}

	ReturnSuccessResponse(w, res)

}
