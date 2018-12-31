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
	switch s {
	case "1":
		return PowerSupplyStatusNoError
	case "2":
		return PowerSupplyStatusGeneralFailure
	case "3":
		return PowerSupplyStatusBISTFailure
	case "4":
		return PowerSupplyStatusFanFailure
	case "5":
		return PowerSupplyStatusTempFailure
	case "6":
		return PowerSupplyStatusInterlockOpen
	case "7":
		return PowerSupplyStatusEPROMFailed
	case "8":
		return PowerSupplyStatusVREFFailed
	case "9":
		return PowerSupplyStatusDACFailed
	case "10":
		return PowerSupplyStatusRAMTestFailed
	case "11":
		return PowerSupplyStatusVoltageChannelFailed
	case "12":
		return PowerSupplyStatusORRingDiodeFailed
	case "13":
		return PowerSupplyStatusBrownOut
	case "14":
		return PowerSupplyStatusGiveUpOnStartup
	case "15":
		return PowerSupplyStatusNVRAMInvalid
	case "16":
		return PowerSupplyStatusCalibrationTableInvalid
	case "17":
		return PowerSupplyStatusNoPowerInput
	default:
		return PowerSupplyStatusUnknown
	}
}

// String converts the PowerSupplyStatus to a human readable string.
func (ps *PowerSupplyStatus) String() string {
	switch *ps {
	case PowerSupplyStatusNoError:
		return "No Error"
	case PowerSupplyStatusGeneralFailure:
		return "General Failure"
	case PowerSupplyStatusBISTFailure:
		return "BIST Failure"
	case PowerSupplyStatusFanFailure:
		return "Fan Failure"
	case PowerSupplyStatusTempFailure:
		return "Temp Failure"
	case PowerSupplyStatusInterlockOpen:
		return "Interlock Open"
	case PowerSupplyStatusEPROMFailed:
		return "EPROM Failed"
	case PowerSupplyStatusVREFFailed:
		return "VREF Failed"
	case PowerSupplyStatusDACFailed:
		return "DAC Failed"
	case PowerSupplyStatusRAMTestFailed:
		return "RAM Test Failed"
	case PowerSupplyStatusVoltageChannelFailed:
		return "Voltage Channel Failed"
	case PowerSupplyStatusORRingDiodeFailed:
		return "ORRing Diode Failed"
	case PowerSupplyStatusBrownOut:
		return "Brown Out"
	case PowerSupplyStatusGiveUpOnStartup:
		return "Give Up On Startup"
	case PowerSupplyStatusNVRAMInvalid:
		return "NVRAM Invalid"
	case PowerSupplyStatusCalibrationTableInvalid:
		return "Calibration Table Invalid"
	case PowerSupplyStatusNoPowerInput:
		return "No Power Input"
	default:
		return "Unknown"
	}
}
