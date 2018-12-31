// +build integration

package hpmib

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newTestingMIB(t *testing.T, generation int) *MIB {
	var mib *MIB
	var err error
	switch generation {
	case 8:
		mib, err = NewMIB(&MIBConfig{
			SNMPConfig{
				Address: "127.0.0.1",
				Port:    1024,
				Version: SNMPVersion3,
				Auth: AuthConfig{
					ContextName:   "proliant-dl380-g8",
					Community:     "proliant-dl380-g8",
					SecurityLevel: SecurityLevelAuthPriv,
					AuthProtocol:  AuthProtocolMD5,
					Username:      "simulator",
					Password:      "auctoritas",
					PrivProtocol:  PrivProtocolDES,
					PrivPassword:  "privatus",
				},
			},
		})
	case 7:
		mib, err = NewMIB(&MIBConfig{
			SNMPConfig{
				Address: "127.0.0.1",
				Port:    1024,
				Version: SNMPVersion3,
				Auth: AuthConfig{
					ContextName:   "proliant-dl380-g7",
					Community:     "proliant-dl380-g7",
					SecurityLevel: SecurityLevelAuthPriv,
					AuthProtocol:  AuthProtocolMD5,
					Username:      "simulator",
					Password:      "auctoritas",
					PrivProtocol:  PrivProtocolDES,
					PrivPassword:  "privatus",
				},
			},
		})
	default:
		t.Fatalf("unrecognized HP generation %d", generation)
	}

	require.NoError(t, err, "failed to initialize the MIB")

	return mib
}

func TestMIB_ArrayAccelerators(t *testing.T) {
	tests := []struct {
		Name       string
		Generation int
		Expected   []ArrayAccelerator
	}{
		{
			Name:       "ProLiant DL380 Generation 7 Array Accelerators",
			Generation: 7,
			Expected: []ArrayAccelerator{
				{
					ID:                 0,
					Status:             StatusOK,
					BatteryStatus:      BatteryStatusOK,
					SerialNumber:       "PBCDF0CRH1K8GA",
					FailedBatterySlots: "",
				},
			},
		},
		{
			Name:       "ProLiant DL380 Accelerator Array Accelerators",
			Generation: 8,
			Expected: []ArrayAccelerator{
				{
					ID:                 0,
					Status:             StatusOK,
					BatteryStatus:      BatteryStatusOK,
					SerialNumber:       "PBKUD0ARH2D0AV",
					FailedBatterySlots: "",
				},
			},
		},
	}

	for _, test := range tests {
		mib := newTestingMIB(t, test.Generation)
		accelerators, err := mib.ArrayAccelerators()
		require.NoError(t, err, "failed to retrieve array accelerator from the MIB")
		assert.Equal(t, test.Expected, accelerators)
	}
}

