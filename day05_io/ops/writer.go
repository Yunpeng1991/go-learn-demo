package ops

type Writer interface {
	WriteFile(data []byte) (err error)

	SaveDiff(diff *DiffData) int
}
