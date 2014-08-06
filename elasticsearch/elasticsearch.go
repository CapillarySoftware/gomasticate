package elasticsearch

import (
	// "errors"
	. "github.com/CapillarySoftware/gomasticate/stomach"
	es "github.com/mattbaird/elastigo/lib"
)

type Elasticsearch struct {
	c *es.Conn
}

func (this *Elasticsearch) Connect(url string) (err error) {
	this.c = es.NewConn()
	this.c.Domain = url
	return
}

func (this *Elasticsearch) IndexDocument(index string, indexType string, id string, doc Document) (err error) {
	_, err = this.c.Index(index, indexType, id, nil, doc)
	this.c.Flush()
	return
}
