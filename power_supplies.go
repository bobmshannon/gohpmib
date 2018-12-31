package hpmib

import (
	"strconv"
)

// PowerSupplyStatus describes the status of a PowerSupply.
type PowerSupplyStatus int

// PowerSupply models a power supply in the HP MIB.
type PowerSupply struct {
	BayNo                 int
	ChassisNo             int
	Condition             Status
	Status                PowerSupplyStatus
	SerialNo              string
	Model                 string
	PowerRatingWatts      int
	PowerConsumptionWatts int
}

// Statuses for power supplies defined by the HP MIB.
const (
	PowerSupplyStatusUnknown                 PowerSupplyStatus = -1
	PowerSupplyStatusNoError                 PowerSupplyStatus = 1
	PowerSupplyStatusGeneralFailure          PowerSupplyStatus = 2
	PowerSupplyStatusBISTFailure             PowerSupplyStatus = 3
	PowerSupplyStatusFanFailure              PowerSupplyStatus = 4
	PowerSupplyStatusTempFailure             PowerSupplyStatus = 5
	PowerSupplyStatusInterlockOpen           PowerSupplyStatus = 6
	PowerSupplyStatusEPROMFailed             PowerSupplyStatus = 7
	PowerSupplyStatusVREFFailed              PowerSupplyStatus = 8
	PowerSupplyStatusDACFailed               PowerSupplyStatus = 9
	PowerSupplyStatusRAMTestFailed           PowerSupplyStatus = 10
	PowerSupplyStatusVoltageChannelFailed    PowerSupplyStatus = 11
	PowerSupplyStatusORRingDiodeFailed       PowerSupplyStatus = 12
	PowerSupplyStatusBrownOut                PowerSupplyStatus = 13
	PowerSupplyStatusGiveUpOnStartup         PowerSupplyStatus = 14
	PowerSupplyStatusNVRAMInvalid            PowerSupplyStatus = 15
	PowerSupplyStatusCalibrationTableInvalid PowerSupplyStatus = 16
	PowerSupplyStatusNoPowerInput            PowerSupplyStatus = 17
)

// Table defined by the HP MIB that contains the status of each power supply.
const (
	cpqHeFltTolPowerSupplyChassis         OID = "1.3.6.1.4.1.232.6.2.9.3.1.1"
	cpqHeFltTolPowerSupplyBay             OID = "1.3.6.1.4.1.232.6.2.9.3.1.2"
	cpqHeFltTolPowerSupplyCondition       OID = "1.3.6.1.4.1.232.6.2.9.3.1.4"
	cpqHeFltTolPowerSupplyStatus          OID = "1.3.6.1.4.1.232.6.2.9.3.1.5"
	cpqHeFltTolPowerSupplyCapacityUsed    OID = "1.3.6.1.4.1.232.6.2.9.3.1.7"
	cpqHeFltTolPowerSupplyCapacityMaximum OID = "1.3.6.1.4.1.232.6.2.9.3.1.8"
	cpqHeFltTolPowerSupplyModel           OID = "1.3.6.1.4.1.232.6.2.9.3.1.10"
	cpqHeFltTolPowerSupplySerialNumber    OID = "1.3.6.1.4.1.232.6.2.9.3.1.11"
)

