package todos

import (
	"../protos"
	"encoding/json"
	"github.com/oklog/ulid"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"upper.io/db.v3"
)

type QueryOptions struct {
	Fields  string `json:"fields,omitempty"`
	Sort    string `json:"sort,omitempty"`
	Filter  string `json:"filter,omitempty"`
	Count   bool   `json:"count,omitempty"`
	Sum     string `json:"sum,omitempty"`
	Context string `json:"context,omitempty"`
	Limit   uint   `json:"limit,omitempty"`
	Page    uint   `json:"page,omitempty"`
}

type DBMeta struct {
	Count       uint64
	CurrentPage uint
	NextPage    uint
	PrevPage    uint
	FirstPage   uint
	LastPage    uint
}

type Hateoas struct {
	Links []*todo.Link
}

// links einem HTS hinzufügen
func (h *Hateoas) AddLink(rel, contenttype, href string, method todo.Link_Method) {
	self := todo.Link{Rel: rel, Href: href, Type: contenttype, Method: method}
	h.Links = append(h.Links, &self)
}

// Erzeuge eine ULID
func GenerateULID() ulid.ULID {
	t := time.Now().UTC()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	newID, _ := ulid.New(ulid.Timestamp(t), entropy)
	return newID
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
	//todo todo.Link_Get,.. nach REST schieben
	var h Hateoas
	h.AddLink("self", "application/json", "http://localhost:8080/todos?page="+strconv.FormatUint(uint64(dbMeta.CurrentPage), 10), todo.Link_GET)
	if dbMeta.PrevPage != 0 {
		h.AddLink("prev", "application/json", "http://localhost:8080/todos?page="+strconv.FormatUint(uint64(dbMeta.CurrentPage-1), 10), todo.Link_GET)
	}
	if dbMeta.NextPage != 0 {
		h.AddLink("next", "application/json", "http://localhost:8080/todos?page="+strconv.FormatUint(uint64(dbMeta.CurrentPage+1), 10), todo.Link_GET)
	}
	h.AddLink("first", "application/json", "http://localhost:8080/todos?page="+strconv.FormatUint(uint64(dbMeta.FirstPage+1), 10), todo.Link_GET)
	h.AddLink("last", "application/json", "http://localhost:8080/todos?page="+strconv.FormatUint(uint64(dbMeta.LastPage), 10), todo.Link_GET)
	h.AddLink("create", "application/json", "http://localhost:8080/todos", todo.Link_POST)
	return h
}

// Query Options für auf das db.Result anwenden.
// fields, sort, limit, page, sind implementiert
// mit der dbMeta kann man sich eine Pagination bauen...
func ApplyRequestOptionsToQuery(res db.Result, options QueryOptions) (db.Result, DBMeta) {
	var meta DBMeta
	if options.Limit != 0 {
		res = res.Paginate(options.Limit)
	} else {
		res = res.Paginate(paginationDefault)
	}

	if options.Fields != "" {
		fields := strings.Split(options.Fields, ",")
		s := make([]interface{}, len(fields))
		for i, field := range fields {
			s[i] = field
		}
		res = res.Select(s...)
	}

	if options.Sort != "" {
		res = res.OrderBy(options.Sort)
	}

	meta.CurrentPage = 1
	if options.Page > 0 {
		meta.CurrentPage = options.Page
		res = res.Page(options.Page)
	}
	pages, _ := res.TotalPages()
	meta.LastPage = pages
	meta.FirstPage = 1
	if meta.CurrentPage < meta.LastPage {
		meta.NextPage = meta.CurrentPage + 1
	}
	if meta.CurrentPage > 1 {
		meta.PrevPage = meta.CurrentPage - 1
	}

	if options.Count {
		meta.Count, _ = res.Count()
	}

	return res, meta
}
