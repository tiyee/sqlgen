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