package http

type RequestBuilder struct {
	r *Request
}

func NewRequestBuilder() *RequestBuilder {
	return &RequestBuilder{
		r: &Request{},
	}
}

func (b *RequestBuilder) DomainName(str string) *RequestBuilder {
	b.r.DomainName = str
	return b
}

func (b *RequestBuilder) UrlPath(str string) *RequestBuilder {
	b.r.UrlPath = str
	return b
}

func (b *RequestBuilder) Method(str string) *RequestBuilder {
	b.r.Method = str
	return b
}

func (b *RequestBuilder) Headers(h map[string]string) *RequestBuilder {
	b.r.Headers = h
	return b
}

func (b *RequestBuilder) Params(p map[string]string) *RequestBuilder {
	b.r.Params = p
	return b
}

func (b *RequestBuilder) Query(q map[string]string) *RequestBuilder {
	b.r.Query = q
	return b
}

func (b *RequestBuilder) Body(body []byte) *RequestBuilder {
	b.r.Body = body
	return b
}

func (b *RequestBuilder) Build() *Request {
	return b.r
}
