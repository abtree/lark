package xmysql

import "strings"

type MysqlWhere struct {
	Query  []string
	Args   []interface{}
	Limit  int
	Offset int
	Sort   string
}

func NewMysqlWhere() *MysqlWhere {
	return &MysqlWhere{
		Query: make([]string, 0),
		Args:  make([]interface{}, 0),
	}
}

func (w *MysqlWhere) SetFilter(query string, value interface{}) {
	w.Query = append(w.Query, query)
	w.Args = append(w.Args, value)
}

func (w *MysqlWhere) AddQuery(query string) {
	w.Query = append(w.Query, query)
}

func (w *MysqlWhere) AddArgs(value interface{}) {
	w.Args = append(w.Args, value)
}

func (w *MysqlWhere) GetQuery() string {
	if len(w.Query) <= 0 {
		return ""
	}
	return strings.Join(w.Query, " AND ")
}

func (w *MysqlWhere) SetSort(sort string) {
	w.Sort = sort
}

func (w *MysqlWhere) SetOffset(offset int) {
	w.Offset = offset
}

func (w *MysqlWhere) SetLimit(limit int) {
	w.Limit = limit
}

func (w *MysqlWhere) Reset() {
	w.Query = make([]string, 0)
	w.Args = make([]interface{}, 0)
	w.Sort = ""
	w.Limit = 0
	w.Offset = 0
}
