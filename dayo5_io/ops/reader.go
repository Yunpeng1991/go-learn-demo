package ops

type Reader interface {
	ReadFile(path string) (data []byte, err error)

	QueryDiff(id int64) *DiffData
}
