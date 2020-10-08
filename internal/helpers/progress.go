package helpers

import (
	"os"
	"sync"
	"sync/atomic"

	"github.com/cheggaaa/pb"
)

type ProgressReader struct {
	ProgBar *pb.ProgressBar
	Fp      *os.File
	Size    int64
	Reads   int64
	SignMap map[int64]struct{}
	mux     sync.Mutex
}

func (r *ProgressReader) Read(p []byte) (int, error) {

	return r.Fp.Read(p)
}

func (r *ProgressReader) ReadAt(p []byte, off int64) (int, error) {
	//https://github.com/aws/aws-sdk-go/issues/2468
	//https://github.com/aws/aws-sdk-go/pull/2456
	//https://github.com/aws/aws-sdk-go/commit/eb7121dfde9200b61572aeafef0eaa0cfb2a7652

	n, err := r.Fp.ReadAt(p, off)
	atomic.AddInt64(&r.Reads, int64(n))
	r.mux.Lock()
	if read := atomic.LoadInt64(&r.Reads); read >= 0 {
		r.ProgBar.Set64(read)
	}
	r.mux.Unlock()

	return n, err
}

func (r *ProgressReader) Seek(offset int64, whence int) (int64, error) {
	return r.Fp.Seek(offset, whence)
}
