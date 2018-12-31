package hpmib

import (
	"strconv"
)

// FanStatus describes the state of a fan.
type FanStatus int

// FanLocale specifies the location of the fan in the system.
type FanLocale int

// FanRedundancy describes the fault tolerance of the fan.
type FanRedundancy int

// Fan models a fan in the HP MIB.
type Fan struct {
	ID         int
	Locale     FanLocale
	Redundancy FanRedundancy
	Status     Status
}

// Table defined by the HP MIB that contains the status of each fan.
const (
	cpqHeFltTolFanIndex     OID = "1.3.6.1.4.1.232.6.2.6.7.1.2"
	cpqHeFltTolFanLocale    OID = "1.3.6.1.4.1.232.6.2.6.7.1.3"
	cpqHeFltTolFanRedundant OID = "1.3.6.1.4.1.232.6.2.6.7.1.7"
	cpqHeFltTolFanCondition OID = "1.3.6.1.4.1.232.6.2.6.7.1.9"
)

// Redundancy states for a fan defined by the HP MIB.
const (
	FanRedundancyUnknown      FanRedundancy = -11
	FanRedundancyOther        FanRedundancy = 1
	FanRedundancyNotRedundant FanRedundancy = 2
	FanRedundancyRedundant    FanRedundancy = 3
)

// Locales for a fan defined by the HP MIB.
const (
	FanLocaleOther           FanLocale = 1
	FanLocaleUnknown         FanLocale = 2
	FanLocaleSystem          FanLocale = 3
	FanLocaleSystemBoard     FanLocale = 4
	FanLocaleIOBoard         FanLocale = 5
	FanLocaleCPU             FanLocale = 6
	FanLocaleMemory          FanLocale = 7
	FanLocaleStorage         FanLocale = 8
	FanLocaleRemovableMedia  FanLocale = 9
	FanLocalePowerSupply     FanLocale = 10
	FanLocaleAmbient         FanLocale = 11
	FanLocaleChassis         FanLocale = 12
	FanLocaleBridgeCard      FanLocale = 13
	FanLocaleManagementBoard FanLocale = 14
	FanLocaleBackplane       FanLocale = 15
	FanLocaleNetworkSlot     FanLocale = 16
	FanLocaleBladeSlot       FanLocale = 17
	FanLocaleVirtual         FanLocale = 18
)

// Fans returns a list of Fans. Returns a non-nil error of the list of Fans
// could not be determined.
func (m *MIB) Fans() ([]Fan, error) {
	fans := []Fan{}

	columns := OIDList{
		cpqHeFltTolFanIndex,
		cpqHeFltTolFanLocale,
		cpqHeFltTolFanRedundant,
		cpqHeFltTolFanCondition,
	}
	table, err := traverseTable(m.snmpClient, columns)
	if err != nil {
		return []Fan{}, err
	}

	for _, row := range table {
		index, err := strconv.Atoi(row[0])
		if err != nil {
			return []Fan{}, err
		}
		locale := parseFanLocale(row[1])
		redundancy := parseFanRedundancy(row[2])
		status := parseStatus(row[3])

		fans = append(fans, Fan{
			ID:         index,
			Status:     status,
			Locale:     locale,
			Redundancy: redundancy,
		})
	}

	return fans, nil
}

func parseFanLocale(s string) FanLocale {
	switch s {
	case "1":
		return FanLocaleOther
	case "3":
		return FanLocaleSystem
	case "4":
		return FanLocaleSystemBoard
	case "5":
		return FanLocaleIOBoard
	case "6":
		return FanLocaleCPU
	case "7":
		return FanLocaleMemory
	case "8":
		return FanLocaleStorage
	case "9":
		return FanLocaleRemovableMedia
	case "10":
		return FanLocalePowerSupply
	case "11":
		return FanLocaleAmbient
	case "12":
		return FanLocaleChassis
	case "13":
		return FanLocaleBridgeCard
	case "14":
		return FanLocaleManagementBoard
	case "15":
		return FanLocaleBackplane
	case "16":
		return FanLocaleNetworkSlot
	case "17":
		return FanLocaleBladeSlot
	case "18":
		return FanLocaleVirtual
	default:
		return FanLocaleUnknown
	}
}

// String converts the FanLocale to a human readable string.
func (f *FanLocale) String() string {
	switch *f {
	case FanLocaleOther:
		return "Other"
	case FanLocaleSystem:
		return "System"
	case FanLocaleSystemBoard:
		return "System Board"
	case FanLocaleIOBoard:
		return "IO Board"
	case FanLocaleCPU:
		return "CPU"
	case FanLocaleMemory:
		return "Memory"
	case FanLocaleStorage:
		return "Storage"
	case FanLocaleRemovableMedia:
		return "Removable Media"
	case FanLocalePowerSupply:
		return "Power Supply"
	case FanLocaleAmbient:
		return "Ambient"
	case FanLocaleChassis:
		return "Chassis"
	case FanLocaleBridgeCard:
		return "Bridge Card"
	case FanLocaleManagementBoard:
		return "Management Board"
	case FanLocaleBackplane:
		return "Backplane"
	case FanLocaleNetworkSlot:
		return "Network Slot"
	case FanLocaleBladeSlot:
		return "Blade Slot"
	case FanLocaleVirtual:
		return "Virtual"
	default:
		return "Unknown"
	}
}

func parseFanRedundancy(s string) FanRedundancy {
	switch s {
	case "1":
		return FanRedundancyOther
	case "2":
		return FanRedundancyNotRedundant
	case "3":
		return FanRedundancyRedundant
	default:
		return FanRedundancyUnknown
	}
}

// String converts the FanRedundancy to a human readable string.
func (f *FanRedundancy) String() string {
	switch *f {
	case FanRedundancyOther:
		return "Other"
	case FanRedundancyNotRedundant:
		return "Not Redundant"
	case FanRedundancyRedundant:
		return "Redundant"
	default:
		return "Unknown"
	}
}
