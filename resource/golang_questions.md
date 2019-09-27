## Q1：为什么下面这两段代码跑起来会出问题？

```golang
package main

import (
	"fmt"
)

type Person struct{
	Age int
}

func (p *Person)HowOld(){
	fmt.Println(24,"years old")
}
func main() {
	Person{}.HowOld()
	
}
```

错误信息：
```
./prog.go:15:10: cannot call pointer method on Person literal
./prog.go:15:10: cannot take the address of Person literal
```

```golang
package main

import (
	"fmt"
)

type Person struct{
	Age int
}

func (p Person)HowOld(){
	fmt.Println(24,"years old")
}
func main() {
	&Person{}.HowOld()
	
}
```
错误信息：
```
./prog.go:15:18: Person literal.HowOld() used as value
```


**因为接收器(receiver)的类型要与实例的类型相同,如果不相同则会报错。**

如果把上面代码稍微改一下，把Howold的方法里的接受改为值类型

```golang
package main

import (
	"fmt"
)

type Person struct{
	Age int
}

// 就是把*Person 改为Person类型
func (p Person)HowOld(){
	fmt.Println(24,"years old")
}
func main() {
	Person{}.HowOld()
	
}
```

这段代码是可以跑的

或者改为这样,也是可以跑的。

```golang
package main

import (
	"fmt"
)

type Person struct{
	Age int
}

func (p *Person)HowOld(){
	fmt.Println(24,"years old")
}
func main() {
	// 或者在这里加一个括号
	(&Person{}).HowOld()
	
}
```


Q2：那为什么创建一个实例后，赋值给新的变量又可以了呢？

比如这段代码

```golang
package main

import (
	"fmt"
)

type Person struct{
	Age int
}

func (p *Person)HowOld(){
	fmt.Println(24,"years old")
}
func main() {
	a:=Person{}
	a.HowOld()
	
}
```

因为在把值赋给变量，再重新调用方法，编译器自动帮你加上了。

编译器把a.HowOld() 改成(&a).HowOld,所以编译通过了。

官方文档里有说到，[点这里](https://tour.golang.org/methods/6) (需科学上网方可访问)
