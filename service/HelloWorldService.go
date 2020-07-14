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
package service

import (
	idlpb "bitbucket.use.dom.carezen.net/users/anushri.veerarajan/repos/hello-world/protogen"
	"context"
)

/*
 * Created @ 15/07/20 12:36 AM
 *
 * @author Anushri Hatti
 */

type HelloWorldServer struct {

}

func (h HelloWorldServer) SayHello(ctx context.Context, request *idlpb.SayHelloRequest) (*idlpb.SayHelloResponse,
	error) {
	name := request.Name
	greeting := "Hello " + name
	return &idlpb.SayHelloResponse{Greeting: greeting}, nil
}






