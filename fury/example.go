/**
* PACKAGE fury
* Name example
* Description TODO
* Author yangqianlei@deltaphone.com.cn
* Date 2023/12/7/007
 */
package main

import (
	"fmt"

	furygo "github.com/alipay/fury/go/fury"
)

func main() {
	type SomeClass struct {
		F1 *SomeClass
		F2 map[string]string
		F3 map[string]string
	}
	fury := furygo.NewFury(true)
	if err := fury.RegisterTagType("example.SomeClass", SomeClass{}); err != nil {
		panic(err)
	}
	value := &SomeClass{F2: map[string]string{"k1": "v1", "k2": "v2"}}
	value.F3 = value.F2
	value.F1 = value
	bytes, err := fury.Marshal(value)
	if err != nil {
		panic(err)
	}
	fmt.Println(value)
	fmt.Println(value.F1.F1.F1)
	fmt.Println(string(bytes))
	var newValue interface{}
	// bytes can be data serialized by other languages.
	if err := fury.Unmarshal(bytes, &newValue); err != nil {
		panic(err)
	}
	fmt.Println(newValue.(*SomeClass).F1)
}
