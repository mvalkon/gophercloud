package shares

import (
	"encoding/json"
	"time"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// Share contains all information associated with an OpenStack Share
type Share struct {
	// The availability zone of the share
	AvailabilityZone string `json:"availability_zone,omitempty"`
	// A description of the share
	Description string `json:"description,omitempty"`
	// DisplayDescription is inherited from BlockStorage API.
	// Both Description and DisplayDescription can be used
	DisplayDescription string `json:"display_description,omitempty"`
	// DisplayName is inherited from BlockStorage API
	// Both DisplayName and Name can be used
	DisplayName string `json:"display_name,omitempty"`
	// Indicates whether a share has replicas or not.
	HasReplicas bool `json:"has_replicas,omitempty"`
	// The host name of the share
	Host string `json:"host,omitempty"`
	// The UUID of the share
	ID string `json:"id,omitempty"`
	// Indicates the visibility of the share
	IsPublic bool `json:"is_public,omitempty"`
	// Share links for pagination
	Links []map[string]string `json:"links,omitempty"`
	// Key, value -pairs of custom metadata
	Metadata map[string]string `json:"metadata,omitempty"`
	// The name of the share
	Name string `json:"name,omitempty"`
	// The UUID of the project to which this share belongs to
	ProjectID string `json:"project_id,omitempty"`
	// The share replication type
	ReplicationType string `json:"replication_type,omitempty"`
	// The UUID of the share network
	ShareNetworkID string `json:"share_network_id,omitempty"`
	// The shared file system protocol
	ShareProto string `json:"share_proto,omitempty"`
	// The UUID of the share server
	ShareServerID string `json:"share_server_id,omitempty"`
	// The UUID of the share type.
	ShareType string `json:"share_type,omitempty"`
	// The name of the share type.
	ShareTypeName string `json:"share_type_name,omitempty"`
	// Size of the share in GB
	Size int `json:"size,omitempty"`
	// UUID of the snapshot from which to create the share
	SnapshotID string `json:"snapshot_id,omitempty"`
	// The share status
	Status string `json:"status,omitempty"`
	// The task state, used for share migration
	TaskState string `json:"task_state,omitempty"`
	// The type of the volume
	VolumeType string `json:"volume_type,omitempty"`
	// The UUID of the consistency group this share belongs to
	ConsistencyGroupID string `json:"consistency_group_id,omitempty"`
	// Used for filtering backends which either support or do not support share snapshots
	SnapshotSupport          bool   `json:"snapshot_support,omitempty"`
	SourceCgsnapshotMemberID string `json:"source_cgsnapshot_member_id,omitempty"`
	// Timestamp when the share was created
	CreatedAt time.Time `json:"-"`
}

func (r *Share) UnmarshalJSON(b []byte) error {
	type tmp Share
	var s struct {
		tmp
		CreatedAt gophercloud.JSONRFC3339MilliNoZ `json:"created_at"`
	}
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*r = Share(s.tmp)

	r.CreatedAt = time.Time(s.CreatedAt)

	return nil
}

type commonResult struct {
	gophercloud.Result
}

// Extract will get the Share object from the commonResult
func (r commonResult) Extract() (*Share, error) {
	var s struct {
		Share *Share `json:"share"`
	}
	err := r.ExtractInto(&s)
	return s.Share, err
}

// SharePage is a pagination.pager that is returned from a call to the List func
type SharePage struct {
	pagination.SinglePageBase
}

// ExtractShares extracts and returns Shares when iterating over a shares.List() call
func ExtractShares(r pagination.Page) ([]Share, error) {
	var s struct {
		Shares []Share `json:"shares"`
	}
	err := (r.(SharePage)).ExtractInto(&s)
	return s.Shares, err
}

// CreateResult contains the result..
type CreateResult struct {
	commonResult
}

// DeleteResult contains the delete results
type DeleteResult struct {
	gophercloud.ErrResult
}

// GetResult contains the get result
type GetResult struct {
	commonResult
}
