package config

// Config structure for the github project
// TriagedColumn id in github
// BlockedColumn id in github
// InProgressColumn id in github
// ProgressColumn id, this is used only for PRs, in github
// DoneColumn id in github
// BacklogReleaseSquadColumn id in github
type Config struct {
	Server struct {
		Address      string `yaml:"address,omitempty"`
		ReadTimeout  int    `yaml:"readTimeout,omitempty"`
		WriteTimeout int    `yaml:"writeTimeout,omitempty"`
	} `yaml:"server,omitempty"`
	Github struct {
		APIURL string `yaml:"apiURL,omitempty"`
		Token  string `yaml:"token,omitempty"`
		Secret string `yaml:"secret,omitempty"`
	} `yaml:"github,omitempty"`
	Caasp struct {
		Repositories []string `yaml:"repositories,omitempty"`
		ProjectRules struct {
			MoveCards []struct {
				Action      string `yaml:"action,omitempty"`
				Match       string `yaml:"match,omitempty"`
				From        []int  `yaml:"from,omitempty"`
				Destination int    `yaml:"destination,omitempty"`
			} `yaml:"moveCards,omitempty"`
		} `yaml:"projectRules,omitempty"`
	} `yaml:"caasp,omitempty"`
}
