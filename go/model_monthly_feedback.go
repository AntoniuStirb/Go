/*
 * DevNest Portal API (OpenAPI 3.0)
 *
 * The DevNest Portal API endpoints definitions based on the OpenAPI 3.0 specification.
 *
 * API version: 1.0.0
 * Contact: contact@devnest.ro
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type MonthlyFeedback struct {

	Id float64 `json:"id,omitempty"`

	Month string `json:"month,omitempty"`

	TeamAtmosphere float64 `json:"teamAtmosphere,omitempty"`

	Workload float64 `json:"workload,omitempty"`

	ClientCommunication float64 `json:"clientCommunication,omitempty"`

	MyInvolvementInTheProject float64 `json:"myInvolvementInTheProject,omitempty"`

	TechnicalResults float64 `json:"technicalResults,omitempty"`

	Comments string `json:"comments,omitempty"`
}
