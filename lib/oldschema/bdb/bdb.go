package bdb

import schema "github.com/tankbusta/gocrackdb-converter/lib/oldschema"

const (
	curCrackTaskVer      float32 = 1.1
	curUserVer           float32 = 1.0
	curEntVer            float32 = 1.0
	curTaskFileVer       float32 = 1.0
	curCrackedHashVer    float32 = 1.0
	curAuditEntryVer     float32 = 1.0
	curEngineFileVer     float32 = 1.0
	curCheckpointFileVer float32 = 1.0
)

var (
	bucketAuditLog    = "audit_log"
	bucketTasks       = "tasks"
	bucketEntName     = "entitlements"
	bucketCheckpoints = "checkpoints"

	bucketTaskFiles   = []string{"files", "task_files"}
	bucketEngineFiles = []string{"files", "engine_files"}

	// Buckets related to entitlements
	bucketEntTaskFiles   = append(bucketTaskFiles, bucketEntName)
	bucketEntTasks       = []string{bucketTasks, bucketEntName}
	bucketEntEngineFiles = append(bucketEngineFiles, bucketEntName)
)

type BoltUser struct {
	ID          int64 `storm:"id,increment"`
	DocVersion  float32
	schema.User `storm:"inline"`
}

type BoltTaskFile struct {
	ID              int64 `storm:"id,increment"`
	DocVersion      float32
	schema.TaskFile `storm:"inline"`
}

type BoltEntitlement struct {
	ID int64 `storm:"id,increment"`
	// UniqueID is the MD5 of UserUUID and EntitleID in schema.EntitlementEntry
	// and is used as a unique check due to uniqueness limitations in bolt/storm
	UniqueID                string `storm:"unique"`
	DocVersion              float32
	schema.EntitlementEntry `storm:"inline"`
}

type BoltAuditLogEntry struct {
	ID                      int64 `storm:"id,increment"`
	DocVersion              float32
	schema.ActivityLogEntry `storm:"inline"`
}

type BoltEngineFile struct {
	ID                int64 `storm:"id,increment"`
	DocVersion        float32
	schema.EngineFile `storm:"inline"`
}

type BoltCheckpointFile struct {
	ID                    string `storm:"id,unique"`
	DocVersion            float32
	schema.CheckpointFile `storm:"inline"`
}
