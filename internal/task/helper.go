package task

import (
	"../proto"
	"encoding/json"
	"github.com/gogo/protobuf/types"
	"github.com/oklog/ulid"
	"strconv"
)

type Hateoas struct {
	Links []*proto.Link
}

func MapTaskToProtoTask(ob1 *Task) *proto.Task {
	var t types.Timestamp
	var q struct{}
	ob2 := proto.Task{ob1.Id.String(), ob1.Title, ob1.Description, proto.Complete(ob1.Completed), &t, &t, q, []byte{}, 0}
	return &ob2
}
func MapProtoTaskToTask(ob1 *proto.Task) *Task {
	id, _ := ulid.Parse(ob1.Id)
	ob2 := Task{id, ob1.Title, ob1.Description, int32(ob1.Completed)}
	return &ob2
}

// links einem HTS hinzufügen
func (h *Hateoas) AddLink(rel, contenttype, href string, method proto.Link_Method) {
	link := proto.Link{Rel: rel, Href: href, Type: contenttype, Method: method}
	h.Links = append(h.Links, &link)
}

// Optionen für Listenelemente kommen aus dem proto als beliebiger Typ daher, jedoch immer in der gleichen nummerierung
// diese werden in die QueryOptions Form gebracht, damit upper sauber damit umgehen kann.
func GetListOptionsFromRequest(options interface{}) QueryOptions {
	tmp, _ := json.Marshal(options)
	var opts QueryOptions
	json.Unmarshal(tmp, &opts)
	return opts
}

// hateoas anhand DBMEta für eine Collection erzeugen
func GenerateCollectionHATEOAS(dbMeta DBMeta) Hateoas {
	//todo Link_Get,.. nach REST schieben
	var h Hateoas
	h.AddLink("self", "application/json", "http://localhost:8888/tasks?page="+strconv.FormatUint(uint64(dbMeta.CurrentPage), 10), proto.Link_GET)
	if dbMeta.PrevPage != 0 {
		h.AddLink("prev", "application/json", "http://localhost:8888/tasks?page="+strconv.FormatUint(uint64(dbMeta.CurrentPage-1), 10), proto.Link_GET)
	}
	if dbMeta.NextPage != 0 {
		h.AddLink("next", "application/json", "http://localhost:8888/tasks?page="+strconv.FormatUint(uint64(dbMeta.CurrentPage+1), 10), proto.Link_GET)
	}
	h.AddLink("first", "application/json", "http://localhost:8888/tasks?page="+strconv.FormatUint(uint64(dbMeta.FirstPage), 10), proto.Link_GET)
	h.AddLink("last", "application/json", "http://localhost:8888/tasks?page="+strconv.FormatUint(uint64(dbMeta.LastPage), 10), proto.Link_GET)
	h.AddLink("create", "application/json", "http://localhost:8888/tasks", proto.Link_POST)
	return h
}

func GenerateEntityHateoas(id string) Hateoas {
	//todo check gegen spec machen
	var h Hateoas
	h.AddLink("self", "application/json", "http://localhost:8888/tasks/"+id, proto.Link_GET)
	h.AddLink("delete", "application/json", "http://localhost:8888/tasks/"+id, proto.Link_DELETE)
	h.AddLink("update", "application/json", "http://localhost:8888/tasks/"+id, proto.Link_PATCH)
	h.AddLink("parent", "application/json", "http://localhost:8888/tasks", proto.Link_GET)
	h.AddLink("complete", "application/json", "http://localhost:8888/tasks"+id+":complete", proto.Link_POST)
	return h
}
