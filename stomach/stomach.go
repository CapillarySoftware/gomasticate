package stomach

//Interface to where we want to store our data
type Stomach interface {
	IndexDocument(string, string, string, Document) (err error)
}

type Document interface {
	String() string
}
