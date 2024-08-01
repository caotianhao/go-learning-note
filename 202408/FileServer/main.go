package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

// Encryptor 定义加密和解密的接口
type Encryptor interface {
	Encrypt(data []byte) ([]byte, error)
	Decrypt(data []byte) ([]byte, error)
}

// AESGCMEncryptor 使用 AES-GCM 加密
type AESGCMEncryptor struct {
	key []byte
}

func NewAESGCMEncryptor(key []byte) *AESGCMEncryptor {
	return &AESGCMEncryptor{key: key}
}

func (e *AESGCMEncryptor) Encrypt(data []byte) ([]byte, error) {
	block, err := aes.NewCipher(e.key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext, nil
}

func (e *AESGCMEncryptor) Decrypt(data []byte) ([]byte, error) {
	// ciphertext, err := base64.URLEncoding.DecodeString(cryptoText)
	// if err != nil {
	// 	return nil, err
	// }

	block, err := aes.NewCipher(e.key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(data) < nonceSize {
		return nil, fmt.Errorf("ciphertext too short")
	}

	nonce, data := data[:nonceSize], data[nonceSize:]
	return gcm.Open(nil, nonce, data, nil)
}

// AESCBCEncryptor 使用 AES-CBC 加密
type AESCBCEncryptor struct {
	key []byte
}

func NewAESCBCEncryptor(key []byte) *AESCBCEncryptor {
	return &AESCBCEncryptor{key: key}
}

// Encrypt 16 + (chunksize/16 + 1)*16
func (e *AESCBCEncryptor) Encrypt(data []byte) ([]byte, error) {
	block, err := aes.NewCipher(e.key)
	if err != nil {
		return nil, err
	}

	data = pkcs7Padding(data, block.BlockSize())
	ciphertext := make([]byte, aes.BlockSize+len(data))
	iv := ciphertext[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], data)

	return ciphertext, nil
}

func (e *AESCBCEncryptor) Decrypt(data []byte) ([]byte, error) {
	// ciphertext, err := base64.URLEncoding.DecodeString(cryptoText)
	// if err != nil {
	// 	return nil, err
	// }

	block, err := aes.NewCipher(e.key)
	if err != nil {
		return nil, err
	}

	if len(data) < aes.BlockSize {
		return nil, fmt.Errorf("ciphertext too short")
	}

	iv := data[:aes.BlockSize]
	data = data[aes.BlockSize:]

	if len(data)%aes.BlockSize != 0 {
		return nil, fmt.Errorf("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(data, data)

	return pkcs7Unpadding(data)
}

// PKCS#7 padding and unpadding
func pkcs7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

func pkcs7Unpadding(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, fmt.Errorf("invalid padding size")
	}
	padding := data[length-1]
	return data[:length-int(padding)], nil
}

var key = []byte("mysecretencryptionkey1234567890a")

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 从请求参数中获取加密算法
	algorithm := r.URL.Query().Get("alg")
	var encryptor Encryptor
	switch algorithm {
	case "aes-gcm":
		encryptor = NewAESGCMEncryptor(key)
	case "aes-cbc":
		encryptor = NewAESCBCEncryptor(key)
	default:
		http.Error(w, "Unsupported algorithm", http.StatusBadRequest)
		return
	}

	// 读取文件
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Failed to read file", http.StatusInternalServerError)
		return
	}
	defer func(file multipart.File) {
		_ = file.Close()
	}(file)

	// 读取文件内容
	fileData, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Failed to read file data", http.StatusInternalServerError)
		return
	}

	// 加密文件内容
	encryptedData, err := encryptor.Encrypt(fileData)
	if err != nil {
		http.Error(w, "Failed to encrypt file", http.StatusInternalServerError)
		return
	}

	// 存储文件
	filePath := strings.TrimPrefix(r.URL.Path, "/upload/")
	if strings.HasSuffix(filePath, "/") {
		http.Error(w, "Invalid file path", http.StatusBadRequest)
		return
	}

	err = os.WriteFile("./"+filePath, encryptedData, 0644)
	if err != nil {
		http.Error(w, "Failed to save file", http.StatusInternalServerError)
		return
	}

	_, _ = w.Write([]byte("File uploaded successfully"))
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 从请求参数中获取加密算法
	algorithm := r.URL.Query().Get("alg")
	var encryptor Encryptor
	switch algorithm {
	case "aes-gcm":
		encryptor = NewAESGCMEncryptor(key)
	case "aes-cbc":
		encryptor = NewAESCBCEncryptor(key)
	default:
		http.Error(w, "Unsupported algorithm", http.StatusBadRequest)
		return
	}

	filePath := strings.TrimPrefix(r.URL.Path, "/download/")
	if strings.HasSuffix(filePath, "/") {
		http.Error(w, "Invalid file path", http.StatusBadRequest)
		return
	}

	// 读取文件
	encryptedData, err := os.ReadFile("./" + filePath)
	if err != nil {
		http.Error(w, "Failed to read file", http.StatusInternalServerError)
		return
	}

	// 解密文件内容
	decryptedData, err := encryptor.Decrypt(encryptedData)
	if err != nil {
		http.Error(w, "Failed to decrypt file", http.StatusInternalServerError)
		return
	}

	// 设置响应头
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filePath))
	w.Header().Set("Content-Type", "application/octet-stream")
	_, _ = w.Write(decryptedData)
}

func main() {
	http.HandleFunc("/upload/", uploadHandler)
	http.HandleFunc("/download/", downloadHandler)

	fmt.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
