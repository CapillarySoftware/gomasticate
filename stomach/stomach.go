package stomach

//Interface to where we want to store our data
type Stomach interface {
	IndexDocument(string, string, Document) (err error)
	Close()
}

type Document interface {
	String() string
	GetId() string
}
