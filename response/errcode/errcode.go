package errcode

type Code string

const (
	ErrorOpenFile        Code = "ERROR_OPEN_FILE"
	ReadFileFailed       Code = "READ_FILE_FAILED"
	UnmarshalFailed      Code = "UNMARSHAL_FAILED"
	ArchiveFailed        Code = "ARCHIVE_FAILED"
	SFTPConnectionFailed Code = "SFTP_CONNECTION_FAILED"
	XMLExtractFailed     Code = "XML_EXTRACT_FAILED"
)
