package rsync

import (
	"crypto/sha256"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"log"
	"os"
	"testing"
)

// calculateFileSHA256 计算并返回给定文件的SHA256哈希值
func calculateFileSHA256(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalln("calculateFileSHA256 err =", err)
		}
	}(file)

	hasher := sha256.New()
	if _, err := io.Copy(hasher, file); err != nil {
		return "", err
	}

	hashBytes := hasher.Sum(nil)
	return fmt.Sprintf("%x", hashBytes), nil
}

func TestRsync(t *testing.T) {
	testCases := []struct {
		Name, BaseFile, NewFile, SigFile, DeltaFile, RebuildFile string
	}{
		{"小文件测试", "base.txt", "new.txt", "base.sig", "new.delta", "rebuild.txt"},
	}
	for _, c := range testCases {
		fmt.Println(c.Name)
		fds := NewFileDeltaStore(16)
		err := fds.CreateSigFile(c.BaseFile, c.SigFile)
		if err != nil {
			t.Fatalf("CreateSigFile err")
		}

		err = fds.CreateDeltaFile(c.SigFile, c.NewFile, c.DeltaFile)
		if err != nil {
			t.Fatalf("CreateDeltaFile err")
		}

		err = fds.RebuildNewFile(c.BaseFile, c.DeltaFile, c.RebuildFile)
		if err != nil {
			t.Fatalf("RebuildNewFile err")
		}

		newFileHash, err := calculateFileSHA256(c.NewFile)
		assert.Equal(t, nil, err)
		rebuildFileHash, err := calculateFileSHA256(c.RebuildFile)
		assert.Equal(t, nil, err)

		assert.Equal(t, newFileHash, rebuildFileHash)
	}
}
