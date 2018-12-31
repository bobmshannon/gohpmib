package hpmib

import (
	"strconv"
)

// MediaType describes a physical drive's media type in the HP MIB.
type MediaType int

// PhysicalDriveStatus describes the state of a physical drive.
type PhysicalDriveStatus int

// SMARTStatus describes the state of a physical disk using S.M.A.R.T. semantics.
type SMARTStatus int

// Media types for physical drives defined by the HP MIB.
const (
	MediaTypeUnknown         MediaType = -1
	MediaTypeOther           MediaType = 1
	MediaTypeRotatingPlatter MediaType = 2
	MediaTypeSolidState      MediaType = 3
	MediaTypeSMR             MediaType = 4
)

// Statuses for physical drives defined by the HP MIB.
const (
	PhysicalDriveStatusUnknown           PhysicalDriveStatus = -1
	PhysicalDriveStatusOther             PhysicalDriveStatus = 1
	PhysicalDriveStatusOK                PhysicalDriveStatus = 2
	PhysicalDriveStatusFailed            PhysicalDriveStatus = 3
	PhysicalDriveStatusPredictiveFailure PhysicalDriveStatus = 4
	PhysicalDriveStatusErasing           PhysicalDriveStatus = 5
	PhysicalDriveStatusEraseDone         PhysicalDriveStatus = 6
	PhysicalDriveStatusEraseQueued       PhysicalDriveStatus = 7
	PhysicalDriveStatusSSDWearOut        PhysicalDriveStatus = 8
	PhysicalDriveStatusNotAuthenticated  PhysicalDriveStatus = 9
)

// S.M.A.R.T. statuses for physical drives defined by the HP MIB.
const (
	SMARTStatusUnknown      SMARTStatus = -1
	SMARTStatusOther        SMARTStatus = 1
	SMARTStatusOK           SMARTStatus = 2
	SMARTStatusReplaceDrive SMARTStatus = 3
)

// Table defined by the HP MIB that contains the status of each physical drive.
const (
	cpqDaPhyDrvCntlrIndex  OID = "1.3.6.1.4.1.232.3.2.5.1.1.1"
	cpqDaPhyDrvIndex       OID = "1.3.6.1.4.1.232.3.2.5.1.1.2"
	cpqDaPhyDrvModel       OID = "1.3.6.1.4.1.232.3.2.5.1.1.3"
	cpqDaPhyDrvStatus      OID = "1.3.6.1.4.1.232.3.2.5.1.1.6"
	cpqDaPhyDrvSize        OID = "1.3.6.1.4.1.232.3.2.5.1.1.45"
	cpqDaPhyDrvSerialNum   OID = "1.3.6.1.4.1.232.3.2.5.1.1.51"
	cpqDaPhyDrvSmartStatus OID = "1.3.6.1.4.1.232.3.2.5.1.1.57"
	cpqDaPhyDrvLocation    OID = "1.3.6.1.4.1.232.3.2.5.1.1.64"
	cpqDaPhyDrvMediaType   OID = "1.3.6.1.4.1.232.3.2.5.1.1.69"
)

// PhysicalDrive models a physical drive in the HP MIB.
type PhysicalDrive struct {
	ID           int
	ControllerID int
	CapacityMB   int
	MediaType    MediaType
	Location     string
	Model        string
	SerialNo     string
	SMARTStatus  SMARTStatus
	Status       PhysicalDriveStatus
}

