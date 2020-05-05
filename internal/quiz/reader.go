package quiz

// Reader represents a reader for reading file
//
// Read receive a io.Reader so the sources can be general as long it implement io.Reader
// Sample case, csv.NewReader return struct that implement io.Reader
type Reader interface {
	Read() ([]*Problem, error)
}
