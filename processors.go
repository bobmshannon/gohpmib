package hpmib

import (
	"strconv"
)

// ProcessorStatus describes the status of the Processor.
type ProcessorStatus int

// Statuses for a processor defined by the HP MIB.
const (
	ProcessorStatusUnknown  ProcessorStatus = 1
	ProcessorStatusOK       ProcessorStatus = 2
	ProcessorStatusDegraded ProcessorStatus = 3
	ProcessorStatusFailed   ProcessorStatus = 4
	ProcessorStatusDisabled ProcessorStatus = 5
)

// ProcessorPowerStatus describes the power status of the Processor.
type ProcessorPowerStatus int

// Power statuses for a processor defined by the HP MIB.
const (
	ProcessorPowerStatusUnknown       ProcessorPowerStatus = 1
	ProcessorPowerStatusLowPowered    ProcessorPowerStatus = 2
	ProcessorPowerStatusNormalPowered ProcessorPowerStatus = 3
	ProcessorPowerStatusHighPowered   ProcessorPowerStatus = 4
)

// Processor models a CPU in the HP MIB.
type Processor struct {
	ID                  int
	Name                string
	MaxClockSpeedHz     int
	CurrentClockSpeedHz int
	PhysicalCores       int
	VirtualCores        int
	Status              ProcessorStatus
	PowerStatus         ProcessorPowerStatus
}

// Table defined by the HP MIB that contains the status of each processor.
const (
	cpqSeCPUUnitIndex      OID = "1.3.6.1.4.1.232.1.2.2.1.1.1"
	cpqSeCPUName           OID = "1.3.6.1.4.1.232.1.2.2.1.1.3"
	cpqSeCPUSpeed          OID = "1.3.6.1.4.1.232.1.2.2.1.1.4"
	cpqSeCPUStatus         OID = "1.3.6.1.4.1.232.1.2.2.1.1.6"
	cpqSeCPUCore           OID = "1.3.6.1.4.1.232.1.2.2.1.1.15"
	cpqSeCPUMaxSpeed       OID = "1.3.6.1.4.1.232.1.2.2.1.1.21"
	cpqSeCPUCoreMaxThreads OID = "1.3.6.1.4.1.232.1.2.2.1.1.25"
	cpqSeCPULowPowerStatus OID = "1.3.6.1.4.1.232.1.2.2.1.1.26"
)

// Processors returns a list of Processors. Returns a non-nil error if the list of Processors
// could not be determined.
func (m *MIB) Processors() ([]Processor, error) {
	processors := []Processor{}

	columns := OIDList{
		cpqSeCPUUnitIndex,
		cpqSeCPUName,
		cpqSeCPUSpeed,
		cpqSeCPUStatus,
		cpqSeCPUCore,
		cpqSeCPUMaxSpeed,
		cpqSeCPUCoreMaxThreads,
		cpqSeCPULowPowerStatus,
	}
	table, err := traverseTable(m.snmpClient, columns)
	if err != nil {
		return []Processor{}, err
	}

	for _, row := range table {
		id, err := strconv.Atoi(row[0])
		if err != nil {
			return []Processor{}, nil
		}
		name := prettifyString(row[1])
		currSpeed, err := strconv.Atoi(row[2])
		if err != nil {
			return []Processor{}, nil
		}
		status := parseProcessorStatus(row[3])
		physicalCores, err := strconv.Atoi(row[4])
		if err != nil {
			return []Processor{}, nil
		}
		maxSpeed, err := strconv.Atoi(row[5])
		if err != nil {
			return []Processor{}, nil
		}
		virtualCores, err := strconv.Atoi(row[6])
		if err != nil {
			return []Processor{}, nil
		}
		powerStatus := parseProcessorPowerStatus(row[7])

		processors = append(processors, Processor{
			ID:                  id,
			Name:                name,
			CurrentClockSpeedHz: currSpeed,
			Status:              status,
			PhysicalCores:       physicalCores,
			MaxClockSpeedHz:     maxSpeed,
			VirtualCores:        virtualCores,
			PowerStatus:         powerStatus,
		})
	}

	return processors, nil
}

func parseProcessorStatus(s string) ProcessorStatus {
	switch s {
	case "2":
		return ProcessorStatusOK
	case "3":
		return ProcessorStatusDegraded
	case "4":
		return ProcessorStatusFailed
	case "5":
		return ProcessorStatusDisabled
	default:
		return ProcessorStatusUnknown

	}
}

// String converts ProcessorStatus to a human readable string.
func (p *ProcessorStatus) String() string {
	switch *p {
	case ProcessorStatusOK:
		return "OK"
	case ProcessorStatusDegraded:
		return "Degraded"
	case ProcessorStatusFailed:
		return "Failed"
	case ProcessorStatusDisabled:
		return "Disabled"
	default:
		return "Unknown"
	}
}

func parseProcessorPowerStatus(s string) ProcessorPowerStatus {
	switch s {
	case "2":
		return ProcessorPowerStatusLowPowered
	case "3":
		return ProcessorPowerStatusNormalPowered
	case "4":
		return ProcessorPowerStatusHighPowered
	default:
		return ProcessorPowerStatusUnknown
	}
}

// String converts ProcessorPowerStatus to a human readable string.
func (p *ProcessorPowerStatus) String() string {
	switch *p {
	case ProcessorPowerStatusLowPowered:
		return "Low Powered"
	case ProcessorPowerStatusNormalPowered:
		return "Normal Powered"
	case ProcessorPowerStatusHighPowered:
		return "High Powered"
	default:
		return "Unknown"
	}
}
