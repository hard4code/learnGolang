#### 符号导出 init方法
    包中所有以大写字母开头的符号将被导出，其余私有
    包中所有init方法在import后被执行

#### import 路径形式 [_ path] | [. path] | [别名 path]
    _ 仅仅为了调用init()，无法通过包名来调用包中的其他函数
    . 包导入之后在调用这个包的函数时，可以省略前缀的包名
    别名 为包定义一个别名,避免相同的包

#### 函数 func (类型限定) funcName (参数) (返回类型)
    类型限定 用于限定调用当前函数的对象类型，指针用于修改，引用用于获取
    
#### defer return 返回值
    在声明时不会执行，在函数return后，按照先进后出原则依次执行defer
    defer一般用于异常处理、释放资源、清理数据、记录日志等
>###### defer、return、返回值三者的执行顺序：
* return最先执行将结果写入返回值
* defer执行收尾工作
* 最后函数携带当前返回值退出

#### panic抛异常 recover抓异常
    位于panic之后的代码不再执行, 但是panic之前的defer延迟函数继续执行
    recover用于panic之前的defer函数中，用于捕获并处理异常