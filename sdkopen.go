package sdkopen_go

import (
	"database/sql"

	"github.com/sdkopen/sdkopen-go/common/env"
	"github.com/sdkopen/sdkopen-go/common/observer"
	"github.com/sdkopen/sdkopen-go/database"
	"github.com/sdkopen/sdkopen-go/messaging"
	"github.com/sdkopen/sdkopen-go/validator"
	"github.com/sdkopen/sdkopen-go/webserver"
)

type SdkOpenOptions struct {
	Database  func() *sql.DB
	Messaging func() *messaging.Provider
	WebServer func() webserver.Server
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
		go messaging.StartConsumer()
	}

	if opts.WebServer != nil {
		webserver.ListenAndServe(opts.WebServer)
	}
}
