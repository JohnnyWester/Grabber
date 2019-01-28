package news

type Topic struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	URL    string `json:"url"`
}

type Archive map[string][]Topic

func New() Archive {
	return make(Archive, 0)
}

func (archive Archive) collect(category string) {
	sources := getSources(category)
	topics := getTopics(sources)

	archive[category] = topics
}

func (archive Archive) result(category string) []Topic {
	return archive[category]
}

func (archive Archive) Serve() {
	for {
		select {
		case category := <-collect:
			archive.collect(category)
		case category := <-request:
			result <- archive.result(category)
		}
	}
}
