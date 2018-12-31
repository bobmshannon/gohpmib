package hpmib

import (
	"strconv"
)

// TemperatureSensorLocale specifies the location of the temperature sensor.
type TemperatureSensorLocale int

// TemperatureSensorThresholdType describes the type of threshold associated with the temperature sensor.
type TemperatureSensorThresholdType int

// TemperatureSensor models a temperature sensor in the HP MIB.
type TemperatureSensor struct {
	ID                    int
	CurrentReadingCelsius int
	Locale                TemperatureSensorLocale
	Status                Status
	Threshold             int
	ThresholdType         TemperatureSensorThresholdType
}

// Locales for a temperature sensor defined by the HP MIB.
const (
	TemperatureSensorLocaleOther          TemperatureSensorLocale = 1
	TemperatureSensorLocaleUnknown        TemperatureSensorLocale = 2
	TemperatureSensorLocaleSystem         TemperatureSensorLocale = 3
	TemperatureSensorLocaleSystemBoard    TemperatureSensorLocale = 4
	TemperatureSensorLocaleIOBoard        TemperatureSensorLocale = 5
	TemperatureSensorLocaleCPU            TemperatureSensorLocale = 6
	TemperatureSensorLocaleMemory         TemperatureSensorLocale = 7
	TemperatureSensorLocaleStorage        TemperatureSensorLocale = 8
	TemperatureSensorLocaleRemovableMedia TemperatureSensorLocale = 9
	TemperatureSensorLocalePowerSupply    TemperatureSensorLocale = 10
	TemperatureSensorLocaleAmbient        TemperatureSensorLocale = 11
	TemperatureSensorLocaleChassis        TemperatureSensorLocale = 12
	TemperatureSensorLocaleBridgeCard     TemperatureSensorLocale = 13
)

// Threshold types for a temperature sensor defined by the HP MIB.
const (
	TemperatureSensorThresholdTypeUnknown    TemperatureSensorThresholdType = -1
	TemperatureSensorThresholdTypeOther      TemperatureSensorThresholdType = 1
	TemperatureSensorThresholdTypeBlowout    TemperatureSensorThresholdType = 5
	TemperatureSensorThresholdTypeCaution    TemperatureSensorThresholdType = 9
	TemperatureSensorThresholdTypeCritical   TemperatureSensorThresholdType = 15
	TemperatureSensorThresholdTypeNoReaction TemperatureSensorThresholdType = 16
)

// Table defined by the HP MIB that contains the status of each temperature sensor.
const (
	cpqHeTemperatureIndex         = "1.3.6.1.4.1.232.6.2.6.8.1.2"
	cpqHeTemperatureLocale        = "1.3.6.1.4.1.232.6.2.6.8.1.3"
	cpqHeTemperatureCelsius       = "1.3.6.1.4.1.232.6.2.6.8.1.4"
	cpqHeTemperatureThreshold     = "1.3.6.1.4.1.232.6.2.6.8.1.5"
	cpqHeTemperatureCondition     = "1.3.6.1.4.1.232.6.2.6.8.1.6"
	cpqHeTemperatureThresholdType = "1.3.6.1.4.1.232.6.2.6.8.1.7"
)

// TemperatureSensors returns a list of Temperature Sensors. Returns a non-nil error if the list of Temperature
// Sensors could not be determined.
func (m *MIB) TemperatureSensors() ([]TemperatureSensor, error) {
	sensors := []TemperatureSensor{}

	columns := OIDList{
		cpqHeTemperatureIndex,
		cpqHeTemperatureLocale,
		cpqHeTemperatureCelsius,
		cpqHeTemperatureThreshold,
		cpqHeTemperatureCondition,
		cpqHeTemperatureThresholdType,
	}
	table, err := traverseTable(m.snmpClient, columns)
	if err != nil {
		return []TemperatureSensor{}, err
	}

	for _, row := range table {
		index, err := strconv.Atoi(row[0])
		if err != nil {
			return []TemperatureSensor{}, err
		}
		locale := parseTemperatureSensorLocale(row[1])
		celsius, err := strconv.Atoi(row[2])
		if err != nil {
			return []TemperatureSensor{}, err
		}
		threshold, err := strconv.Atoi(row[3])
		if err != nil {
			return []TemperatureSensor{}, err
		}
		status := parseStatus(row[4])
		thresholdType := parseTemperatureSensorThresholdType(row[5])

		sensors = append(sensors, TemperatureSensor{
			ID:                    index,
			Locale:                locale,
			CurrentReadingCelsius: celsius,
			Threshold:             threshold,
			Status:                status,
			ThresholdType:         thresholdType,
		})
	}

	return sensors, nil
}

