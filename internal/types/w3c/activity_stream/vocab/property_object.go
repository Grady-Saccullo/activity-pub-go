package vocab

import "net/url"

type PropertyObject interface {
	GetIRI() *url.URL

	SetIRI(url.URL)

	IsIRI() bool

	Object

	IsObject() bool

	Activity

	IsActivity() bool
}
