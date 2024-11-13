package main

import (
	"fmt"
	"log"
	"os/user"
	"strings"
)

func main() {
	log.SetFlags(19)

	fmt.Println("巫术与黑魔法，奥丁与八足天马，漫漫的黑色荒原，" +
		"终年白雪皑皑的国度，少年负剑跋涉千里，收集亡灵的瓦基丽，" +
		"中土大陆打造着指环的炼金工匠，红发绿甲弓箭手的箭翎是凤凰羽，" +
		"暗夜精灵为他的子民们吟唱不死的咒语，三头四臂的狗给涂满鱼腥草的洞穴守门，" +
		"莱戈拉斯的金长发在阳光下依旧熠熠生辉。")

	info, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	username := info.Username
	tmp := strings.Split(username, "\\")

	fmt.Println("Hello, " + tmp[1] + "!")
}
