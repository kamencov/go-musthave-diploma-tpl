package main

import "testing"

func TestConfigs_Parsed(t *testing.T) {
	tests := []struct {
		name           string
		runAddress     string
		logLevel       string
		addrConDB      string
		accrualAddress string
	}{
		{
			name:           "Successful_parsing",
			runAddress:     ":8080",
			logLevel:       "info",
			addrConDB:      "",
			accrualAddress: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf := NewConfig()
			conf.RunAddress = tt.runAddress
			conf.LogLevel = tt.logLevel
			conf.AddrConDB = tt.addrConDB
			conf.AccrualSystemAddress = tt.accrualAddress

			conf.Parsed()

			if conf.RunAddress != tt.runAddress {
				t.Errorf("expected RunAddress '%s', got '%s'", tt.runAddress, conf.RunAddress)
			}
			if conf.LogLevel != tt.logLevel {
				t.Errorf("expected LogLevel '%s', got '%s'", tt.logLevel, conf.LogLevel)
			}
			if conf.AddrConDB != tt.addrConDB {
				t.Errorf("expected AddrConDB '%s', got '%s'", tt.addrConDB, conf.AddrConDB)
			}
			if conf.AccrualSystemAddress != tt.accrualAddress {
				t.Errorf("expected AccrualSystemAddress '%s', got '%s'", tt.accrualAddress, conf.AccrualSystemAddress)
			}
		})
	}
}
