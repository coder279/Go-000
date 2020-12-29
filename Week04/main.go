package main

//多核CPU如何保证原子自增?

/**
 * 缺点:main函数是无法感知到go func的生命周期
 */
//func main() {
//	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//		fmt.Fprintln(w,"hello wolrd")
//	})
//	//不断的创建Goroutine,main退出所有线程也就退出
//	go func() {
//		if err := http.ListenAndServe(":8080",nil);err != nil {
//			log.Fatal(err) //底层调用os.Exit(1)直接对于程序进行退出,导致defer无法执行
//		}
//	}()
//空的select语句永远保持阻塞
//	select {}
//}

//第二种
//自己去做，不使用goroutine,单线程没有问题
//func main(){
//	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//		fmt.Fprintln(w,"Hello world")
//	})
//	if err := http.ListenAndServe(":8080",nil);err != nil {
//		log.Fatal(err)
//	}
//}

//第三种
//需要开一个goroutine去做另一件事的时候
//1.需要知道goroutine什么时候结束
//2.需要知道goroutine怎么结束

//func main(){
//	mux := http.NewServeMux()
//	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//		fmt.Fprintln(w,"Hello World")
//	})
//	go http.ListenAndServe("localhost:8081",http.DefaultServeMux)
//	http.ListenAndServe("0.0.0.0:8080",mux)
//}

//func serverApp(){
//	mux := http.NewServeMux()
//	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//		fmt.Fprintln(w,"Hello Lichen")
//	})
//	http.ListenAndServe("0.0.0.0:8080",mux)
//	time.Sleep(10*time.Second)
//	panic("异常")
//}
//
//func serverDebug(){
//	http.ListenAndServe("127.0.0.1:8001",http.DefaultServeMux)
//}
////写goroutine的人一定是调用者
////搞清楚什么时候停止,知道什么时候停止才能启动一个goroutine
//func main()  {
//	go serverDebug()
//	serverApp()
//}

//go平滑退出

//func main(){
//	done := make(chan error,2)
//	stop := make(chan struct{}) //0 size
//	go func() {
//		done <- serverDebug(stop)
//	}()
//
//	go func() {
//		done <- serverApp(stop)
//	}()
//
//	var stopped bool
//	for i:=0;i<cap(done);i++{
//		if err := <-done;err!=nil{
//			fmt.Println("error:%v",err)
//		}
//		if !stopped{
//			stopped = true
//			close(stop)
//		}
//	}
//}
//
//func serve(addr string,handler http.Handler,stop <-chan struct{})error{
//	s := http.Server{
//		Addr:addr,
//		Handler:handler,
//	}
//	go func() {
//		<-stop
//		s.Shutdown(context.Background())
//	}()
//	return s.ListenAndServe()
//}


/**
 * main函数控制goroutine的启动与否，将并行交给调用者
 */
//func main(){
//	done := make(chan error,1)
//	stop := make(chan struct{})
//	go func() {
//		done<-ServerApp(stop) //退出会给done发送一个信号
//	}()
//
//	for i:=0;i<cap(done);i++{
//		<-done //接收到信号 退出
//		close(stop)
//	}
//}
//
//func ServerApp(stop chan struct{}) error{
//	go func() {
//		<-stop
//		http.Shutdown()
//	}()
//	return http.ListenAndServe()
//}

//goroutine泄露
//func leak(){
//	ch := make(chan int)
//
//	go func() {
//		val := <-ch
//		fmt.Println(val)
//	}()
//}

//代码超时控制
//type result struct {
//	record string
//	err error
//}
//func Process(term string)error{
//	ctx,cancle := context.WithTimeout(context.Background(),100*time.Millisecond)
//	defer cancle()
//	ch := make(chan result)
//	go func() {
//		record,err := search(term)
//		ch <- result{record,err}
//	}()
//	select {
//		case <-ctx.Done():
//			return errors.New("search cancel")
//		case result := <-ch:
//			if result.err != nil{
//				return result.err
//			}
//			fmt.Println("Received:",result.record)
//			return nil
//	}
//}

func main(){
	var a Tracker
	a.Shutdown()
}