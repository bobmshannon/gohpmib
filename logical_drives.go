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

var (
	faultToleranceIDMappings = map[string]FaultTolerance{
		"1":  FaultToleranceOther,
		"2":  FaultToleranceNone,
		"3":  FaultToleranceMirroring,
		"4":  FaultToleranceDataGuard,
		"5":  FaultToleranceDistributedDataGuard,
		"7":  FaultToleranceAdvancedDataGuard,
		"8":  FaultToleranceRAID50,
		"9":  FaultToleranceRAID10,
		"10": FaultToleranceRAID1ADM,
		"11": FaultToleranceRAID10ADM,
	}
	faultToleranceHumanMappings = map[FaultTolerance]string{
		FaultToleranceOther:                "Other",
		FaultToleranceNone:                 "None",
		FaultToleranceMirroring:            "Mirroring",
		FaultToleranceDataGuard:            "Data Guard",
		FaultToleranceDistributedDataGuard: "Distributed Data Guard",
		FaultToleranceAdvancedDataGuard:    "Advanced Data Guard",
		FaultToleranceRAID50:               "RAID-50",
		FaultToleranceRAID10:               "RAID-10",
		FaultToleranceRAID1ADM:             "RAID-1 ADM",
		FaultToleranceRAID10ADM:            "RAID-10 ADM",
	}
	logicalDriveStatusIDMappings = map[string]LogicalDriveStatus{
		"2":  LogicalDriveStatusOK,
		"3":  LogicalDriveStatusFailed,
		"4":  LogicalDriveStatusUnconfigured,
		"5":  LogicalDriveStatusRecovering,
		"6":  LogicalDriveStatusReadyForRebuild,
		"7":  LogicalDriveStatusRebuilding,
		"8":  LogicalDriveStatusWrongDrive,
		"9":  LogicalDriveStatusBadConnect,
		"10": LogicalDriveStatusOverheating,
		"11": LogicalDriveStatusShutdown,
		"12": LogicalDriveStatusExpanding,
		"13": LogicalDriveStatusNotAvailable,
		"14": LogicalDriveStatusQueuedForExpansion,
		"15": LogicalDriveStatusMultipathAccessDegraded,
		"16": LogicalDriveStatusErasing,
	}
	logicalDriveStatusHumanMappings = map[LogicalDriveStatus]string{
		LogicalDriveStatusOK:                      "OK",
		LogicalDriveStatusFailed:                  "Failed",
		LogicalDriveStatusUnconfigured:            "Unconfigured",
		LogicalDriveStatusRecovering:              "Recovering",
		LogicalDriveStatusReadyForRebuild:         "Ready For Rebuild",
		LogicalDriveStatusRebuilding:              "Rebuilding",
		LogicalDriveStatusWrongDrive:              "Wrong Drive",
		LogicalDriveStatusBadConnect:              "Bad Connect",
		LogicalDriveStatusOverheating:             "Overheating",
		LogicalDriveStatusShutdown:                "Shutdown",
		LogicalDriveStatusExpanding:               "Expanding",
		LogicalDriveStatusNotAvailable:            "Not Available",
		LogicalDriveStatusQueuedForExpansion:      "Queued For Expansion",
		LogicalDriveStatusMultipathAccessDegraded: "Multipath Access Degraded",
		LogicalDriveStatusErasing:                 "Erasing",
	}
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
	tolerance, ok := faultToleranceIDMappings[s]
	if !ok {
		return FaultToleranceUnknown
	}
	return tolerance
}

// String converts the FaultTolerance to a human readable string.
func (f *FaultTolerance) String() string {
	s, ok := faultToleranceHumanMappings[*f]
	if !ok {
		return "Unknown"
	}
	return s
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
	status, ok := logicalDriveStatusIDMappings[s]
	if !ok {
		return LogicalDriveStatusUnknown
	}
	return status
}

// String converts the LogicalDriveStatus to a human readable string.
func (l *LogicalDriveStatus) String() string {
	s, ok := logicalDriveStatusHumanMappings[*l]
	if !ok {
		return "Unknown"
	}
	return s
}
