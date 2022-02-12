// Code generated by sqlc. DO NOT EDIT.

package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type LogLevel string

const (
	LogLevelTrace LogLevel = "trace"
	LogLevelDebug LogLevel = "debug"
	LogLevelInfo  LogLevel = "info"
	LogLevelWarn  LogLevel = "warn"
	LogLevelError LogLevel = "error"
)

func (e *LogLevel) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = LogLevel(s)
	case string:
		*e = LogLevel(s)
	default:
		return fmt.Errorf("unsupported scan type for LogLevel: %T", src)
	}
	return nil
}

type LogSource string

const (
	LogSourceProvisionerDaemon LogSource = "provisioner_daemon"
	LogSourceProvisioner       LogSource = "provisioner"
)

func (e *LogSource) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = LogSource(s)
	case string:
		*e = LogSource(s)
	default:
		return fmt.Errorf("unsupported scan type for LogSource: %T", src)
	}
	return nil
}

type LoginType string

const (
	LoginTypeBuiltIn LoginType = "built-in"
	LoginTypeSaml    LoginType = "saml"
	LoginTypeOIDC    LoginType = "oidc"
)

func (e *LoginType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = LoginType(s)
	case string:
		*e = LoginType(s)
	default:
		return fmt.Errorf("unsupported scan type for LoginType: %T", src)
	}
	return nil
}

type ParameterDestinationScheme string

const (
	ParameterDestinationSchemeNone                ParameterDestinationScheme = "none"
	ParameterDestinationSchemeEnvironmentVariable ParameterDestinationScheme = "environment_variable"
	ParameterDestinationSchemeProvisionerVariable ParameterDestinationScheme = "provisioner_variable"
)

func (e *ParameterDestinationScheme) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = ParameterDestinationScheme(s)
	case string:
		*e = ParameterDestinationScheme(s)
	default:
		return fmt.Errorf("unsupported scan type for ParameterDestinationScheme: %T", src)
	}
	return nil
}

type ParameterScope string

const (
	ParameterScopeOrganization ParameterScope = "organization"
	ParameterScopeProject      ParameterScope = "project"
	ParameterScopeImportJob    ParameterScope = "import_job"
	ParameterScopeUser         ParameterScope = "user"
	ParameterScopeWorkspace    ParameterScope = "workspace"
)

func (e *ParameterScope) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = ParameterScope(s)
	case string:
		*e = ParameterScope(s)
	default:
		return fmt.Errorf("unsupported scan type for ParameterScope: %T", src)
	}
	return nil
}

type ParameterSourceScheme string

const (
	ParameterSourceSchemeNone ParameterSourceScheme = "none"
	ParameterSourceSchemeData ParameterSourceScheme = "data"
)

func (e *ParameterSourceScheme) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = ParameterSourceScheme(s)
	case string:
		*e = ParameterSourceScheme(s)
	default:
		return fmt.Errorf("unsupported scan type for ParameterSourceScheme: %T", src)
	}
	return nil
}

type ParameterTypeSystem string

const (
	ParameterTypeSystemNone ParameterTypeSystem = "none"
	ParameterTypeSystemHCL  ParameterTypeSystem = "hcl"
)

func (e *ParameterTypeSystem) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = ParameterTypeSystem(s)
	case string:
		*e = ParameterTypeSystem(s)
	default:
		return fmt.Errorf("unsupported scan type for ParameterTypeSystem: %T", src)
	}
	return nil
}

type ProvisionerJobType string

const (
	ProvisionerJobTypeProjectVersionImport ProvisionerJobType = "project_version_import"
	ProvisionerJobTypeWorkspaceProvision   ProvisionerJobType = "workspace_provision"
)

func (e *ProvisionerJobType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = ProvisionerJobType(s)
	case string:
		*e = ProvisionerJobType(s)
	default:
		return fmt.Errorf("unsupported scan type for ProvisionerJobType: %T", src)
	}
	return nil
}

type ProvisionerStorageMethod string

const (
	ProvisionerStorageMethodFile ProvisionerStorageMethod = "file"
)

func (e *ProvisionerStorageMethod) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = ProvisionerStorageMethod(s)
	case string:
		*e = ProvisionerStorageMethod(s)
	default:
		return fmt.Errorf("unsupported scan type for ProvisionerStorageMethod: %T", src)
	}
	return nil
}

type ProvisionerType string

