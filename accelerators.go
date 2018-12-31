package hpmib

import (
	"strconv"
)

// Table defined by the HP MIB that contains the status of each Array Accelerator.
const (
	cpqDaAccelCntlrIndex      = "1.3.6.1.4.1.232.3.2.2.2.1.1"
	cpqDaAccelBattery         = "1.3.6.1.4.1.232.3.2.2.2.1.6"
	cpqDaAccelCondition       = "1.3.6.1.4.1.232.3.2.2.2.1.9"
	cpqDaAccelSerialNumber    = "1.3.6.1.4.1.232.3.2.2.2.1.11"
	cpqDaAccelFailedBatteries = "1.3.6.1.4.1.232.3.2.2.2.1.15"
)

// ArrayAccelerator models an Array Accelerator in the HP MIB.
type ArrayAccelerator struct {
	ID                 int
	Status             Status
	BatteryStatus      BatteryStatus
	SerialNumber       string
	FailedBatterySlots string
}

// ArrayAccelerators returns a list of Array Accelerators. Returns a non-nil error of the list of ArrayAccelerators
// could not be determined.
func (m *MIB) ArrayAccelerators() ([]ArrayAccelerator, error) {
	accelerators := []ArrayAccelerator{}

	columns := OIDList{
		cpqDaAccelCntlrIndex,
		cpqDaAccelCondition,
		cpqDaAccelBattery,
		cpqDaAccelSerialNumber,
		cpqDaAccelFailedBatteries,
	}
	table, err := traverseTable(m.snmpClient, columns)
	if err != nil {
		return []ArrayAccelerator{}, err
	}

	for _, row := range table {
		index, err := strconv.Atoi(row[0])
		if err != nil {
			return []ArrayAccelerator{}, err
		}
		status := parseStatus(row[1])
		battStatus := parseBatteryStatus(row[2])
		serialNo := prettifyString(row[3])
		failedSlots := prettifyString(row[4])

		accelerators = append(accelerators, ArrayAccelerator{
			ID:                 index,
			FailedBatterySlots: failedSlots,
			Status:             status,
			SerialNumber:       serialNo,
			BatteryStatus:      battStatus,
		})
	}

	return accelerators, nil
}

// parseBatteryStatus takes an SNMP value and determines the battery status as defined by the HP MIB.
func parseBatteryStatus(s string) BatteryStatus {
	switch s {
	case "1":
		return BatteryStatusOther
	case "2":
		return BatteryStatusOK
	case "3":
		return BatteryStatusCharging
	case "4":
		return BatteryStatusFailed
	case "5":
		return BatteryStatusDegraded
	case "6":
		return BatteryStatusNotPresent
	case "7":
		return BatteryStatusCapacitorFailed
	default:
		return BatteryStatusUnknown
	}
}

// String converts the BatteryStatus to a human readable string.
func (s *BatteryStatus) String() string {
	switch *s {
	case BatteryStatusOther:
		return "Other"
	case BatteryStatusOK:
		return "OK"
	case BatteryStatusCharging:
		return "Charging"
	case BatteryStatusDegraded:
		return "Degraded"
	case BatteryStatusFailed:
		return "Failed"
	case BatteryStatusNotPresent:
		return "Not Present"
	case BatteryStatusCapacitorFailed:
		return "Capacitor Failed"
	default:
		return "Unknown"
	}
}
