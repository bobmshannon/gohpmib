package hpmib

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/soniah/gosnmp"
)

// OIDs defined by the HP MIB that each contain an integer describing the status of a sub-system.
//    COMPONENT          SUB-SYSTEM
//    Processors         (232.1.2.2.4:cpqSeCpuCondition)
//    Memory             (232.6.2.14.4:cpqHeResilientMemCondition)
//    Cooling            (232.6.2.6.4:cpqHeThermalSystemFanStatus)
//    Sensors            (232.6.2.6.3:cpqHeThermalTempStatus)
//    Power              (232.6.2.9.1:cpqHeFltTolPwrSupplyCondition)
//    ProLiant Logs      (232.6.2.11.2:cpqHeEventLogCondition)
//    ASR                (232.6.2.5.17:cpqHeAsrCondition)
//    Drive Array        (232.3.1.3:cpqDaMibCondition)
//    SCSI               (232.5.1.3:cpqScsiMibCondition)
//    Storage Enclosures (232.8.1.3:cpqSsMibCondition)
//    IDE                (232.14.1.3:cpqIdeMibCondition)
//    FC                 (232.16.1.3:cpqFcaMibCondition)
//    Networks           (232.18.1.3:cpqNicMibCondition)
//    MP                 (232.9.1.3:cpqSm2MibCondition)
//    HW/BIOS            (232.6.2.16.1:cpqHeHWBiosCondition)
//    Battery            (232.6.2.17.1:cpqHeSysBackupBatteryCondition)
//    iSCSI              (232.169.1.3:cpqiScsiMibCondition)
const (
	cpqHeAsrCondition              OID = "1.3.6.1.4.1.232.6.2.5.17"
	cpqHeSysBackupBatteryCondition OID = "1.3.6.1.4.1.232.6.2.17.1"
	cpqDaCntlrOverallCondition     OID = "1.3.6.1.4.1.232.3.2.2.1.1.6"
	cpqDaMibCondition              OID = "1.3.6.1.4.1.232.3.1.3"
	cpqSsMibCondition              OID = "1.3.6.1.4.1.232.8.1.3"
	cpqHeThermalSystemFanStatus    OID = "1.3.6.1.4.1.232.6.2.6.4"
	cpqHeResilientMemCondition     OID = "1.3.6.1.4.1.232.6.2.14.4"
	cpqHeFltTolPwrSupplyCondition  OID = "1.3.6.1.4.1.232.6.2.9.1"
	cpqSeCPUCondition              OID = "1.3.6.1.4.1.232.1.2.2.4"
	cpqHeThermalTempStatus         OID = "1.3.6.1.4.1.232.6.2.6.3"
	cpqHePowerMeterCurrReading     OID = "1.3.6.1.4.1.232.6.2.15.3"
)

// ASRStatus returns the status of the advanced server recovery sub-system.
// Returns a non-nil error if the status could not be determined.
func (m *MIB) ASRStatus() (Status, error) {
	return getStatusSummary(m.snmpClient, cpqHeAsrCondition)
}

// BackupBatteryStatus returns the status of the battery backup sub-system.
// Returns a non-nil error if the status could not be determined.
func (m *MIB) BackupBatteryStatus() (Status, error) {
	return getStatusSummary(m.snmpClient, cpqHeSysBackupBatteryCondition)
}

// ControllerStatus returns the overall status of the storage controllers.
// Returns a non-nil error if the status could not be determined.
func (m *MIB) ControllerStatus() (Status, error) {
	return getStatusSummary(m.snmpClient, cpqDaCntlrOverallCondition)
}

// DriveArrayStatus returns the overall status of the drive arrays.
// Returns a non-nil error if the status could not be determined.
func (m *MIB) DriveArrayStatus() (Status, error) {
	return getStatusSummary(m.snmpClient, cpqDaMibCondition)
}

// EnclosureStatus returns the overall status of the physical enclosure.
func (m *MIB) EnclosureStatus() (Status, error) {
	return getStatusSummary(m.snmpClient, cpqSsMibCondition)
}