func parseTemperatureSensorLocale(s string) TemperatureSensorLocale {
	switch s {
	case "1":
		return TemperatureSensorLocaleOther
	case "2":
		return TemperatureSensorLocaleUnknown
	case "3":
		return TemperatureSensorLocaleSystem
	case "4":
		return TemperatureSensorLocaleSystemBoard
	case "5":
		return TemperatureSensorLocaleIOBoard
	case "6":
		return TemperatureSensorLocaleCPU
	case "7":
		return TemperatureSensorLocaleMemory
	case "8":
		return TemperatureSensorLocaleStorage
	case "9":
		return TemperatureSensorLocaleRemovableMedia
	case "10":
		return TemperatureSensorLocalePowerSupply
	case "11":
		return TemperatureSensorLocaleAmbient
	case "12":
		return TemperatureSensorLocaleChassis
	case "13":
		return TemperatureSensorLocaleBridgeCard
	default:
		return TemperatureSensorLocaleUnknown
	}
}

// String converts the TemperatureSensorLocale to a human readable string.
func (t *TemperatureSensorLocale) String() string {
	switch *t {
	case TemperatureSensorLocaleOther:
		return "Other"
	case TemperatureSensorLocaleUnknown:
		return "Unknown"
	case TemperatureSensorLocaleSystem:
		return "System"
	case TemperatureSensorLocaleSystemBoard:
		return "System Board"
	case TemperatureSensorLocaleIOBoard:
		return "IO Board"
	case TemperatureSensorLocaleCPU:
		return "CPU"
	case TemperatureSensorLocaleMemory:
		return "Memory"
	case TemperatureSensorLocaleStorage:
		return "Storage"
	case TemperatureSensorLocaleRemovableMedia:
		return "Removable Media"
	case TemperatureSensorLocalePowerSupply:
		return "Power Supply"
	case TemperatureSensorLocaleAmbient:
		return "Ambient"
	case TemperatureSensorLocaleChassis:
		return "Chassis"
	case TemperatureSensorLocaleBridgeCard:
		return "Bridge Card"
	default:
		return "Unknown"
	}
}

func parseTemperatureSensorThresholdType(s string) TemperatureSensorThresholdType {
	switch s {
	case "1":
		return TemperatureSensorThresholdTypeOther
	case "5":
		return TemperatureSensorThresholdTypeBlowout
	case "9":
		return TemperatureSensorThresholdTypeCaution
	case "15":
		return TemperatureSensorThresholdTypeCritical
	case "16":
		return TemperatureSensorThresholdTypeNoReaction
	default:
		return TemperatureSensorThresholdTypeUnknown
	}
}

// String converts the TemperatureSensorThresholdType to a human readable string.
func (t *TemperatureSensorThresholdType) String() string {
	switch *t {
	case TemperatureSensorThresholdTypeOther:
		return "Other"
	case TemperatureSensorThresholdTypeBlowout:
		return "Blowout"
	case TemperatureSensorThresholdTypeCaution:
		return "Caution"
	case TemperatureSensorThresholdTypeCritical:
		return "Critical"
	case TemperatureSensorThresholdTypeNoReaction:
		return "No Reaction"
	default:
		return "Unknown"
	}
}
