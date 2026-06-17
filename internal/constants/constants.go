package constants

import "errors"

var (
	ErrUploadNotFound = errors.New("upload not found")
	ErrFlowNotFound   = errors.New("flow not found")
	ErrDuplicateHash  = errors.New("duplicate file hash (use force=true to re-parse)")
	ErrFileTooLarge   = errors.New("file exceeds maximum size")
)
