# patterns_go

golang解释设计模式


桥接模式 - JDBC的扩展\MySQL-JDBC\oracle-JDBC...  -- 侧重多模块组装

抽象工厂模式 - 文件日志\redis日志\Elastic日志...   -- 侧重行为的抽象定义，同样的抽象多个具体的实现

适配器模式 - 数据查找 - MySQL的搜索语句和elastic的搜索语句DSL-data search language不一样的通过一个适配器统一搜索语句

责任链模式 - token鉴权比如各种授权方式的鉴权handler校验





晋升指南


" 类命名多用 proxy\adapter\chain之类的，就会显得很有水准 "



“设计模式不拘泥于实现形式，它的思想更多的是引领我们在编码的时候，尽可能的少面向过程，尽可能的解决模块耦合的问题” 