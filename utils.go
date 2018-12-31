package hpmib

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/soniah/gosnmp"
)

// traverseTable traverses the table with the provided columns by issuing multiple GETNEXT
// requests to the SNMP agent until one or more results are returned that are outside of the
// tablespace.
func traverseTable(client *gosnmp.GoSNMP, columns OIDList) ([][]string, error) {
	table := [][]string{}
	for i := range table {
		table[i] = make([]string, 0, len(columns))
	}

	rootOIDs := columns.Strings()
	currentOIDs := columns.Strings()

traverse:
	for {
		res, err := client.GetNext(currentOIDs)
		if err != nil {
			return [][]string{}, nil
		}
		row := []string{}
		currentOIDs = []string{}
		for i, v := range res.Variables {
			currentOIDs = append(currentOIDs, v.Name)
			if !strings.HasPrefix(v.Name, "."+rootOIDs[i]) {
				break traverse
			}
			switch v.Type {
			case gosnmp.OctetString:
				row = append(row, string(v.Value.([]byte)))
			case gosnmp.Integer:
				row = append(row, strconv.Itoa(v.Value.(int)))
			default:
				return [][]string{}, fmt.Errorf("unknown variable type encountered for OID %s", v.Name)
			}
		}
		if len(row) != len(columns) {
			return [][]string{}, fmt.Errorf("expected number of columns in the row %d to equal the number of columns %d in the table", len(row), len(columns))
		}
		table = append(table, row)
	}

	return table, nil
}

// prettifyString removes redundant whitespace characters from a string.
func prettifyString(s string) string {
	return strings.TrimSpace(strings.Join(strings.Fields(s), " "))
}
