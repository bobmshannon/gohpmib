package hpmib

import (
	"strconv"
)

// Controller models a storage controller in the HP MIB.
type Controller struct {
	ID       int
	SlotNo   int
	SerialNo string
	Status   Status
	Location string
}

// Table defined by the HP MIB that contains the status of each controller.
const (
	cpqDaCntlrIndex        = "1.3.6.1.4.1.232.3.2.2.1.1.1"
	cpqDaCntlrSlot         = "1.3.6.1.4.1.232.3.2.2.1.1.5"
	cpqDaCntlrCondition    = "1.3.6.1.4.1.232.3.2.2.1.1.6"
	cpqDaCntlrSerialNumber = "1.3.6.1.4.1.232.3.2.2.1.1.15"
	cpqDaCntlrHwLocation   = "1.3.6.1.4.1.232.3.2.2.1.1.20"
)

// Controllers returns a list of Controllers. Returns a non-nil error of the list of Controllers
// could not be determined.
func (m *MIB) Controllers() ([]Controller, error) {
	controllers := []Controller{}

	columns := OIDList{
		cpqDaCntlrIndex,
		cpqDaCntlrSlot,
		cpqDaCntlrCondition,
		cpqDaCntlrSerialNumber,
		cpqDaCntlrHwLocation,
	}
	table, err := traverseTable(m.snmpClient, columns)
	if err != nil {
		return []Controller{}, err
	}

	for _, row := range table {
		index, err := strconv.Atoi(row[0])
		if err != nil {
			return []Controller{}, err
		}
		slot, err := strconv.Atoi(row[1])
		if err != nil {
			return []Controller{}, err
		}
		status := parseStatus(row[2])
		serialNo := prettifyString(row[3])
		location := prettifyString(row[4])

		controllers = append(controllers, Controller{
			ID:       index,
			SlotNo:   slot,
			Status:   status,
			SerialNo: serialNo,
			Location: location,
		})
	}

	return controllers, nil
}
