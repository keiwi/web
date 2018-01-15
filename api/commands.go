package api

import (
	"encoding/json"
	"net/http"

	"github.com/keiwi/utils"
	"github.com/keiwi/web/models"
)

// CommandJSON - json data expected for creating a new command
type CommandJSON struct {
	Command     string `json:"command"`
	Namn        string `json:"namn"`
	Description string `json:"description"`
	Format      string `json:"format"`
}

// CreateCommand - Handler for creating a new command
func (api *API) CreateCommand(w http.ResponseWriter, res *http.Request) {
	decoder := json.NewDecoder(res.Body)
	jsondata := CommandJSON{}

	w.Header().Set("Content-Type", "application/json")
	if err := decoder.Decode(&jsondata); err != nil {
		utils.Log.Error(err.Error())
		outputJSON(w, false, "An internal error occured", nil)
		return
	}

	if jsondata.Command == "" {
		outputJSON(w, false, "Command is missing", nil)
		return
	}

	if jsondata.Namn == "" {
		outputJSON(w, false, "Namn is missing", nil)
		return
	}

	if jsondata.Description == "" {
		outputJSON(w, false, "Description is missing", nil)
		return
	}

	cmd := &models.Command{
		Command:     jsondata.Command,
		Namn:        jsondata.Namn,
		Description: jsondata.Description,
		Format:      jsondata.Format,
	}

	if err := api.commands.Create(cmd); err != nil {
		utils.Log.Error(err.Error())
		outputJSON(w, false, "An internal error occured", nil)
		return
	}

	outputJSON(w, true, "Successfully created the command", cmd)
}

// EditCommand modifies an existing command in the database
func (api *API) EditCommand(w http.ResponseWriter, res *http.Request) {
	decoder := json.NewDecoder(res.Body)
	jsondata := EditJSON{}

	w.Header().Set("Content-Type", "application/json")
	if err := decoder.Decode(&jsondata); err != nil {
		utils.Log.Error(err.Error())
		outputJSON(w, false, "An internal error occured", nil)
		return
	}

	cmd, err := api.commands.Find(jsondata.ID)
	if err != nil {
		utils.Log.Error(err.Error())
		outputJSON(w, false, "Can't find a command with this ID", nil)
		return
	}

	v, ok := jsondata.Value.(string)
	if !ok {
		outputJSON(w, false, "Value is not a string", nil)
		return
	}

	switch jsondata.Option {
	case "command", "Command":
		cmd.Command = v
	case "name", "Name", "namn", "Namn":
		cmd.Namn = v
	case "description", "Description":
		cmd.Description = v
	case "format", "Format":
		cmd.Format = v
	default:
		outputJSON(w, false, "Please provide a correct column", nil)
		return
	}

	if err = api.commands.Save(cmd); err != nil {
		utils.Log.Error(err.Error())
		outputJSON(w, false, "An internal error occured when saving the command", nil)
		return
	}
	outputJSON(w, true, "Successfully saved the changes for the command", cmd)
}

// DeleteCommand deletes a specific command from the database
func (api *API) DeleteCommand(w http.ResponseWriter, res *http.Request) {
	decoder := json.NewDecoder(res.Body)
	jsondata := IDOptions{}

	w.Header().Set("Content-Type", "application/json")
	if err := decoder.Decode(&jsondata); err != nil {
		utils.Log.Error(err.Error())
		outputJSON(w, false, "An internal error occured", nil)
		return
	}

	err := api.commands.DeleteWithID(jsondata.ID)
	if err != nil {
		utils.Log.Error(err.Error())
		outputJSON(w, false, "An internal error occured when deleting the command", nil)
		return
	}
	outputJSON(w, true, "Successfully deleted the command", nil)
}

// GetCommands returns an array of all the commands in the database
func (api *API) GetCommands(w http.ResponseWriter, res *http.Request) {
	cmds, err := api.commands.FindAll()

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		utils.Log.Error(err.Error())
		outputJSON(w, false, "An internal error occured when trying to find all commands", nil)
		return
	}

	js, _ := json.Marshal(cmds)
	if _, err := w.Write(js); err != nil {
		utils.Log.Fatal(err.Error())
	}
}
