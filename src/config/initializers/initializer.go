package initializers

type Initializer struct {
	Environment   string
	Brand         string
	Version       string
	Mongo         Mongo
	Application   Application
	Elasticsearch Elasticsearch
	Sharepoint    Sharepoint
	DataDog       DataDog
}