const (
	ProvisionerTypeEcho      ProvisionerType = "echo"
	ProvisionerTypeTerraform ProvisionerType = "terraform"
)

func (e *ProvisionerType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = ProvisionerType(s)
	case string:
		*e = ProvisionerType(s)
	default:
		return fmt.Errorf("unsupported scan type for ProvisionerType: %T", src)
	}
	return nil
}

type UserStatus string

const (
	UserstatusActive         UserStatus = "active"
	UserstatusDormant        UserStatus = "dormant"
	UserstatusDecommissioned UserStatus = "decommissioned"
)

func (e *UserStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = UserStatus(s)
	case string:
		*e = UserStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for UserStatus: %T", src)
	}
	return nil
}

type WorkspaceTransition string

const (
	WorkspaceTransitionStart  WorkspaceTransition = "start"
	WorkspaceTransitionStop   WorkspaceTransition = "stop"
	WorkspaceTransitionDelete WorkspaceTransition = "delete"
)

func (e *WorkspaceTransition) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = WorkspaceTransition(s)
	case string:
		*e = WorkspaceTransition(s)
	default:
		return fmt.Errorf("unsupported scan type for WorkspaceTransition: %T", src)
	}
	return nil
}

type APIKey struct {
	ID               string    `db:"id" json:"id"`
	HashedSecret     []byte    `db:"hashed_secret" json:"hashed_secret"`
	UserID           string    `db:"user_id" json:"user_id"`
	Application      bool      `db:"application" json:"application"`
	Name             string    `db:"name" json:"name"`
	LastUsed         time.Time `db:"last_used" json:"last_used"`
	ExpiresAt        time.Time `db:"expires_at" json:"expires_at"`
	CreatedAt        time.Time `db:"created_at" json:"created_at"`
	UpdatedAt        time.Time `db:"updated_at" json:"updated_at"`
	LoginType        LoginType `db:"login_type" json:"login_type"`
	OIDCAccessToken  string    `db:"oidc_access_token" json:"oidc_access_token"`
	OIDCRefreshToken string    `db:"oidc_refresh_token" json:"oidc_refresh_token"`
	OIDCIDToken      string    `db:"oidc_id_token" json:"oidc_id_token"`
	OIDCExpiry       time.Time `db:"oidc_expiry" json:"oidc_expiry"`
	DevurlToken      bool      `db:"devurl_token" json:"devurl_token"`
}

