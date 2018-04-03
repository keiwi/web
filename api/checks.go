package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/keiwi/utils"
	"github.com/keiwi/utils/models"
)

// DeleteCheck deletes a specific check from the database
func (api *API) DeleteCheck(w http.ResponseWriter, res *http.Request) {
	decoder := json.NewDecoder(res.Body)
	jsondata := IDOptions{}

	w.Header().Set("Content-Type", "application/json")
	if err := decoder.Decode(&jsondata); err != nil {
		utils.Log.Error(err.Error())
		outputJSON(w, false, "An internal error occured", nil)
		return
	}

	err := api.handler.Checks.DeleteWithID(jsondata.ID)
	if err != nil {
		utils.Log.Error(err.Error())
		outputJSON(w, false, "An internal error occured when deleting the check", nil)
		return
	}

	outputJSON(w, true, "Successfully deleted the check", nil)
}

// GetChecks returns an array of all the checks in the database
func (api *API) GetChecks(w http.ResponseWriter, res *http.Request) {
	checks, err := api.handler.Checks.FindAll()

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		utils.Log.Error(err.Error())
		outputJSON(w, false, "An internal error occured when trying to find all checks", nil)
		return
	}

	js, _ := json.Marshal(checks)
	if _, err := w.Write(js); err != nil {
		utils.Log.Fatal(err.Error())
	}
}

// GetCheckWithID returns a check if ID exists
func (api *API) GetCheckWithID(w http.ResponseWriter, res *http.Request) {
	decoder := json.NewDecoder(res.Body)
	jsondata := IDOptions{}

	w.Header().Set("Content-Type", "application/json")
	if err := decoder.Decode(&jsondata); err != nil {
		utils.Log.Error(err.Error())
		outputJSON(w, false, "An internal error occured", nil)
		return
	}

	check, err := api.handler.Checks.Find(jsondata.ID)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		utils.Log.Error(err.Error())
		outputJSON(w, false, "An internal error occured when trying to find the check", nil)
		return
	}

	js, _ := json.Marshal(check)
	if _, err := w.Write(js); err != nil {
		utils.Log.Error(err.Error())
	}
}

// ClientCommandID is the expected data when trying to find checks with specific client id and command id
type ClientCommandID struct {
	ClientID  string   `json:"client_id"`
	CommandID []string `json:"command_id"`
}

// GetWithClientIDAndCommandID tries to find checks with specific client id and command id
func (api *API) GetWithClientIDAndCommandID(w http.ResponseWriter, res *http.Request) {
	decoder := json.NewDecoder(res.Body)
	jsondata := ClientCommandID{}
	w.Header().Set("Content-Type", "application/json")
	if err := decoder.Decode(&jsondata); err != nil {
		utils.Log.Error(err.Error())
		outputJSON(w, false, "An internal error occured", nil)
		return
	}

	var checks []models.Check
	for _, cmd := range jsondata.CommandID {
		dataChecks, err := api.handler.Checks.FindWithClientIDAndCommandID(jsondata.ClientID, cmd)
		if err != nil {
			utils.Log.Error(err.Error())
			outputJSON(w, false, "An internal error occured when trying to find the checks", nil)
			return
		}
		if len(dataChecks) >= 1 {
			checks = append(checks, dataChecks[0])
		}
	}

	js, _ := json.Marshal(checks)
	if _, err := w.Write(js); err != nil {
		utils.Log.Fatal(err.Error())
	}
}

// ChecksBetweenDateClient is the expected data when trying to find checks between dates with a specific client
type ChecksBetweenDateClient struct {
	ClientID  string `json:"client_id"`
	CommandID string `json:"command_id"`
	From      string `json:"from"`
	To        string `json:"to"`
	Max       int    `json:"max"`
}

// GetWithChecksBetweenDateClient tries to find checks between dates with a specific client
func (api *API) GetWithChecksBetweenDateClient(w http.ResponseWriter, res *http.Request) {
	decoder := json.NewDecoder(res.Body)
	jsondata := ChecksBetweenDateClient{}
	w.Header().Set("Content-Type", "application/json")
	if err := decoder.Decode(&jsondata); err != nil {
		utils.Log.Error(err.Error())
		outputJSON(w, false, "An internal error occured", nil)
		return
	}

	// TODO: Check the time layout
	from, err := time.Parse("2006-01-02 15:04:05", jsondata.From)
	if err != nil {
		utils.Log.Error(err.Error())
		outputJSON(w, false, "An internal error occured", nil)
		return
	}

	to, err := time.Parse("2006-01-02 15:04:05", jsondata.To)
	if err != nil {
		utils.Log.Error(err.Error())
		outputJSON(w, false, "An internal error occured", nil)
		return
	}

	checks, err := api.handler.Checks.GetChecksBetweenDateClient(from, to, jsondata.CommandID, jsondata.ClientID, jsondata.Max)
	if err != nil {
		utils.Log.Error(err.Error())
		outputJSON(w, false, "An internal error occured", nil)
		return
	}

	js, _ := json.Marshal(checks)
	if _, err := w.Write(js); err != nil {
		utils.Log.Fatal(err.Error())
	}
}

// SendManualCheck tries to send a manual check to a specific client
func (api *API) SendManualCheck(w http.ResponseWriter, res *http.Request) {
	// decoder := json.NewDecoder(res.Body)

}

// EditClient modifies an existing client in the database
/* func (api *API) EditClient(w http.ResponseWriter, res *http.Request) {
	decoder := json.NewDecoder(res.Body)
	jsondata := EditJSON{}

	w.Header().Set("Content-Type", "application/json")
	if err := decoder.Decode(&jsondata); err != nil {
		log.Print(err)
		outputJSON(w, false, "An internal error occured", nil)
		return
	}

	client, err := api.clients.Find(jsondata.ID)
	if err != nil {
		log.Print(err)
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
		log.Print(err)
		outputJSON(w, false, "An internal error occured when saving the client", nil)
		return
	}
	outputJSON(w, true, "Successfully saved the changes for the client", client)
} */