// FanStatus returns the status of the fan(s) in the system.
// Returns StatusOther if fan status detection is not supported by the system,
// StatusOK if all fans are operating properly, StatusDegraded if a non-required
// fan is not operating properly, or StatusFailed if a required fan is not operating properly.
// Returns a non-nil error if the status could not be determined.
func (m *MIB) FanStatus() (Status, error) {
	return getStatusSummary(m.snmpClient, cpqHeThermalSystemFanStatus)
}

// MemoryStatus returns the status of the advanced memory protection sub-system.
// Returns a non-nil error if the status could not be determined.
func (m *MIB) MemoryStatus() (Status, error) {
	return getStatusSummary(m.snmpClient, cpqHeResilientMemCondition)
}

// PowerSupplyStatus returns the status of the fault tolerant power supply sub-system.
// Returns StatusOther if power supply status detection is not supported by the system,
// StatusOK if all power supplies are operating properly, StatusDegraded if one or more
// power supplies are operating in a degraded state, or StatusFailed if one or more power
// supplies have failed.
// Returns a non-nil error if the status could not be determined.
func (m *MIB) PowerSupplyStatus() (Status, error) {
	return getStatusSummary(m.snmpClient, cpqHeFltTolPwrSupplyCondition)
}

// PowerMeterReading returns the current power meter reading in Watts.
// Returns -1 if power meter is not supported by the server or an error
// if the power meter reading could not be determined.
func (m *MIB) PowerMeterReading() (int, error) {
	res, err := m.snmpClient.GetNext([]string{string(cpqHePowerMeterCurrReading)})
	if err != nil {
		return -1, err
	}
	if len(res.Variables) == 0 {
		return -1, ErrNoResultsReturned
	}
	if res.Variables[0].Type != gosnmp.Integer {
		return -1, ErrExpectedInteger
	}
	watts := res.Variables[0].Value.(int)
	return watts, nil
}

// ProcessorStatus returns the status of the processor sub-system.
// Returns a non-nil error if the status could not be determined.
func (m *MIB) ProcessorStatus() (Status, error) {
	return getStatusSummary(m.snmpClient, cpqSeCPUCondition)
}

// TemperatureSensorStatus returns the status of the system's temperature sensors.
// Returns StatusOther if temperature sensor status detection is not supported by the system,
// StatusOK if all temperature sensors are within normal operating range, StatusDegraded if one
// or more temperature sensors are outside a normal operating range, or StatusFailed if one or
// more temperature sensors detect a condition that could permanently damage the system.
// Returns a non-nil error if the status could not be determined.
func (m *MIB) TemperatureSensorStatus() (Status, error) {
	return getStatusSummary(m.snmpClient, cpqHeThermalTempStatus)
}

// getStatusSummary fetches the provided OID defined by the MIB whose value is expected to contain an
// integer that describes the overall status of a sub-system.
func getStatusSummary(client *gosnmp.GoSNMP, oid OID) (Status, error) {
	res, err := client.GetNext([]string{string(oid)})
	if err != nil {
		return StatusUnknown, err
	}
	if res.Error == gosnmp.NoSuchName {
		return StatusUnknown, fmt.Errorf("OID %s not found", oid.String())
	}
	if len(res.Variables) != 1 {
		return StatusUnknown, fmt.Errorf("expected only 1 variable in response but got %d", len(res.Variables))
	}
	if res.Variables[0].Type == gosnmp.NoSuchObject || res.Variables[0].Type == gosnmp.NoSuchInstance {
		return StatusOther, nil
	}
	if res.Variables[0].Type != gosnmp.Integer {
		return StatusUnknown, errors.New("expected variable type to be an integer")
	}
	return parseStatus(fmt.Sprintf("%v", res.Variables[0].Value)), nil
}

// parseStatus takes an SNMP value and determines the status as defined by the HP MIB.
func parseStatus(s string) Status {
	switch s {
	case "1":
		return StatusOther
	case "2":
		return StatusOK
	case "3":
		return StatusDegraded
	case "4":
		return StatusFailed
	default:
		return StatusUnknown
	}
}

// String converts the Status to a human readable string.
func (s *Status) String() string {
	switch *s {
	case StatusOther:
		return "Other"
	case StatusOK:
		return "OK"
	case StatusDegraded:
		return "Degraded"
	case StatusFailed:
		return "Failed"
	default:
		return "Unknown"
	}
}
