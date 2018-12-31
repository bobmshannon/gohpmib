package hpmib

import (
	"strconv"
	"strings"
)

// LogicalDriveStatus describes the state of a logical drive.
type LogicalDriveStatus int

// FaultTolerance describes the fault tolerance of a logical drive.
type FaultTolerance int

// LogicalDrive models a logical drive in the HP MIB.
type LogicalDrive struct {
	// ID is the index of this logical drive.
	ID int
	// Name is this logical drive's name that is presented to the OS.
	Name string
	// AvailableSpares contains a list of physical drive IDs that that are available spares for this logical drive.
	AvailableSpares []string
	// ControllerID is the index of this logical drives' controller.
	ControllerID int
	// CapacityMB is the total capacity of this logical drive in megabytes.
	CapacityMB int
	// Condition represents the overall condition of this logical drive and any associated physical drives.
	Condition Status
	// Status represents the current status of this logical drive.
	Status LogicalDriveStatus
	// FaultTolerance is the fault tolerance mode of this logical drive.
	FaultTolerance FaultTolerance
}

// Fault tolerance modes for logical drives defined by the HP MIB.
const (
	FaultToleranceUnknown              FaultTolerance = -1
	FaultToleranceOther                FaultTolerance = 1
	FaultToleranceNone                 FaultTolerance = 2
	FaultToleranceMirroring            FaultTolerance = 3
	FaultToleranceDataGuard            FaultTolerance = 4
	FaultToleranceDistributedDataGuard FaultTolerance = 5
	FaultToleranceAdvancedDataGuard    FaultTolerance = 7
	FaultToleranceRAID50               FaultTolerance = 8
	FaultToleranceRAID10               FaultTolerance = 9
	FaultToleranceRAID1ADM             FaultTolerance = 10
	FaultToleranceRAID10ADM            FaultTolerance = 11
)

// Statuses for logical drives defined by the HP MIB.
const (
	LogicalDriveStatusOK                      LogicalDriveStatus = 2
	LogicalDriveStatusFailed                  LogicalDriveStatus = 3
	LogicalDriveStatusUnconfigured            LogicalDriveStatus = 4
	LogicalDriveStatusRecovering              LogicalDriveStatus = 5
	LogicalDriveStatusReadyForRebuild         LogicalDriveStatus = 6
	LogicalDriveStatusRebuilding              LogicalDriveStatus = 7
	LogicalDriveStatusWrongDrive              LogicalDriveStatus = 8
	LogicalDriveStatusBadConnect              LogicalDriveStatus = 9
	LogicalDriveStatusOverheating             LogicalDriveStatus = 10
	LogicalDriveStatusShutdown                LogicalDriveStatus = 11
	LogicalDriveStatusExpanding               LogicalDriveStatus = 12
	LogicalDriveStatusNotAvailable            LogicalDriveStatus = 13
	LogicalDriveStatusQueuedForExpansion      LogicalDriveStatus = 14
	LogicalDriveStatusMultipathAccessDegraded LogicalDriveStatus = 15
	LogicalDriveStatusErasing                 LogicalDriveStatus = 16
	LogicalDriveStatusUnknown                 LogicalDriveStatus = -1
)

// Table defined by the HP MIB that contains the status of each logical drive.
const (
	cpqDaLogDrvCntlrIndex  OID = "1.3.6.1.4.1.232.3.2.3.1.1.1"
	cpqDaLogDrvIndex       OID = "1.3.6.1.4.1.232.3.2.3.1.1.2"
	cpqDaLogDrvFaultTol    OID = "1.3.6.1.4.1.232.3.2.3.1.1.3"
	cpqDaLogDrvStatus      OID = "1.3.6.1.4.1.232.3.2.3.1.1.4"
	cpqDaLogDrvAvailSpares OID = "1.3.6.1.4.1.232.3.2.3.1.1.8"
	cpqDaLogDrvSize        OID = "1.3.6.1.4.1.232.3.2.3.1.1.9"
	cpqDaLogDrvCondition   OID = "1.3.6.1.4.1.232.3.2.3.1.1.11"
	cpqDaLogDrvOsName      OID = "1.3.6.1.4.1.232.3.2.3.1.1.14"
)

// LogicalDrives returns a list of Logical Drives. Returns a non-nil error if the list of Logical
// Drives could not be determined.
func (m *MIB) LogicalDrives() ([]LogicalDrive, error) {
	logicalDrives := []LogicalDrive{}

	columns := OIDList{
		cpqDaLogDrvCntlrIndex,
		cpqDaLogDrvIndex,
		cpqDaLogDrvFaultTol,
		cpqDaLogDrvStatus,
		cpqDaLogDrvAvailSpares,
		cpqDaLogDrvSize,
		cpqDaLogDrvCondition,
		cpqDaLogDrvOsName,
	}
	table, err := traverseTable(m.snmpClient, columns)
	if err != nil {
		return []LogicalDrive{}, err
	}

	for _, row := range table {
		cntlrIndex, err := strconv.Atoi(row[0])
		if err != nil {
			return []LogicalDrive{}, err
		}
		index, err := strconv.Atoi(row[1])
		if err != nil {
			return []LogicalDrive{}, err
		}
		faultTol := parseFaultTolerance(row[2])
		condition := parseStatus(row[3])
		availSpares := parseAvailableSpares(row[4])
		size, err := strconv.Atoi(row[5])
		if err != nil {
			return []LogicalDrive{}, err
		}
		status := parseLogicalDriveStatus(row[6])
		osName := row[7]
		logicalDrives = append(logicalDrives, LogicalDrive{
			ControllerID:    cntlrIndex,
			ID:              index,
			FaultTolerance:  faultTol,
			Condition:       condition,
			AvailableSpares: availSpares,
			Status:          status,
			CapacityMB:      size,
			Name:            osName,
		})
	}

	return logicalDrives, nil
}

