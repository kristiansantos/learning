package initializers

type Initializer struct {
	Environment   string
	Version       string
	Mongo         Mongo
	Application   Application
	Elasticsearch Elasticsearch
}
