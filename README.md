# patterns_go

golang解释设计模式


桥接模式 - JDBC的扩展\MySQL-JDBC\oracle-JDBC...  -- 侧重多模块组装

抽象工厂模式 - 文件日志\redis日志\Elastic日志...   -- 侧重行为的抽象定义，同样的抽象多个具体的实现


适配器模式 - Spring框架中DispatcherServlet类的doDispatch方法中也能看到适配器模式的身影，它的作用是传递request的Controller并执行相应的方法、返回ModeView对象。Spring定义了适配接口，每个Controller有自己对应的适配器实现类，从而避免了大量if/else代码
