// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// GENERATED, DO NOT EDIT.

package model

import "context"

// ClientSet is an interface that allows getting all clients.
type ClientSet interface {
	// Applications returns the client for interacting with the Application builders.
	Applications() *ApplicationClient

	// ApplicationResources returns the client for interacting with the ApplicationResource builders.
	ApplicationResources() *ApplicationResourceClient

	// ApplicationRevisions returns the client for interacting with the ApplicationRevision builders.
	ApplicationRevisions() *ApplicationRevisionClient

	// Connectors returns the client for interacting with the Connector builders.
	Connectors() *ConnectorClient

	// Environments returns the client for interacting with the Environment builders.
	Environments() *EnvironmentClient

	// Modules returns the client for interacting with the Module builders.
	Modules() *ModuleClient

	// Projects returns the client for interacting with the Project builders.
	Projects() *ProjectClient

	// Roles returns the client for interacting with the Role builders.
	Roles() *RoleClient

	// Settings returns the client for interacting with the Setting builders.
	Settings() *SettingClient

	// Subjects returns the client for interacting with the Subject builders.
	Subjects() *SubjectClient

	// Tokens returns the client for interacting with the Token builders.
	Tokens() *TokenClient

	// WithTx gives a new transactional client in the callback function,
	// if already in a transaction, this will keep in the same transaction.
	WithTx(context.Context, func(tx *Tx) error) error

	// Dialect returns the dialect name of the driver.
	Dialect() string
}

// ApplicationClientGetter is an interface that allows getting ApplicationClient.
type ApplicationClientGetter interface {
	// Applications returns the client for interacting with the Application builders.
	Applications() *ApplicationClient
}

// ApplicationResourceClientGetter is an interface that allows getting ApplicationResourceClient.
type ApplicationResourceClientGetter interface {
	// ApplicationResources returns the client for interacting with the ApplicationResource builders.
	ApplicationResources() *ApplicationResourceClient
}

// ApplicationRevisionClientGetter is an interface that allows getting ApplicationRevisionClient.
type ApplicationRevisionClientGetter interface {
	// ApplicationRevisions returns the client for interacting with the ApplicationRevision builders.
	ApplicationRevisions() *ApplicationRevisionClient
}

// ConnectorClientGetter is an interface that allows getting ConnectorClient.
type ConnectorClientGetter interface {
	// Connectors returns the client for interacting with the Connector builders.
	Connectors() *ConnectorClient
}

// EnvironmentClientGetter is an interface that allows getting EnvironmentClient.
type EnvironmentClientGetter interface {
	// Environments returns the client for interacting with the Environment builders.
	Environments() *EnvironmentClient
}

// ModuleClientGetter is an interface that allows getting ModuleClient.
type ModuleClientGetter interface {
	// Modules returns the client for interacting with the Module builders.
	Modules() *ModuleClient
}

// ProjectClientGetter is an interface that allows getting ProjectClient.
type ProjectClientGetter interface {
	// Projects returns the client for interacting with the Project builders.
	Projects() *ProjectClient
}

// RoleClientGetter is an interface that allows getting RoleClient.
type RoleClientGetter interface {
	// Roles returns the client for interacting with the Role builders.
	Roles() *RoleClient
}

// SettingClientGetter is an interface that allows getting SettingClient.
type SettingClientGetter interface {
	// Settings returns the client for interacting with the Setting builders.
	Settings() *SettingClient
}

// SubjectClientGetter is an interface that allows getting SubjectClient.
type SubjectClientGetter interface {
	// Subjects returns the client for interacting with the Subject builders.
	Subjects() *SubjectClient
}

// TokenClientGetter is an interface that allows getting TokenClient.
type TokenClientGetter interface {
	// Tokens returns the client for interacting with the Token builders.
	Tokens() *TokenClient
}
