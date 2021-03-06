### 工厂方法模式（Factory Method pattern）

##### 创建模式（Creational Pattern）


在讨论工厂方法模式之前，我们首先要讨论的一定是简单工厂(Simple Factory)，严格的说，简单工厂不是Gof23设计模式中的一种。

首先我们看看不使用简单工厂或者工厂方法之前，我们要创建对象的方式:

假设我们有汽车类型接口及其两种实现：一个benz，一个bwm

汽车类型car：

```go
type Car interface {
	Drive() string
}

type CarType struct {
	CarName string
	Size    int
}
```

两种实现：
```go
type Benz struct {
	CarType CarType
}

func (b Benz) Drive() string {
	return b.CarType.CarName
}

type Bmw struct {
	CarType CarType
}

func (b Bmw) Drive() string {
	return b.CarType.CarName
}
```

对两种汽车类型的调用：
```go
func main() {
	carTypeBenz := CarType{
		CarName: "benz",
		Size:    10,
	}

	benz := Benz{CarType: carTypeBenz}
	fmt.Println(benz.Drive())

	carTypeBmw := CarType{
		CarName: "bmw",
		Size:    9,
	}

	bmw := Benz{CarType: carTypeBmw}
	fmt.Println(bmw.Drive())
}
```

以上代码是没有使用任何工厂模式的用法，虽然并没有什么bug，但是如果代码中这样的接口实例化和使用存在多处:


```go
	carTypeBenz := CarType{
		CarName: "benz",
		Size:    10,
	}

	benz := Benz{CarType: carTypeBenz}
	fmt.Println(benz.Drive())

	carTypeBmw := CarType{
		CarName: "bmw",
		Size:    9,
	}

	bmw := Benz{CarType: carTypeBmw}
	fmt.Println(bmw.Drive())
 
    ...

	carTypeBenz := CarType{
		CarName: "benz",
		Size:    10,
	}

	benz := Benz{CarType: carTypeBenz}
	fmt.Println(benz.Drive())

	carTypeBmw := CarType{
		CarName: "bmw",
		Size:    9,
	}

	bmw := Benz{CarType: carTypeBmw}
	fmt.Println(bmw.Drive())

    ...
```

假设上面的代码片段大量存在于各种方法和函数中，正常情况下没有任何问题。但是假如汽车接口一旦发生了变化，比如对原有汽车结构体中的Size进行了拆分，分为了两个字段（wheelbase和length）,那么这个重构就会造成所有调用点都要改动，都要将size改成两个参数。

```go
	carTypeBmw := CarType{
		CarName: "bmw",
		Size:    9,
	}

>>

	carTypeBmw := CarType{
		CarName: "bmw",
		wheelbase:    2840,
		length      : 4704,
	}
```

所有使用该汽车对象的地方都需要改动（有100处使用，就要改动100处）。

所以，上述代码也没有实现代码的复用，创建汽车类型的代码在很多地方重复，这是造成如果使用了100处就要改动这100个地方的原因。

我们使用简单工厂模式对上述代码进行改动。

增加一个工厂函数，用于产生汽车对象。
```go
func GetCar(carType string) Car {
	carType1 := CarType{
		carName: carType,
		size:    10,
	}
	if carType == "benz" {
		benz := Benz{CarType: carType1}
		return benz
	} else if carType == "bmw" {
		bmw := Bmw{CarType: carType1}
		return bmw
	} else {
		return nil
	}
}
```

而新的客户端代码则变成了

```go
	car1 := simpleFactory.GetCar("bmw")
	fmt.Println(car1.Drive())

	car2 := simpleFactory.GetCar("benz")
	fmt.Println(car2.Drive())
```

从上面的代码可以看出，GetCar是一个工厂函数，其负责了汽车对象的创建，这样就解决了对象的创建和使用分离的问题，工厂创建对象，客户端使用对象。回到上面的问题：即便有100处代码使用了汽车类型，调用点不受任何影响，不会发生改动。只是工厂函数内部改动即可：

```go
func GetCar(carType string) Car {
	carType1 := CarType{
		carName: carType,
		wheelbase:    2840,
		length      : 4704,
	}
	if carType == "benz" {
		benz := Benz{CarType: carType1}
		return benz
	} else if carType == "bmw" {
		bmw := Bmw{CarType: carType1}
		return bmw
	} else {
		return nil
	}
}
```

这样一来，代码质量就高了很多。

上述改动，只是简单的使用了简单工厂对对象的创建进行了处理，就达到了很好的效果。

似乎看起来，简单工厂已经很好的解决了很多问题，为什么还会有工厂方法模式的存在？

简单工厂的问题在于，上述代码中的GetCar函数承担了太多职责，一旦有很多种类型的汽车要创建，那个 if 代码块就会不断膨胀，维护起来会非常困难。

此时我们使用工厂方法模式解决该问题。工厂方法和简单工厂的差异是在于：简单工厂是在上面所提到的工厂函数中进行判断和决策，以绝对要创建的对象类型，这也是简单工厂的问题之一，如果让工厂函数承担了所有的判断逻辑，一旦对象的类型增多的时候，判断逻辑部分的代码会变得过于复杂。而工厂方法模式，***则是将判断逻辑部分转移到客户端代码中。*** 这是工厂方法和简单工厂最大的差异之一。

#### 定义

***Exposes a method for creating objects, allowing subclasses to control the actual creation process.***

***定义一个用于创建对象的接口，让子类决定实例化哪一个类。工厂方法模式使一个类的实例化延迟到其子类。***

![Factory_Method Pattern UML](https://github.com/nox60/go-design-pattern/blob/master/images/factory_method_patern.png)

### 优点

将对象的创建和使用完全良好分离，很好的提高了代码的重用度。

### 例子分析

##### 一. 汽车案例

汽车案例的例子在上面的简单工厂中我们已经看到。也讨论了简单工厂的问题。

此处，我们为每一种对象（汽车类型）增加一个构建工厂方法。

首先是工厂接口

```go
type Factory interface {
	GetCar() Car
}
```

然后每个汽车对象均有一个创建工厂，来创建对应的汽车对象：

```go
type BenzFactory struct {
}

func (*BenzFactory) GetCar(carName string) Car {
	cartype := CarType{
		CarName:  carName,
		CarBrand: "benz",
		Size:     10,
	}

	benz := Benz{
		CarType: cartype,
	}

	return benz
}
```

```go
type BmwFactory struct {
}

func (*BmwFactory) GetCar(carName string) Car {
	cartype := CarType{
		CarName:  carName,
		CarBrand: "bmw",
		Size:     10,
	}

	bmw := Bmw{
		CarType: cartype,
	}

	return bmw
}
```

这样，在客户端代码中，调用不同的工厂生产不同的对象：

```go
	bmwFactory := factoryMethod.BmwFactory{}
	bmw1 := bmwFactory.GetCar("bmw1")
	fmt.Println(bmw1.Drive())

	benzFactory := factoryMethod.BenzFactory{}
	benz1 := benzFactory.GetCar("my benz")
	fmt.Println(benz1.Drive())
```

### 总结
工厂方法模式是代码封装性的一个良好体现，很好的将对象的创建和使用隔离开。在开发中合理的使用，会体会到开闭原则的真正意义。

### 参考文档：

https://en.wikipedia.org/wiki/Visitor_pattern

https://dzone.com/articles/design-patterns-visitor

https://blog.51cto.com/zero01/2067822

https://zhidao.baidu.com/question/131823823.html

https://blog.csdn.net/qq_34337272/article/details/80472071