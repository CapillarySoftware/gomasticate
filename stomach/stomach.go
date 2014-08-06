package stomach

//Interface to where we want to store our data
type Stomach interface {
	IndexDocument(Document) (err error)
}

type Document interface {
	String() string
	GetIndex() string
	GetIndexType() string
	GetId() string
}
