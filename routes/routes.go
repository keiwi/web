package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/keiwi/utils"
	"github.com/keiwi/web/api"
	"github.com/keiwi/web/auth"
	"github.com/urfave/negroni"
)

// NewRoutes builds the routes for the api
func NewRoutes(api *api.API) *mux.Router {
	utils.Log.Debug("Initializing API routes")
	mux := mux.NewRouter()

	// Static files
	mux.PathPrefix("/static").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("./client/dist/static"))))

	// api
	a := mux.PathPrefix("/api/v1").Subrouter()

	// users
	u := a.PathPrefix("/user").Subrouter()
	u.HandleFunc("/signup", api.UserSignup).Methods("POST")
	u.HandleFunc("/login", api.UserLogin).Methods("POST")
	u.Handle("/info", negroni.New(
		negroni.HandlerFunc(auth.JwtMiddleware.HandlerWithNext),
		negroni.Wrap(http.HandlerFunc(api.UserInfo)),
	)) // Example where the user need to be logged in

	// commands
	cmd := a.PathPrefix("/commands").Subrouter()
	cmd.HandleFunc("/create", api.CreateCommand).Methods("POST")
	cmd.HandleFunc("/delete", api.DeleteCommand).Methods("POST")
	cmd.HandleFunc("/edit", api.EditCommand).Methods("POST")
	cmd.HandleFunc("/get", api.GetCommands).Methods("POST")

	// groups
	g := a.PathPrefix("/groups").Subrouter()
	g.HandleFunc("/create", api.CreateGroup).Methods("POST")
	g.HandleFunc("/delete/id", api.DeleteGroupID).Methods("POST")
	g.HandleFunc("/delete/name", api.DeleteGroupName).Methods("POST")
	g.HandleFunc("/edit", api.EditGroup).Methods("POST")
	g.HandleFunc("/get", api.GetGroups).Methods("POST")
	g.HandleFunc("/exists", api.HasGroup).Methods("POST")

	// clients
	cl := a.PathPrefix("/clients").Subrouter()
	cl.HandleFunc("/create", api.CreateClient).Methods("POST")
	cl.HandleFunc("/delete", api.DeleteClient).Methods("POST")
	cl.HandleFunc("/edit", api.EditClient).Methods("POST")
	cl.HandleFunc("/get/all", api.GetClients).Methods("POST")
	cl.HandleFunc("/get/id", api.GetClientWithID).Methods("POST")

	// checks
	ch := a.PathPrefix("/checks").Subrouter()
	ch.HandleFunc("/delete", api.DeleteCheck).Methods("POST")
	ch.HandleFunc("/get/all", api.GetChecks).Methods("POST")
	ch.HandleFunc("/get/id", api.GetCheckWithID).Methods("POST")
	ch.HandleFunc("/get/client-cmd", api.GetWithClientIDAndCommandID).Methods("POST")
	ch.HandleFunc("/get/checks-date-client", api.GetWithChecksBetweenDateClient).Methods("POST")

	// catch-all for serving index.html
	mux.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./client/dist/index.html")
	})

	return mux
}
