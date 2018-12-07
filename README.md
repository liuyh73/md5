## MD5
### golang环境配置
首先，需要安装配置golang环境，[具体步骤](https://blog.csdn.net/liuyh73/article/details/82874981)
### 下载安装md5包

```bash
go get -d github.com/liuyh73/md5
```

### 使用md5包
```go
package main

import (
	"fmt"

	"github.com/liuyh73/md5"
)

func main() {
	fmt.Println("The quick brown fox jumps over the lazy dog: ", md5.Encrypt("The quick brown fox jumps over the lazy dog")) //9e107d9d372bb6826bd81d3542a419d6
	fmt.Println("The quick brown fox jumps over the lazy cog: ", md5.Encrypt("The quick brown fox jumps over the lazy cog")) // 1055d3e698d289f2af8663725127bd4b
	fmt.Println(": ", md5.Encrypt(""))                                                                                       // d41d8cd98f00b204e9800998ecf8427e
}
```

[博客链接](https://blog.csdn.net/liuyh73/article/details/84873413)