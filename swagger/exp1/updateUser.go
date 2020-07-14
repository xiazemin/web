func UpdateUser(w http.ResponseWriter, r *http.Request) {
defer r.Body.Close()
// decode body data into user struct
decoder := json.NewDecoder(r.Body)
user := models.User{}
err := decoder.Decode(&user)
if err != nil {
WriteResponse(w, ErrorResponseCode, "user data is invalid, please check!", nil)
return
}

// check if user exists
data, err := models.GetUserById(user.Id)
if err != nil {
logrus.Warn(err)
WriteResponse(w, ErrorResponseCode, "query user failed", nil)
return
}
if data.Id == 0 {
WriteResponse(w, ErrorResponseCode, "user not exists, no need to update", nil)
return
}

// update
_, err = models.Update(user)
if err != nil {
WriteResponse(w, ErrorResponseCode, "update user data failed, please try again!", nil)
return
}
WriteResponse(w, SuccessResponseCode, "update user data success!", nil)
}
