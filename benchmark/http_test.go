package benchmark
//
//import (
//	"encoding/json"
//	"io/ioutil"
//	"net/http"
//	"net/url"
//	"testing"
//)
//
//const (
//	ginAddress = "localhost:8081"
//)
//
//func BenchmarkHTTP(b *testing.B) {
//	uri := url.URL{
//		Scheme:   "http",
//		Host:     ginAddress,
//		Path:     "/http",
//		RawQuery: "name=tony",
//	}
//	var err error
//	var response *http.Response
//
//	for n := 0; n < b.N; n++ {
//		if response, err = http.Get(uri.String()); err != nil {
//			b.Fatalf("http request failed: %v", err)
//		}
//		if response == nil {
//			b.Fatalf("http response is wrong: %v", response)
//		}
//
//		var body []byte
//		if body, err = ioutil.ReadAll(response.Body); err != nil {
//			b.Fatalf("http body is wrong: %v", err)
//		}
//		if err = response.Body.Close(); err != nil {
//			b.Fatalf("http body close fail: %v", err)
//		}
//		type Body struct {
//			Message string
//		}
//		var obj Body
//		err = json.Unmarshal(body, &obj)
//	}
//	//fmt.Println(string(body))
//}
