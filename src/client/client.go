/**
 * Created by chaolinding on 2018/3/19.
 */

package client

import (
	"google.golang.org/grpc"
	"log"
	"proto"
	"context"
)

func ClientStart(){
	conn, err := grpc.Dial("127.0.0.1:8021", grpc.WithInsecure())//不使用证书
	if err != nil {
		log.Fatal( err )
	}

	client := myproto.NewAddressBookStoreClient(conn)
	resp, err := client.AddPerson(context.TODO(), new(myproto.AddPersonRequest))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("id==========",resp.GetId())
}