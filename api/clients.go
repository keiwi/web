package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/keiwi/utils"
	"github.com/keiwi/utils/models"
	"gopkg.in/mgo.v2/bson"
)

// ClientJSON - json data expected for creating a new client
type ClientJSON struct {
	IP   string `json:"ip"`
	Name string `json:"name"`
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

	if jsondata.Name == "" {
		outputJSON(w, false, "Name is missing", nil)
		return
	}

	client := &models.Client{
		Name: jsondata.Name,
		IP:   jsondata.IP,
	}

	if err := api.handler.Clients.Create(client); err != nil {
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

	err := api.handler.Clients.DeleteWithID(jsondata.ID)
	if err != nil {
		utils.Log.Error(err.Error())
		outputJSON(w, false, "An internal error occured when deleting the client", nil)
		return
	}

	outputJSON(w, true, "Successfully deleted the client", nil)
}

// GetClients returns an array of all the clients in the database
func (api *API) GetClients(w http.ResponseWriter, res *http.Request) {
	clients, err := api.handler.Clients.FindAll()

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

	client, err := api.handler.Clients.Find(jsondata.ID)

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

	client, err := api.handler.Clients.Find(jsondata.ID)
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

	updates := bson.M{}
	switch jsondata.Option {
	case "name", "Name", "namn", "Namn":
		updates["name"] = v
		client.Name = v
	case "ip", "IP":
		updates["ip"] = v
		client.IP = v
	case "groups", "Groups":
		del, add := seperateGroups(objectIDArrayToString(client.GroupIDs), v)
		for _, d := range del {
			if d == "" {
				continue
			}
			for i, g := range client.GroupIDs {
				id := g.Hex()
				if d == id {
					client.GroupIDs = append(client.GroupIDs[:i], client.GroupIDs[i+1:]...)
					break
				}
			}
		}
		for _, a := range add {
			if a == "" {
				continue
			}
			group, err := api.handler.Groups.Find(a)
			if err != nil || group == nil {
				utils.Log.Error(err.Error())
				outputJSON(w, false, "Can't find a group with the id "+a, nil)
				return
			}
			client.GroupIDs = append(client.GroupIDs, bson.ObjectIdHex(a))
		}
		updates["group_ids"] = client.GroupIDs
	default:
		outputJSON(w, false, "Please provide a correct column", nil)
		return
	}

	if err = api.handler.Clients.Save(jsondata.ID, utils.Updates{"$set": updates}); err != nil {
		utils.Log.Error(err.Error())
		outputJSON(w, false, "An internal error occured when saving the client", nil)
		return
	}
	outputJSON(w, true, "Successfully saved the changes for the client", client)
}

func objectIDArrayToString(list []bson.ObjectId) string {
	length := len(list)
	out := ""
	for i, v := range list {
		if i >= length-1 {
			out += v.Hex()
		} else {
			out += v.Hex() + ", "
		}
	}
	return out
}

func seperateGroups(old, new string) (deleted, added []string) {
	splitOld := strings.Split(old, ",")
	splitNew := strings.Split(new, ",")

	deleted = findDifference(splitOld, splitNew)
	added = findDifference(splitNew, splitOld)
	return
}

func findDifference(a1, a2 []string) []string {
	var out []string
	for _, i1 := range a1 {
		m := false
		for _, i2 := range a2 {
			if i1 == i2 {
				m = true
				break
			}
		}
		if m {
			continue
		}
		out = append(out, i1)
	}
	return out
}
