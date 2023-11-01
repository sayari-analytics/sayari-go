package sdk

import "net/url"

func EncodeRecordID(recordID string) string {
	return url.QueryEscape(recordID)
}
