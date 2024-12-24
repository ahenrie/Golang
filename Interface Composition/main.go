package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
)

func main() {
	payload := []byte("hello high value software engineer")
	hashAndBroadcast(newHashReader(payload))
}

type hashReader struct {
	*bytes.Reader
	buf *bytes.Buffer
}

func newHashReader(b []byte) *hashReader {
	return &hashReader{
		Reader: bytes.NewReader(b),
		buf:    bytes.NewBuffer(b),
	}
}
func (h *hashReader) hash() string {
	return hex.EncodeToString(h.buf.Bytes())
}

func hashAndBroadcast(r io.Reader) error {
	hash := r.(*hashReader).hash()
	fmt.Println(hash)

	return broadcast(r)
}

func broadcast(r io.Reader) error {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	fmt.Println("String of the bytes: ", string(b))

	return nil
}
