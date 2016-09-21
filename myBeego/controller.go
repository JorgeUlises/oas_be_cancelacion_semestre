package myBeego

import (
	"log"
	//"reflect"
	// "bytes"
	// "io"
	// "strconv"

	copyBeego "github.com/astaxie/beego"
	copyContext "github.com/astaxie/beego/context"
)

type BeegoInput struct {
	copyContext.BeegoInput
}

// type BeegoOutput struct {
// 	copyContext.BeegoOutput
// }
//

//type Context copyContext.Context
// type Context struct {
// 	copyContext.Context
// }

//var Controller *copyBeego.Controller
//http://golangtutorials.blogspot.com.co/2011/06/anonymous-fields-in-structs-like-object.html
//https://play.golang.org/p/1iYzdoFqIC
type Controller struct {
	copyBeego.Controller
}

func (c *Controller) Prepare() {
	log.Println("hi")
}

// func (input *Controller) Ctx.Input.Query(key string) string {
//   log.Println("hi2")
//   if val := input.Param(key); val != "" {
// 		return val
// 	}
// 	if input.Context.Request.Form == nil {
// 		input.Context.Request.ParseForm()
// 	}
//   return input.Context.Request.Form.Get(key)
//   //return c.Ctx
// }

// func (output *BeegoOutput) Body(content []byte) error {
// 	log.Println("Hola mundo")
// 	var encoding string
// 	var buf = &bytes.Buffer{}
// 	if output.EnableGzip {
// 		encoding = copyContext.ParseEncoding(output.Context.Request)
// 	}
// 	if b, n, _ := copyContext.WriteBody(encoding, buf, content); b {
// 		output.Header("Content-Encoding", n)
// 	} else {
// 		output.Header("Content-Length", strconv.Itoa(len(content)))
// 	}
// 	// Write status code if it has been set manually
// 	// Set it to 0 afterwards to prevent "multiple response.WriteHeader calls"
// 	if output.Status != 0 {
// 		output.Context.ResponseWriter.WriteHeader(output.Status)
// 		output.Status = 0
// 	} else {
// 		output.Context.ResponseWriter.Started = true
// 	}
// 	io.Copy(output.Context.ResponseWriter, buf)
// 	return nil
// }

// // Query returns input data item string by a given string.
// func (input *BeegoInput) Query(key string) string {
//   log.Println("Hola mundo")
// 	if val := input.Param(key); val != "" {
// 		return val
// 	}
// 	if input.Context.Request.Form == nil {
// 		input.Context.Request.ParseForm()
// 	}
// 	return input.Context.Request.Form.Get(key)
// }
