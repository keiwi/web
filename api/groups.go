package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/keiwi/utils"
	"github.com/keiwi/web/models"
)

// RenameGroupOptions is the values that
// can be sent in the POST request body
// when renaming a group
type RenameGroupOptions struct {
	NewName string `json:"new_name"`
	OldName string `json:"old_name"`
}

// RenameGroup will rename all instances of OldName to NewName
// it will return an error if NewName already exists
func (api *API) RenameGroup(w http.ResponseWriter, res *http.Request) {
	decoder := json.NewDecoder(res.Body)
	jsondata := RenameGroupOptions{}

	w.Header().Set("Content-Type", "application/json")
	if err := decoder.Decode(&jsondata); err != nil {
		utils.Log.Error(err.Error())
		outputJSON(w, false, "An internal error occured", nil)
		return
	}

	if jsondata.NewName == "" {
		outputJSON(w, false, "Please provide a new name for the group", nil)
		return
	}
	if jsondata.OldName == "" {
		outputJSON(w, false, "Please provide the name of the group you want to rename", nil)
		return
	}

	if ok := api.groups.ExistsName(jsondata.NewName); !ok {
		outputJSON(w, false, "There is already an existing group with this name", nil)
		return
	}

	count, err := api.groups.UpdateName(jsondata.NewName, jsondata.OldName)
	if err != nil {
		utils.Log.Error(err.Error())
		outputJSON(w, false, "An internal error occured", nil)
	}

	outputJSON(w, true, fmt.Sprintf("Renamed %d group instances in the database", count), count)
}

// AddGroupOptions is the values that
// can be sent in the POST request body
// when adding a command to a group
type AddGroupOptions struct {
	GroupName string `json:"group_name"`
	CommandID int    `json:"command_id"`
	Delay     int    `json:"delay"`
	StopError bool   `json:"stop_error"`
}

// CreateGroup will add a new command to a specific group
func (api *API) CreateGroup(w http.ResponseWriter, res *http.Request) {
	decoder := json.NewDecoder(res.Body)
	jsondata := AddGroupOptions{CommandID: 1}

	w.Header().Set("Content-Type", "application/json")
	if err := decoder.Decode(&jsondata); err != nil {
		utils.Log.Error(err.Error())
		outputJSON(w, false, "An internal error occured", nil)
		return
	}

	if jsondata.CommandID == -1 {
		outputJSON(w, false, "Please provide a command id", nil)
		return
	}
	if jsondata.GroupName == "" {
		outputJSON(w, false, "Please provide a group name", nil)
		return
	}

	group := &models.Group{
		CommandID: jsondata.CommandID,
		GroupName: jsondata.GroupName,
		NextCheck: jsondata.Delay,
		StopError: jsondata.StopError,
	}

	err := api.groups.Create(group)
	if err != nil {
		utils.Log.Error(err.Error())
		outputJSON(w, false, "An internal error occured", nil)
	}

	outputJSON(w, true, "Successfully created added the command to the group", group)
}

// EditGroup will modify an existing group command
func (api *API) EditGroup(w http.ResponseWriter, res *http.Request) {
	decoder := json.NewDecoder(res.Body)
	jsondata := EditJSON{}

	w.Header().Set("Content-Type", "application/json")
	if err := decoder.Decode(&jsondata); err != nil {
		utils.Log.Error(err.Error())
		outputJSON(w, false, "An internal error occured", nil)
		return
	}

	group, err := api.groups.Find(jsondata.ID)
	if err != nil {
		utils.Log.Error(err.Error())
		outputJSON(w, false, "Can't find a command with this ID", nil)
		return
	}

	switch jsondata.Option {
	case "command_id", "CommandID":
		id, e := convertToInt(jsondata.Value)
		if e != nil {
			outputJSON(w, false, "ID is not a number", nil)
			log.Print(e)
			return
		}
		if id > 2147483647 || id < 0 {
			outputJSON(w, false, "Please send a valid number", nil)
			return
		}
		group.CommandID = int(id)
	case "next_check", "NextCheck":
		next, e := convertToInt(jsondata.Value)
		if e != nil {
			outputJSON(w, false, "Value is not a number", nil)
			log.Print(e)
			return
		}
		if next > 2147483647 || next < 0 {
			outputJSON(w, false, "Please send a valid number", nil)
			return
		}
		group.NextCheck = int(next)
	case "stop_error", "StopError":
		stop, ok := jsondata.Value.(bool)
		if !ok {
			outputJSON(w, false, "Value is not a boolean", nil)
			return
		}
		group.StopError = stop
	default:
		outputJSON(w, false, "Please provide a correct column", nil)
		return
	}

	if err = api.groups.Save(group); err != nil {
		utils.Log.Error(err.Error())
		outputJSON(w, false, "An internal error occured when saving the group", nil)
		return
	}
	outputJSON(w, true, "Successfully saved the changes for the group", group)
}

