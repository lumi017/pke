/*
 * Pipeline API
 *
 * Pipeline v0.3.0 swagger
 *
 * API version: master
 * Contact: info@banzaicloud.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pipeline

type NodePoolsPke struct {
	Name string `json:"name"`
	Roles []string `json:"roles"`
	Provider string `json:"provider"`
	ProviderConfig map[string]interface{} `json:"providerConfig"`
	Hosts []PkeHosts `json:"hosts"`
}
