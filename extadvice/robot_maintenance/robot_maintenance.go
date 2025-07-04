package robot_maintenance

import (
	"embed"
	"github.com/steadybit/advice-kit/go/advice_kit_api"
	"github.com/steadybit/extension-kit/extbuild"
	"github.com/steadybit/extension-kit/extutil"
	"github.com/steadybit/extension-scaffold/extadvice/advice_common"
	"github.com/steadybit/extension-scaffold/extrobots"
)

const RobotMaintenanceID = "com.steadybit.extension_scaffold.advice.robot-maintenance"

//go:embed *
var RobotMaintenanceContent embed.FS

func GetAdviceDescriptionRobotMaintenance() advice_kit_api.AdviceDefinition {
	return advice_kit_api.AdviceDefinition{
		Id:                        RobotMaintenanceID,
		Label:                     "Robot Maintenance",
		Version:                   extbuild.GetSemverVersionStringOrUnknown(),
		Icon:                      "data:image/svg+xml,%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20viewBox%3D%220%200%20800.55%20856.85%22%3E%3Cpath%20fill%3D%22currentColor%22%20class%3D%22st0%22%20d%3D%22M670.38%20608.27l-71.24-46.99-59.43%2099.27-69.12-20.21-60.86%2092.89%203.12%2029.24%20330.9-60.97-19.22-206.75-54.15%20113.52zm-308.59-89.14l53.09-7.3c8.59%203.86%2014.57%205.33%2024.87%207.95%2016.04%204.18%2034.61%208.19%2062.11-5.67%206.4-3.17%2019.73-15.36%2025.12-22.31l217.52-39.46%2022.19%20268.56-372.65%2067.16-32.25-268.93zm404.06-96.77l-21.47%204.09L703.13.27.27%2081.77l86.59%20702.68%2082.27-11.94c-6.57-9.38-16.8-20.73-34.27-35.26-24.23-20.13-15.66-54.32-1.37-75.91%2018.91-36.48%20116.34-82.84%20110.82-141.15-1.98-21.2-5.35-48.8-25.03-67.71-.74%207.85.59%2015.41.59%2015.41s-8.08-10.31-12.11-24.37c-4-5.39-7.14-7.11-11.39-14.31-3.03%208.33-2.63%2017.99-2.63%2017.99s-6.61-15.62-7.68-28.8c-3.92%205.9-4.91%2017.11-4.91%2017.11s-8.59-24.62-6.63-37.88c-3.92-11.54-15.54-34.44-12.25-86.49%2021.45%2015.03%2068.67%2011.46%2087.07-15.66%206.11-8.98%2010.29-33.5-3.05-81.81-8.57-30.98-29.79-77.11-38.06-94.61l-.99.71c4.36%2014.1%2013.35%2043.66%2016.8%2057.99%2010.44%2043.47%2013.24%2058.6%208.34%2078.64-4.17%2017.42-14.17%2028.82-39.52%2041.56-25.35%2012.78-58.99-18.32-61.12-20.04-24.63-19.62-43.68-51.63-45.81-67.18-2.21-17.02%209.81-27.24%2015.87-41.16-8.67%202.48-18.34%206.88-18.34%206.88s11.54-11.94%2025.77-22.27c5.89-3.9%209.35-6.38%2015.56-11.54-8.99-.15-16.29.11-16.29.11s14.99-8.1%2030.53-14c-11.37-.5-22.25-.08-22.25-.08s33.45-14.96%2059.87-25.94c18.17-7.45%2035.92-5.25%2045.89%209.17%2013.09%2018.89%2026.84%2029.15%2055.98%2035.51%2017.89-7.93%2023.33-12.01%2045.81-18.13%2019.79-21.76%2035.33-24.58%2035.33-24.58s-7.71%207.07-9.77%2018.18c11.22-8.84%2023.52-16.22%2023.52-16.22s-4.76%205.88-9.2%2015.22l1.03%201.53c13.09-7.85%2028.48-14.04%2028.48-14.04s-4.4%205.56-9.56%2012.76c9.87-.08%2029.89.42%2037.66%201.3%2045.87%201.01%2055.39-48.99%2072.99-55.26%2022.04-7.87%2031.89-12.63%2069.45%2024.26%2032.23%2031.67%2057.41%2088.36%2044.91%20101.06-10.48%2010.54-31.16-4.11-54.08-32.68-12.11-15.13-21.27-33.01-25.56-55.74-3.62-19.18-17.71-30.31-17.71-30.31S520%2092.95%20520%20109.01c0%208.77%201.1%2041.56%2015.16%2059.96-1.39%202.69-2.04%2013.31-3.58%2015.34-16.36-19.77-51.49-33.92-57.22-38.09%2019.39%2015.89%2063.96%2052.39%2081.08%2087.37%2016.19%2033.08%206.65%2063.4%2014.84%2071.25%202.33%202.25%2034.82%2042.73%2041.07%2063.07%2010.9%2035.45.65%2072.7-13.62%2095.81l-39.85%206.21c-5.83-1.62-9.76-2.43-14.99-5.46%202.88-5.1%208.61-17.82%208.67-20.44l-2.25-3.95c-12.4%2017.57-33.18%2034.63-50.44%2044.43-22.59%2012.8-48.63%2010.83-65.58%205.58-48.11-14.84-93.6-47.35-104.57-55.89%200%200-.34%206.82%201.73%208.35%2012.13%2013.68%2039.92%2038.43%2066.78%2055.68l-57.26%206.3%2027.07%20210.78c-12%201.72-13.87%202.56-27.01%204.43-11.58-40.91-33.73-67.62-57.94-83.18-21.35-13.72-50.8-16.81-78.99-11.23l-1.81%202.1c19.6-2.04%2042.74.8%2066.51%2015.85%2023.33%2014.75%2042.13%2052.85%2049.05%2075.79%208.86%2029.32%2014.99%2060.68-8.86%2093.92-16.97%2023.63-66.51%2036.69-106.53%208.44%2010.69%2017.19%2025.14%2031.25%2044.59%2033.9%2028.88%203.92%2056.29-1.09%2075.16-20.46%2016.11-16.56%2024.65-51.19%2022.4-87.66l25.49-3.7%209.2%2065.46%20421.98-50.81-34.43-335.8zM509.12%20244.59c-1.18%202.69-3.03%204.45-.25%2013.2l.17.5.44%201.13%201.16%202.62c5.01%2010.24%2010.51%2019.9%2019.7%2024.83%202.38-.4%204.84-.67%207.39-.8%208.63-.38%2014.08.99%2017.54%202.85.31-1.72.38-4.24.19-7.95-.67-12.97%202.57-35.03-22.36-46.64-9.41-4.37-22.61-3.02-27.01%202.43.8.1%201.52.27%202.08.46%206.65%202.33%202.14%204.62.95%207.37m69.87%20121.02c-3.27-1.8-18.55-1.09-29.29.19-20.46%202.41-42.55%209.51-47.39%2013.29-8.8%206.8-4.8%2018.66%201.7%2023.53%2018.23%2013.62%2034.21%2022.75%2051.08%2020.53%2010.36-1.36%2019.49-17.76%2025.96-32.64%204.43-10.25%204.43-21.31-2.06-24.9M397.85%20260.65c5.77-5.48-28.74-12.68-55.52%205.58-19.75%2013.47-20.38%2042.35-1.47%2058.72%201.89%201.62%203.45%202.77%204.91%203.71%205.52-2.6%2011.81-5.23%2019.05-7.58%2012.23-3.97%2022.4-6.02%2030.76-7.11%204-4.47%208.65-12.34%207.49-26.59-1.58-19.33-16.23-16.26-5.22-26.73%22%2F%3E%3C%2Fsvg%3E",
		Tags:                      &[]string{"robot", "maintenance"},
		AssessmentQueryApplicable: "target.type=\"" + extrobots.TargetType + "\"",
		Status: advice_kit_api.AdviceDefinitionStatus{
			ActionNeeded: advice_kit_api.AdviceDefinitionStatusActionNeeded{
				AssessmentQuery: "robot.needs.maintenance = \"true\"",
				Description: advice_kit_api.AdviceDefinitionStatusActionNeededDescription{
					Instruction: advice_common.ReadAdviceFile(RobotMaintenanceContent, "instructions.md"),
					Motivation:  advice_common.ReadAdviceFile(RobotMaintenanceContent, "motivation.md"),
					Summary:     advice_common.ReadAdviceFile(RobotMaintenanceContent, "action_needed_summary.md"),
				},
			},
			Implemented: advice_kit_api.AdviceDefinitionStatusImplemented{
				Description: advice_kit_api.AdviceDefinitionStatusImplementedDescription{
					Summary: advice_common.ReadAdviceFile(RobotMaintenanceContent, "implemented.md"),
				},
			},
			ValidationNeeded: advice_kit_api.AdviceDefinitionStatusValidationNeeded{
				Description: advice_kit_api.AdviceDefinitionStatusValidationNeededDescription{
					Summary: advice_common.ReadAdviceFile(RobotMaintenanceContent, "validation_needed.md"),
				},
				Validation: extutil.Ptr([]advice_kit_api.Validation{
					{
						Id:               RobotMaintenanceID + ".experiment-1",
						Type:             "EXPERIMENT",
						Name:             "Robot Maintenance",
						ShortDescription: "Check how ${target.attr('robot.name')} behaves when running a maintenance task.",
						Experiment:       extutil.Ptr(advice_kit_api.Experiment(advice_common.ReadAdviceFile(RobotMaintenanceContent, "experiment_robot_maintenance.json"))),
					},
				}),
			},
		},
	}
}
