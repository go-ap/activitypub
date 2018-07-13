package activitypub

type (
	// IRI is a Internationalized Resource Identifiers (IRIs) RFC3987
	IRI URI

	// URI is a Uniform Resource Identifier (URI) RFC3986
	URI string
)

func (u URI) String() string {
	return string(u)
}
func (i IRI) String() string {
	return string(i)
}

func (u URI) GetLink() URI {
	return u
}
func (i IRI) GetLink() URI {
	return URI(i)
}
