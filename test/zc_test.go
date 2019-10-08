/**
 * @Author: DollarKiller
 * @Description: 杂项测试
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 15:38 2019-09-24
 */
package test

import (
	"log"
	"ninemanga-reptile/utils"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

func TestOne(t *testing.T) {
	str := "The Prince's Private Child Chapitre 01 "
	reg := `(\d+)`
	compile := regexp.MustCompile(reg)
	submatch := compile.FindAllStringSubmatch(str, -1)
	i, e := strconv.Atoi(submatch[0][1])
	if e == nil {
		log.Println(i)
	}
}

func TestTwo(t *testing.T) {
	utils.GetPage("10/23")
}

func TestSSc(t *testing.T) {
	url := "http://fr.ninemanga.com/chapter/The%20Prince%27s%20Private%20Child/16548-10-1.html"
	index := strings.LastIndex(url, "-")

	url = url[:index+1] + "2" + url[index+2:]
	log.Println(url)
}
