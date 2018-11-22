package task

import (
	"../task"
	"encoding/json"
	"github.com/gogo/protobuf/types"
	"strconv"
)

type Hateoas struct {
	Links []*Link
}

func MapTaskToProtoTask(ob1 *task.Task) *Task {
	var t types.Timestamp
	var q struct{}
	ob2 := Task{ob1.Id, ob1.Title, ob1.Description, Complete(ob1.Completed), &t, &t, q, []byte{}, 0}
	return &ob2
}
func MapProtoTaskToTask(ob1 *Task) *task.Task {
	ob2 := task.Task{ob1.Id, ob1.Title, ob1.Description, int32(ob1.Completed)}
	return &ob2
}

// links einem HTS hinzufügen
func (h *Hateoas) AddLink(rel, contenttype, href string, method Link_Method) {
	self := Link{Rel: rel, Href: href, Type: contenttype, Method: method}
	h.Links = append(h.Links, &self)
}

// Optionen für Listenelemente kommen aus dem proto als beliebiger Typ daher, jedoch immer in der gleichen nummerierung
// diese werden in die QueryOptions Form gebracht, damit upper sauber damit umgehen kann.
func GetListOptionsFromRequest(options interface{}) task.QueryOptions {
	tmp, _ := json.Marshal(options)
	var opts task.QueryOptions
	json.Unmarshal(tmp, &opts)
	return opts
}

// hateoas anhand DBMEta für eine Collection erzeugen
func GenerateCollectionHATEOAS(dbMeta task.DBMeta) Hateoas {
	//todo Link_Get,.. nach REST schieben
	var h Hateoas
	h.AddLink("self", "application/json", "http://localhost:8080/tasks?page="+strconv.FormatUint(uint64(dbMeta.CurrentPage), 10), Link_GET)
	if dbMeta.PrevPage != 0 {
		h.AddLink("prev", "application/json", "http://localhost:8080/tasks?page="+strconv.FormatUint(uint64(dbMeta.CurrentPage-1), 10), Link_GET)
	}
	if dbMeta.NextPage != 0 {
		h.AddLink("next", "application/json", "http://localhost:8080/tasks?page="+strconv.FormatUint(uint64(dbMeta.CurrentPage+1), 10), Link_GET)
	}
	h.AddLink("first", "application/json", "http://localhost:8080/tasks?page="+strconv.FormatUint(uint64(dbMeta.FirstPage+1), 10), Link_GET)
	h.AddLink("last", "application/json", "http://localhost:8080/tasks?page="+strconv.FormatUint(uint64(dbMeta.LastPage), 10), Link_GET)
	h.AddLink("create", "application/json", "http://localhost:8080/tasks", Link_POST)
	return h
}

func GenerateEntityHateoas(id string) Hateoas {
	//todo check gegen spec machen
	var h Hateoas
	h.AddLink("self", "application/json", "http://localhost:8080/tasks/"+id, Link_GET)
	h.AddLink("delete", "application/json", "http://localhost:8080/tasks/"+id, Link_DELETE)
	h.AddLink("update", "application/json", "http://localhost:8080/tasks/"+id, Link_PATCH)
	h.AddLink("parent", "application/json", "http://localhost:8080/tasks", Link_GET)
	h.AddLink("complete", "application/json", "http://localhost:8080/tasks"+id+":complete", Link_POST)
	return h
}
