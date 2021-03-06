// This source code file is AUTO-GENERATED by github.com/taskcluster/jsonschema2go

package pulse

import (
	"encoding/json"
	"errors"
)

type (
	// Namespace creation request
	//
	// See http://schemas.taskcluster.net/pulse/v1/namespace-request.json#
	NamespaceCreationRequest struct {

		// The contact information which will be handed off to the notification service
		//
		// See http://schemas.taskcluster.net/pulse/v1/namespace-request.json#/properties/contact
		Contact json.RawMessage `json:"contact"`
	}

	// Namespace creation response
	//
	// See http://schemas.taskcluster.net/pulse/v1/namespace-response.json#
	NamespaceCreationResponse struct {

		// The contact information which will be handed off to the notification service
		//
		// See http://schemas.taskcluster.net/pulse/v1/namespace-response.json#/properties/contact
		Contact json.RawMessage `json:"contact"`

		// The name of the namespace created
		//
		// See http://schemas.taskcluster.net/pulse/v1/namespace-response.json#/properties/namespace
		Namespace string `json:"namespace"`

		// The password created for authentication
		//
		// See http://schemas.taskcluster.net/pulse/v1/namespace-response.json#/properties/password
		Password string `json:"password"`

		// The username created for authentication
		//
		// See http://schemas.taskcluster.net/pulse/v1/namespace-response.json#/properties/username
		Username string `json:"username"`
	}

	// Request to post a message on IRC.
	//
	// See http://schemas.taskcluster.net/pulse/v1/irc-request.json#
	PostIRCMessageRequest struct {

		// The contact method (eg. irc)
		//
		// Possible values:
		//   * "irc"
		//
		// See http://schemas.taskcluster.net/pulse/v1/irc-request.json#/properties/method
		Method string `json:"method"`

		// Details for the contact method (eg. irc channel)
		//
		// See http://schemas.taskcluster.net/pulse/v1/irc-request.json#/properties/payload
		Payload struct {

			// Channel to post the message in. Please note that you **must** supply
			// either `user` or `channel`, you cannot supply both.
			//
			// Syntax:     ^[#&][^ ,\u0007]{1,199}$
			//
			// See http://schemas.taskcluster.net/pulse/v1/irc-request.json#/properties/payload/properties/channel
			Channel string `json:"channel,omitempty"`

			// User to post the message to. Please note that you **must** supply
			// either `user` or `channel`, you cannot supply both.
			//
			// Syntax:     ^[A-Za-z\[\]\\~_\^{|}][A-Za-z0-9\-\[\]\\~_\^{|}]{0,254}$
			// Min length: 1
			// Max length: 255
			//
			// See http://schemas.taskcluster.net/pulse/v1/irc-request.json#/properties/payload/properties/user
			User string `json:"user,omitempty"`
		} `json:"payload"`
	}

	// An array of RabbitMQ exchanges containing the details of RabbitMQ exchanges
	//
	// See http://schemas.taskcluster.net/pulse/v1/exchanges-response.json#
	RabbitMQExchanges []struct {

		// See http://schemas.taskcluster.net/pulse/v1/exchanges-response.json#/items/properties/arguments
		Arguments struct {
		} `json:"arguments,omitempty"`

		// Whether or not the exchange deletes when all queues are finished using it
		//
		// See http://schemas.taskcluster.net/pulse/v1/exchanges-response.json#/items/properties/auto-delete
		AutoDelete bool `json:"auto-delete,omitempty"`

		// Whether or not the exchange survives broker restart
		//
		// See http://schemas.taskcluster.net/pulse/v1/exchanges-response.json#/items/properties/durable
		Durable bool `json:"durable,omitempty"`

		// The exchange's name
		//
		// See http://schemas.taskcluster.net/pulse/v1/exchanges-response.json#/items/properties/name
		Name string `json:"name,omitempty"`

		// The exchange's type
		//
		// See http://schemas.taskcluster.net/pulse/v1/exchanges-response.json#/items/properties/type
		Type string `json:"type,omitempty"`

		// The exchange's vhost
		//
		// See http://schemas.taskcluster.net/pulse/v1/exchanges-response.json#/items/properties/vhost
		Vhost string `json:"vhost,omitempty"`
	}

	// Rabbit overview response
	//
	// See http://schemas.taskcluster.net/pulse/v1/rabbit-overview.json#
	RabbitOverviewResponse struct {

		// The name of the cluster
		//
		// Min length: 1
		//
		// See http://schemas.taskcluster.net/pulse/v1/rabbit-overview.json#/properties/cluster_name
		Cluster_Name string `json:"cluster_name"`

		// The version of the management
		//
		// Min length: 1
		//
		// See http://schemas.taskcluster.net/pulse/v1/rabbit-overview.json#/properties/management_version
		Management_Version string `json:"management_version"`

		// The version of RabbitMQ
		//
		// Min length: 1
		//
		// See http://schemas.taskcluster.net/pulse/v1/rabbit-overview.json#/properties/rabbitmq_version
		Rabbitmq_Version string `json:"rabbitmq_version"`
	}

	// Request to send an email
	//
	// See http://schemas.taskcluster.net/pulse/v1/email-request.json#
	SendEmailRequest struct {

		// The contact method (eg. email)
		//
		// Possible values:
		//   * "email"
		//
		// See http://schemas.taskcluster.net/pulse/v1/email-request.json#/properties/method
		Method string `json:"method"`

		// Details for the contact method (eg. email address)
		//
		// See http://schemas.taskcluster.net/pulse/v1/email-request.json#/properties/payload
		Payload struct {

			// E-mail address to which the message should be sent
			//
			// See http://schemas.taskcluster.net/pulse/v1/email-request.json#/properties/payload/properties/address
			Address string `json:"address"`

			// Reply-to e-mail (this property is optional)
			//
			// See http://schemas.taskcluster.net/pulse/v1/email-request.json#/properties/payload/properties/replyTo
			ReplyTo string `json:"replyTo,omitempty"`
		} `json:"payload"`
	}

	// See http://schemas.taskcluster.net/pulse/v1/irc-request.json#/properties/payload/oneOf[0]
	Var json.RawMessage

	// See http://schemas.taskcluster.net/pulse/v1/irc-request.json#/properties/payload/oneOf[1]
	Var1 json.RawMessage

	// See http://schemas.taskcluster.net/pulse/v1/namespace-request.json#/properties/contact/oneOf[0]
	Var2 PostIRCMessageRequest

	// See http://schemas.taskcluster.net/pulse/v1/namespace-request.json#/properties/contact/oneOf[1]
	Var3 SendEmailRequest

	// See http://schemas.taskcluster.net/pulse/v1/namespace-response.json#/properties/contact/oneOf[0]
	Var4 PostIRCMessageRequest

	// See http://schemas.taskcluster.net/pulse/v1/namespace-response.json#/properties/contact/oneOf[1]
	Var5 SendEmailRequest
)

// MarshalJSON calls json.RawMessage method of the same name. Required since
// Var is of type json.RawMessage...
func (this *Var) MarshalJSON() ([]byte, error) {
	x := json.RawMessage(*this)
	return (&x).MarshalJSON()
}

// UnmarshalJSON is a copy of the json.RawMessage implementation.
func (this *Var) UnmarshalJSON(data []byte) error {
	if this == nil {
		return errors.New("Var: UnmarshalJSON on nil pointer")
	}
	*this = append((*this)[0:0], data...)
	return nil
}

// MarshalJSON calls json.RawMessage method of the same name. Required since
// Var1 is of type json.RawMessage...
func (this *Var1) MarshalJSON() ([]byte, error) {
	x := json.RawMessage(*this)
	return (&x).MarshalJSON()
}

// UnmarshalJSON is a copy of the json.RawMessage implementation.
func (this *Var1) UnmarshalJSON(data []byte) error {
	if this == nil {
		return errors.New("Var1: UnmarshalJSON on nil pointer")
	}
	*this = append((*this)[0:0], data...)
	return nil
}
