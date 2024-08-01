package rsync

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type FileDeltaStore struct {
	BlockSize int
}

func NewFileDeltaStore(blockSize int) *FileDeltaStore {
	return &FileDeltaStore{BlockSize: blockSize}
}

func (fds *FileDeltaStore) CreateSigFile(baseFile, sigFile string) error {
	rs := RSync{BlockSize: fds.BlockSize}
	open, _ := os.Open(baseFile)
	defer func(srcReader *os.File) {
		err := srcReader.Close()
		if err != nil {
			log.Fatalln("CreateSigFile1", err)
		}
	}(open)

	mySig := make([]BlockHash, 0, 16)
	err := rs.CreateSignature(open, func(bl BlockHash) error {
		mySig = append(mySig, bl)
		return nil
	})
	if err != nil {
		return err
	}

	create, _ := os.Create(sigFile)
	defer func(sf *os.File) {
		err := sf.Close()
		if err != nil {
			log.Fatalln("CreateSigFile2", err)
		}
	}(create)

	enc := json.NewEncoder(create)
	err = enc.Encode(mySig)
	if err != nil {
		return err
	}

	return nil
}

func (fds *FileDeltaStore) CreateDeltaFile(sigFile, newFile, deltaFile string) error {
	rs := RSync{BlockSize: fds.BlockSize}

	open, err := os.Open(sigFile)
	if err != nil {
		return err
	}
	defer func(sf *os.File) {
		err := sf.Close()
		if err != nil {
			log.Fatalln("CreateDeltaFile1", err)
		}
	}(open)
	decoder := json.NewDecoder(open)
	sig := make([]BlockHash, 0, 16)
	err = decoder.Decode(&sig)
	if err != nil {
		return err
	}

	create, err := os.Create(deltaFile)
	if err != nil {
		return err
	}
	encoder := json.NewEncoder(create)
	newOpen, _ := os.Open(newFile)

	wo := func(op Operation) error {
		err = encoder.Encode(op)
		if err != nil {
			log.Fatalln("CreateDeltaFile2", err)
		}
		//opsOut <- op
		return nil
	}

	done := make(chan struct{})

	go func() {
		defer func(df *os.File) {
			err := df.Close()
			if err != nil {
				log.Fatalln("CreateDeltaFile3", err)
			}
		}(create)
		//defer close(opsOut)
		err := rs.CreateDelta(newOpen, sig, wo)
		if err != nil {
			return
		}
		done <- struct{}{}
	}()
	<-done
	return nil
}

func (fds *FileDeltaStore) RebuildNewFile(baseFile, deltaFile, rebuildFile string) error {
	rs := RSync{BlockSize: fds.BlockSize}
	srcReader, err := os.Open(baseFile)
	if err != nil {
		return err
	}
	defer func(srcReader *os.File) {
		err := srcReader.Close()
		if err != nil {
			log.Fatalln("RebuildNewFile1", err)
		}
	}(srcReader)
	_, err = srcReader.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}

	srcWriter, err := os.Create(rebuildFile)
	if err != nil {
		return err
	}
	df, err := os.Open(deltaFile)
	if err != nil {
		return err
	}
	decoder := json.NewDecoder(df)

	readOps := make(chan Operation)
	go func() {
		defer func(df *os.File) {
			err := df.Close()
			if err != nil {
				log.Fatalln("RebuildNewFile2", err)
			}
		}(df)
		defer close(readOps)
		for {
			var op Operation
			err = decoder.Decode(&op)
			if err != nil {
				break
			}
			readOps <- op
		}
	}()

	return rs.ApplyDelta(srcWriter, srcReader, readOps)
}
