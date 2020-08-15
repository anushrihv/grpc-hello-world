/*
 * Copyright 2006-2020 (c) Care.com, Inc.
 * 1400 Main Street, Waltham, MA, 02451, U.S.A.
 * All rights reserved.
 *
 * This software is the confidential and proprietary information of
 * Care.com, Inc. ("Confidential Information").  You shall not disclose
 * such Confidential Information and shall use it only in accordance with
 * the terms of an agreement between you and CZen.
 */
package main

import (
	idlpb "bitbucket.use.dom.carezen.net/users/anushri.veerarajan/repos/hello-world/protogen"
	"bitbucket.use.dom.carezen.net/users/anushri.veerarajan/repos/hello-world/service"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

/*
 * Created @ 15/07/20 1:34 AM
 *
 * @author Anushri Hatti
 */

func main(){
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error occurred while listening")
		return
	}
	grpcServer := grpc.NewServer()
	idlpb.RegisterHelloWorldServer(grpcServer, &service.HelloWorldServer{})
	fmt.Println("Listening at port 8080")
	if err = grpcServer.Serve(listener); err != nil {
		fmt.Println("Error occurred while serving ", err)
	}
}