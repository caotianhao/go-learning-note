package main

import (
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	// 创建 P-256 椭圆曲线实例
	p256 := elliptic.P256()

	// 选择一个私钥
	privateKey, _ := generatePrivateKey(p256)

	// 计算对应的公钥
	publicKeyX, publicKeyY := p256.ScalarBaseMult(privateKey.Bytes())

	// 打印公钥的 X 和 Y 坐标
	fmt.Printf("Public Key (X): %x\n", publicKeyX)
	fmt.Printf("Public Key (Y): %x\n", publicKeyY)

	// 计算私钥的模逆
	inverse := new(big.Int).ModInverse(publicKeyX, publicKeyY)

	// 打印私钥的模逆
	fmt.Printf("Private Key Inverse: %x\n", inverse)
}

// 生成一个随机私钥
func generatePrivateKey(curve elliptic.Curve) (*big.Int, error) {
	maxNum := new(big.Int).Sub(curve.Params().N, big.NewInt(1))
	privateKey, err := rand.Int(rand.Reader, maxNum)
	if err != nil {
		return nil, err
	}
	privateKey.Add(privateKey, big.NewInt(1)) // 确保私钥大于0
	return privateKey, nil
}
