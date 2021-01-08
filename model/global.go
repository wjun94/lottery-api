package model

import (
	"job-api/utils"
)

var Utils = new(utils.Utils)

type SQLError struct {
	Number  uint16
	Message string
}
