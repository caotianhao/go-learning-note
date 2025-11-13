package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

//给定一个字符串数组，计算并返回默克尔树的根哈希值。
//
//函数签名
//
//function calculateMerkleRootByHash(data: string[]): string;
//哈希函数
//// 简化的哈希函数（面试专用）
//function simpleHash(data: string): string {
//  let hash = 0;
//  for (let i = 0; i < data.length; i++) {
//    const char = data.charCodeAt(i);
//    hash = ((hash << 5) - hash) + char;
//    hash = hash & hash; // 转为32位整数
//  }
//  return Math.abs(hash).toString(16);
//}
//输入输出示例
//// 输入：4个字符串
//const input1 = ["apple", "banana", "cherry", "date"];
//console.log(calculateMerkleRootByHash(input1));
//// 输出：根哈希值（具体值取决于哈希函数实现）
//
//// 输入：3个字符串（奇数）
//const input2 = ["hello", "world", "test"];
//console.log(calculateMerkleRootByHash(input2));

func simpleHash(data string) string {
	hash := sha256.New()
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))
}

func calculateMerkleRootByHash(data []string) string {
	if len(data) == 0 {
		return ""
	}

	if len(data) == 1 {
		return simpleHash(data[0])
	}

	next := make([]string, 0)
	for i := 0; i < len(data); i += 2 {
		cur := data[i]
		if i+1 < len(data) {
			next = append(next, simpleHash(data[i+1]+data[i]))
		} else {
			next = append(next, simpleHash(cur+cur))
		}
	}

	//queue := list.New()
	//for _, v := range data {
	//
	//}

	return calculateMerkleRootByHash(next)
}

func calculateMerkleRoot(data []string) string {
	data2 := make([]string, 0)
	for _, datum := range data {
		data2 = append(data2, simpleHash(datum))
	}
	return calculateMerkleRootByHash(data2)
}

func main() {
	input := []string{"apple", "banana", "cherry", "date"}
	input2 := []string{"apple", "banana", "cherry"}
	fmt.Println(calculateMerkleRoot(input))
	fmt.Println(calculateMerkleRoot(input2))
}
