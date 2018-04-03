package api

import (
	"encoding/json"
	"net/http"

	"github.com/keiwi/utils"
	"github.com/keiwi/utils/models"
	"github.com/keiwi/web/auth"
)

// UserJSON - json data expected for login/signup
type UserJSON struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UserSignup - Handler for signing up a user
func (api *API) UserSignup(w http.ResponseWriter, res *http.Request) {
	decoder := json.NewDecoder(res.Body)
	jsondata := UserJSON{}
	err := decoder.Decode(&jsondata)

	if err != nil || jsondata.Username == "" || jsondata.Password == "" {
		http.Error(w, "Missing username or password", http.StatusBadRequest)
		return
	}

	if api.handler.Users.HasUser(jsondata.Username) {
		http.Error(w, "Username already exists", http.StatusBadRequest)
		return
	}

	if api.handler.Users.HasUserByEmail(jsondata.Email) {
		http.Error(w, "A user with this email already exists", http.StatusBadRequest)
		return
	}

	user, err := api.handler.Users.AddUser(jsondata.Username, jsondata.Email, jsondata.Password)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		utils.Log.WithField("error", err).Error("error adding user")
		return
	}
	jsontoken := auth.GetJSONToken(user)

	w.Header().Set("Conent-Type", "application/json")
	if _, err := w.Write([]byte(jsontoken)); err != nil {
		utils.Log.Fatal(err.Error())
	}
}

// UserLogin - Handler for login a user
func (api *API) UserLogin(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	jsondata := UserJSON{}
	err := decoder.Decode(&jsondata)

	if err != nil || jsondata.Username == "" || jsondata.Password == "" {
		http.Error(w, "Missing username or password", http.StatusBadRequest)
		return
	}

	user, err := api.handler.Users.FindUser(jsondata.Username)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		utils.Log.WithField("error", err).Error("error adding user")
		return
	}

	if user.Username == "" {
		user, err = api.handler.Users.FindUserByEmail(jsondata.Username)
		if err != nil {
			http.Error(w, "Internal error", http.StatusInternalServerError)
			utils.Log.WithField("error", err).Error("error adding user")
			return
		}

		if user.Username == "" {
			http.Error(w, "Username or email not found", http.StatusBadRequest)
			return
		}
	}

	if !api.handler.Users.CheckPassword(user.Password, jsondata.Password) {
		http.Error(w, "Bad password", http.StatusBadRequest)
		return
	}

	jsontoken := auth.GetJSONToken(user)
	w.Header().Set("Conent-Type", "application/json")
	if _, err := w.Write([]byte(jsontoken)); err != nil {
		utils.Log.Fatal(err.Error())
	}
}

// GetUserFromContext - return User reference from header token
func (api *API) GetUserFromContext(req *http.Request) *models.User {
	/*userclaims := auth.GetUserClaimsFromContext(req)
	user := api.handler.Users.FindUserByUUID(userclaims["uuid"].(string))
	return user */
	return nil
}

// UserInfo - example to get
func (api *API) UserInfo(w http.ResponseWriter, req *http.Request) {
	user := api.GetUserFromContext(req)
	js, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(js); err != nil {
		utils.Log.Fatal(err.Error())
	}
}
