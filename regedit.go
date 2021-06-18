package main

import (
	"errors"
	"fmt"

	"golang.org/x/sys/windows/registry"
)

func getRegValue(k registry.Key, rv string) (string, error) {
	val, valType, err := k.GetStringValue(rv)
	if err != nil {
		if errors.Is(err, registry.ErrNotExist) {
			return "", nil
		}
		return "", fmt.Errorf("could not get %s: %w", rv, err)
	}
	if valType != registry.SZ {
		return "", fmt.Errorf("want REG_SZ but got REG_EXPAND_SZ %s", rv)
	}
	return val, nil
}

func setOrDeleteRegValue(k registry.Key, rvKey, rvValue string) error {
	if rvValue == "" {
		// delete if exists
		_, _, err := k.GetStringValue(rvKey)
		if err != nil {
			if errors.Is(err, registry.ErrNotExist) {
				return nil
			}
			return fmt.Errorf("could not get %s: %w", rvKey, err)
		}
		if err := k.DeleteValue(rvKey); err != nil {
			return fmt.Errorf("could not delete %s: %w", rvKey, err)
		}
	} else {
		// set value
		if err := k.SetStringValue(rvKey, rvValue); err != nil {
			return fmt.Errorf("could not write value %s to %s: %w", rvValue, rvKey, err)
		}

	}
	return nil
}
