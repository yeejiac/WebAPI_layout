package internal

type EchoServer struct{}

// func Start_grpc()
// {
// 	apiListener, err := net.Listen("tcp", ":1203")
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	else{
// 		log.Println(" grpc.Serve bind success: ")
// 	}

// 	grpc := grpc.NewServer()

// 	reflection.Register(grpc)
// 	if err := grpc.Serve(apiListener); err != nil {
// 		log.Fatal(" grpc.Serve Error: ", err)
// 		return
// 	}
// 	else{
// 		log.Println(" grpc.Serve regist success: ")
// 	}
// }
