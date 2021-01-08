package model

// SQLError sql报错
type SQLError struct {
	Number  uint16
	Message string
}
