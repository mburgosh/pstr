package main

type App struct {
	Name        string `json:"name"`
	Group       string `json:"group"`
	Type        string `json:"type"`
	Scalability struct {
		MinReplicas                    int `json:"minReplicas"`
		MaxReplicas                    int `json:"maxReplicas"`
		TargetCPUUtilizationPercentage int `json:"targetCPUUtilizationPercentage"`
	} `json:"scalability"`
	ImageName      string `json:"imageName"`
	DockerContext  string `json:"dockerContext,omitempty"`
	DockerFilePath string `json:"dockerFilePath,omitempty"`
	Ports          struct {
		External struct {
			HTTP int `json:"http"`
		} `json:"external"`
		Internal struct {
			HTTP int `json:"http"`
		} `json:"internal"`
	} `json:"ports"`
	Readiness struct {
		Path string `json:"path,omitempty"`
		Port int `json:"port,omitempty"`
		Exec struct {
			Command []string `json:"command,omitempty"`
		} `json:"exec,omitempty"`
		InitialDelaySeconds int    `json:"initialDelaySeconds"`
		PeriodSeconds       int    `json:"periodSeconds"`
		FailureThreshold    int    `json:"failureThreshold"`
		TimeoutSeconds      int    `json:"timeoutSeconds"`
		SuccessThreshold    int    `json:"successThreshold"`
	} `json:"readiness"`
	Health struct {
		Path string `json:"path,omitempty"`
		Port int `json:"port,omitempty"`
		Exec struct {
			Command []string `json:"command,omitempty"`
		} `json:"exec,omitempty"`
		InitialDelaySeconds int    `json:"initialDelaySeconds"`
		PeriodSeconds       int    `json:"periodSeconds"`
		FailureThreshold    int    `json:"failureThreshold"`
		TimeoutSeconds      int    `json:"timeoutSeconds"`
		SuccessThreshold    int    `json:"successThreshold"`
	} `json:"health"`
	PreStop struct {
		HTTPGet struct {
			Path string `json:"path,omitempty"`
			Port int    `json:"port,omitempty"`
		} `json:"httpGet,omitempty"`
		Exec struct {
			Command []string `json:"command,omitempty"`
		} `json:"exec,omitempty"`
	} `json:"preStop"`
}

type Ingresstmpl struct {
	Application_name string
	Deploy           App
	Namespace        string
	Hostnames        []string
	Build_id         string
}

type Autoscalertmpl struct {
	Deploy      App
	Deploy_name string
	Namespace   string
}

type Servicetmpl struct {
	Application_name string
	Namespace        string
	Cluster_ip       string
	Build_id         string
	Deploy           App
}

type Deploytmpl struct {
	Application_name           string
	Bamboo_deploy_release      string
	Build_id                   string
	Build_nr                   string
	CONSUL_APPLICATION         string
	CONSUL_ENVIRONMENT         string
	CONSUL_PASSWORD            string
	CONSUL_URL                 string
    CONSUL_FULL_URL            string
	CONSUL_USERNAME            string
	Deploy                     App
	Deploy_name                string
	Git_repo                   string
	Namespace                  string
	NEW_RELIC_LICENSE_KEY      string
	NEW_RELIC_API_URL          string
	NEW_RELIC_API_KEY_PASSWORD string
	Ssh_key                    string
}
