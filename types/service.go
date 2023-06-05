package types

type Service struct {
	BaseObject
	// The name of the service
	Name string `json:"name"`
	// The user-provided description of the service.
	Description string `json:"description"`
	// Time in seconds that an incident is automatically resolved if left open for that long.
	// Value is null if the feature is disabled. Value must not be negative. Setting this field to 0, null (or unset in POST request) will disable the feature.
	// Default:  14400
	AutoResolveTimeout int `json:"auto_resolve_timeout"`
	// Time in seconds that an incident changes to the Triggered State after being Acknowledged. Value is null if the feature is disabled. Value must not be negative. Setting this field to 0, null (or unset in POST request) will disable the feature.
	// Default: 1800
	AcknowledgementTimeout int `json:"acknowledgement_timeout"`
	// The date/time when this service was created
	CreateAt string `json:"create_at"`
	// The current state of the Service. Valid statuses are:Show all...
	//    Allowed values: active warning critical maintenance disabled
	//    Default: active
	Status ServiceStatus `json:"status"`
	// The date/time when the most recent incident was created for this service.
	LastIncidentTimestamp string              `json:"last_incident_timestamp"`
	EscalationPolicy      BaseObjectReference `json:"escalation_policy"`
	// Response plays associated with this service.
	ResponsePlay BaseObjectReference `json:"response_play"`
	// The set of teams associated with this service.
	Teams []BaseObjectReference `json:"teams"`
	// An array containing Integration objects that belong to this service.
	// If integrations is passed as an argument, these are full objects - otherwise, these are references.
	Integrations        []BaseObjectReference `json:"integrations"`
	IncidentUrgencyRule IncidentUrgencyRule   `json:"incident_urgency_rule"`
	SupportHours        SupportHours          `json:"support_hours"`
	// An array containing scheduled actions for the service.
	ScheduleActions []ScheduleAction `json:"schedule_actions"`
	// The array of Add-ons associated with this service.
	Addons []Addon `json:"addons"`
	// Whether a service creates only incidents, or both alerts and incidents. A service must create alerts in order to enable incident merging.
	//
	// "create_incidents" - The service will create one incident and zero alerts for each incoming event.
	// "create_alerts_and_incidents" - The service will create one incident and one associated alert for each incoming event.
	// Allowed values: create_incidents  create_alerts_and_incidents
	AlertCreation string `json:"alert_creation"`
	// Defines how alerts on this service will be automatically grouped into incidents.
	// Note that the alert grouping features are available only on certain plans. To turn grouping off set the type to null.
	AlertGroupingParameters AlertGroupingParameters `json:"alert_grouping_parameters"`
	// Defines how alerts on this service will be automatically grouped into incidents. Note that the alert grouping features are available only on certain plans. There are three available options:
	//
	// null - No alert grouping on the service. Each alert will create a separate incident;
	// "time" - All alerts within a specified duration will be grouped into the same incident. This duration is set in the alert_grouping_timeout setting (described below). Available on Standard, Enterprise, and Event Intelligence plans;
	// "intelligent" - Alerts will be intelligently grouped based on a machine learning model that looks at the alert summary, timing, and the history of grouped alerts. Available on Enterprise and Event Intelligence plans
	// Allowed values: time intelligent
	AlertGrouping string `json:"alert_grouping"`
	// The duration in minutes within which to automatically group incoming alerts.
	// This setting applies only when alert_grouping is set to time.
	// To continue grouping alerts until the Incident is resolved, set this value to 0.
	AlertGroupingTimeout int `json:"alert_grouping_timeout"`
	// Defines how alerts on this service are automatically suspended for a period of time before triggering,
	// when identified as likely being transient.
	// Note that automatically pausing notifications is only available on certain plans.
	//
	//Example: {"enabled":true,"timeout":300}
	AutoPauseNotificationsParameters AutoPauseNotificationsParameters `json:"auto_pause_notifications_parameters"`
}

type CreateAndUpdateServicePayload struct {
	Name                   string `json:"name"`
	Description            string `json:"description"`
	AutoResolve            bool   `json:"auto_resolve_timeout"`
	AcknowledgementTimeout int    `json:"acknowledgement_timeout"`
}

type GetServiceResp struct {
	Service *Service `json:"service"`
}

type ServiceStatus string

const (
	ServiceStatusActive      ServiceStatus = "active"
	ServiceStatusWarning     ServiceStatus = "warning"
	ServiceStatusCritical    ServiceStatus = "critical"
	ServiceStatusMaintenance ServiceStatus = "maintenance"
	ServiceStatusDisabled    ServiceStatus = "disabled"
)

func (ss ServiceStatus) String() string {
	return string(ss)
}