type File struct {
	Hash      string    `db:"hash" json:"hash"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	CreatedBy string    `db:"created_by" json:"created_by"`
	Mimetype  string    `db:"mimetype" json:"mimetype"`
	Data      []byte    `db:"data" json:"data"`
}

type License struct {
	ID        int32           `db:"id" json:"id"`
	License   json.RawMessage `db:"license" json:"license"`
	CreatedAt time.Time       `db:"created_at" json:"created_at"`
}

type Organization struct {
	ID                     string    `db:"id" json:"id"`
	Name                   string    `db:"name" json:"name"`
	Description            string    `db:"description" json:"description"`
	CreatedAt              time.Time `db:"created_at" json:"created_at"`
	UpdatedAt              time.Time `db:"updated_at" json:"updated_at"`
	Default                bool      `db:"default" json:"default"`
	AutoOffThreshold       int64     `db:"auto_off_threshold" json:"auto_off_threshold"`
	CpuProvisioningRate    float32   `db:"cpu_provisioning_rate" json:"cpu_provisioning_rate"`
	MemoryProvisioningRate float32   `db:"memory_provisioning_rate" json:"memory_provisioning_rate"`
	WorkspaceAutoOff       bool      `db:"workspace_auto_off" json:"workspace_auto_off"`
}

type OrganizationMember struct {
	OrganizationID string    `db:"organization_id" json:"organization_id"`
	UserID         string    `db:"user_id" json:"user_id"`
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
	UpdatedAt      time.Time `db:"updated_at" json:"updated_at"`
	Roles          []string  `db:"roles" json:"roles"`
}

type ParameterSchema struct {
	ID                       uuid.UUID                  `db:"id" json:"id"`
	CreatedAt                time.Time                  `db:"created_at" json:"created_at"`
	JobID                    uuid.UUID                  `db:"job_id" json:"job_id"`
	Name                     string                     `db:"name" json:"name"`
	Description              string                     `db:"description" json:"description"`
	DefaultSourceScheme      ParameterSourceScheme      `db:"default_source_scheme" json:"default_source_scheme"`
	DefaultSourceValue       string                     `db:"default_source_value" json:"default_source_value"`
	AllowOverrideSource      bool                       `db:"allow_override_source" json:"allow_override_source"`
	DefaultDestinationScheme ParameterDestinationScheme `db:"default_destination_scheme" json:"default_destination_scheme"`
	AllowOverrideDestination bool                       `db:"allow_override_destination" json:"allow_override_destination"`
	DefaultRefresh           string                     `db:"default_refresh" json:"default_refresh"`
	RedisplayValue           bool                       `db:"redisplay_value" json:"redisplay_value"`
	ValidationError          string                     `db:"validation_error" json:"validation_error"`
	ValidationCondition      string                     `db:"validation_condition" json:"validation_condition"`
	ValidationTypeSystem     ParameterTypeSystem        `db:"validation_type_system" json:"validation_type_system"`
	ValidationValueType      string                     `db:"validation_value_type" json:"validation_value_type"`
}

type ParameterValue struct {
	ID                uuid.UUID                  `db:"id" json:"id"`
	Name              string                     `db:"name" json:"name"`
	CreatedAt         time.Time                  `db:"created_at" json:"created_at"`
	UpdatedAt         time.Time                  `db:"updated_at" json:"updated_at"`
	Scope             ParameterScope             `db:"scope" json:"scope"`
	ScopeID           string                     `db:"scope_id" json:"scope_id"`
	SourceScheme      ParameterSourceScheme      `db:"source_scheme" json:"source_scheme"`
	SourceValue       string                     `db:"source_value" json:"source_value"`
	DestinationScheme ParameterDestinationScheme `db:"destination_scheme" json:"destination_scheme"`
}

type Project struct {
	ID              uuid.UUID       `db:"id" json:"id"`
	CreatedAt       time.Time       `db:"created_at" json:"created_at"`
	UpdatedAt       time.Time       `db:"updated_at" json:"updated_at"`
	OrganizationID  string          `db:"organization_id" json:"organization_id"`
	Name            string          `db:"name" json:"name"`
	Provisioner     ProvisionerType `db:"provisioner" json:"provisioner"`
	ActiveVersionID uuid.UUID       `db:"active_version_id" json:"active_version_id"`
}

type ProjectImportJobResource struct {
	ID         uuid.UUID           `db:"id" json:"id"`
	CreatedAt  time.Time           `db:"created_at" json:"created_at"`
	JobID      uuid.UUID           `db:"job_id" json:"job_id"`
	Transition WorkspaceTransition `db:"transition" json:"transition"`
	Type       string              `db:"type" json:"type"`
	Name       string              `db:"name" json:"name"`
}

type ProjectVersion struct {
	ID          uuid.UUID `db:"id" json:"id"`
	ProjectID   uuid.UUID `db:"project_id" json:"project_id"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
	Name        string    `db:"name" json:"name"`
	Description string    `db:"description" json:"description"`
	ImportJobID uuid.UUID `db:"import_job_id" json:"import_job_id"`
}

type ProvisionerDaemon struct {
	ID           uuid.UUID         `db:"id" json:"id"`
	CreatedAt    time.Time         `db:"created_at" json:"created_at"`
	UpdatedAt    sql.NullTime      `db:"updated_at" json:"updated_at"`
	Name         string            `db:"name" json:"name"`
	Provisioners []ProvisionerType `db:"provisioners" json:"provisioners"`
}

type ProvisionerJob struct {
	ID             uuid.UUID                `db:"id" json:"id"`
	CreatedAt      time.Time                `db:"created_at" json:"created_at"`
	UpdatedAt      time.Time                `db:"updated_at" json:"updated_at"`
	StartedAt      sql.NullTime             `db:"started_at" json:"started_at"`
	CancelledAt    sql.NullTime             `db:"cancelled_at" json:"cancelled_at"`
	CompletedAt    sql.NullTime             `db:"completed_at" json:"completed_at"`
	Error          sql.NullString           `db:"error" json:"error"`
	OrganizationID string                   `db:"organization_id" json:"organization_id"`
	InitiatorID    string                   `db:"initiator_id" json:"initiator_id"`
	Provisioner    ProvisionerType          `db:"provisioner" json:"provisioner"`
	StorageMethod  ProvisionerStorageMethod `db:"storage_method" json:"storage_method"`
	StorageSource  string                   `db:"storage_source" json:"storage_source"`
	Type           ProvisionerJobType       `db:"type" json:"type"`
	Input          json.RawMessage          `db:"input" json:"input"`
	WorkerID       uuid.NullUUID            `db:"worker_id" json:"worker_id"`
}

