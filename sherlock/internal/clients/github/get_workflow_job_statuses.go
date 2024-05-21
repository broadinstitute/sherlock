package github

import (
	"context"
	"fmt"
	"github.com/google/go-github/v58/github"
	"net/http"
)

type JobPartial struct {
	Name   string
	Status string
}

func GetWorkflowJobStatuses(ctx context.Context, owner string, repo string, runID uint, attemptNumber uint) (map[int64]JobPartial, error) {
	statuses := make(map[int64]JobPartial)
	if isEnabled() {
		// The client library doesn't have /repos/{owner}/{repo}/actions/runs/{run_id}/attempts/{attempt_number}/jobs,
		// so we implement it ourselves like how the client library implements
		// /repos/{owner}/{repo}/actions/runs/{run_id}/jobs.
		url := fmt.Sprintf("repos/%s/%s/actions/runs/%d/attempts/%d/jobs", owner, repo, runID, attemptNumber)
		request, err := client.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return nil, err
		}
		jobs := new(github.Jobs)
		response, err := client.Do(ctx, request, jobs)
		if err != nil {
			return nil, err
		} else if response.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("unexpected status code %d", response.StatusCode)
		}
		for _, job := range jobs.Jobs {
			var status string
			if job.Conclusion != nil {
				status = *job.Conclusion
			} else if job.Status != nil {
				status = *job.Status
			}
			if job.ID != nil && job.Name != nil && status != "" {
				statuses[*job.ID] = JobPartial{
					Name:   *job.Name,
					Status: status,
				}
			}
		}
	}
	return statuses, nil
}

func FilterToProblematicJobStatuses(statuses map[int64]JobPartial) map[int64]JobPartial {
	problematicStatuses := make(map[int64]JobPartial)
	for id, job := range statuses {
		if job.Status == "failure" || job.Status == "cancelled" || job.Status == "timed_out" || job.Status == "action_required" {
			problematicStatuses[id] = job
		}
	}
	return problematicStatuses
}
