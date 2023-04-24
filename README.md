# sqlgen
一个go语言小工具，可以利用lex和ast, 把建表语句转换成go struct，

## tutorial

执行命名即可生成

`go run cmd/sqlgen.go --in database/test.sql --out models/shop_order.go --pkg  models`

本项目内置了一个简单的orm，可以根据实际需要更改。

因为orm用到了generic，所以go version >=1.8, 如果不需要这个orm，则可以降低go版本


## usage

参数说明
* `--in` sql语句保存文件
* `--out` 输出文件，如果不传，咋保存到`os.stdout`
* `--pkg` 输出文件的package名称

## v4版本

默认版本用的是tidb的parser，包含依赖较多，并不是很满意，所以开发了基于yacc的版本，但是v3只是初级版，ast的抽象程序较低，代码质量也不高，但是为了方便其他人学习，还是保留了下来。

## v5版本

v5是一个完整的版本,采用goyacc解析sql语句生成ast，然后根据ast用go自带的编译器生成go代码

生成parse文件`go run golang.org/x/tools/cmd/goyacc -l -o v5/out.go v5/lexer.y`

如果本地没有安装tools，记得先安装`go get -u golang.org/x/tools`

生成go文件`go run cmd/v5.go --in database/test.sql --out models/shop_order.go --pkg  models`