func TestMIB_Controllers(t *testing.T) {
	tests := []struct {
		Name       string
		Expected   []Controller
		Generation int
	}{
		{
			Name:       "ProLiant DL380 Generation 7 Controllers",
			Generation: 7,
			Expected: []Controller{
				{
					ID:       0,
					SlotNo:   0,
					SerialNo: "500143801756DC50",
					Status:   StatusOK,
					Location: "Slot 0",
				},
			},
		},
		{
			Name:       "ProLiant DL380 Generation 8 Controllers",
			Generation: 8,
			Expected: []Controller{
				{
					ID:       0,
					SlotNo:   0,
					SerialNo: "50014380210B3DD0",
					Status:   StatusOK,
					Location: "Slot 0",
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			mib := newTestingMIB(t, test.Generation)
			controllers, err := mib.Controllers()
			require.NoError(t, err, "failed to retrieve controllers from the MIB")
			assert.Equal(t, test.Expected, controllers)
		})
	}
}

func TestMIB_LogicalDrives(t *testing.T) {
	tests := []struct {
		Name       string
		Expected   []LogicalDrive
		Generation int
	}{
		{
			Name:       "ProLiant DL380 Generation 7 Logical Drives",
			Generation: 7,
			Expected: []LogicalDrive{
				{
					ID:              1,
					Name:            "/dev/sda",
					AvailableSpares: []string{},
					ControllerID:    0,
					CapacityMB:      953837,
					Condition:       StatusOK,
					Status:          LogicalDriveStatusOK,
					FaultTolerance:  FaultToleranceMirroring,
				},
				{
					ID:              2,
					Name:            "/dev/sdb",
					AvailableSpares: []string{},
					CapacityMB:      953837,
					Condition:       StatusOK,
					Status:          LogicalDriveStatusOK,
					FaultTolerance:  FaultToleranceMirroring,
				},
			},
		},
		{
			Name:       "ProLiant DL380 Generation 8 Logical Drives",
			Generation: 8,
			Expected: []LogicalDrive{
				{
					ID:              1,
					Name:            "/dev/sda",
					AvailableSpares: []string{},
					ControllerID:    0,
					CapacityMB:      572293,
					Condition:       StatusOK,
					Status:          LogicalDriveStatusOK,
					FaultTolerance:  FaultToleranceMirroring,
				},
				{
					ID:              2,
					Name:            "/dev/sdb",
					AvailableSpares: []string{},
					CapacityMB:      572293,
					Condition:       StatusOK,
					Status:          LogicalDriveStatusOK,
					FaultTolerance:  FaultToleranceMirroring,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			mib := newTestingMIB(t, test.Generation)
			logicalDrives, err := mib.LogicalDrives()
			require.NoError(t, err, "failed to retrieve logical drives from the MIB")
			assert.Equal(t, test.Expected, logicalDrives)
		})
	}
}

func TestMIB_PhysicalDrives(t *testing.T) {
	tests := []struct {
		Name       string
		Expected   []PhysicalDrive
		Generation int
	}{
		{
			Name:       "ProLiant DL380 Generation 7 Physical Drives",
			Generation: 7,
			Expected: []PhysicalDrive{
				{
					ID:           0,
					ControllerID: 0,
					CapacityMB:   953869,
					MediaType:    MediaTypeRotatingPlatter,
					Location:     "Port 1I Box 1 Bay 1",
					Model:        "ATA ST91000640NS",
					SerialNo:     "9XG3P718",
					SMARTStatus:  SMARTStatusOK,
					Status:       PhysicalDriveStatusOK,
				},
				{
					ID:           1,
					ControllerID: 0,
					CapacityMB:   953869,
					MediaType:    MediaTypeRotatingPlatter,
					Location:     "Port 1I Box 1 Bay 2",
					Model:        "ATA ST91000640NS",
					SerialNo:     "9XG3V13S",
					SMARTStatus:  SMARTStatusOK,
					Status:       PhysicalDriveStatusOK,
				},
				{
					ID:           2,
					ControllerID: 0,
					CapacityMB:   953869,
					MediaType:    MediaTypeRotatingPlatter,
					Location:     "Port 1I Box 1 Bay 3",
					Model:        "ATA ST91000640NS",
					SerialNo:     "9XG3TS6B",
					SMARTStatus:  SMARTStatusOK,
					Status:       PhysicalDriveStatusOK,
				},
				{
					ID:           3,
					ControllerID: 0,
					CapacityMB:   953869,
					MediaType:    MediaTypeRotatingPlatter,
					Location:     "Port 1I Box 1 Bay 4",
					Model:        "ATA ST91000640NS",
					SerialNo:     "9XG3WE7M",
					SMARTStatus:  SMARTStatusOK,
					Status:       PhysicalDriveStatusOK,
				},
			},
		},
		{
			Name:       "ProLiant DL380 Generation 8 Physical Drives",
			Generation: 8,
			Expected: []PhysicalDrive{
				{
					ID:           8,
					ControllerID: 0,
					CapacityMB:   572325,
					MediaType:    MediaTypeRotatingPlatter,
					Location:     "Port 1I Box 1 Bay 1",
					Model:        "HP EG0600FBLSH",
					SerialNo:     "6XR49KJK0000M334J34K",
					SMARTStatus:  SMARTStatusOK,
					Status:       PhysicalDriveStatusOK,
				},
				{
					ID:           9,
					ControllerID: 0,
					CapacityMB:   572325,
					MediaType:    MediaTypeRotatingPlatter,
					Location:     "Port 1I Box 1 Bay 2",
					Model:        "HP EG0600FBLSH",
					SerialNo:     "6XR49LC90000B236LZEM",
					SMARTStatus:  SMARTStatusOK,
					Status:       PhysicalDriveStatusOK,
				},
				{
					ID:           10,
					ControllerID: 0,
					CapacityMB:   572325,
					MediaType:    MediaTypeRotatingPlatter,
					Location:     "Port 1I Box 1 Bay 3",
					Model:        "HP EG0600FBDSR",
					SerialNo:     "EA01PD31TYY11310",
					SMARTStatus:  SMARTStatusOK,
					Status:       PhysicalDriveStatusOK,
				},
				{
					ID:           11,
					ControllerID: 0,
					CapacityMB:   572325,
					MediaType:    MediaTypeRotatingPlatter,
					Location:     "Port 1I Box 1 Bay 4",
					Model:        "HP EG0600FBDSR",
					SerialNo:     "EA01PD31U0BK1310",
					SMARTStatus:  SMARTStatusOK,
					Status:       PhysicalDriveStatusOK,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			mib := newTestingMIB(t, test.Generation)
			physicalDrives, err := mib.PhysicalDrives()
			require.NoError(t, err, "failed to retrieve physical drives from the MIB")
			assert.Equal(t, test.Expected, physicalDrives)
		})
	}
}

func TestMIB_Fans(t *testing.T) {
	tests := []struct {
		Name       string
		Expected   []Fan
		Generation int
	}{
		{
			Name:       "ProLiant DL380 Generation 7 Fans",
			Generation: 7,
			Expected: []Fan{
				{
					ID:         1,
					Locale:     FanLocaleSystem,
					Redundancy: FanRedundancyRedundant,
					Status:     StatusOK,
				},
				{
					ID:         2,
					Locale:     FanLocaleSystem,
					Redundancy: FanRedundancyRedundant,
					Status:     StatusOK,
				},
				{
					ID:         3,
					Locale:     FanLocaleSystem,
					Redundancy: FanRedundancyRedundant,
					Status:     StatusOK,
				},
				{
					ID:         4,
					Locale:     FanLocaleSystem,
					Redundancy: FanRedundancyRedundant,
					Status:     StatusOK,
				},
				{
					ID:         5,
					Locale:     FanLocaleSystem,
					Redundancy: FanRedundancyRedundant,
					Status:     StatusOK,
				},
				{
					ID:         6,
					Locale:     FanLocaleSystem,
					Redundancy: FanRedundancyRedundant,
					Status:     StatusOK,
				},
			},
		},
		{
			Name:       "ProLiant DL380 Generation 8 Fans",
			Generation: 8,
			Expected: []Fan{
				{
					ID:         1,
					Locale:     FanLocaleSystem,
					Redundancy: FanRedundancyRedundant,
					Status:     StatusOK,
				},
				{
					ID:         2,
					Locale:     FanLocaleSystem,
					Redundancy: FanRedundancyRedundant,
					Status:     StatusOK,
				},
				{
					ID:         3,
					Locale:     FanLocaleSystem,
					Redundancy: FanRedundancyRedundant,
					Status:     StatusOK,
				},
				{
					ID:         4,
					Locale:     FanLocaleSystem,
					Redundancy: FanRedundancyRedundant,
					Status:     StatusOK,
				},
				{
					ID:         5,
					Locale:     FanLocaleSystem,
					Redundancy: FanRedundancyRedundant,
					Status:     StatusOK,
				},
				{
					ID:         6,
					Locale:     FanLocaleSystem,
					Redundancy: FanRedundancyRedundant,
					Status:     StatusOK,
				},
			},
		},
	}

	for _, test := range tests {
		mib := newTestingMIB(t, test.Generation)
		fans, err := mib.Fans()
		require.NoError(t, err, "failed to retrieve fans from the MIB")
		assert.Equal(t, test.Expected, fans)
	}
}

func TestMIB_MemoryModules(t *testing.T) {
	tests := []struct {
		Name       string
		Expected   []MemoryModule
		Generation int
	}{
		{
			Name:       "ProLiant DL380 Generation 7 Memory Modules",
			Generation: 7,
			Expected: []MemoryModule{
				{
					CPUNumber:    1,
					ModuleNumber: 1,
					SizeKB:       0,
					PartNo:       "",
					Status:       MemoryModuleStatusNotPresent,
				},
				{
					CPUNumber:    1,
					ModuleNumber: 2,
					SizeKB:       16777216,
					PartNo:       "",
					Status:       MemoryModuleStatusGood,
				},
				{
					CPUNumber:    1,
					ModuleNumber: 3,
					SizeKB:       16777216,
					PartNo:       "",
					Status:       MemoryModuleStatusGood,
				},
				{
					CPUNumber:    1,
					ModuleNumber: 4,
					SizeKB:       0,
					PartNo:       "",
					Status:       MemoryModuleStatusNotPresent,
				},
				{
					CPUNumber:    1,
					ModuleNumber: 5,
					SizeKB:       16777216,
					PartNo:       "",
					Status:       MemoryModuleStatusGood,
				},
				{
					CPUNumber:    1,
					ModuleNumber: 6,
					SizeKB:       16777216,
					PartNo:       "",
					Status:       MemoryModuleStatusGood,
				},
				{
					CPUNumber:    1,
					ModuleNumber: 7,
					SizeKB:       0,
					PartNo:       "",
					Status:       MemoryModuleStatusNotPresent,
				},
				{
					CPUNumber:    1,
					ModuleNumber: 8,
					SizeKB:       16777216,
					PartNo:       "",
					Status:       MemoryModuleStatusGood,
				},
				{
					CPUNumber:    1,
					ModuleNumber: 9,
					SizeKB:       16777216,
					PartNo:       "",
					Status:       MemoryModuleStatusGood,
				},
				{
					CPUNumber:    2,
					ModuleNumber: 1,
					SizeKB:       0,
					PartNo:       "",
					Status:       MemoryModuleStatusNotPresent,
				},
				{
					CPUNumber:    2,
					ModuleNumber: 2,
					SizeKB:       16777216,
					PartNo:       "",
					Status:       MemoryModuleStatusGood,
				},
				{
					CPUNumber:    2,
					ModuleNumber: 3,
					SizeKB:       16777216,
					PartNo:       "",
					Status:       MemoryModuleStatusGood,
				},
				{
					CPUNumber:    2,
					ModuleNumber: 4,
					SizeKB:       0,
					PartNo:       "",
					Status:       MemoryModuleStatusNotPresent,
				},
				{
					CPUNumber:    2,
					ModuleNumber: 5,
					SizeKB:       16777216,
					PartNo:       "",
					Status:       MemoryModuleStatusGood,
				},
				{
					CPUNumber:    2,
					ModuleNumber: 6,
					SizeKB:       16777216,
					PartNo:       "",
					Status:       MemoryModuleStatusGood,
				},
				{
					CPUNumber:    2,
					ModuleNumber: 7,
					SizeKB:       0,
					PartNo:       "",
					Status:       MemoryModuleStatusNotPresent,
				},
				{
					CPUNumber:    2,
					ModuleNumber: 8,
					SizeKB:       16777216,
					PartNo:       "",
					Status:       MemoryModuleStatusGood,
				},
				{
					CPUNumber:    2,
					ModuleNumber: 9,
					SizeKB:       16777216,
					PartNo:       "",
					Status:       MemoryModuleStatusGood,
				},
			},
		},
		{
			Name:       "ProLiant DL380 Generation 8 Memory Modules",
			Generation: 8,
			Expected: []MemoryModule{
				{
					CPUNumber:    1,
					ModuleNumber: 1,
					SizeKB:       16777216,
					PartNo:       "",
					Status:       MemoryModuleStatusGood,
				},
				{
					CPUNumber:    1,
					ModuleNumber: 2,
					SizeKB:       16777216,
					PartNo:       "",
					Status:       MemoryModuleStatusGood,
				},
				{
					CPUNumber:    1,
					ModuleNumber: 3,
					SizeKB:       0,
					PartNo:       "",
					Status:       MemoryModuleStatusNotPresent,
				},
				{
					CPUNumber:    1,
					ModuleNumber: 4,
					SizeKB:       16777216,
					PartNo:       "",
					Status:       MemoryModuleStatusGood,
				},
				{
					CPUNumber:    1,
					ModuleNumber: 5,
					SizeKB:       16777216,
					PartNo:       "",
					Status:       MemoryModuleStatusGood,
				},
				{
					CPUNumber:    1,
					ModuleNumber: 6,
					SizeKB:       0,
					PartNo:       "",
					Status:       MemoryModuleStatusNotPresent,
				},
				{
					CPUNumber:    1,
					ModuleNumber: 7,
					SizeKB:       0,
					PartNo:       "",
					Status:       MemoryModuleStatusNotPresent,
				},
				{
					CPUNumber:    1,
					ModuleNumber: 8,
					SizeKB:       16777216,
					PartNo:       "",
					Status:       MemoryModuleStatusGood,
				},
				{
					CPUNumber:    1,
					ModuleNumber: 9,
					SizeKB:       16777216,
					PartNo:       "",
					Status:       MemoryModuleStatusGood,
				},
				{
					CPUNumber:    1,
					ModuleNumber: 10,
					SizeKB:       0,
					PartNo:       "",
					Status:       MemoryModuleStatusNotPresent,
				},
				{
					CPUNumber:    1,
					ModuleNumber: 11,
					SizeKB:       16777216,
					PartNo:       "",
					Status:       MemoryModuleStatusGood,
				},
				{
					CPUNumber:    1,
					ModuleNumber: 12,
					SizeKB:       16777216,
					PartNo:       "",
					Status:       MemoryModuleStatusGood,
				},
				{
					CPUNumber:    2,
					ModuleNumber: 1,
					SizeKB:       16777216,
					PartNo:       "",
					Status:       MemoryModuleStatusGood,
				},
				{
					CPUNumber:    2,
					ModuleNumber: 2,
					SizeKB:       16777216,
					PartNo:       "",
					Status:       MemoryModuleStatusGood,
				},
				{
					CPUNumber:    2,
					ModuleNumber: 3,
					SizeKB:       0,
					PartNo:       "",
					Status:       MemoryModuleStatusNotPresent,
				},
				{
					CPUNumber:    2,
					ModuleNumber: 4,
					SizeKB:       16777216,
					PartNo:       "",
					Status:       MemoryModuleStatusGood,
				},
				{
					CPUNumber:    2,
					ModuleNumber: 5,
					SizeKB:       16777216,
					PartNo:       "",
					Status:       MemoryModuleStatusGood,
				},
				{
					CPUNumber:    2,
					ModuleNumber: 6,
					SizeKB:       0,
					PartNo:       "",
					Status:       MemoryModuleStatusNotPresent,
				},
				{
					CPUNumber:    2,
					ModuleNumber: 7,
					SizeKB:       0,
					PartNo:       "",
					Status:       MemoryModuleStatusNotPresent,
				},
				{
					CPUNumber:    2,
					ModuleNumber: 8,
					SizeKB:       16777216,
					PartNo:       "",
					Status:       MemoryModuleStatusGood,
				},
				{
					CPUNumber:    2,
					ModuleNumber: 9,
					SizeKB:       16777216,
					PartNo:       "",
					Status:       MemoryModuleStatusGood,
				},
				{
					CPUNumber:    2,
					ModuleNumber: 10,
					SizeKB:       0,
					PartNo:       "",
					Status:       MemoryModuleStatusNotPresent,
				},
				{
					CPUNumber:    2,
					ModuleNumber: 11,
					SizeKB:       16777216,
					PartNo:       "",
					Status:       MemoryModuleStatusGood,
				},
				{
					CPUNumber:    2,
					ModuleNumber: 12,
					SizeKB:       16777216,
					PartNo:       "",
					Status:       MemoryModuleStatusGood,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			mib := newTestingMIB(t, test.Generation)
			modules, err := mib.MemoryModules()
			require.NoError(t, err, "failed to retrieve memory modules from the MIB")
			assert.Equal(t, test.Expected, modules)
		})
	}
}

func TestMIB_PowerSupplies(t *testing.T) {
	tests := []struct {
		Name       string
		Expected   []PowerSupply
		Generation int
	}{
		{
			Name:       "ProLiant DL380 Generation 7 Power Supplies",
			Generation: 7,
			Expected: []PowerSupply{
				{
					ChassisNo:             0,
					BayNo:                 1,
					Condition:             StatusOK,
					Status:                PowerSupplyStatusNoError,
					SerialNo:              "5AQNB0C4D1A5RQ",
					Model:                 "512327-B21",
					PowerRatingWatts:      750,
					PowerConsumptionWatts: 105,
				},
				{
					ChassisNo:             0,
					BayNo:                 2,
					Condition:             StatusOK,
					Status:                PowerSupplyStatusNoError,
					SerialNo:              "5AQNB0C4D1A5RO",
					Model:                 "512327-B21",
					PowerRatingWatts:      750,
					PowerConsumptionWatts: 35,
				},
			},
		},
		{
			Name:       "ProLiant DL380 Generation 8 Power Supplies",
			Generation: 8,
			Expected: []PowerSupply{
				{
					ChassisNo:             0,
					BayNo:                 1,
					Condition:             StatusOK,
					Status:                PowerSupplyStatusNoError,
					SerialNo:              "5BXRF0BLL2T0DQ",
					Model:                 "656363-B21",
					PowerRatingWatts:      750,
					PowerConsumptionWatts: 50,
				},
				{
					ChassisNo:             0,
					BayNo:                 2,
					Condition:             StatusOK,
					Status:                PowerSupplyStatusNoError,
					SerialNo:              "5BXRF0BLL2T0DT",
					Model:                 "656363-B21",
					PowerRatingWatts:      750,
					PowerConsumptionWatts: 75,
				},
			},
		},
	}

	for _, test := range tests {
		mib := newTestingMIB(t, test.Generation)
		powerSupplies, err := mib.PowerSupplies()
		require.NoError(t, err, "failed to retrieve power supplies from the MIB")
		assert.Equal(t, test.Expected, powerSupplies)
	}
}

func TestMIB_PowerMeterReading(t *testing.T) {
	tests := []struct {
		Name       string
		Expected   int
		Generation int
	}{
		{
			Name:       "ProLiant DL380 Generation 7 Power Meter Reading",
			Generation: 7,
			Expected:   144,
		},
		{
			Name:       "ProLiant DL380 Generation 8 Power Meter Reading",
			Generation: 8,
			Expected:   130,
		},
	}

	for _, test := range tests {
		mib := newTestingMIB(t, test.Generation)
		watts, err := mib.PowerMeterReading()
		require.NoError(t, err, "failed to retrieve power meter reading from the MIB")
		assert.Equal(t, test.Expected, watts)
	}
}

func TestMIB_Processors(t *testing.T) {
	tests := []struct {
		Name       string
		Expected   []Processor
		Generation int
	}{
		{
			Name:       "ProLiant DL380 Generation 7 Processors",
			Generation: 7,
			Expected: []Processor{
				{
					ID:                  0,
					Name:                "Intel(R) Xeon(R) CPU X5650 @ 2.67GHz",
					MaxClockSpeedHz:     4800,
					CurrentClockSpeedHz: 2667,
					PhysicalCores:       6,
					VirtualCores:        12,
					Status:              ProcessorStatusOK,
					PowerStatus:         ProcessorPowerStatusUnknown,
				},
				{
					ID:                  1,
					Name:                "Intel(R) Xeon(R) CPU X5650 @ 2.67GHz",
					MaxClockSpeedHz:     4800,
					CurrentClockSpeedHz: 2667,
					PhysicalCores:       6,
					VirtualCores:        12,
					Status:              ProcessorStatusOK,
					PowerStatus:         ProcessorPowerStatusUnknown,
				},
			},
		},
		{
			Name:       "ProLiant DL380 Generation 8 Processors",
			Generation: 8,
			Expected: []Processor{
				{
					ID:                  0,
					Name:                "Intel(R) Xeon(R) CPU E5-2670 0 @ 2.60GHz",
					MaxClockSpeedHz:     4800,
					CurrentClockSpeedHz: 2600,
					PhysicalCores:       8,
					VirtualCores:        16,
					Status:              ProcessorStatusOK,
					PowerStatus:         ProcessorPowerStatusUnknown,
				},
				{
					ID:                  1,
					Name:                "Intel(R) Xeon(R) CPU E5-2670 0 @ 2.60GHz",
					MaxClockSpeedHz:     4800,
					CurrentClockSpeedHz: 2600,
					PhysicalCores:       8,
					VirtualCores:        16,
					Status:              ProcessorStatusOK,
					PowerStatus:         ProcessorPowerStatusUnknown,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			mib := newTestingMIB(t, test.Generation)
			processors, err := mib.Processors()
			require.NoError(t, err, "failed to retrieve processors from the MIB")
			assert.Equal(t, test.Expected, processors)
		})
	}
}

func TestMIB_Model(t *testing.T) {
	tests := []struct {
		Name       string
		Generation int
		Expected   string
	}{
		{
			Name:       "ProLiant DL380 Generation 7 Model",
			Generation: 7,
			Expected:   "ProLiant DL380 G7",
		},
		{
			Name:       "ProLiant DL380 Generation 8 Model",
			Generation: 8,
			Expected:   "ProLiant DL380p Gen8",
		},
	}

	for _, test := range tests {
		mib := newTestingMIB(t, test.Generation)
		model, err := mib.Model()
		require.NoError(t, err, "failed to retrieve model from the MIB")
		assert.Equal(t, test.Expected, model)
	}
}

func TestMIB_SerialNumber(t *testing.T) {
	tests := []struct {
		Name       string
		Generation int
		Expected   string
	}{
		{
			Name:       "ProLiant DL380 Generation 7 Serial Number",
			Generation: 7,
			Expected:   "CZ21470BB8",
		},
		{
			Name:       "ProLiant DL380 Generation 8 Serial Number",
			Generation: 8,
			Expected:   "USE31629DN",
		},
	}

	for _, test := range tests {
		mib := newTestingMIB(t, test.Generation)
		serialNo, err := mib.SerialNumber()
		require.NoError(t, err, "failed to serial number from the MIB")
		assert.Equal(t, test.Expected, serialNo)
	}
}

func TestMIB_TemperatureSensors(t *testing.T) {
	tests := []struct {
		Name       string
		Expected   []TemperatureSensor
		Generation int
	}{
		{
			Name:       "ProLiant DL380 Generation 7 Temperature Sensors",
			Generation: 7,
			Expected: []TemperatureSensor{
				{
					ID:                    1,
					CurrentReadingCelsius: 26,
					Locale:                TemperatureSensorLocaleAmbient,
					Threshold:             41,
					ThresholdType:         TemperatureSensorThresholdTypeCaution,
					Status:                StatusOK,
				},
				{
					ID:                    2,
					CurrentReadingCelsius: 40,
					Locale:                TemperatureSensorLocaleCPU,
					Threshold:             82,
					ThresholdType:         TemperatureSensorThresholdTypeCaution,
					Status:                StatusOK,
				},
				{
					ID:                    3,
					CurrentReadingCelsius: 40,
					Locale:                TemperatureSensorLocaleCPU,
					Threshold:             82,
					ThresholdType:         TemperatureSensorThresholdTypeCaution,
					Status:                StatusOK,
				},
				{
					ID:                    4,
					CurrentReadingCelsius: 41,
					Locale:                TemperatureSensorLocaleMemory,
					Threshold:             87,
					ThresholdType:         TemperatureSensorThresholdTypeCaution,
					Status:                StatusOK,
				},
				{
					ID:                    5,
					CurrentReadingCelsius: 42,
					Locale:                TemperatureSensorLocaleMemory,
					Threshold:             87,
					ThresholdType:         TemperatureSensorThresholdTypeCaution,
					Status:                StatusOK,
				},
				{
					ID:                    6,
					CurrentReadingCelsius: 42,
					Locale:                TemperatureSensorLocaleMemory,
					Threshold:             87,
					ThresholdType:         TemperatureSensorThresholdTypeCaution,
					Status:                StatusOK,
				},
				{
					ID:                    7,
					CurrentReadingCelsius: 43,
					Locale:                TemperatureSensorLocaleMemory,
					Threshold:             87,
					ThresholdType:         TemperatureSensorThresholdTypeCaution,
					Status:                StatusOK,
				},
				{
					ID:                    8,
					CurrentReadingCelsius: 51,
					Locale:                TemperatureSensorLocaleSystem,
					Threshold:             90,
					ThresholdType:         TemperatureSensorThresholdTypeCaution,
					Status:                StatusOK,
				},
				{
					ID:                    9,
					CurrentReadingCelsius: 45,
					Locale:                TemperatureSensorLocaleSystem,
					Threshold:             65,
					ThresholdType:         TemperatureSensorThresholdTypeCaution,
					Status:                StatusOK,
				},
				{
					ID:                    10,
					CurrentReadingCelsius: 50,
					Locale:                TemperatureSensorLocaleSystem,
					Threshold:             90,
					ThresholdType:         TemperatureSensorThresholdTypeCaution,
					Status:                StatusOK,
				},
				{
					ID:                    11,
					CurrentReadingCelsius: 43,
					Locale:                TemperatureSensorLocaleSystem,
					Threshold:             70,
					ThresholdType:         TemperatureSensorThresholdTypeCaution,
					Status:                StatusOK,
				},
				{
					ID:                    12,
					CurrentReadingCelsius: 57,
					Locale:                TemperatureSensorLocaleSystem,
					Threshold:             90,
					ThresholdType:         TemperatureSensorThresholdTypeCaution,
					Status:                StatusOK,
				},
				{
					ID:                    19,
					CurrentReadingCelsius: 32,
					Locale:                TemperatureSensorLocaleSystem,
					Threshold:             70,
					ThresholdType:         TemperatureSensorThresholdTypeCaution,
					Status:                StatusOK,
				},
				{
					ID:                    20,
					CurrentReadingCelsius: 36,
					Locale:                TemperatureSensorLocaleSystem,
					Threshold:             70,
					ThresholdType:         TemperatureSensorThresholdTypeCaution,
					Status:                StatusOK,
				},
				{
					ID:                    21,
					CurrentReadingCelsius: 40,
					Locale:                TemperatureSensorLocaleSystem,
					Threshold:             80,
					ThresholdType:         TemperatureSensorThresholdTypeCaution,
					Status:                StatusOK,
				},
				{
					ID:                    22,
					CurrentReadingCelsius: 39,
					Locale:                TemperatureSensorLocaleSystem,
					Threshold:             80,
					ThresholdType:         TemperatureSensorThresholdTypeCaution,
					Status:                StatusOK,
				},
				{
					ID:                    23,
					CurrentReadingCelsius: 48,
					Locale:                TemperatureSensorLocaleSystem,
					Threshold:             77,
					ThresholdType:         TemperatureSensorThresholdTypeCaution,
					Status:                StatusOK,
				},
				{
					ID:                    24,
					CurrentReadingCelsius: 44,
					Locale:                TemperatureSensorLocaleSystem,
					Threshold:             70,
					ThresholdType:         TemperatureSensorThresholdTypeCaution,
					Status:                StatusOK,
				},
				{
					ID:                    25,
					CurrentReadingCelsius: 41,
					Locale:                TemperatureSensorLocaleSystem,
					Threshold:             70,
					ThresholdType:         TemperatureSensorThresholdTypeCaution,
					Status:                StatusOK,
				},
				{
					ID:                    26,
					CurrentReadingCelsius: 42,
					Locale:                TemperatureSensorLocaleSystem,
					Threshold:             70,
					ThresholdType:         TemperatureSensorThresholdTypeCaution,
					Status:                StatusOK,
				},
				{
					ID:                    29,
					CurrentReadingCelsius: 40,
					Locale:                TemperatureSensorLocaleStorage,
					Threshold:             60,
					ThresholdType:         TemperatureSensorThresholdTypeCaution,
					Status:                StatusOK,
				},
				{
					ID:                    30,
					CurrentReadingCelsius: 74,
					Locale:                TemperatureSensorLocaleSystem,
					Threshold:             110,
					ThresholdType:         TemperatureSensorThresholdTypeCaution,
					Status:                StatusOK,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			mib := newTestingMIB(t, test.Generation)
			sensors, err := mib.TemperatureSensors()
			require.NoError(t, err, "failed to retrieve temperature sensors from the MIB")
			assert.Equal(t, test.Expected, sensors)
		})
	}
}

func TestMIB_ASRStatus(t *testing.T) {
	tests := []struct {
		Name       string
		Expected   Status
		Generation int
	}{
		{
			Name:       "ProLiant DL380 Generation 7 ASR Status",
			Expected:   StatusOK,
			Generation: 7,
		},
		{
			Name:       "ProLiant DL380 Generation 8 ASR Status",
			Expected:   StatusOK,
			Generation: 8,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			mib := newTestingMIB(t, test.Generation)
			asrStatus, err := mib.ASRStatus()
			require.NoError(t, err, "failed to retrieve overall status of the advanced server recovery sub-system from the MIB")
			assert.Equal(t, test.Expected, asrStatus)
		})
	}
}

func TestMIB_BackupBatteryStatus(t *testing.T) {
	tests := []struct {
		Name       string
		Expected   Status
		Generation int
	}{
		{
			Name:       "ProLiant DL380 Generation 7 Backup Battery Status",
			Expected:   StatusOther,
			Generation: 7,
		},
		{
			Name:       "ProLiant DL380 Generation 8 Backup Battery Status",
			Expected:   StatusOther,
			Generation: 8,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			mib := newTestingMIB(t, test.Generation)
			battStatus, err := mib.BackupBatteryStatus()
			require.NoError(t, err, "failed to retrieve overall status of the backup battery sub-system from the MIB")
			assert.Equal(t, test.Expected, battStatus)
		})
	}
}

func TestMIB_ControllerStatus(t *testing.T) {
	tests := []struct {
		Name       string
		Expected   Status
		Generation int
	}{
		{
			Name:       "ProLiant DL380 Generation 7 Controller Status",
			Expected:   StatusOK,
			Generation: 7,
		},
		{
			Name:       "ProLiant DL380 Generation 8 Controller Status",
			Expected:   StatusOK,
			Generation: 8,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			mib := newTestingMIB(t, test.Generation)
			controllerStatus, err := mib.ControllerStatus()
			require.NoError(t, err, "failed to retrieve overall status of the controller sub-system from the MIB")
			assert.Equal(t, test.Expected, controllerStatus)
		})
	}
}

func TestMIB_DriveArrayStatus(t *testing.T) {
	tests := []struct {
		Name       string
		Expected   Status
		Generation int
	}{
		{
			Name:       "ProLiant DL380 Generation 7 Drive Array Status",
			Expected:   StatusOK,
			Generation: 7,
		},
		{
			Name:       "ProLiant DL380 Generation 8 Drive Array Status",
			Expected:   StatusOK,
			Generation: 8,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			mib := newTestingMIB(t, test.Generation)
			daStatus, err := mib.DriveArrayStatus()
			require.NoError(t, err, "failed to retrieve overall status of the drive array sub-system from the MIB")
			assert.Equal(t, test.Expected, daStatus)
		})
	}
}

func TestMIB_EnclosureStatus(t *testing.T) {
	tests := []struct {
		Name       string
		Expected   Status
		Generation int
	}{
		{
			Name:       "ProLiant DL380 Generation 7 Enclosure Status",
			Expected:   StatusOK,
			Generation: 7,
		},
		{
			Name:       "ProLiant DL380 Generation 8 Enclosure Status",
			Expected:   StatusOK,
			Generation: 8,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			mib := newTestingMIB(t, test.Generation)
			enclosureStatus, err := mib.EnclosureStatus()
			require.NoError(t, err, "failed to retrieve overall status of the enclosure from the MIB")
			assert.Equal(t, test.Expected, enclosureStatus)
		})
	}
}

func TestMIB_FanStatus(t *testing.T) {
	tests := []struct {
		Name       string
		Expected   Status
		Generation int
	}{
		{
			Name:       "ProLiant DL380 Generation 7 Fan Status",
			Expected:   StatusOK,
			Generation: 7,
		},
		{
			Name:       "ProLiant DL380 Generation 8 Fan Status",
			Expected:   StatusOK,
			Generation: 8,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			mib := newTestingMIB(t, test.Generation)
			fanStatus, err := mib.FanStatus()
			require.NoError(t, err, "failed to retrieve overall status of the fan sub-system from the MIB")
			assert.Equal(t, test.Expected, fanStatus)
		})
	}
}

