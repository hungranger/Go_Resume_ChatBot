package injector

import (
	"github.com/google/go-cloud/wire"
	"github.com/hungranger/Go_Resume_ChatBot/pkg/config"
	"github.com/hungranger/Go_Resume_ChatBot/pkg/server"
)

var provideConfig = wire.NewSet(config.ProvideIniConfig)
var provideServer = wire.NewSet(server.ProvideServer)

// compose providers in one big set
var providers = wire.NewSet(provideConfig, provideServer)