// parseFaultTolerance returns the FaultTolerance that corresponds with the given string ID.
// Returns FaultToleranceUnknown if the given string ID cannot be not accounted for.
func parseFaultTolerance(s string) FaultTolerance {
	switch s {
	case "1":
		return FaultToleranceOther
	case "2":
		return FaultToleranceNone
	case "3":
		return FaultToleranceMirroring
	case "4":
		return FaultToleranceDataGuard
	case "5":
		return FaultToleranceDistributedDataGuard
	case "7":
		return FaultToleranceAdvancedDataGuard
	case "8":
		return FaultToleranceRAID50
	case "9":
		return FaultToleranceRAID10
	case "10":
		return FaultToleranceRAID1ADM
	case "11":
		return FaultToleranceRAID10ADM
	default:
		return FaultToleranceUnknown
	}
}

// String converts the FaultTolerance to a human readable string.
func (f *FaultTolerance) String() string {
	switch *f {
	case FaultToleranceOther:
		return "Other"
	case FaultToleranceNone:
		return "None"
	case FaultToleranceMirroring:
		return "Mirroring"
	case FaultToleranceDataGuard:
		return "Data Guard"
	case FaultToleranceDistributedDataGuard:
		return "Distributed Data Guard"
	case FaultToleranceAdvancedDataGuard:
		return "Advanced Data Guard"
	case FaultToleranceRAID50:
		return "RAID-50"
	case FaultToleranceRAID10:
		return "RAID-10"
	case FaultToleranceRAID1ADM:
		return "RAID-1 ADM"
	case FaultToleranceRAID10ADM:
		return "RAID-10 ADM"
	default:
		return "Unknown"
	}
}

// parseAvailableSpares returns the list of spares from the given space delimited string.
// Returns an empty list of there are no available spares.
func parseAvailableSpares(s string) []string {
	spares := strings.Split(s, " ")
	if len(spares) == 1 && spares[0] == "" {
		return []string{}
	}
	return spares
}

// parseFaultTolerance returns the LogicalDriveStatus that corresponds with the given string.
// Returns LogicalDriveStatusUnknown if the given string ID cannot be not accounted for.
func parseLogicalDriveStatus(s string) LogicalDriveStatus {
	switch s {
	case "2":
		return LogicalDriveStatusOK
	case "3":
		return LogicalDriveStatusFailed
	case "4":
		return LogicalDriveStatusUnconfigured
	case "5":
		return LogicalDriveStatusRecovering
	case "6":
		return LogicalDriveStatusReadyForRebuild
	case "7":
		return LogicalDriveStatusRebuilding
	case "8":
		return LogicalDriveStatusWrongDrive
	case "9":
		return LogicalDriveStatusBadConnect
	case "10":
		return LogicalDriveStatusOverheating
	case "11":
		return LogicalDriveStatusShutdown
	case "12":
		return LogicalDriveStatusExpanding
	case "13":
		return LogicalDriveStatusNotAvailable
	case "14":
		return LogicalDriveStatusQueuedForExpansion
	case "15":
		return LogicalDriveStatusMultipathAccessDegraded
	case "16":
		return LogicalDriveStatusErasing
	default:
		return LogicalDriveStatusUnknown
	}
}

// String converts the LogicalDriveStatus to a human readable string.
func (l *LogicalDriveStatus) String() string {
	switch *l {
	case LogicalDriveStatusOK:
		return "OK"
	case LogicalDriveStatusFailed:
		return "Failed"
	case LogicalDriveStatusUnconfigured:
		return "Unconfigured"
	case LogicalDriveStatusRecovering:
		return "Recovering"
	case LogicalDriveStatusReadyForRebuild:
		return "Ready For Rebuild"
	case LogicalDriveStatusRebuilding:
		return "Rebuilding"
	case LogicalDriveStatusWrongDrive:
		return "Wrong Drive"
	case LogicalDriveStatusBadConnect:
		return "Bad Connect"
	case LogicalDriveStatusOverheating:
		return "Overheating"
	case LogicalDriveStatusShutdown:
		return "Shutdown"
	case LogicalDriveStatusExpanding:
		return "Expanding"
	case LogicalDriveStatusNotAvailable:
		return "Not Available"
	case LogicalDriveStatusQueuedForExpansion:
		return "Queued For Expansion"
	case LogicalDriveStatusMultipathAccessDegraded:
		return "Multipath Access Degraded"
	case LogicalDriveStatusErasing:
		return "Erasing"
	default:
		return "Unknown"
	}
}
