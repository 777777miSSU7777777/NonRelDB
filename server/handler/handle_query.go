package handler

import (
	"strings"
	"NonRelDB/server/storage/query"
)

func HandleQuery(q string) string {
	p := strings.Split(q, " ")
	if len(p) == 2 {
		if strings.ToLower(p[0]) == "get" {
			return query.Get(p[1])
		} else if strings.ToLower(p[0]) == "del" {
			return query.Del(p[1])
		} else if strings.ToLower(p[0]) == "keys" {
			return query.Keys(p[1])
		}
	} else if len(p) >= 3 {
		if strings.ToLower(p[0]) == "set" {
			return query.Set(p[1], strings.Join(p[2:]," "))
		}
	}
	return "Undefined query"
}
