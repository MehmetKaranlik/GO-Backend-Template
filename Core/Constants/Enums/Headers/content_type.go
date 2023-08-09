package Headers

type ContentType string

const (
	ContentTypeKey             = "Content-Type"
	JSON           ContentType = "application/json"
	HTML           ContentType = "text/html"
)

func (self ContentType) Value() string {
	return string(self)
}
