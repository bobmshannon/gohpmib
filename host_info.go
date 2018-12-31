package hpmib

import (
	"github.com/soniah/gosnmp"
)

// OIDs defined by the HP MIB that describe the properties of the system.
const (
	cpqSiSysSerialNum OID = "1.3.6.1.4.1.232.2.2.2.1"
	cpqSiProductName  OID = "1.3.6.1.4.1.232.2.2.4.2"
)

// SerialNumber returns the serial number of the server.
// Returns a non-nil error if the serial number could not be determined.
func (m *MIB) SerialNumber() (string, error) {
	res, err := m.snmpClient.GetNext([]string{string(cpqSiSysSerialNum)})
	if err != nil {
		return "", err
	}
	if len(res.Variables) == 0 {
		return "", ErrNoResultsReturned
	}
	if res.Variables[0].Type != gosnmp.OctetString {
		return "", ErrExpectedOctetString
	}
	serialNo := string(res.Variables[0].Value.([]byte))
	return prettifyString(serialNo), nil
}

// Model returns the model of the server.
// Returns a non-nil error if the model could not be determined.
func (m *MIB) Model() (string, error) {
	res, err := m.snmpClient.GetNext([]string{string(cpqSiProductName)})
	if err != nil {
		return "", err
	}
	if len(res.Variables) == 0 {
		return "", ErrNoResultsReturned
	}
	if res.Variables[0].Type != gosnmp.OctetString {
		return "", ErrExpectedOctetString
	}
	model := string(res.Variables[0].Value.([]byte))
	return prettifyString(model), nil
}
