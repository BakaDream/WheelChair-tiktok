package Storage

type Storage interface {
	UploadFile()
}

func NewStorage() {
	switch gobal.StorageType {
	case "local":
		return &Local{}
	case "tencent-cos":
		return &TencentCos{}

	}
}
