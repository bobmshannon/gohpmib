package hpmib

import "strconv"

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

var (
	fanLocaleIDMappings = map[string]FanLocale{
		"1":  FanLocaleOther,
		"2":  FanLocaleUnknown,
		"3":  FanLocaleSystem,
		"4":  FanLocaleSystemBoard,
		"5":  FanLocaleIOBoard,
		"6":  FanLocaleCPU,
		"7":  FanLocaleMemory,
		"8":  FanLocaleStorage,
		"9":  FanLocaleRemovableMedia,
		"10": FanLocalePowerSupply,
		"11": FanLocaleAmbient,
		"12": FanLocaleChassis,
		"13": FanLocaleBridgeCard,
		"14": FanLocaleManagementBoard,
		"15": FanLocaleBackplane,
		"16": FanLocaleNetworkSlot,
		"17": FanLocaleBladeSlot,
		"18": FanLocaleVirtual,
	}
	fanLocaleHumanMappings = map[FanLocale]string{
		FanLocaleOther:           "Other",
		FanLocaleUnknown:         "Unknown",
		FanLocaleSystem:          "System",
		FanLocaleSystemBoard:     "System Board",
		FanLocaleIOBoard:         "IO Board",
		FanLocaleCPU:             "CPU",
		FanLocaleMemory:          "Memory",
		FanLocaleStorage:         "Storage",
		FanLocaleRemovableMedia:  "Removable Media",
		FanLocalePowerSupply:     "Power Supply",
		FanLocaleAmbient:         "Ambient",
		FanLocaleChassis:         "Chassis",
		FanLocaleBridgeCard:      "Bridge Card",
		FanLocaleManagementBoard: "Management Board",
		FanLocaleBackplane:       "Backplane",
		FanLocaleNetworkSlot:     "Network Slot",
		FanLocaleBladeSlot:       "Blade Slot",
		FanLocaleVirtual:         "Virtual",
	}
	fanRedundancyIDMappings = map[string]FanRedundancy{
		"1": FanRedundancyOther,
		"2": FanRedundancyNotRedundant,
		"3": FanRedundancyRedundant,
	}
	fanRedundancyHumanMappings = map[FanRedundancy]string{
		FanRedundancyOther:        "Other",
		FanRedundancyNotRedundant: "Not Redundant",
		FanRedundancyRedundant:    "Redundant",
	}
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
	locale, ok := fanLocaleIDMappings[s]
	if !ok {
		return FanLocaleUnknown
	}
	return locale
}

// String converts the FanLocale to a human readable string.
func (f *FanLocale) String() string {
	s, ok := fanLocaleHumanMappings[*f]
	if !ok {
		return "Unknown"
	}
	return s
}

func parseFanRedundancy(s string) FanRedundancy {
	redundancy, ok := fanRedundancyIDMappings[s]
	if !ok {
		return FanRedundancyUnknown
	}
	return redundancy
}

// String converts the FanRedundancy to a human readable string.
func (f *FanRedundancy) String() string {
	s, ok := fanRedundancyHumanMappings[*f]
	if !ok {
		return "Unknown"
	}
	return s
}
