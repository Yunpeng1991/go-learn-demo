## 接口定义
```go

package main

func main(){
}

type Retriever interface{ Get(url string)string }

```

## 接口实现
1\主要是类型就可以实现接口  
2\必须实现接口的所有func才算实现接口
```go
package main

func main(){
}

//declare function
type Adder func() int  

//implements io.Read
func (adder Adder) Read(data []byte)(int,error){ return 0,nil }

```


## 接口组合