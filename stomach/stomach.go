package stomach

//Interface to where we want to store our data
type Stomach interface {
	IndexDocument(string, string, Document) (err error)
	IndexDocumentDynamic(index string, indexType string, doc interface{}, id string) error
	Close()
}

type Document interface {
	String() string
	GetId() string
}
