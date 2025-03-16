# prototype
prototype 是一个用户克隆Go语言里任何变量的工具。

## 安装
```go
go get -u github.com/go-leo/prototype
```

# bool类型

| 源    |    目标     | 说明                                                                   |
|------|---------|----------------------------------------------------------------------|
| bool | bool    | 类型相同，直接复制                                                            |
| bool | any     | 底层也是bool，直接复制                                                        |
| bool | string  | strconv.FormatBool 转成字符串                                             |
| bool | int     | true:1, false:0                                                      |
| bool | uint    | true:1, false:0                                                      |
| bool | float   | true:1, false:0                                                      |
| bool | pointer | 解引用后再复制                                                              |
| bool | struct  | sql.NullXXX、wrapperspb.XXXValue、structpb.Value、anypb.Any，底层还是基础类型，同上 |
| bool | ClonerFrom  | 自定义克隆                                                                |

# int类型

| 源   | 目标         | 说明                                                                    |
|-----|------------|-----------------------------------------------------------------------|
| int | int        | 类型相同，直接复制                                                             |
| int | any        | 底层也是int，直接复制                                                          |
| int | uint       | 类型转化                                                     |
| int | float      | 类型转化                                                      |
| int | bool       | 0:false, !0:true                                                      |
| int | string     | strconv.FormatInt 转成字符串                                               |
| int | pointer    | 解引用后再复制                                                               |
| int | struct     | sql.NullXXX、wrapperspb.XXXValue、structpb.Value、anypb.Any，底层还是基础类型，同上  |
| int | ClonerFrom | 自定义克隆                                                                 |

# uint类型

| 源   | 目标         | 说明                                                                  |
|-----|------------|---------------------------------------------------------------------|
| uint | uint       | 类型相同，直接复制                                                             |
| uint | any        | 底层也是uint，直接复制                                                       |
| uint | int        | 类型转化                                                    |
| uint | float      | 类型转化                                                         |
| uint | bool       | 0:false, !0:true                                                    |
| uint | string     | strconv.FormatIUint 转成字符串                                           |
| uint | pointer    | 解引用后再复制                                                             |
| uint | struct     | sql.NullXXX、wrapperspb.XXXValue、structpb.Value、anypb.Any，底层还是基础类型，同上 |
| uint | ClonerFrom | 自定义克隆                                                               |

# float类型

| 源   | 目标         | 说明                                                              |
|-----|------------|-----------------------------------------------------------------|
| float | float      | 类型相同，直接复制                                                       |
| float | any        | 底层也是float，直接复制                                                       |
| float | uint       | 类型转化                                                           |
| float | int        | 类型转化                                                    |
| float | bool       | 0:false, !0:true                                                |
| float | string     | strconv.FormatFloat 转成字符串                                       |
| float | pointer    | 解引用后再复制                                                         |
| float | struct     | sql.NullXXX、wrapperspb.XXXValue、structpb.Value、anypb.Any，底层还是基础类型，同上 |
| float | ClonerFrom | 自定义克隆                                                           |

# string类型

| 源   | 目标        | 说明                                                                   |
|-----|-----------|----------------------------------------------------------------------|
| string | string    | 类型相同，直接复制                                                            |
| string | any       | 底层也是string，直接复制                                                      |
| string | []byte    | 类型转化                                                                 |
| string | []rune    | 类型转化                                                                 |
| string | int       | strconv.ParseInt                                                     |
| string | uint      | strconv.ParseUint                                                    |
| string | float     | strconv.ParseFloat                                                   |
| string | bool      | strconv.ParseBool                                                     |
| string | pointer   | 解引用后再复制                                                              |
| string | struct    | sql.NullXXX、wrapperspb.XXXValue、structpb.Value、anypb.Any，底层还是基础类型，同上 |
| string | ClonerFrom | 自定义克隆                                                                |



# []byte类型

| 源   | 目标        | 说明                                                                   |
|-----|-----------|----------------------------------------------------------------------|
| []byte | []byte    | 类型相同，直接复制                                                            |
| []byte | string    | 类型相同，直接复制                                                            |
| []byte | any       | base64后再转                                                            |
| []byte | int       | 先转出string，然后在转                                                       |
| []byte | uint      | 先转出string，然后在转                                                |
| []byte | float     | 先转出string，然后在转                                                   |
| []byte | bool      | 先转出string，然后在转                                                  |
| []byte | pointer   | 解引用后再复制                                                              |
| []byte | struct    | sql.NullXXX、wrapperspb.XXXValue、structpb.Value、anypb.Any，底层还是基础类型，同上 |
| []byte | ClonerFrom | 自定义克隆                                                                |


# time.Time类型

| 源   | 目标        | 说明                                                                 |
|-----|-----------|--------------------------------------------------------------------|
| time.Time | time.Time    | 类型相同，直接复制                                                          |
| time.Time | struct    | sql.NullXXX、wrapperspb.XXXValue、structpb.Value、anypb.Any，底层还是一样，同上 |
| time.Time | string    | TimeToString 函数转成string                                            |
| time.Time | []byte    | TimeToString 函数转成string，在转[]byte                                   |
| time.Time | any       | 还是time.Time，类型相同，直接复制                                              |
| time.Time | int       | TimeToInt转成int，然后在转                                                |
| time.Time | uint      | TimeToInt转成int，然后在转                                                     |
| time.Time | float     | TimeToInt转成int，然后在转                                                     |
| time.Time | pointer   | 解引用后再复制                                                            |
| time.Time | ClonerFrom | 自定义克隆                                                              |

slice类型

| 源         | 目标         | 说明                  |
|-----------|------------|---------------------|
| []byte    | x          | 走 bytes克隆           |
| slice     | any        | 底层还是slice，类型相同，直接复制 |
| slice     | slice      | 元素类型转化，然后转成slice    |
| slice     | array      | 元素类型转化，然后转成array    |
| slice    | pointer    | 解引用后再复制 |
| slice | ClonerFrom | 自定义克隆               |

array类型

| 源         | 目标         | 说明                  |
|-----------|------------|---------------------|
| array     | array      | 元素类型转化，然后转成array    |
| array     | slice      | 元素类型转化，然后转成slice    |
| array     | any        | 底层还是slice，类型相同，直接复制 |
| array    | pointer    | 解引用后再复制 |
| array | ClonerFrom | 自定义克隆               |

map类型

| 源     | 目标         | 说明                    |
|-------|------------|-----------------------|
| map   | map        | key和value类型转化，然后转成map |
| array | any        | 底层一样，类型相同，直接复制        |
| array | struct     | 按key找到字段，根据字段类型然后复制   |
| array | pointer    | 解引用后再复制               |
| array | ClonerFrom | 自定义克隆                 |

struct 类型

| 源              | 目标         | 说明        |
|----------------|------------|-----------|
| sql.XXX        | X          | 底层类型复制    |
| wrapperspb.XXX | X          | 底层类型复制    |
| anypb.Any      | X          | 接触原类型然后复制 |
| array          | pointer    | 解引用后再复制   |
| array          | ClonerFrom | 自定义克隆     |


# pointer类型
| 克隆类型       | 目标 | 说明          |
|------------|----|-------------|
| 浅克隆，并且类型相同 | X  | 指针赋值        |
| 深克隆或者类型不同  | X  | 追溯到非指针类型在复制 |


# interface类型

| 源      | 目标 | 说明          |
|--------|----|-------------|
| ClonerTo | X  | 自定义克隆       |
| 其他接口   | X  | 追溯到非接口类型在复制 |