func TestMIB_MemoryStatus(t *testing.T) {
	tests := []struct {
		Name       string
		Expected   Status
		Generation int
	}{
		{
			Name:       "ProLiant DL380 Generation 7 Fan Status",
			Expected:   StatusOK,
			Generation: 7,
		},
		{
			Name:       "ProLiant DL380 Generation 8 Fan Status",
			Expected:   StatusOK,
			Generation: 8,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			mib := newTestingMIB(t, test.Generation)
			memStatus, err := mib.MemoryStatus()
			require.NoError(t, err, "failed to retrieve overall status of the memory sub-system from the MIB")
			assert.Equal(t, test.Expected, memStatus)
		})
	}
}

func TestMIB_ProcessorStatus(t *testing.T) {
	tests := []struct {
		Name       string
		Expected   Status
		Generation int
	}{
		{
			Name:       "ProLiant DL380 Generation 7 Processor Status",
			Expected:   StatusOK,
			Generation: 7,
		},
		{
			Name:       "ProLiant DL380 Generation 8 Processor Status",
			Expected:   StatusOK,
			Generation: 8,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			mib := newTestingMIB(t, test.Generation)
			processorStatus, err := mib.ProcessorStatus()
			require.NoError(t, err, "failed to retrieve overall status of the processor sub-system from the MIB")
			assert.Equal(t, test.Expected, processorStatus)
		})
	}
}

func TestMIB_PowerSupplyStatus(t *testing.T) {
	tests := []struct {
		Name       string
		Expected   Status
		Generation int
	}{
		{
			Name:       "ProLiant DL380 Generation 7 Power Supply Status",
			Expected:   StatusOK,
			Generation: 7,
		},
		{
			Name:       "ProLiant DL380 Generation 8 Power Supply Status",
			Expected:   StatusOK,
			Generation: 8,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			mib := newTestingMIB(t, test.Generation)
			psStatus, err := mib.PowerSupplyStatus()
			require.NoError(t, err, "failed to retrieve overall status of the power supply sub-system from the MIB")
			assert.Equal(t, test.Expected, psStatus)
		})
	}
}
