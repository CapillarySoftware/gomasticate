package elasticsearch

//elasticsearch wrapper class
import (
	// "errors"
	. "github.com/CapillarySoftware/gomasticate/stomach"
	es "github.com/mattbaird/elastigo/lib"
)

type Elasticsearch struct {
	c *es.Conn
}

//Connect to elasticsearch
func (this *Elasticsearch) Connect(url string) (err error) {
	this.c = es.NewConn()
	this.c.Domain = url
	return
}

//Index documents into elasticsearch
func (this *Elasticsearch) IndexDocument(index string, indexType string, id string, doc Document) (err error) {
	_, err = this.c.Index(index, indexType, id, nil, doc)
	this.c.Flush()
	return
}
