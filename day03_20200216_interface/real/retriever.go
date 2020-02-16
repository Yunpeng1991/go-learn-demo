package real

import "fmt"

type Retriever struct {
	Contents string
}

/**implements Stringer*/
/**like the java toString()*/
func (r *Retriever) String() string {
	return fmt.Sprintf("Retriever : {Contents=%s}", r.Contents)
}

/**implements the Reader**/
/**see the io.Reader**/
func (r *Retriever) Read(p []byte) (n int, err error) {
	return 0, nil
}

/**implements the Writer**/
/**see the os.file,io.Writer**/
func (r *Retriever) Write(p []byte) (n int, err error) {
	return 0, nil
}
