package main

import "fmt"
//import "net/url"
import "net/http"
//import "strings"
import "log"

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RequestURI)
//	PraseForm()	
	

	if r.FormValue("resolution") != "" {
		
		fmt.Printf("%s is %s", "resolution", r.FormValue("resolution"))
		fmt.Println("Yes\n");
	} else {
		fmt.Println("No\n");
		return
	}

	fmt.Println("GoodBoy\n");

//	fmt.Println(r.Proto)
//	// output:HTTP/1.1
//	fmt.Println(r.TLS)
//	// output: <nil>
//	fmt.Println(r.Host)
//	// output: localhost:9090
//	fmt.Println(r.RequestURI)
//	// output: /index?id=1
//
//	scheme := "http://"
//	if r.TLS != nil {
//		scheme = "https://"
//	}
////	fmt.Println(strings.Join([]string{scheme, r.Host, r.RequestURI}, ""))
    // output: http://localhost:9090/index?id=1

//我们将解析这个 URL 示例，它包含了一个 scheme，认证信息，主机名，端口，路径，查询参数和片段。
    //s := "postgres://user:pass@host.com:5432/path?k=v&g=123#f"
//解析这个 URL 并确保解析没有出错。
//	u := r.RequestURI
////    u, err := url.Parse(s)
////    if err != nil {
////        panic(err)
////    }
////直接访问 scheme。
//    fmt.Println(u.Scheme)
////User 包含了所有的认证信息，这里调用 Username和 Password 来获取独立值。
//    fmt.Println(u.User)
//    fmt.Println(u.User.Username())
//    p, _ := u.User.Password()
//    fmt.Println(p)
////Host 同时包括主机名和端口信息，如过端口存在的话，使用 strings.Split() 从 Host 中手动提取端口。
//    fmt.Println(u.Host)
//    h := strings.Split(u.Host, ":")
//    fmt.Println(h[0])
//    fmt.Println(h[1])
//这里我们提出路径和查询片段信息。
//    fmt.Println(u.Path)
//    fmt.Println(u.Fragment)
////要得到字符串中的 k=v 这种格式的查询参数，可以使用 RawQuery 函数。你也可以将查询参数解析为一个map。已解析的查询参数 map 以查询字符串为键，对应值字符串切片为值，所以如何只想得到一个键对应的第一个值，将索引位置设置为 [0] 就行了。
////    fmt.Println(u.RawQuery)
////    m, _ := url.ParseQuery(u.RawQuery)
////    fmt.Println(m)
////    fmt.Println(m["k"][0])
////
////	if m["g"][0] != "" {
////		fmt.Println("Yes!")
////    	fmt.Println(m["g"][0])
////	} else {
////		fmt.Println("No!")
////	}
}

func main() {
    http.HandleFunc("/", sayhelloName) //设置访问的路由
	err := http.ListenAndServe(":2000", nil) //设置监听的端口
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

