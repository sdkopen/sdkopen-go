package restserver

type HttpStatusCode int

const (
	StatusContinue                      HttpStatusCode = 100
	StatusSwitchingProtocols            HttpStatusCode = 101
	StatusProcessing                    HttpStatusCode = 102
	StatusEarlyHints                    HttpStatusCode = 103
	StatusOK                            HttpStatusCode = 200
	StatusCreated                       HttpStatusCode = 201
	StatusAccepted                      HttpStatusCode = 202
	StatusNonAuthoritativeInfo          HttpStatusCode = 203
	StatusNoContent                     HttpStatusCode = 204
	StatusResetContent                  HttpStatusCode = 205
	StatusPartialContent                HttpStatusCode = 206
	StatusMultiStatus                   HttpStatusCode = 207
	StatusAlreadyReported               HttpStatusCode = 208
	StatusIMUsed                        HttpStatusCode = 226
	StatusMultipleChoices               HttpStatusCode = 300
	StatusMovedPermanently              HttpStatusCode = 301
	StatusFound                         HttpStatusCode = 302
	StatusSeeOther                      HttpStatusCode = 303
	StatusNotModified                   HttpStatusCode = 304
	StatusUseProxy                      HttpStatusCode = 305
	StatusTemporaryRedirect             HttpStatusCode = 307
	StatusPermanentRedirect             HttpStatusCode = 308
	StatusBadRequest                    HttpStatusCode = 400
	StatusUnauthorized                  HttpStatusCode = 401
	StatusPaymentRequired               HttpStatusCode = 402
	StatusForbidden                     HttpStatusCode = 403
	StatusNotFound                      HttpStatusCode = 404
	StatusMethodNotAllowed              HttpStatusCode = 405
	StatusNotAcceptable                 HttpStatusCode = 406
	StatusProxyAuthRequired             HttpStatusCode = 407
	StatusRequestTimeout                HttpStatusCode = 408
	StatusConflict                      HttpStatusCode = 409
	StatusGone                          HttpStatusCode = 410
	StatusLengthRequired                HttpStatusCode = 411
	StatusPreconditionFailed            HttpStatusCode = 412
	StatusRequestEntityTooLarge         HttpStatusCode = 413
	StatusRequestURITooLong             HttpStatusCode = 414
	StatusUnsupportedMediaType          HttpStatusCode = 415
	StatusRequestedRangeNotSatisfiable  HttpStatusCode = 416
	StatusExpectationFailed             HttpStatusCode = 417
	StatusTeapot                        HttpStatusCode = 418
	StatusMisdirectedRequest            HttpStatusCode = 421
	StatusUnprocessableEntity           HttpStatusCode = 422
	StatusLocked                        HttpStatusCode = 423
	StatusFailedDependency              HttpStatusCode = 424
	StatusTooEarly                      HttpStatusCode = 425
	StatusUpgradeRequired               HttpStatusCode = 426
	StatusPreconditionRequired          HttpStatusCode = 428
	StatusTooManyRequests               HttpStatusCode = 429
	StatusRequestHeaderFieldsTooLarge   HttpStatusCode = 431
	StatusUnavailableForLegalReasons    HttpStatusCode = 451
	StatusInternalServerError           HttpStatusCode = 500
	StatusNotImplemented                HttpStatusCode = 501
	StatusBadGateway                    HttpStatusCode = 502
	StatusServiceUnavailable            HttpStatusCode = 503
	StatusGatewayTimeout                HttpStatusCode = 504
	StatusHTTPVersionNotSupported       HttpStatusCode = 505
	StatusVariantAlsoNegotiates         HttpStatusCode = 506
	StatusInsufficientStorage           HttpStatusCode = 507
	StatusLoopDetected                  HttpStatusCode = 508
	StatusNotExtended                   HttpStatusCode = 510
	StatusNetworkAuthenticationRequired HttpStatusCode = 511
)

var httpStatusText = map[HttpStatusCode]string{
	StatusContinue:                      "Continue",
	StatusSwitchingProtocols:            "Switching Protocols",
	StatusProcessing:                    "Processing",
	StatusEarlyHints:                    "Early Hints",
	StatusOK:                            "OK",
	StatusCreated:                       "Created",
	StatusAccepted:                      "Accepted",
	StatusNonAuthoritativeInfo:          "Non Authoritative Info",
	StatusNoContent:                     "No Content",
	StatusResetContent:                  "Reset Content",
	StatusPartialContent:                "Partial Content",
	StatusMultiStatus:                   "Multi Status",
	StatusAlreadyReported:               "Already Reported",
	StatusIMUsed:                        "IM Used",
	StatusMultipleChoices:               "Multiple Choices",
	StatusMovedPermanently:              "Moved Permanently",
	StatusFound:                         "Found",
	StatusSeeOther:                      "See Other",
	StatusNotModified:                   "Not Modified",
	StatusUseProxy:                      "Use Proxy",
	StatusTemporaryRedirect:             "Temporary Redirect",
	StatusPermanentRedirect:             "Permanent Redirect",
	StatusBadRequest:                    "Bad Request",
	StatusUnauthorized:                  "Unauthorized",
	StatusPaymentRequired:               "Payment Required",
	StatusForbidden:                     "Forbidden",
	StatusNotFound:                      "Not Found",
	StatusMethodNotAllowed:              "Method Not Allowed",
	StatusNotAcceptable:                 "Not Acceptable",
	StatusProxyAuthRequired:             "Proxy Auth Required",
	StatusRequestTimeout:                "Request Timeout",
	StatusConflict:                      "Conflict",
	StatusGone:                          "Gone",
	StatusLengthRequired:                "Length Required",
	StatusPreconditionFailed:            "Precondition Failed",
	StatusRequestEntityTooLarge:         "Request Entity Too Large",
	StatusRequestURITooLong:             "Request URI Too Long",
	StatusUnsupportedMediaType:          "Unsupported Media Type",
	StatusRequestedRangeNotSatisfiable:  "Requested Range Not Satisfiable",
	StatusExpectationFailed:             "Expectation Failed",
	StatusTeapot:                        "Teapot",
	StatusMisdirectedRequest:            "Misdirected Request",
	StatusUnprocessableEntity:           "Unprocessable Entity",
	StatusLocked:                        "Locked",
	StatusFailedDependency:              "Failed Dependency",
	StatusTooEarly:                      "Too Early",
	StatusUpgradeRequired:               "Upgrade Required",
	StatusPreconditionRequired:          "Precondition Required",
	StatusTooManyRequests:               "Too Many Requests",
	StatusRequestHeaderFieldsTooLarge:   "Request Header Fields Too Large",
	StatusUnavailableForLegalReasons:    "Unavailable For Legal Reasons",
	StatusInternalServerError:           "Internal Server Error",
	StatusNotImplemented:                "Not Implemented",
	StatusBadGateway:                    "Bad Gateway",
	StatusServiceUnavailable:            "Service Unavailable",
	StatusGatewayTimeout:                "Gateway Timeout",
	StatusHTTPVersionNotSupported:       "HTTP Version Not Supported",
	StatusVariantAlsoNegotiates:         "Variant Also Negotiates",
	StatusInsufficientStorage:           "Insufficient Storage",
	StatusLoopDetected:                  "Loop Detected",
	StatusNotExtended:                   "Not Extended",
	StatusNetworkAuthenticationRequired: "Network Authentication Required",
}

func (sc HttpStatusCode) Int() int {
	return int(sc)
}

func (sc HttpStatusCode) String() string {
	if text, ok := httpStatusText[sc]; ok {
		return text
	}
	return ""
}
