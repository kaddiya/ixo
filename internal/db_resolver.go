package internal

import (
	"errors"
	"strings"
)

const (
	PGSQL = "POSTGRES"
)

func ResolveDB(connString string) (string, error) {
	if strings.Contains(connString, "pgsql") {
		return PGSQL, nil
	}
	return "", errors.New("Unsupported Database")
}
