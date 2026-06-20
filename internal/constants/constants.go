package constants

import "errors"

var (
	ErrUploadNotFound = errors.New("upload not found")
	ErrFlowNotFound   = errors.New("flow not found")
	ErrFileTooLarge   = errors.New("file exceeds maximum size")
)

const (
	BatchPending  = "pending"
	BatchRunning  = "running"
	BatchDone     = "done"
	BatchError    = "error"
)
