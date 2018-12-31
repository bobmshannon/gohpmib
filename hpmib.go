package hpmib

import (
	"fmt"

	"github.com/soniah/gosnmp"
)

// AuthProtocol is the auth protocol to use for SNMPv3 connections.
type AuthProtocol string

// Auth protocols for SNMPv3 connections.
const (
	// AuthProtocolMD5 models the MD5 auth protocol for SNMPv3 connections.
	AuthProtocolMD5 AuthProtocol = "MD5"
	// AuthProtocolSHA models the SHA auth protocol for SNMPv3 connections.
	AuthProtocolSHA AuthProtocol = "SHA"
)

// PrivProtocol is the priv protocol to use for SNMPv3 connections.
type PrivProtocol string

// Priv protocols for SNMPv3 connections.
const (
	// PrivProtocolAES models the AES priv protocol for SNMPv3 connections.
	PrivProtocolAES PrivProtocol = "AES"
	// PrivProtocolDES models the DES priv protocol for SNMPv3 connections.
	PrivProtocolDES PrivProtocol = "DES"
)

// SecurityLevel is the security level to use for SNMPv3 connections.
type SecurityLevel string

// Security levels for SNMPv3 connections.
const (
	// SecurityLevelAuthPriv models the authPriv security level for SNMPv3 connections.
	SecurityLevelAuthPriv SecurityLevel = "authPriv"
	// SecurityLevelAuthNoPriv models the authNoPriv security level for SNMPv3 connections.
	SecurityLevelAuthNoPriv SecurityLevel = "authNoPriv"
	// SecurityLevelNoAuthNoPriv models the noAuthNoPriv securitylevel for SNMPv3 connections.
	SecurityLevelNoAuthNoPriv SecurityLevel = "noAuthNoPriv"
)

// SNMPVersion is the protocol version to use for SNMP connections.
type SNMPVersion string

// Versions of SNMP.
const (
	// SNMPVersion3 models version 3 of of the SNMP protocol.
	SNMPVersion3 SNMPVersion = "3"
	// SNMPVersion2c models version 2c of of the SNMP protocol.
	SNMPVersion2c SNMPVersion = "2c"
	// SNMPVersion1 models version 1 of of the SNMP protocol.
	SNMPVersion1 SNMPVersion = "1"
)

// Status describes the state of a device.
type Status int

// Statuses defined by the HP MIB.
const (
	// StatusUnknown is used if the status cannot be determined.
	StatusUnknown Status = -1
	// StatusOther represents the "Other" state defined by the HP MIB.
	StatusOther Status = 1
	// StatusOK represents the "OK" state defined by the HP MIB.
	StatusOK Status = 2
	// StatusDegraded represents the "Degraded" state defined by the HP MIB.
	StatusDegraded Status = 3
	// StatusFailed represents the "Failed" state defined by the HP MIB.
	StatusFailed Status = 4
)

// BatteryStatus descibes the state of a battery.
type BatteryStatus int

const (
	// BatteryStatusUnknown indicates that the battery status cannot be determined.
	BatteryStatusUnknown BatteryStatus = -1
	// BatteryStatusOther indicates that the instrument agent does not recognize battery status.
	BatteryStatusOther BatteryStatus = 1
	// BatteryStatusOK indicates that the battery is fully charged.
	BatteryStatusOK BatteryStatus = 2
	// BatteryStatusCharging indicates that the battery power is less than 75% and is recharging.
	BatteryStatusCharging BatteryStatus = 3
	// BatteryStatusFailed indicates that the battery has failed.
	BatteryStatusFailed BatteryStatus = 4
	// BatteryStatusDegraded indicates that the battery is below the sufficient voltage level and has not been recharged.
	BatteryStatusDegraded BatteryStatus = 5
	// BatteryStatusNotPresent indicates that there is no battery installed.
	BatteryStatusNotPresent BatteryStatus = 6
	// BatteryStatusCapacitorFailed indicates that the battery capacitor failed.
	BatteryStatusCapacitorFailed BatteryStatus = 7
)

// OID represents an object ID
type OID string

// OIDList represents a list of object IDs
type OIDList []OID

// A StatusChecker queries the HP MIB for device status information.
type StatusChecker interface {
	ArrayAccelerators() ([]ArrayAccelerator, error)
	ASRStatus() (Status, error)
	BackupBatteryStatus() (Status, error)
	Controllers() ([]Controller, error)
	ControllerStatus() (Status, error)
	DriveArrayStatus() (Status, error)
	EnclosureStatus() (Status, error)
	Fans() ([]Fan, error)
	FanStatus() (Status, error)
	LogicalDrives() ([]LogicalDrive, error)
	MemoryModules() ([]MemoryModule, error)
	MemoryStatus() (Status, error)
	Model() (string, error)
	PhysicalDrives() ([]PhysicalDrive, error)
	PowerMeterReading() (int, error)
	PowerSupplies() ([]PowerSupply, error)
	PowerSupplyStatus() (Status, error)
	Processors() ([]Processor, error)
	ProcessorStatus() (Status, error)
	SerialNumber() (string, error)
	TemperatureSensors() ([]TemperatureSensor, error)
	TemperatureSensorStatus() (Status, error)
}

// MIB implements StatusChecker.
type MIB struct {
	snmpClient *gosnmp.GoSNMP
}

// MIBConfig is used to configure the HP MIB.
type MIBConfig struct {
	SNMPConfig `yaml:"snmp"`
}

