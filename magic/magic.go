package magic

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"os"
)

type Header struct {
	Header uint64
}

// MagicNumbers check first 8 bytes of file (image, binary etc.)
func MagicNumbers(file string) (string, error) {
	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer f.Close()
	gp, err := getPayload(f)
	if err != nil {
		return "", err
	}
	mb, err := getMagicBytes(gp)
	if err != nil {
		return "", err
	}
	return mb, nil
}

func getPayload(file *os.File) (*bytes.Reader, error) {
	stats, err := file.Stat()
	if err != nil {
		return nil, err
	}
	var size = stats.Size()
	b := make([]byte, size)
	bufR := bufio.NewReader(file)
	_, err = bufR.Read(b)
	if err != nil {
		return nil, err
	}
	bReader := bytes.NewReader(b)
	return bReader, nil
}

func getMagicBytes(data *bytes.Reader) (string, error) {
	var header Header
	if err := binary.Read(data, binary.BigEndian, &header.Header); err != nil {
		return "", err
	}
	bArr := make([]byte, 8)
	binary.BigEndian.PutUint64(bArr, header.Header)
	return string(bArr[1:4]), nil
}
