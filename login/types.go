// This source code file is AUTO-GENERATED by github.com/taskcluster/jsonschema2go

package login

type (
	// A response containing temporary credentials.
	//
	// See http://schemas.taskcluster.net/login/v1/credentials-response.json#
	CredentialsResponse struct {

		// Syntax:     ^[a-zA-Z0-9_-]{22,66}$
		//
		// See http://schemas.taskcluster.net/login/v1/credentials-response.json#/properties/accessToken
		AccessToken string `json:"accessToken"`

		// See http://schemas.taskcluster.net/login/v1/credentials-response.json#/properties/certificate
		Certificate string `json:"certificate"`

		// Syntax:     ^[A-Za-z0-9@/:._-]+$
		//
		// See http://schemas.taskcluster.net/login/v1/credentials-response.json#/properties/clientId
		ClientID string `json:"clientId"`
	}

	// See http://schemas.taskcluster.net/login/v1/persona-request.json#
	PersonaAssertionRequest struct {

		// The Persona assertion from `navigator.id.get`
		//
		// See http://schemas.taskcluster.net/login/v1/persona-request.json#/properties/assertion
		Assertion string `json:"assertion"`

		// The audience against which to verify the assertion, in the format
		// `https://site.com:443`.  This must be from a whitelist of sites
		// configured in the login service.
		//
		// See http://schemas.taskcluster.net/login/v1/persona-request.json#/properties/audience
		Audience string `json:"audience"`
	}
)
