/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 18:48 2019-09-24
 */
package test

import (
	"io/ioutil"
	"ninemanga-reptile/utils"
	"testing"
)

func TestDowx(t *testing.T) {
	chrome := utils.StartChrome("http://fr.ninemanga.com/chapter/The%20Prince%27s%20Private%20Child/16548-10-1.html")
	chrome = utils.StartChrome("https://www.google.com/search?q=goog&rlz=1C5CHFA_enUS868&gbv=1&sei=lsqKXe_QNsHV-gSWk7WYDQ")
	chrome = utils.StartChrome("http://fr.ninemanga.com/chapter/The%20Prince%27s%20Private%20Child/16548-10-1.html")
	chrome = utils.StartChrome("https://www.google.com/search?q=goog&rlz=1C5CHFA_enUS868&gbv=1&sei=lsqKXe_QNsHV-gSWk7WYDQ")
	ioutil.WriteFile("he.html", []byte(chrome), 00666)
}
