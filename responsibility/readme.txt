责任链模式


审批财务报销:

//  Leader 直接上级只能审核500元以下的报销
// 	Director 总监只能审核5000元以下的报销
// 	CFO 首席财务官可以审核任意金额的报销


调用的方式非常优美

handlerChain := HandlerChain{}
handlerChain.AddHandler(leader) // 添加leader审批实现
handlerChain.AddHandler(director) // 添加director审批实现
handlerChain.AddHandler(cfo) // 添加director审批实现

request := Request{
	Name:   "kay",
	Amount: 200,
}
handlerChain.HandleRequest(request)

调用方无需考虑会交给谁处理，责任链会自动转移处理handler


在代码之中实际的应用就像是那个 auth 模块的鉴权 ，当前端传过来一个token的时候，校验 token是否有效

token授权有很多种，短信、密码、rsa证书，那么依据责任链模式将token依次传递给各自的实现判定是否有效 ～

