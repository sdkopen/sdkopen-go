package sdkopen_go

import (
	"database/sql"

	"github.com/sdkopen/sdkopen-go/common/env"
	"github.com/sdkopen/sdkopen-go/database"
	"github.com/sdkopen/sdkopen-go/messaging"
	"github.com/sdkopen/sdkopen-go/observer"
	"github.com/sdkopen/sdkopen-go/restserver"
	"github.com/sdkopen/sdkopen-go/validator"
)

type SdkOpenOptions struct {
	Database   func() *sql.DB
	Messaging  func() *messaging.Provider
	RestServer func() restserver.Server
}

func init() {
	validator.Initialize()
	observer.Initialize()
}

func Initialize(opts *SdkOpenOptions) {
	env.Load()

	if opts.Database != nil {
		database.Initialize(opts.Database)
	}

	if opts.Messaging != nil {
		messaging.Initialize(opts.Messaging())
	}

	if opts.RestServer != nil {
		restserver.ListenAndServe(opts.RestServer)
	}
}
