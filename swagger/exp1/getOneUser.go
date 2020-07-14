// Package classification User API.
//
// The purpose of this service is to provide an application
// that is using plain go code to define an API
//
//      Host: localhost
//      Version: 0.0.1
//
// swagger:meta
package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"user.server/models"
	"github.com/Sirupsen/logrus"
)
// swagger:parameters getSingleUser
type GetUserParam struct {
	// an id of user info
	//
	// Required: true
	// in: path
	Id int `json:"id"`
}

func GetOneUser(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /users/{id} users getSingleUser
	//
	// get a user by userID
	//
	// This will show a user info
	//
	//     Responses:
	//       200: UserResponse
	decoder := json.NewDecoder(r.Body)
	var param GetUserParam
	err := decoder.Decode(&param)
	if err != nil {
		WriteResponse(w, ErrorResponseCode, "request param is invalid, please check!", nil)
		return
	}

	// get user from db
	user, err := models.GetOne(strconv.Itoa(param.Id))
	if err != nil {
		logrus.Warn(err)
		WriteResponse(w, ErrorResponseCode, "failed", nil)
		return
	}
	WriteResponse(w, SuccessResponseCode, "success", user)
}


// User Info
//
// swagger:response UserResponse
type UserWapper struct {
	// in: body
	Body ResponseMessage
}

type ResponseMessage struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
