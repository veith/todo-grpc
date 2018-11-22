package task

import (
	"github.com/oklog/ulid"
	"math/rand"
	"strings"
	"time"
	"upper.io/db.v3"
)

// Anfrageoptionen für upper
type QueryOptions struct {
	Fields  string `json:"fields,omitempty"`
	Sort    string `json:"sort,omitempty"`
	Filter  string `json:"filter,omitempty"`
	Count   bool   `json:"count,omitempty"`
	Sum     string `json:"sum,omitempty"`
	Context string `json:"context,omitempty"`
	Limit   uint   `json:"limit,omitempty"`
	Page    uint   `json:"page,omitempty"`
	Cursor  uint   `json:"cursor,omitempty"` // for cursor pagination
}

type DBMeta struct {
	Count       uint64
	CurrentPage uint
	NextPage    uint
	PrevPage    uint
	FirstPage   uint
	LastPage    uint
}

type Task struct {
	Id          string `json:"id,omitempty" db:"id,pk,omitempty"`
	Title       string `json:"title,omitempty" db:"title,omitempty"`
	Description string `json:"description,omitempty" db:"description,omitempty"`
	Completed   int32  `json:"completed,omitempty" db:"completed"`
}

// Erzeuge eine ULID
func GenerateULID() ulid.ULID {
	t := time.Now().UTC()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	newID, _ := ulid.New(ulid.Timestamp(t), entropy)
	return newID
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
