# bulkdo

a bulkdo tool, just for test and fun

design:

## readItems

1. 创建 bulkdo 文件
2. 创建 readItems 方法
3. 编写测试方法
4. 完成 readItems 方法
5. 根据异常测试用例，完善 readitems 方法

## parseCommands

1. 创建 parseCommands 方法声明，接受两个参数，1 个是模板字符串 reader，一个是 items，返回 commands string slice, error
2. 编写测试方法
3. 完成 parseCommands

## execCommands

1. 创建 execCommands 方法声明，接受一个参数，commands string slice, 返回 output string silce 和 error
2. 编写测试方法
3. 完成 execCommands


## BulkDo 和 命令行

首先是 BulkDo 方法 （其实就是将前面的方法组合一下就ok了）
1. 创建 BulkDo 方法声明，接受 2 个参数，一个是模板 reader， 一个是参数的 reader，返回 2 个值， output string slice 和 error
2. 编写测试方法
3. 完成 BulkDo

然后是命令行， 命令行的运行方式是：
```
bulkdo myecho1
```

读取 ~/.bulkdo/ 下的命令行模板，和当前路径下的 .bulkdoitems 参数， 然后调用 BulkDo 方法


## 总结
总结一下，就是从一个文件中读取批量命令的参数，填充到模板中去，再执行它