package CDN

type ICDNService[T any] interface {
	connect() T
	UploadFile(file []byte, fileName string) (string, error)
	disconnect() error
}
