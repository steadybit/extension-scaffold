package robots

import (
	"github.com/steadybit/discovery-kit/go/discovery_kit_api"
	"github.com/steadybit/extension-kit/exthttp"
	"github.com/steadybit/extension-kit/extutil"
	"github.com/steadybit/extension-scaffold/extconfig"
	"net/http"
)

func RegisterRobotDiscoveryHandlers() {
	exthttp.RegisterHttpHandler("/robot/discovery", exthttp.GetterAsHandler(getRobotDiscoveryDescription))
	exthttp.RegisterHttpHandler("/robot/discovery/target-description", exthttp.GetterAsHandler(getRobotTargetDescription))
	exthttp.RegisterHttpHandler("/robot/discovery/attribute-descriptions", exthttp.GetterAsHandler(getRobotAttributeDescriptions))
	exthttp.RegisterHttpHandler("/robot/discovery/discovered-targets", getDiscoveredRobots)
}

func getRobotDiscoveryDescription() discovery_kit_api.DiscoveryDescription {
	return discovery_kit_api.DiscoveryDescription{
		Id:         robotTargetId,
		RestrictTo: extutil.Ptr(discovery_kit_api.LEADER),
		Discover: discovery_kit_api.DescribingEndpointReferenceWithCallInterval{
			Method:       "GET",
			Path:         "/robot/discovery/discovered-targets",
			CallInterval: extutil.Ptr("10m"),
		},
	}
}

func getRobotTargetDescription() discovery_kit_api.TargetDescription {
	return discovery_kit_api.TargetDescription{
		Id:       robotTargetId,
		Label:    discovery_kit_api.PluralLabel{One: "Robot", Other: "Robots"},
		Category: extutil.Ptr("example"),
		Version:  "1.0.0-SNAPSHOT",
		Icon:     extutil.Ptr(robotIcon),
		Table: discovery_kit_api.Table{
			Columns: []discovery_kit_api.Column{
				{Attribute: "steadybit.label"},
				{Attribute: "robot.reportedBy"},
			},
			OrderBy: []discovery_kit_api.OrderBy{
				{
					Attribute: "steadybit.label",
					Direction: "ASC",
				},
			},
		},
	}
}

func getRobotAttributeDescriptions() discovery_kit_api.AttributeDescriptions {
	return discovery_kit_api.AttributeDescriptions{
		Attributes: []discovery_kit_api.AttributeDescription{
			{
				Attribute: "robot.reportedBy",
				Label: discovery_kit_api.PluralLabel{
					One:   "Reported by",
					Other: "Reported by",
				},
			},
		},
	}
}

func getDiscoveredRobots(w http.ResponseWriter, r *http.Request, _ []byte) {
	targets := make([]discovery_kit_api.Target, len(extconfig.Config.RobotNames))
	for i, name := range extconfig.Config.RobotNames {
		targets[i] = discovery_kit_api.Target{
			Id:         name,
			TargetType: robotTargetId,
			Label:      name,
			Attributes: map[string][]string{"robot.reportedBy": {"extension-scaffold"}},
		}
	}
	exthttp.WriteBody(w, discovery_kit_api.DiscoveredTargets{Targets: targets})
}
