package main

import (
	"fmt"
)

const (
	east  = "Purvavideha"
	west  = "Aparagodaniya"
	south = "Jambudvipa"
	north = "Uttarakuru"
)

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
	info     string
}

type Trie struct {
	root *TrieNode
}

func NewTrieNode() *TrieNode {
	return &TrieNode{children: make(map[rune]*TrieNode)}
}

func NewTrie() *Trie {
	return &Trie{root: NewTrieNode()}
}

func (t *Trie) Insert(word, info string) {
	node := t.root
	for _, char := range word {
		if _, exists := node.children[char]; !exists {
			node.children[char] = NewTrieNode()
		}
		node = node.children[char]
	}
	node.isEnd = true
	node.info = info
}

func (t *Trie) Search(prefix string) []string {
	node := t.root
	for _, char := range prefix {
		if _, exists := node.children[char]; !exists {
			return []string{}
		}
		node = node.children[char]
	}
	return t.collect(node, prefix)
}

func (t *Trie) collect(node *TrieNode, prefix string) []string {
	var results []string
	if node.isEnd {
		results = append(results, fmt.Sprintf("%s: %s", prefix, node.info))
	}
	for char, child := range node.children {
		results = append(results, t.collect(child, prefix+string(char))...)
	}
	return results
}

func printContinents(trie *Trie) {
	entries := trie.Search("")
	for _, entry := range entries {
		fmt.Println(entry)
	}
}

func main() {
	trie := NewTrie()

	trie.Insert(east, "Dongseung Shinjoo: Son Oh Gong chulbal ji, Ryeongsan oe")
	trie.Insert(west, "Seou Nyeok Hoju: Ryeongsan sodji, chogeungui choejong mokjeokji")
	trie.Insert(south, "Nam Seonbuju: Ingan chaeseok ji, Dang Samjangui kohyang")
	trie.Insert(north, "Bukgu Rooju: Jungsun soju, Sinseon cheonseol ji ji")

	trie.Insert("Sun Wukong", "Son Oh Gong: Chogeung daeumui ryeongdui, cham-yeom cheol ma")
	trie.Insert("Tang Sanzang", "Dang Samjang: Jabi shim-jang, chogeungui balchija")
	trie.Insert("Zhu Bajie", "Joo Palk-gae: Mot-boleun ikyung, chogeung teamwon")
	trie.Insert("Sha Wujing", "Sa Ohjeong: Chogeung daeumui jungsangwon")
	trie.Insert("Bai Longma", "Baek Yongma: Malhyeong, Dang Seung ui hoong-seo saengsin")

	results := trie.Search("Sun")
	for _, result := range results {
		fmt.Println(result)
	}

	printContinents(trie)
}
