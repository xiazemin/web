package service

import "user.server/models"

// swagger:parameters UpdateUserResponseWrapper
type UpdateUserRequest struct {
	// in: body
	Body models.User
}

// Update User Info
//
// swagger:response UpdateUserResponseWrapper
type UpdateUserResponseWrapper struct {
	// in: body
	Body ResponseMessage
}

// swagger:route POST /users users UpdateUserResponseWrapper
//
// Update User
//
// This will update user info
//
//     Responses:
//       200: UpdateUserResponseWrapper
