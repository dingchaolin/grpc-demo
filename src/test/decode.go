/**
 * Created by chaolinding on 2018/3/16.
 */

package test

import (
	"proto"
	"github.com/golang/protobuf/proto"
	"log"
	"fmt"
)

func Code()[]byte{
	pn := []*myproto.PhoneNumber{
		&myproto.PhoneNumber{
			Number: "15600287030",
			Type: myproto.PhoneType_MOBILE,
		},
		&myproto.PhoneNumber{
			Number: "18810042351",
			Type: myproto.PhoneType_HOME,
		},

	}

	p := &myproto.Person{
		Id: 13,
		Name:  "dcl",
		Email: "dcl@163.com",
		Phones: pn,
	}
	// 进行编码
	data, err := proto.Marshal(p)
	if err != nil {
		fmt.Println("marshaling error: ", err)
	}

	return data
}

func Decode( buf []byte)myproto.Person{
	var p myproto.Person

	err := proto.Unmarshal(buf, &p)
	if err != nil {
		log.Fatal( err )
	}
	return p
}