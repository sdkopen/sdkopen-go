package sdkopen_go

import (
	"github.com/sdkopen/sdkopen-go/observer"
	"github.com/sdkopen/sdkopen-go/restserver"
	"github.com/sdkopen/sdkopen-go/validator"
)

func init() {
	validator.Initialize()
	observer.Initialize()
}

func Initialize() {
	restserver.ListenAndServe(restserver.CreateChiServer)
}