type ProvisionerJobLog struct {
	ID        uuid.UUID `db:"id" json:"id"`
	JobID     uuid.UUID `db:"job_id" json:"job_id"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	Source    LogSource `db:"source" json:"source"`
	Level     LogLevel  `db:"level" json:"level"`
	Output    string    `db:"output" json:"output"`
}

type User struct {
	ID                  string     `db:"id" json:"id"`
	Email               string     `db:"email" json:"email"`
	Name                string     `db:"name" json:"name"`
	Revoked             bool       `db:"revoked" json:"revoked"`
	LoginType           LoginType  `db:"login_type" json:"login_type"`
	HashedPassword      []byte     `db:"hashed_password" json:"hashed_password"`
	CreatedAt           time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt           time.Time  `db:"updated_at" json:"updated_at"`
	TemporaryPassword   bool       `db:"temporary_password" json:"temporary_password"`
	AvatarHash          string     `db:"avatar_hash" json:"avatar_hash"`
	SshKeyRegeneratedAt time.Time  `db:"ssh_key_regenerated_at" json:"ssh_key_regenerated_at"`
	Username            string     `db:"username" json:"username"`
	DotfilesGitUri      string     `db:"dotfiles_git_uri" json:"dotfiles_git_uri"`
	Roles               []string   `db:"roles" json:"roles"`
	Status              UserStatus `db:"status" json:"status"`
	Relatime            time.Time  `db:"relatime" json:"relatime"`
	GpgKeyRegeneratedAt time.Time  `db:"gpg_key_regenerated_at" json:"gpg_key_regenerated_at"`
	Decomissioned       bool       `db:"_decomissioned" json:"_decomissioned"`
	Shell               string     `db:"shell" json:"shell"`
}

type Workspace struct {
	ID        uuid.UUID `db:"id" json:"id"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	OwnerID   string    `db:"owner_id" json:"owner_id"`
	ProjectID uuid.UUID `db:"project_id" json:"project_id"`
	Name      string    `db:"name" json:"name"`
}

type WorkspaceAgent struct {
	ID                  uuid.UUID       `db:"id" json:"id"`
	WorkspaceResourceID uuid.UUID       `db:"workspace_resource_id" json:"workspace_resource_id"`
	CreatedAt           time.Time       `db:"created_at" json:"created_at"`
	UpdatedAt           time.Time       `db:"updated_at" json:"updated_at"`
	InstanceMetadata    json.RawMessage `db:"instance_metadata" json:"instance_metadata"`
	ResourceMetadata    json.RawMessage `db:"resource_metadata" json:"resource_metadata"`
}

type WorkspaceHistory struct {
	ID               uuid.UUID           `db:"id" json:"id"`
	CreatedAt        time.Time           `db:"created_at" json:"created_at"`
	UpdatedAt        time.Time           `db:"updated_at" json:"updated_at"`
	WorkspaceID      uuid.UUID           `db:"workspace_id" json:"workspace_id"`
	ProjectVersionID uuid.UUID           `db:"project_version_id" json:"project_version_id"`
	Name             string              `db:"name" json:"name"`
	BeforeID         uuid.NullUUID       `db:"before_id" json:"before_id"`
	AfterID          uuid.NullUUID       `db:"after_id" json:"after_id"`
	Transition       WorkspaceTransition `db:"transition" json:"transition"`
	Initiator        string              `db:"initiator" json:"initiator"`
	ProvisionerState []byte              `db:"provisioner_state" json:"provisioner_state"`
	ProvisionJobID   uuid.UUID           `db:"provision_job_id" json:"provision_job_id"`
}

type WorkspaceResource struct {
	ID                  uuid.UUID     `db:"id" json:"id"`
	CreatedAt           time.Time     `db:"created_at" json:"created_at"`
	WorkspaceHistoryID  uuid.UUID     `db:"workspace_history_id" json:"workspace_history_id"`
	Type                string        `db:"type" json:"type"`
	Name                string        `db:"name" json:"name"`
	WorkspaceAgentToken string        `db:"workspace_agent_token" json:"workspace_agent_token"`
	WorkspaceAgentID    uuid.NullUUID `db:"workspace_agent_id" json:"workspace_agent_id"`
}
