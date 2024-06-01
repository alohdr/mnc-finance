package entity

type (
	Language struct {
		Language       string        `json:"language"`
		Appeared       *int          `json:"appeared"`
		Created        []string      `json:"created"`
		Functional     *bool         `json:"functional"`
		ObjectOriented *bool         `json:"object-oriented"`
		Relation       RelationChild `json:"relation"`
	}

	RelationChild struct {
		InfluencedBy []string `json:"influenced-by"`
		Influences   []string `json:"influences"`
	}
)
