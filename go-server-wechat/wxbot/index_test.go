package wxbot

import (
	"encoding/xml"
	"fmt"
	"testing"
)

type PublicMsg struct {
	MsgType      string `form:"MsgType" json:"MsgType" uri:"MsgType" xml:"MsgType"`
	Content      string `form:"Content" json:"Content" uri:"Content" xml:"Content"`
	ToUserName   string `form:"ToUserName" json:"ToUserName" uri:"ToUserName" xml:"ToUserName"`
	FromUserName string `form:"FromUserName" json:"FromUserName" uri:"FromUserName" xml:"FromUserName"`
}

func TestXML(t *testing.T) {
	teststr := ` <xml>
    <ToUserName><![CDATA[gh_b68073381519]]></ToUserName>
    <FromUserName><![CDATA[oBRxus-mD2Ogr3qJgKnm06Ks7Vl4]]></FromUserName>
    <CreateTime>1673426271</CreateTime>
    <MsgType><![CDATA[text]]></MsgType>
    <Content><![CDATA[哈哈哈]]></Content>
    <MsgId>23957705098404427</MsgId>
    <Encrypt><![CDATA[oLXRiv5Tn3pwl0EQRpQWYNcZs4g+6uDKz6BgWMai4GraJEadGxUVLWQudReadwurSOtoRARy9FQarqDWvyyfeCG5wHW1SxYzU9pj+mNoMzbZCjSABKkcwy/w3w++FqyvIUDge0UK2wgcv6vYJdzMk0wlerOp4oRZ6+/UaNOwZ9UbEkMgx1wjyGa5et5q/lkw7kG1RBfsdWIPr0tOq0m5Pz8lkDikkbVKVx0KgE33HnF9GrV3EFudlCvmCR7FI/19VEx+RZf1WbsZw587xJ9FUOIiV/pxQybPTbxW/zQAECtA36GzCxQPS33OLaeIfU7yMg13VjTXBDnIz99s3a+SEb9R1OC1I4xfVkVaFzPFnccoQcoFJ7wwKZ8nmRh54WtusGs2fQvypJ+5jzSjf/Vv6E9rwU0XoQZXbJA7FBnYAno=]]></Encrypt>
</xml>`

	var publicMsg PublicMsg

	if err := xml.Unmarshal([]byte(teststr), &publicMsg); err != nil {

	}
	fmt.Println(publicMsg)

}