// AuthConfig is used to configure authentication and authorization for the SNMP connection.
type AuthConfig struct {
	// Community specifies the community name to use with the SNMP connection.
	Community string `yaml:"community"`
	// SecurityLevel specifies the security level to use for SNMPv3 connections. Supported security levels are "authPriv", "authNoPriv", and "noAuthNoPriv".
	SecurityLevel SecurityLevel `yaml:"security-level,omitempty"`
	// Username specifies the username use for SNMPv3 connections.
	Username string `yaml:"username,omitempty"`
	// Password specifies the password use for SNMPv3 connections.
	Password string `yaml:"password,omitempty"`
	// AuthProtocol specifies the auth protocol to use for SNMPv3 connections. Supported auth protocols are "SHA" and "MD5".
	AuthProtocol AuthProtocol `yaml:"auth-protocol,omitempty"`
	// PrivProtocol specifies the priv protocol to use for SNMPv3 connections. Supported priv protocols are "DES" and "AES".
	PrivProtocol PrivProtocol `yaml:"priv-protocol,omitempty"`
	// PrivPassword specifies the priv password to use for SNMPv3 connections.
	PrivPassword string `yaml:"priv-password,omitempty"`
	// ContextName specifies the context name to use for SNMPv3 connections.
	ContextName string `yaml:"context-name,omitempty"`
}

// SNMPConfig is used to configure the connection to the SNMP agent.
type SNMPConfig struct {
	// Auth contains parameters used for authentication and authorization.
	Auth AuthConfig `yaml:"auth"`
	// Address specifies the address of the SNMP agent.
	Address string `yaml:"address"`
	// Port specifies the port that the SNMP agent is listening on.
	Port int `yaml:"port"`
	// Version specifies the SNMP protocol version to use. Supported versions are "1", "2c", and "3".
	Version SNMPVersion `yaml:"version"`
}

// NewMIB returns a new HP MIB.
func NewMIB(cfg *MIBConfig) (*MIB, error) {
	c := gosnmp.Default
	c.Community = cfg.Auth.Community
	c.Target = cfg.Address
	c.Port = uint16(cfg.Port)
	c.ContextName = cfg.Auth.ContextName
	switch cfg.Version {
	case SNMPVersion1:
		c.Version = gosnmp.Version1
	case SNMPVersion2c:
		c.Version = gosnmp.Version2c
	case SNMPVersion3:
		c.Version = gosnmp.Version3
		var err error
		c, err = configureSNMPClientWithAuth(cfg.Auth, c)
		if err != nil {
			return nil, err
		}
	}
	if err := c.Connect(); err != nil {
		return nil, err
	}
	return &MIB{
		snmpClient: c,
	}, nil
}

// Connect creates a new socket to be used by the SNMP client.
func (m *MIB) Connect() error {
	return m.snmpClient.Connect()
}

// Close closes the socket used by the SNMP client.
func (m *MIB) Close() error {
	return m.snmpClient.Conn.Close()
}

// configureSNMPClientWithAuth configures the SNMPv3 client using the provided authentication configuration.
func configureSNMPClientWithAuth(authCfg AuthConfig, client *gosnmp.GoSNMP) (*gosnmp.GoSNMP, error) {
	client.SecurityModel = gosnmp.UserSecurityModel
	auth, priv := false, false
	switch authCfg.SecurityLevel {
	case SecurityLevelAuthPriv:
		if authCfg.PrivPassword == "" {
			return nil, fmt.Errorf("priv password must be specified when using authPriv")
		}
		if authCfg.PrivProtocol != PrivProtocolAES && authCfg.PrivProtocol != PrivProtocolDES {
			return nil, fmt.Errorf("priv protocol must be AES or DES when using authPriv")
		}
		client.MsgFlags = gosnmp.AuthPriv
		auth = true
		priv = true
	case SecurityLevelAuthNoPriv:
		if authCfg.Password == "" {
			return nil, fmt.Errorf("password is required when using authNoPriv")
		}
		if authCfg.AuthProtocol != AuthProtocolMD5 && authCfg.AuthProtocol != AuthProtocolSHA {
			return nil, fmt.Errorf("auth protocol must be SHA or MD5 when using authNoPriv")
		}
		client.MsgFlags = gosnmp.AuthNoPriv
		auth = true
	case SecurityLevelNoAuthNoPriv:
		if authCfg.Username == "" {
			return nil, fmt.Errorf("username is required when using noAuthNoPriv")
		}
		client.MsgFlags = gosnmp.NoAuthNoPriv
	}
	securityParams := &gosnmp.UsmSecurityParameters{
		UserName: authCfg.Username,
	}
	if auth {
		securityParams.AuthenticationPassphrase = authCfg.Password
		switch authCfg.AuthProtocol {
		case AuthProtocolSHA:
			securityParams.AuthenticationProtocol = gosnmp.SHA
		case AuthProtocolMD5:
			securityParams.AuthenticationProtocol = gosnmp.MD5
		}
	}
	if priv {
		securityParams.PrivacyPassphrase = authCfg.PrivPassword
		switch authCfg.PrivProtocol {
		case PrivProtocolDES:
			securityParams.PrivacyProtocol = gosnmp.DES
		case PrivProtocolAES:
			securityParams.PrivacyProtocol = gosnmp.AES
		}
	}
	client.SecurityParameters = securityParams
	return client, nil
}

// String converts the OID to a string.
func (o *OID) String() string {
	return string(*o)
}

// Strings converts the list of OIDs to a list of strings.
func (o *OIDList) Strings() []string {
	l := make([]string, 0, len(*o))
	for _, id := range *o {
		l = append(l, id.String())
	}
	return l
}
