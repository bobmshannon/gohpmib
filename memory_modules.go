package hpmib

import (
	"strconv"
)

// MemoryModule models a Memory Module in the HP MIB.
type MemoryModule struct {
	CPUNumber    int
	ModuleNumber int
	SizeKB       int
	PartNo       string
	Status       MemoryModuleStatus
}

// MemoryModuleStatus describes the state of the Memory Module.
type MemoryModuleStatus int

// Statuses for a Memory Module defined by the HP MIB.
const (
	MemoryModuleStatusUnknown      MemoryModuleStatus = -1
	MemoryModuleStatusOther        MemoryModuleStatus = 1
	MemoryModuleStatusNotPresent   MemoryModuleStatus = 2
	MemoryModuleStatusPresent      MemoryModuleStatus = 3
	MemoryModuleStatusGood         MemoryModuleStatus = 4
	MemoryModuleStatusAdd          MemoryModuleStatus = 5
	MemoryModuleStatusUpgrade      MemoryModuleStatus = 6
	MemoryModuleStatusMissing      MemoryModuleStatus = 7
	MemoryModuleStatusDoesNotMatch MemoryModuleStatus = 8
	MemoryModuleStatusNotSupported MemoryModuleStatus = 9
	MemoryModuleStatusBadConfig    MemoryModuleStatus = 10
	MemoryModuleStatusDegraded     MemoryModuleStatus = 11
	MemoryModuleStatusSpare        MemoryModuleStatus = 12
	MemoryModuleStatusPartial      MemoryModuleStatus = 13
)

// Table defined by the HP MIB that contains the status of each memory module.
const (
	cpqHeResMem2CpuNum       = "1.3.6.1.4.1.232.6.2.14.13.1.3"
	cpqHeResMem2ModuleNum    = "1.3.6.1.4.1.232.6.2.14.13.1.5"
	cpqHeResMem2ModuleSize   = "1.3.6.1.4.1.232.6.2.14.13.1.6"
	cpqHeResMem2ModulePartNo = "1.3.6.1.4.1.232.6.2.14.13.1.10"
	cpqHeResMem2ModuleStatus = "1.3.6.1.4.1.232.6.2.14.13.1.19"
)

// MemoryModules returns a list of Memory Modules. Returns a non-nil error if the list of Memory Modules
// could not be determined.
func (m *MIB) MemoryModules() ([]MemoryModule, error) {
	modules := []MemoryModule{}

	columns := OIDList{
		cpqHeResMem2CpuNum,
		cpqHeResMem2ModuleNum,
		cpqHeResMem2ModuleSize,
		cpqHeResMem2ModulePartNo,
		cpqHeResMem2ModuleStatus,
	}
	table, err := traverseTable(m.snmpClient, columns)
	if err != nil {
		return []MemoryModule{}, err
	}

	for _, row := range table {
		cpuNumber, err := strconv.Atoi(row[0])
		if err != nil {
			return []MemoryModule{}, nil
		}
		moduleNumber, err := strconv.Atoi(row[1])
		if err != nil {
			return []MemoryModule{}, nil
		}
		sizeKB, err := strconv.Atoi(row[2])
		if err != nil {
			return []MemoryModule{}, nil
		}
		partNo := prettifyString(row[3])
		status := parseMemoryModuleStatus(row[4])

		modules = append(modules, MemoryModule{
			CPUNumber:    cpuNumber,
			ModuleNumber: moduleNumber,
			SizeKB:       sizeKB,
			PartNo:       partNo,
			Status:       status,
		})
	}

	return modules, nil
}

// String converts MemoryModuleStatus to a human readable string.
func (m *MemoryModuleStatus) String() string {
	switch *m {
	case MemoryModuleStatusOther:
		return "Other"
	case MemoryModuleStatusNotPresent:
		return "Not Present"
	case MemoryModuleStatusPresent:
		return "Present"
	case MemoryModuleStatusGood:
		return "Good"
	case MemoryModuleStatusAdd:
		return "Add"
	case MemoryModuleStatusUpgrade:
		return "Upgrade"
	case MemoryModuleStatusMissing:
		return "Missing"
	case MemoryModuleStatusDoesNotMatch:
		return "Does Not Match"
	case MemoryModuleStatusNotSupported:
		return "Not Supported"
	case MemoryModuleStatusBadConfig:
		return "Bad Config"
	case MemoryModuleStatusDegraded:
		return "Degraded"
	case MemoryModuleStatusSpare:
		return "Spare"
	case MemoryModuleStatusPartial:
		return "Partial"
	default:
		return "Unknown"
	}
}

func parseMemoryModuleStatus(s string) MemoryModuleStatus {
	switch s {
	case "1":
		return MemoryModuleStatusOther
	case "2":
		return MemoryModuleStatusNotPresent
	case "3":
		return MemoryModuleStatusPresent
	case "4":
		return MemoryModuleStatusGood
	case "5":
		return MemoryModuleStatusAdd
	case "6":
		return MemoryModuleStatusUpgrade
	case "7":
		return MemoryModuleStatusMissing
	case "8":
		return MemoryModuleStatusDoesNotMatch
	case "9":
		return MemoryModuleStatusNotSupported
	case "10":
		return MemoryModuleStatusBadConfig
	case "11":
		return MemoryModuleStatusDegraded
	case "12":
		return MemoryModuleStatusSpare
	case "13":
		return MemoryModuleStatusPartial
	default:
		return MemoryModuleStatusUnknown
	}
}