// DeleteGroupID deletes a specific command from the database
func (api *API) DeleteGroupID(w http.ResponseWriter, res *http.Request) {
	decoder := json.NewDecoder(res.Body)
	jsondata := IDOptions{}

	w.Header().Set("Content-Type", "application/json")
	if err := decoder.Decode(&jsondata); err != nil {
		utils.Log.Error(err.Error())
		outputJSON(w, false, "An internal error occured", nil)
		return
	}

	err := api.groups.DeleteWithID(jsondata.ID)
	if err != nil {
		utils.Log.Error(err.Error())
		outputJSON(w, false, "An internal error occured when deleting the command", nil)
		return
	}
	outputJSON(w, true, "Successfully deleted the command", nil)
}

// DeleteGroupNameOptions is the values that
// can be sent in the POST request body
// when deleting a command
type DeleteGroupNameOptions struct {
	Name string `json:"name"`
}

// DeleteGroupName deletes a specific command from the database
func (api *API) DeleteGroupName(w http.ResponseWriter, res *http.Request) {
	decoder := json.NewDecoder(res.Body)
	jsondata := DeleteGroupNameOptions{}

	w.Header().Set("Content-Type", "application/json")
	if err := decoder.Decode(&jsondata); err != nil {
		utils.Log.Error(err.Error())
		outputJSON(w, false, "An internal error occured", nil)
		return
	}

	if jsondata.Name == "" {
		outputJSON(w, false, "Please provide a group name", nil)
		return
	}

	err := api.groups.DeleteWithName(jsondata.Name)
	if err != nil {
		utils.Log.Error(err.Error())
		outputJSON(w, false, "An internal error occured when deleting the command", nil)
		return
	}
	outputJSON(w, true, "Successfully deleted the command", nil)
}

// GetGroups will return a JSON array of all groups in the database
func (api *API) GetGroups(w http.ResponseWriter, res *http.Request) {
	groups, err := api.groups.FindAll()

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		utils.Log.Error(err.Error())
		outputJSON(w, false, "An internal error occured when trying to find all groups", nil)
		return
	}

	js, _ := json.Marshal(groups)
	if _, err := w.Write(js); err != nil {
		utils.Log.Fatal(err.Error())
	}
}

// GroupName is the expected data when
// checking if a group exists
type GroupName struct {
	Name string `json:"name"`
}

// HasGroup will return if the group name exists in the database
func (api *API) HasGroup(w http.ResponseWriter, res *http.Request) {
	decoder := json.NewDecoder(res.Body)
	jsondata := GroupName{}

	w.Header().Set("Content-Type", "application/json")
	if err := decoder.Decode(&jsondata); err != nil {
		utils.Log.Error(err.Error())
		outputJSON(w, false, "An internal error occured", nil)
		return
	}

	if jsondata.Name == "" {
		outputJSON(w, false, "Please provide a group name", nil)
		return
	}

	js, _ := json.Marshal(api.groups.HasGroup(jsondata.Name))
	if _, err := w.Write(js); err != nil {
		utils.Log.Fatal(err.Error())
	}
}