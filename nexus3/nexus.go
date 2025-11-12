package nexus3

import (
	"github.com/gcroucher/go-nexus-client/nexus3/pkg/blobstore"
	"github.com/gcroucher/go-nexus-client/nexus3/pkg/capability"
	"github.com/gcroucher/go-nexus-client/nexus3/pkg/cleanup"
	"github.com/gcroucher/go-nexus-client/nexus3/pkg/client"
	"github.com/gcroucher/go-nexus-client/nexus3/pkg/readonly"
	"github.com/gcroucher/go-nexus-client/nexus3/pkg/repository"
	"github.com/gcroucher/go-nexus-client/nexus3/pkg/security"
)

const (
	// ContentTypeApplicationJSON ...
	ContentTypeApplicationJSON = "application/json"
	// ContentTypeTextPlain ...
	ContentTypeTextPlain = "text/plain"
	basePath             = "service/rest/"
)

type NexusClient struct {
	client *client.Client

	// API Services
	BlobStore     *blobstore.BlobStoreService
	Capability    *capability.CapabilityService
	Repository    *repository.RepositoryService
	RoutingRule   *RoutingRuleService
	CleanupPolicy *cleanup.CleanupPolicyService
	Security      *security.SecurityService
	Script        *ScriptService
	ReadOnly      *readonly.ReadOnlyService
	MailConfig    *MailConfigService
}

// NewClient returns an instance of client that implements the Client interface
func NewClient(config client.Config) *NexusClient {
	client := client.NewClient(config)
	return &NexusClient{
		client:        client,
		BlobStore:     blobstore.NewBlobStoreService(client),
		Capability:    capability.NewCapabilityService(client),
		Repository:    repository.NewRepositoryService(client),
		RoutingRule:   NewRoutingRuleService(client),
		CleanupPolicy: cleanup.NewCleanupPolicyService(client),
		Security:      security.NewSecurityService(client),
		Script:        NewScriptService(client),
		ReadOnly:      readonly.NewReadOnlyService(client),
		MailConfig:    NewMailConfigService(client),
	}
}
