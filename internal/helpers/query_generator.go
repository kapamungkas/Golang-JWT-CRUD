package helpers

import (
	"fmt"
	"strings"
)

func CreateQueryOrderAndLimit(option map[string]string) string {
	query := ""

	fmt.Println(option)
	for i, v := range option {
		if i == "order_by" {
			splitter := strings.Split(v, "|")
			if len(splitter) != 0 {
				query += " ORDER BY " + splitter[0] + " " + splitter[1]
			}
		}

		if i == "limit" {
			splitter := strings.Split(v, "|")
			if len(splitter) != 0 {
				query += " LIMIT " + splitter[0] + ", " + splitter[1]
			}
		}
	}

	return query
}
