package schemas

import (
	"strings"

	goGitlab "github.com/xanzy/go-gitlab"
)

// Job ..
type Job struct {
	Name                  string
	Stage                 string
	Status                string
	TagList               string
	FailureReason         string
	Runner                Runner
	ID                    int
	Timestamp             float64
	DurationSeconds       float64
	QueuedDurationSeconds float64
	ArtifactSize          float64
}

// Runner ..
type Runner struct {
	Description string
}

// Jobs ..
type Jobs map[string]Job

// NewJob ..
func NewJob(gj goGitlab.Job) Job {
	var (
		artifactSize float64
		timestamp    float64
	)

	for _, artifact := range gj.Artifacts {
		artifactSize += float64(artifact.Size)
	}

	if gj.CreatedAt != nil {
		timestamp = float64(gj.CreatedAt.Unix())
	}

	return Job{
		ID:                    gj.ID,
		Name:                  gj.Name,
		Stage:                 gj.Stage,
		Timestamp:             timestamp,
		DurationSeconds:       gj.Duration,
		QueuedDurationSeconds: gj.QueuedDuration,
		Status:                gj.Status,
		TagList:               strings.Join(gj.TagList, ","),
		ArtifactSize:          artifactSize,
		FailureReason:         gj.FailureReason,

		Runner: Runner{
			Description: gj.Runner.Description,
		},
	}
}
