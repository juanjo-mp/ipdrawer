/*
 * server/serverpb/server.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * OpenAPI spec version: version not set
 *
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package apiclient

type ServerpbGetNetworkResponse struct {
	Network string `json:"network,omitempty"`

	DefaultGateways []string `json:"default_gateways,omitempty"`

	Broadcast string `json:"broadcast,omitempty"`

	Netmask string `json:"netmask,omitempty"`

	Tags []ModelTag `json:"tags,omitempty"`
}