/**
 * Created by chaolinding on 2018/3/16.
 */

package server

import (
	"net"
	"log"
	"google.golang.org/grpc"
	"proto"
	"sync"
	"golang.org/x/net/context"
)

type addressBookStoreServer struct {
	mutex sync.Mutex
	id int32
	book *myproto.AddressBook
}

func newAddressBookStoreServer()*addressBookStoreServer{
	return &addressBookStoreServer{
		id: 0,
		book: new(myproto.AddressBook),
	}
}

/*
ctx context.Context 为了控制超时 等信息...
 */
func (s* addressBookStoreServer)AddPerson(ctx context.Context, req *myproto.AddPersonRequest)(*myproto.AddPersonResponse, error){
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.id++
	log.Printf("add called: %v", req)
	return &myproto.AddPersonResponse{
		Id: s.id,
	}, nil
}

func ServerStart(){
	l, err := net.Listen("tcp", ":8021")
	if err != nil{
		log.Fatal( "grpc server start failed !")
	}
	server := grpc.NewServer()

	myproto.RegisterAddressBookStoreServer(server,newAddressBookStoreServer())

	server.Serve(l)
}