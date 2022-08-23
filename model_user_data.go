/*
 * DevCycle Bucketing API
 *
 * Documents the DevCycle Bucketing API which provides and API interface to User Bucketing and for generated SDKs.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package devcycle

type UserData struct {
	// Unique id to identify the user
	UserId string `json:"user_id"`
	// User's email used to identify the user on the dashboard / target audiences
	Email string `json:"email,omitempty"`
	// User's name used to identify the user on the dashboard / target audiences
	Name string `json:"name,omitempty"`
	// User's language in ISO 639-1 format
	Language string `json:"language,omitempty"`
	// User's country in ISO 3166 alpha-2 format
	Country string `json:"country,omitempty"`
	// App Version of the running application
	AppVersion string `json:"appVersion,omitempty"`
	// App Build number of the running application
	AppBuild string `json:"appBuild,omitempty"`
	// User's custom data to target the user with, data will be logged to DevCycle for use in dashboard.
	CustomData map[string]interface{} `json:"customData,omitempty"`
	// User's custom data to target the user with, data will not be logged to DevCycle only used for feature bucketing.
	PrivateCustomData map[string]interface{} `json:"privateCustomData,omitempty"`
	// Date the user was created, Unix epoch timestamp format
	CreatedDate float64 `json:"createdDate,omitempty"`
	// Date the user was created, Unix epoch timestamp format
	LastSeenDate float64 `json:"lastSeenDate,omitempty"`
	// Platform the Client SDK is running on
	Platform string `json:"platform,omitempty"`
	// Version of the platform the Client SDK is running on
	PlatformVersion string `json:"platformVersion,omitempty"`
	// User's device model
	DeviceModel string `json:"deviceModel,omitempty"`
	// DevCycle SDK type
	SdkType string `json:"sdkType,omitempty"`
	// DevCycle SDK Version
	SdkVersion string `json:"sdkVersion,omitempty"`
}
