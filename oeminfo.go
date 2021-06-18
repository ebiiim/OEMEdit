package main

import (
	"fmt"

	"golang.org/x/sys/windows/registry"
)

type OEMInfo struct {
	Logo         string `yaml:"Logo"`
	Manufacturer string `yaml:"Manufacturer"`
	Model        string `yaml:"Model"`
	SupportHours string `yaml:"SupportHours"`
	SupportPhone string `yaml:"SupportPhone"`
	SupportURL   string `yaml:"SupportURL"`
}

func NewOEMInfo(logo, mfr, model, hours, phone, url string) OEMInfo {
	return OEMInfo{
		Logo:         logo,
		Manufacturer: mfr,
		Model:        model,
		SupportHours: hours,
		SupportPhone: phone,
		SupportURL:   url,
	}
}

const (
	// Registry keys
	RKOEMInfo = `SOFTWARE\Microsoft\Windows\CurrentVersion\OEMInformation`
	// Registry values (all REG_SZ)
	RVLogo  = "Logo"
	RVMfr   = "Manufacturer"
	RVModel = "Model"
	RVHours = "SupportHours"
	RVPhone = "SupportPhone"
	RVURL   = "SupportURL"
)

func GetOEMInfo() (_ OEMInfo, err error) {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, RKOEMInfo, registry.QUERY_VALUE)
	if err != nil {
		return OEMInfo{}, fmt.Errorf("could not open registry key %s: %w", RKOEMInfo, err)
	}
	defer func() {
		if cErr := k.Close(); cErr != nil {
			err = fmt.Errorf("could not close registry key %s (the original err: %v): %w", RKOEMInfo, err, cErr)
		}
	}()

	logo, err1 := getRegValue(k, RVLogo)
	mfr, err2 := getRegValue(k, RVMfr)
	model, err3 := getRegValue(k, RVModel)
	hours, err4 := getRegValue(k, RVHours)
	phone, err5 := getRegValue(k, RVPhone)
	url, err6 := getRegValue(k, RVURL)

	if err1 != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil || err6 != nil {
		return OEMInfo{}, fmt.Errorf("could not get registry values (%v, %v, %v, %v, %v, %v)", err1, err2, err3, err4, err5, err6)
	}
	return NewOEMInfo(logo, mfr, model, hours, phone, url), nil
}

func SetOEMInfo(o OEMInfo) (err error) {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, RKOEMInfo, registry.QUERY_VALUE|registry.SET_VALUE)
	if err != nil {
		return fmt.Errorf("could not open registry key %s: %w", RKOEMInfo, err)
	}
	defer func() {
		if cErr := k.Close(); cErr != nil {
			err = fmt.Errorf("could not close registry key %s (the original err: %v): %w", RKOEMInfo, err, cErr)
		}
	}()

	err1 := setOrDeleteRegValue(k, RVLogo, o.Logo)
	err2 := setOrDeleteRegValue(k, RVMfr, o.Manufacturer)
	err3 := setOrDeleteRegValue(k, RVModel, o.Model)
	err4 := setOrDeleteRegValue(k, RVHours, o.SupportHours)
	err5 := setOrDeleteRegValue(k, RVPhone, o.SupportPhone)
	err6 := setOrDeleteRegValue(k, RVURL, o.SupportURL)

	if err1 != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil || err6 != nil {
		return fmt.Errorf("could not set registry values (%v, %v, %v, %v, %v, %v)", err1, err2, err3, err4, err5, err6)
	}
	return nil
}
