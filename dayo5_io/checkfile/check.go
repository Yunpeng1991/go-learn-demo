package checkfile

import "learn-demo/dayo5_io/ops"

type CompareFile struct {
	data []byte

	compareData []byte

	differ bool

	differData *ops.DiffData
}

func (cf *CompareFile) ReadFile(path string) (data []byte, err error) {
	return nil, nil
}

func (cf *CompareFile) WriteFile(data []byte) error {
	return nil
}

func (cf *CompareFile) QueryDiff(id int64) *ops.DiffData {
	return nil
}

func (cf *CompareFile) SaveDiff(diff *ops.DiffData) int {
	return 0
}