var (
	powerSupplyStatusIDMappings = map[string]PowerSupplyStatus{
		"1":  PowerSupplyStatusNoError,
		"2":  PowerSupplyStatusGeneralFailure,
		"3":  PowerSupplyStatusBISTFailure,
		"4":  PowerSupplyStatusFanFailure,
		"5":  PowerSupplyStatusTempFailure,
		"6":  PowerSupplyStatusInterlockOpen,
		"7":  PowerSupplyStatusEPROMFailed,
		"8":  PowerSupplyStatusVREFFailed,
		"9":  PowerSupplyStatusDACFailed,
		"10": PowerSupplyStatusRAMTestFailed,
		"11": PowerSupplyStatusVoltageChannelFailed,
		"12": PowerSupplyStatusORRingDiodeFailed,
		"13": PowerSupplyStatusBrownOut,
		"14": PowerSupplyStatusGiveUpOnStartup,
		"15": PowerSupplyStatusNVRAMInvalid,
		"16": PowerSupplyStatusCalibrationTableInvalid,
		"17": PowerSupplyStatusNoPowerInput,
	}
	powerSupplyStatusHumanMappings = map[PowerSupplyStatus]string{
		PowerSupplyStatusNoError:                 "No Error",
		PowerSupplyStatusGeneralFailure:          "General Failure",
		PowerSupplyStatusBISTFailure:             "BIST Failure",
		PowerSupplyStatusFanFailure:              "Fan Failure",
		PowerSupplyStatusTempFailure:             "Temp Failure",
		PowerSupplyStatusInterlockOpen:           "Interlock Open",
		PowerSupplyStatusEPROMFailed:             "EEPROM Failed",
		PowerSupplyStatusVREFFailed:              "VREF Failed",
		PowerSupplyStatusDACFailed:               "DAC Failed",
		PowerSupplyStatusRAMTestFailed:           "Ram Test Failed",
		PowerSupplyStatusVoltageChannelFailed:    "Voltage Channel Failed",
		PowerSupplyStatusORRingDiodeFailed:       "ORRing Diode Failed",
		PowerSupplyStatusBrownOut:                "Brown Out",
		PowerSupplyStatusGiveUpOnStartup:         "Give Up On Startup",
		PowerSupplyStatusNVRAMInvalid:            "NVRAM Invalid",
		PowerSupplyStatusCalibrationTableInvalid: "Calibration Table Invalid",
		PowerSupplyStatusNoPowerInput:            "No Power Input",
	}
)

// PowerSupplies returns a list of Power Supplies. Returns a non-nil error if the list of Power
// Supplies could not be determined.
func (m *MIB) PowerSupplies() ([]PowerSupply, error) {
	powerSupplies := []PowerSupply{}

	columns := OIDList{
		cpqHeFltTolPowerSupplyChassis,
		cpqHeFltTolPowerSupplyBay,
		cpqHeFltTolPowerSupplyCondition,
		cpqHeFltTolPowerSupplyStatus,
		cpqHeFltTolPowerSupplyModel,
		cpqHeFltTolPowerSupplySerialNumber,
		cpqHeFltTolPowerSupplyCapacityMaximum,
		cpqHeFltTolPowerSupplyCapacityUsed,
	}
	table, err := traverseTable(m.snmpClient, columns)
	if err != nil {
		return []PowerSupply{}, err
	}

	for _, row := range table {
		chassisNo, err := strconv.Atoi(row[0])
		if err != nil {
			return []PowerSupply{}, err
		}
		bayNo, err := strconv.Atoi(row[1])
		if err != nil {
			return []PowerSupply{}, err
		}
		condition := parseStatus(row[2])
		status := parsePowerSupplyStatus(row[3])
		model := prettifyString(row[4])
		serialNo := prettifyString(row[5])
		ratingWatts, err := strconv.Atoi(row[6])
		if err != nil {
			return []PowerSupply{}, err
		}
		consumptionWatts, err := strconv.Atoi(row[7])
		if err != nil {
			return []PowerSupply{}, err
		}

		powerSupplies = append(powerSupplies, PowerSupply{
			ChassisNo:             chassisNo,
			BayNo:                 bayNo,
			Condition:             condition,
			Status:                status,
			Model:                 model,
			SerialNo:              serialNo,
			PowerRatingWatts:      ratingWatts,
			PowerConsumptionWatts: consumptionWatts,
		})
	}

	return powerSupplies, nil
}

func parsePowerSupplyStatus(s string) PowerSupplyStatus {
	status, ok := powerSupplyStatusIDMappings[s]
	if !ok {
		return PowerSupplyStatusUnknown
	}
	return status
}

// String converts the PowerSupplyStatus to a human readable string.
func (ps *PowerSupplyStatus) String() string {
	s, ok := powerSupplyStatusHumanMappings[*ps]
	if !ok {
		return "Unknown"
	}
	return s
}
