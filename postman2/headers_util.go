package postman2

import (
	"strings"

	"github.com/andrewcretin/swaggman/openapi3"
	oas3 "github.com/getkin/kin-openapi/openapi3"
	"github.com/grokify/gotilla/net/httputilmore"
	"github.com/grokify/gotilla/type/stringsutil"
)

const DefaultMediaTypePreferences string = `multipart/form-data,application/json,application/x-www-form-urlencoded,application/xml,text/plain`

func DefaultMediaTypePreferencesSlice() []string {
	return strings.Split(DefaultMediaTypePreferences, ",")
}

func AppendPostmanHeaderValueLower(headers []Header, headerName string, options, preferenceOrder []string) ([]Header, string) {
	headerName = strings.TrimSpace(headerName)
	headerValue := stringsutil.SliceChooseOnePreferredLowerTrimSpace(options, preferenceOrder)
	if len(headerName) > 0 && len(headerValue) > 0 {
		headers = append(headers, Header{
			Key:   headerName,
			Value: headerValue})
	}
	return headers, headerValue
}

func AddOperationReqResMediaTypeHeaders(
	headers []Header,
	operation *oas3.Operation,
	reqPreferences []string,
	resPreferences []string) ([]Header, string, string) {
	headers, reqMediaType := AppendPostmanHeaderValueLower(
		headers,
		httputilmore.HeaderContentType,
		openapi3.OperationRequestMediaTypes(operation),
		reqPreferences,
	)
	headers, resMediaType := AppendPostmanHeaderValueLower(
		headers,
		httputilmore.HeaderAccept,
		openapi3.OperationResponseMediaTypes(operation),
		resPreferences,
	)
	return headers, reqMediaType, resMediaType
}
