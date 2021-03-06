### 装饰模式（Decorator pattern）
 
##### 创建模式（Creational Pattern）
 
 
#### 定义
 
***Allows for the dynamic wrapping of objects in order to modify their existing responsibilities and behaviors.***
 
***动态地给一个对象添加一些额外的职责或者功能。***
 
![Decorator Pattern UML](https://github.com/nox60/go-design-pattern/blob/master/images/decorator_method.png)

装饰器模式在日常研发中心是一个出场率非常高的模式，因为有太多的场景适用于该模式了。

装饰器模式常常用来动态的为已存在的功能增加新的功能，而不用对原有代码进行任何修改以造成影响。

比如，我们现在有一个订票功能OrderTicket。但是我们打算给他增加一些功能，比如订票之后发短信通知，或者发邮件通知。

听起来，似乎新建两个函数，SendMsgAfterOrderTicket（订票之后发短信），和SendEmailAfterOrderTicket（订票之后发邮件）。

这样我们在新的函数中实现发短信或者发邮件的功能，先调用原有的OrderTicket实现订票，然后在订票之后发送短信或者邮件。

但是如果我们需要既发送短信又发送邮件呢？如果在增加一个自动推送手机APP消息呢？如果希望只是选择性的选择其中的一种或者两种或者全部的通知呢？

重写的方式势必会比较多。如果继续增加通知方式，带来的更多麻烦则更不用说了。

这就是装饰器模式显示作用的地方了。

### 优点
 
1. 良好的给原有的对象增加了新的功能并不会入侵原有代码。
2. 能很好的组织新增加的功能，比如选择新功能的数量或者执行顺序，增加到原有功能的前面或者后面。
 
### 例子分析
 
##### 购票案例

1. 现在有一个订票功能，希望能够动态的增加一些其他功能，比如在订单完成之前短信通知用户，或者在订单完成之后邮件发送订单详情给用户，以及订单完成之后微信通知用户一些信息等。
2. 以上功能需要灵活的排列组合，有时候只需要短信，有时候需要短信+邮件，有时候是短信+微信信息。
3. 原有的订票功能已经实现，新的功能不能影响到原有功能。

代码如下:

原订单接口:
```go
type OrderTicketI interface {
	PlaceOrder(order Order)
}

type Order struct {
	UserId    string
	ProductId string
}
```

原订单实现
```go
type OrderTicketImplement struct {
}

func (p *OrderTicketImplement) PlaceOrder(order Order) {
	fmt.Println("完成订单， 产品ID：" + order.ProductId + ", 用户ID：" + order.UserId)
}
```

包装类，订单之前发送短信：
```go
type SendMsgBeforeOrder struct {
	Function OrderTicketI
}

func (s *SendMsgBeforeOrder) PlaceOrder(order Order) {
	sendMsg()
	s.Function.PlaceOrder(order)
}

func sendMsg() {
	fmt.Println("发送短信")
}
```

包装类，订单之后发送邮件:
```go
type SendEmailAfterOrder struct {
	Function OrderTicketI
}

func (s *SendEmailAfterOrder) PlaceOrder(order Order) {
	s.Function.PlaceOrder(order)
	sendEmail()
}

func sendEmail() {
	fmt.Println("发送邮件")
}
```

包装类，订单之后发送微信消息:
```go
type SendWeChatMessageAfterOrder struct {
	Function OrderTicketI
}

func (s *SendWeChatMessageAfterOrder) PlaceOrder(order Order) {
	s.Function.PlaceOrder(order)
	sendWeChatMessage()
}

func sendWeChatMessage() {
	fmt.Println("发送微信信息")
}
```

调用例子
```go
func main() {
	order := Order{
		UserId:    "10099202",
		ProductId: "09988779",
	}
	placeOrder := OrderTicketImplement{}
	sendMsgBeforeOrder := SendMsgBeforeOrder{}
	sendEmailAfterOrder := SendEmailAfterOrder{}
	sendWeChatMessageAfterOrder := SendWeChatMessageAfterOrder{}

	//例子1 完成订单之前先发送短信，订单完成之后发送邮件
	//使用短信发送功能包装原有的订票功能
	sendMsgBeforeOrder.Decorator = &placeOrder

	//使用邮件发送功能包装上面的短信发送功能(短信发送功能已经包装了原有的订票功能)，这样一来，就会实现先发送短信，然后订票，然后发送邮件
	sendEmailAfterOrder.Decorator = &sendMsgBeforeOrder
	sendEmailAfterOrder.PlaceOrder(order)

	fmt.Println()

	order2 := Order{
		UserId:    "10098547",
		ProductId: "09989948",
	}
	//例子2 完成订单之前先发送短信，订单完成之后发送邮件和微信消息
	//使用短信发送功能包装原有的订票功能
	sendMsgBeforeOrder.Decorator = &placeOrder

	//使用邮件发送功能包装上面的短信发送功能(短信发送功能已经包装了原有的订票功能)，这样一来，就会实现先发送短信，然后订票，然后发送邮件
	sendEmailAfterOrder.Decorator = &sendMsgBeforeOrder

	//使用微信消息功能包装上面的邮件发送功能(邮件发送功能已经在原有的短信包装了的订票功能上进行了包装)
	sendWeChatMessageAfterOrder.Decorator = &sendEmailAfterOrder
	sendWeChatMessageAfterOrder.PlaceOrder(order2)
}
```

输出结果：
```go
发送短信
完成订单， 产品ID：09988779, 用户ID：10099202
发送邮件

发送短信
完成订单， 产品ID：09989948, 用户ID：10098547
发送邮件
发送微信信息
```

### 总结
通过上面的例子可以看出，装饰模式的特点就体现在装饰（包装）上，可以很好的给原有的功能进行包装以增加新的功能，而且对原有的功能没有影响。

最重要的是，可以灵活的实现新增加的功能的选择，很大程度提高了代码的复用。
 
### 参考文档：
https://www.tutorialspoint.com/design_pattern/decorator_pattern.htm
