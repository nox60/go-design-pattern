### 备忘录模式（Memento Pattern）

##### 行为模式（Behavioural Pattern）

#### 定义

***Allows for capturing and externalizing an object's internal state so that it can be restored later, all without violating encapsulation.***

***在不破坏封装性的前提下，捕获一个对象的内部状态，并在该对象之外保存着这个状态。这样以后就可将该对象恢复到原先保存的状态。***

![Memento Pattern UML](https://github.com/nox60/go-design-pattern/blob/master/images/memento_pattern.png)

备忘录模式的核心其实就是reload这个操作。

在游戏过程中，我们经常有先保存当前游戏状态，然后去打怪或者做其他事情，一旦发现状况不对，则读取刚才的存档已避免损失这样的操作。

这就是备忘录模式的典型场景，我们所保存的存档，就是备忘录。

### 优点
1. 给用户提供了一种可以恢复状态的机制，可以使用能够比较方便地回到某个历史的状态。
2. 实现了信息的封装，使得用户不需要关心状态的保存细节。

### 例子分析

##### 游戏案例

1. 游戏程序，有体力HP，耐力SP，魔法力MP三个值。
2. 游戏玩家可以存档。
3. 一旦游戏角色GAME OVER，可以读取最近的存档重新开始。

代码如下：

备忘录对象，用来保存游戏状态（在现实中往往保存在数据库或者缓存中）：
```go
type Memento struct {
	HP int
	SP int
	MP int
}
```

游戏程序：
```go
type Game struct {
	HP int
	SP int
	MP int
}


func (g *Game) Play(HP int, SP int, MP int) {
	g.HP += HP
	g.SP += SP
	g.MP += MP
}

func (g *Game) Save() Memento {
	return Memento{
		HP: g.HP,
		SP: g.SP,
		MP: g.MP,
	}
}

func (g *Game) Load(m Memento) {
	g.HP = m.HP
	g.SP = m.SP
	g.MP = m.MP
}

func (g *Game) Status() {
	fmt.Printf("当前生命值 HP: %d，当前耐力: %d, 当前魔法力: %d\n", g.HP, g.SP, g.MP)
}
```

测试程序：
```go
func main() {
	game := &Game{
		HP: 10,
		SP: 8,
		MP: 6,
	}

	game.Status()
	progress := game.Save()
	game.Play(-10, -8, -6)
	game.Status()
	fmt.Println("游戏Game Over, 重新读档")
	game.Load(progress)
	game.Status()
}
```

输出
```go
当前生命值 HP: 10，当前耐力: 8, 当前魔法力: 6
当前生命值 HP: 0，当前耐力: 0, 当前魔法力: 0
游戏Game Over, 重新读档
当前生命值 HP: 10，当前耐力: 8, 当前魔法力: 6
```

### 参考文档：
