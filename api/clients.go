package api

import (
	"encoding/json"
	"net/http"

	"github.com/keiwi/utils"
	"github.com/keiwi/web/models"
)

// ClientJSON - json data expected for creating a new client
type ClientJSON struct {
	GroupNames string `json:"group_names"`
	IP         string `json:"ip"`
	Namn       string `json:"namn"`
}

// CreateClient - Handler for creating a new client
func (api *API) CreateClient(w http.ResponseWriter, res *http.Request) {
	decoder := json.NewDecoder(res.Body)
	jsondata := ClientJSON{}

	w.Header().Set("Content-Type", "application/json")
	if err := decoder.Decode(&jsondata); err != nil {
		utils.Log.Error(err.Error())
		outputJSON(w, false, "An internal error occured", nil)
		return
	}

	if jsondata.IP == "" {
		outputJSON(w, false, "IP is missing", nil)
		return
	}

	if jsondata.Namn == "" {
		outputJSON(w, false, "Namn is missing", nil)
		return
	}

	client := &models.Client{
		GroupNames: jsondata.GroupNames,
		Namn:       jsondata.Namn,
		IP:         jsondata.IP,
	}

	if err := api.clients.Create(client); err != nil {
		utils.Log.Error(err.Error())
		outputJSON(w, false, "An internal error occured", nil)
		return
	}

	outputJSON(w, true, "Successfully created the client", client)
}

// DeleteClient deletes a specific client from the database
func (api *API) DeleteClient(w http.ResponseWriter, res *http.Request) {
	decoder := json.NewDecoder(res.Body)
	jsondata := IDOptions{}

	w.Header().Set("Content-Type", "application/json")
	if err := decoder.Decode(&jsondata); err != nil {
		utils.Log.Error(err.Error())
		outputJSON(w, false, "An internal error occured", nil)
		return
	}

	err := api.clients.DeleteWithID(jsondata.ID)
	if err != nil {
		utils.Log.Error(err.Error())
		outputJSON(w, false, "An internal error occured when deleting the client", nil)
		return
	}
	outputJSON(w, true, "Successfully deleted the client", nil)
}

// GetClients returns an array of all the clients in the database
func (api *API) GetClients(w http.ResponseWriter, res *http.Request) {
	clients, err := api.clients.FindAll()

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		utils.Log.Error(err.Error())
		outputJSON(w, false, "An internal error occured when trying to find all clients", nil)
		return
	}

	js, _ := json.Marshal(clients)
	if _, err := w.Write(js); err != nil {
		utils.Log.Fatal(err.Error())
	}
}

// GetClientWithID returns a client if ID exists
func (api *API) GetClientWithID(w http.ResponseWriter, res *http.Request) {
	decoder := json.NewDecoder(res.Body)
	jsondata := IDOptions{}

	w.Header().Set("Content-Type", "application/json")
	if err := decoder.Decode(&jsondata); err != nil {
		utils.Log.Error(err.Error())
		outputJSON(w, false, "An internal error occured", nil)
		return
	}

	client, err := api.clients.Find(jsondata.ID)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		utils.Log.Error(err.Error())
		outputJSON(w, false, "An internal error occured when trying to find the client", nil)
		return
	}

	js, _ := json.Marshal(client)
	if _, err := w.Write(js); err != nil {
		utils.Log.Fatal(err.Error())
	}
}

// EditClient modifies an existing client in the database
func (api *API) EditClient(w http.ResponseWriter, res *http.Request) {
	decoder := json.NewDecoder(res.Body)
	jsondata := EditJSON{}

	w.Header().Set("Content-Type", "application/json")
	if err := decoder.Decode(&jsondata); err != nil {
		utils.Log.Error(err.Error())
		outputJSON(w, false, "An internal error occured", nil)
		return
	}

	client, err := api.clients.Find(jsondata.ID)
	if err != nil {
		utils.Log.Error(err.Error())
		outputJSON(w, false, "Can't find a client with this ID", nil)
		return
	}

	v, ok := jsondata.Value.(string)
	if !ok {
		outputJSON(w, false, "Value is not a string", nil)
		return
	}

	switch jsondata.Option {
	case "group_names", "GroupNames":
		client.GroupNames = v
	case "name", "Name", "namn", "Namn":
		client.Namn = v
	case "ip", "IP":
		client.IP = v
	default:
		outputJSON(w, false, "Please provide a correct column", nil)
		return
	}

	if err = api.clients.Save(client); err != nil {
		utils.Log.Error(err.Error())
		outputJSON(w, false, "An internal error occured when saving the client", nil)
		return
	}
	outputJSON(w, true, "Successfully saved the changes for the client", client)
}
