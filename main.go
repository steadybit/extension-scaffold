package main

import (
	"github.com/rs/zerolog/log"
	"github.com/steadybit/action-kit/go/action_kit_api/v2"
	"github.com/steadybit/discovery-kit/go/discovery_kit_api"
	"github.com/steadybit/event-kit/go/event_kit_api"
	"github.com/steadybit/extension-kit/exthttp"
	"github.com/steadybit/extension-kit/extlogging"
	"github.com/steadybit/extension-scaffold/extconfig"
	"github.com/steadybit/extension-scaffold/version"
)

func main() {
	// Most Steadybit extensions leverage zerolog. To encourage persistent logging setups across extensions,
	// you may leverage the extlogging package to initialize zerolog. Among others, this package supports
	// configuration of active log levels and the log format (JSON or plain text).
	extlogging.InitZeroLog()

	// Many extensions require some form of configuration. These calls exist to parse and validate the
	// configuration obtained from environment variables.
	extconfig.ParseConfiguration()
	extconfig.ValidateConfiguration()

	log.Info().Msgf("Starting with version %s", version.Get())

	// This call registers a handler for the extension's root path. This is the path initially accessed
	// by the Steadybit agent to obtain the extension's capabilities.
	exthttp.RegisterHttpHandler("/", exthttp.GetterAsHandler(getExtensionList))

	exthttp.Listen(exthttp.ListenOpts{
		// This is the default port under which your extension is accessible.
		// The port can be configured externally through the
		// STEADYBIT_EXTENSION_PORT environment variable.
		Port: 7070,
	})
}

// ExtensionListResponse exists to merge the possible root path responses supported by the
// various extension kits. In this case, the response for ActionKit, DiscoveryKit and EventKit.
type ExtensionListResponse struct {
	action_kit_api.ActionList       `json:",inline"`
	discovery_kit_api.DiscoveryList `json:",inline"`
	event_kit_api.EventListenerList `json:",inline"`
}

func getExtensionList() ExtensionListResponse {
	return ExtensionListResponse{
		// See this document to learn more about the action list:
		// https://github.com/steadybit/action-kit/blob/main/docs/action-api.md#action-list
		ActionList: action_kit_api.ActionList{
			Actions: []action_kit_api.DescribingEndpointReference{
				{
					"GET",
					"/robot/actions/echo",
				},
			},
		},

		// See this document to learn more about the discovery list:
		// https://github.com/steadybit/discovery-kit/blob/main/docs/discovery-api.md#index-response
		DiscoveryList: discovery_kit_api.DiscoveryList{
			Discoveries: []discovery_kit_api.DescribingEndpointReference{
				{
					"GET",
					"/robot/discoveries",
				},
			},
			TargetTypes: []discovery_kit_api.DescribingEndpointReference{
				{
					"GET",
					"/robot/discoveries/type",
				},
			},
			TargetAttributes: []discovery_kit_api.DescribingEndpointReference{
				{
					"GET",
					"/robot/discoveries/attributes",
				},
			},
		},

		// See this document to learn more about the event listener list:
		// https://github.com/steadybit/event-kit/blob/main/docs/event-api.md#event-listeners-list
		EventListenerList: event_kit_api.EventListenerList{
			EventListeners: []event_kit_api.EventListener{
				{
					Method:   "POST",
					Path:     "/events/all",
					ListenTo: []string{"*"},
				},
			},
		},
	}
}
