// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2022 Steadybit GmbH

package extevents

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"github.com/steadybit/event-kit/go/event_kit_api"
	extension_kit "github.com/steadybit/extension-kit"
	"github.com/steadybit/extension-kit/exthttp"
	"net/http"
)

func RegisterEventListenerHandlers() {
	exthttp.RegisterHttpHandler("/events/all", onEvent)
}

func onEvent(w http.ResponseWriter, r *http.Request, body []byte) {
	var event event_kit_api.EventRequestBody
	err := json.Unmarshal(body, &event)
	if err != nil {
		exthttp.WriteError(w, extension_kit.ToError("Failed to decode event request body", err))
		return
	}

	log.Info().Msgf("Received event %s", event.EventName)

	exthttp.WriteBody(w, "{}")
}
