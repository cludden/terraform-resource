package models

type OutRequest struct {
	Source Source    `json:"source"`
	Params OutParams `json:"params"`
}

type OutResponse struct {
	Version  Version  `json:"version"`
	Metadata Metadata `json:"metadata"`
}

type OutParams struct {
	EnvName            string `json:"env_name"`
	EnvNameFile        string `json:"env_name_file"`
	GenerateRandomName bool   `json:"generate_random_name"`
	Action             string `json:"action,omitempty"` // optional
	PlanName           string `json:"plan_name,omitempty"`
	DiscardPlan        bool   `json:"discard_plan"`
	Terraform
}

const (
	DestroyAction = "destroy"
)
