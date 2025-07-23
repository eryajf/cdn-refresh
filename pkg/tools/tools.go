package tools

// AliGetUrls 单次提交多个 URL 或多个目录时，使用换行符（\n）或（\r\n）分隔
func AliGetUrls(urls []string) string {
	var result string
	for _, url := range urls {
		result += url + "\n"
	}
	return result
}

// AliGetRefreshType
func AliGetRefreshType(rtype string) string {
	switch rtype {
	case "url":
		return "file"
	case "path":
		return "directory"
	default:
		return ""
	}
}

// TcGetRefreshType
func TcGetRefreshType(rtype string) string {
	switch rtype {
	case "url":
		return "purge_url"
	case "path":
		return "purge_prefix"
	default:
		return ""
	}
}

// StringSliceToInterfaceSlice converts []*string to []any
func StringSliceToInterfaceSlice(strs []*string) []any {
	result := make([]any, len(strs))
	for i, s := range strs {
		result[i] = s
	}
	return result
}