// PhysicalDrives returns a list of Physical Drives. Returns a non-nil error of the list of Physical
// Drives could not be determined.
func (m *MIB) PhysicalDrives() ([]PhysicalDrive, error) {
	physicalDrives := []PhysicalDrive{}

	columns := OIDList{
		cpqDaPhyDrvCntlrIndex,
		cpqDaPhyDrvIndex,
		cpqDaPhyDrvModel,
		cpqDaPhyDrvStatus,
		cpqDaPhyDrvSize,
		cpqDaPhyDrvSerialNum,
		cpqDaPhyDrvSmartStatus,
		cpqDaPhyDrvLocation,
		cpqDaPhyDrvMediaType,
	}
	table, err := traverseTable(m.snmpClient, columns)
	if err != nil {
		return []PhysicalDrive{}, err
	}

	for _, row := range table {
		cntlrIndex, err := strconv.Atoi(row[0])
		if err != nil {
			return []PhysicalDrive{}, err
		}
		index, err := strconv.Atoi(row[1])
		if err != nil {
			return []PhysicalDrive{}, err
		}
		model := prettifyString(row[2])
		status := parsePhysicalDriveStatus(row[3])
		size, err := strconv.Atoi(row[4])
		if err != nil {
			return []PhysicalDrive{}, err
		}
		serialNo := prettifyString(row[5])
		smartStatus := parseSMARTStatus(row[6])
		location := prettifyString(row[7])
		mediaType := parseMediaType(row[8])

		physicalDrives = append(physicalDrives, PhysicalDrive{
			ControllerID: cntlrIndex,
			ID:           index,
			Model:        model,
			Status:       status,
			CapacityMB:   size,
			SerialNo:     serialNo,
			SMARTStatus:  smartStatus,
			Location:     location,
			MediaType:    mediaType,
		})
	}

	return physicalDrives, nil
}

func parsePhysicalDriveStatus(s string) PhysicalDriveStatus {
	switch s {
	case "1":
		return PhysicalDriveStatusOther
	case "2":
		return PhysicalDriveStatusOK
	case "3":
		return PhysicalDriveStatusFailed
	case "4":
		return PhysicalDriveStatusPredictiveFailure
	case "5":
		return PhysicalDriveStatusErasing
	case "6":
		return PhysicalDriveStatusEraseDone
	case "7":
		return PhysicalDriveStatusEraseQueued
	case "8":
		return PhysicalDriveStatusSSDWearOut
	case "9":
		return PhysicalDriveStatusNotAuthenticated
	default:
		return PhysicalDriveStatusUnknown
	}
}

// String converts the PhysicalDriveStatus to a human readable string.
func (p *PhysicalDriveStatus) String() string {
	switch *p {
	case PhysicalDriveStatusOther:
		return "Other"
	case PhysicalDriveStatusOK:
		return "OK"
	case PhysicalDriveStatusFailed:
		return "Failed"
	case PhysicalDriveStatusPredictiveFailure:
		return "Predictive Failure"
	case PhysicalDriveStatusErasing:
		return "Erasing"
	case PhysicalDriveStatusEraseDone:
		return "Erase Done"
	case PhysicalDriveStatusEraseQueued:
		return "Erase Queued"
	case PhysicalDriveStatusSSDWearOut:
		return "SSD Wear Out"
	case PhysicalDriveStatusNotAuthenticated:
		return "Not Authenticated"
	default:
		return "Unknown"
	}
}

func parseMediaType(s string) MediaType {
	switch s {
	case "1":
		return MediaTypeOther
	case "2":
		return MediaTypeRotatingPlatter
	case "3":
		return MediaTypeSolidState
	case "4":
		return MediaTypeSMR
	default:
		return MediaTypeUnknown
	}
}

// String converts the MediaType to a human readable string.
func (m *MediaType) String() string {
	switch *m {
	case MediaTypeOther:
		return "Other"
	case MediaTypeRotatingPlatter:
		return "Rotating Platter"
	case MediaTypeSolidState:
		return "Solid State"
	case MediaTypeSMR:
		return "SMR"
	default:
		return "Unknown"
	}
}

func parseSMARTStatus(s string) SMARTStatus {
	switch s {
	case "1":
		return SMARTStatusOther
	case "2":
		return SMARTStatusOK
	case "3":
		return SMARTStatusReplaceDrive
	default:
		return SMARTStatusUnknown
	}
}

// String converts the SMARTStatus to a human readable string.
func (s *SMARTStatus) String() string {
	switch *s {
	case SMARTStatusOther:
		return "Other"
	case SMARTStatusOK:
		return "OK"
	case SMARTStatusReplaceDrive:
		return "Replace Drive"
	default:
		return "Unknown"
	}
}
