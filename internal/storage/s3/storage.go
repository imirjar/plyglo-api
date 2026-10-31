package s3

type Storage struct {
}

func New(dbconn string) *Storage {
	return &Storage{}
}
