package elasticsearch

//elasticsearch wrapper class
import (
	// "errors"
	// "fmt"
	. "github.com/CapillarySoftware/gomasticate/stomach"
	es "github.com/mattbaird/elastigo/lib"
	"sync"
)

type Elasticsearch struct {
	c  *es.Conn
	wg *sync.WaitGroup
	// bulk *es.BulkIndexer
}

//Connect to elasticsearch
func (this *Elasticsearch) Connect(url string) (err error) {
	this.c = es.NewConn()
	this.c.Domain = url
	return
}

//Index documents into elasticsearch
func (this *Elasticsearch) IndexDocument(index string, indexType string, doc Document) (err error) {
	_, err = this.c.Index(index, indexType, doc.GetId(), nil, doc)
	// now := time.Now().UTC()
	// this.bulk.Index(index, indexType, "", doc.GetId(), &now, doc, false)
	return
}

func (this *Elasticsearch) Close() {
}
