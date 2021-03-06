package test1

/*
	go语言的函数使用func关键字，特点如下：
		·无需前置申明
		·不支持命名嵌套定义: func xxx(){ func yyy(){}}
		·不支持同名重载函数
		·不支持默认参数
		·支持不定长参数
		·支持多返回
		·支持命名返回值
		·支持匿名函数和闭包
	go函数没有做尾递归优化，但是go的栈大小是GB级别的，因此不太会造成栈溢出
	然而仍然要关心拷贝栈的复制成本

	函数建议命名规则：
		·通常是动词、介词、名词的组合
		·避免不必要的缩写，printError 比 printErr 更好一些
		·避免使用类型关键字，buildUserStruct看上去会很别扭
		·避免歧义，不能有多种用途解释造成误解
		·避免智能通过大小写区分的同名函数
		·避免使用数字，除非是专有名词，例如UTF8
		·避免添加作用域提示前缀
		·统一使用驼峰（camelCase）/帕斯卡（PascalCase）命名方式
		·使用习惯用语，比如 init表示初始化，is/has表示返回布尔值
		·使用反义词命名行为相反的函数，比如get/set，min/max

	go的所有参数都是值传递，若想使用引用传递，需要使用指针参数
	注意：被复制的指针会延长目标对象的声明周期，还可能会导致它被分配到堆上，这样性能消耗就需要加上堆内存分配和垃圾回收的成本
	在栈上复制小对象只需要很少的指令即可完成，远比运行时进行堆内存分配要快的多。
	并发编程也提倡尽可能使用不可变对象，可以消除数据同步等麻烦。
	是否需要对原有对象状态进行修改，需要按具体情况具体分析。
*/
import (
	"log"
	"unsafe"
)

// Calculate 表示一种运算函数类型
type Calculate func(x int32, y int32) int32

/*
	定义一个结构体，里面的属性都是函数
*/
type MultiCalculate struct {
	Add   func(a int, b int) int
	Minus func(a int, b int) int
}

var add Calculate = func(x, y int32) int32 {
	log.Printf("%d + %d = %d \n", x, y, x+y)
	return x + y
}

func run(c Calculate, x int32, y int32) {
	c(x, y)
}

// FuncTest 方法用于测试函数类型
func FuncTest() {
	log.Printf("%T,%d \n", add, unsafe.Sizeof(add))
	run(add, 5, 6)

	mc := MultiCalculate{
		Add: func(a int, b int) int {
			return a + b
		},
		Minus: func(a int, b int) int {
			return a - b
		},
	}
	log.Printf("a + b = %d \n", mc.Add(1, 2))
	log.Printf("a - b = %d \n", mc.Minus(1, 2))
	ClosureTest()
	DeferTest()
}

/*
	闭包测试，闭包是指，外部函数可以访问当前函数的内部函数的变量
	在这里就是 ClosureTest 调用 ClosureFunc 所返回的匿名函数，还能正确的访问到x
	闭包得以实现，是因为，ClosureFunc返回的不仅仅是函数，还有x的变量的指针
	*注意*
	Go语言中函数的return不是原子操作，在底层是分为两步来执行
	第一步：返回值赋值
	defer
	第二步：真正的RET返回
	函数中如果存在defer，那么defer执行的时机是在第一步和第二步之间
*/
func ClosureFunc(x int) (func(), []func()) {
	log.Printf("outter -> x -> %p \n", &x)
	var s []func()
	// for循环的i是复用的，因此i的地址永远是一个不变
	for i := 0; i < 3; i++ {
		log.Printf("for -> i -> %p \n", &i)
		// x每次都会分配一个新的地址来放值，如果不使用x 返回的函数列表中的函数所调用的i都将为最终值
		j := i
		s = append(s, func() {
			x += i
			log.Printf("i -> %p,%d \n", &i, i)
			log.Printf("j -> %p,%d \n", &j, j)
			log.Printf("x -> %p,%d \n", &x, x)
		})
	}
	return func() {
		log.Println(x)
	}, s
}

func ClosureTest() {
	f1, f2 := ClosureFunc(5)
	f1()
	for _, f := range f2 {
		f()
	}
	// 上面for循环中将上下文环境中的x进行了+i操作，因此影响了下面f1()函数的值
	f1()

}

/*
	延迟调用
	return和panic语句都会终止当前函数流程，引发延迟调用
	由于return不是ret汇编指令，因此会先更新返回值
	执行顺序：x=100(return 100),call defer(x+=100),ret(return x)
	(x int) 是带命名返回值
*/
func DeferFunc() (x int) {
	defer log.Printf("defer1: x->[%p,%d] \n", &x, x)
	// 匿名函数作用域被隔离，也就是说相对外部作用域完全独立
	defer func() {
		log.Printf("defer2: x->[%p,%d] \n", &x, x)
		x += 100
	}()
	// 语句块作用域不隔离。
	{
		log.Printf("defer3: x->[%p,%d] \n", &x, x)
	}
	/*
		这里的return其实做了两件事
		1. 将x赋值为100
		2. reutrn x
		所以当赋值完成以后，开始执行defer，最终导致return的x值在第二个defer中被修改后返回
	*/
	return 100
}

/*
	延迟调用时注册的是调用，必须提供执行所需的参数（哪怕为空），
	所以参数再注册时就被复制并缓存起来。
	因此如果对参数比较敏感，可以使用闭包或者指针
*/
func DeferParamTest() {
	x, y := 1, 2
	defer func(a int) {
		// 打印 a = 1 y = 202 因为 y为DeferParamTest函数闭包引用，因此在整个闭包调用过程中保持一致
		// 而a在注册调用时，已经被复制为注册时x的值，也就是 1，因此a永远是1
		log.Printf("DeferParamTest() defer: x = %d, y = %d", a, y)
	}(x)

	x += 100
	y += 200
	log.Printf("DeferParamTest() main: x = %d, y = %d", x, y)
}

func DeferTest() {
	log.Printf("test: %d \n", DeferFunc())
}

/*
	如果不使用 FormatString 命令，则 Format参数签名将长到没法看
	type xxx func(a,b) (c,d) 表达式表示“函数签名”
	如果两个函数的“函数签名”相同，那么表示这两个函数的函数的函数名，参数和返回值(类型，个数，顺序)都相同。
*/
type FormatString func(string, ...interface{}) (string error)

func Format(f FormatString, s string, a ...interface{}) (string error) {
	return f(s, a)
}

/*
	返回局部变量指针是安全的
	go编译器会做变量逃逸分析，如果指针没有发生逃逸，则仍然会在栈中分配
	如果发生了逃逸，则会在堆中分配
*/
func ReturnPointTest() *int {
	a := 0x001
	return &a
}

/*
	向外传递参数，可以return，也可以使用二级指针
	SecondLevelPointTest的参数就是**int就是二级指针
	建议使用返回值，也就是return
*/
func SecondLevelPointUseTest() {
	var p *int
	SecondLevelPointTest(&p)
	println(p)
}

func SecondLevelPointTest(p **int) {
	x := 100
	*p = &x
}

/*
	如果参数过多可以使用符合结构类型作为参数，以实现可选参数
	对于多个地方公用的参数，可以创建initOptions来初始化参数返回参数对象指针，以免多个地方做初始化，便于代码复用
*/

type options struct {
	name     string
	password string
}

func MultiParameterTest(opts *options) {
	log.Printf("Parameter name=%s password=%s", opts.name, opts.password)
}

/*
	可变长参数本质就是一个slice
	由于是slice，因此参数复制的是切片，而不是底层数组
	对于底层数组的改变会影响参数
*/
func VariableLengthParameterTest(opts ...int) {
	for i, opt := range opts {
		opts[i] += 100 // 改变了外层数组的值
		opt += 1000    // 不会改变外层数组的值，因为是一个副本
	}
}

func VariableLengthParameterUseTest() {
	a := []int{1, 2, 3}
	VariableLengthParameterTest(a...)
	log.Printf("%+v", a)
}

/*
	go返回值可以有多个，并且可以对返回值进行命名
	对返回值命名可以让返回值含义更加清晰
	****注意 如果是结构体返回值变量，变量默认值为nil，必须做初始化
	命名返回值可以当局部变量使用，由retur隐式返回
*/
func ReturnTest(a int, b int) (total int, average int) {
	// 返回值
	total = a + b
	{
		// 返回参数average被该处定义的同名average参数遮蔽，因此无法隐式返回
		// 因此在这里需要显示声明返回参数
		// 如果写成 average = (a + b) / 2 则使用的是返回变量，就可以直接用return隐式返回
		average := (a + b) / 2
		return total, average
	}

	// return
}

/*
	函数签名：就是表示了函数的 出参，入参的类型和数量
	可以利用函数签名来实现 函数泛型
*/
func FuncSignatureTest() {
	// := 号的左边（add）是函数变量，右边是函数字面值
	add := func(x, y int) int { return x + y }
	x := mathCalc(1, 2, add)
	log.Printf("计算结果是：%d", x)

	minus := func(x, y int) int { return x - y }
	x = mathCalc(1, 2, Caculater(minus))
	log.Printf("计算结果是：%d", x)
	/*
		// 不符合函数签名的，会产生编译错误
		add1 := func(x, y, z int) int { return x + y + z }
		x := mathCalc(1, 2, add1)
	*/
	// 使用了函数签名Caculater强转的函数，就自动实现了 icalcer 的方法
	imathCalc(1, 2, Caculater(add))
}

// 定义一个接口，接口中有两个方法
type icalcer interface {
	caculate(int, int) int
}

// 定义一个函数签名，表示函数有两个int入参和一个int出参
type Caculater func(int, int) int

// Caculater 绑定一个方法叫 caculate
func (c Caculater) caculate(a int, b int) int {
	return c(a, b)
}

/*
	函数签名的应用1
	下面是一个数学计算使用的通用方法
	c Caculater 就变成了一个泛型，只要符合 Caculater 函数签名格式的函数都可以作为参数
	例如，写个 add(a int,b int) int 或者  minus(a int,b int) int 都可以作为参数
	这样就提高了灵活性，比如本来是用加法计算的，我现在想换成减法，就不需要改动 mathCalc 中的代码，只需要新写一个minus方法，并作为参数即可
	又例如 c Caculater 是一个加密算法的函数签名格式，那么 我现在用的是 加密方法a 想换成b 直接将b作为参数传递即可
*/
func mathCalc(a int, b int, c Caculater) int {
	return c(a, b)
}

/*
	函数签名的应用2
	如果使用了函数签名Caculater强转的函数，就自动实现了 icalcer 的方法
*/
func imathCalc(a int, b int, c icalcer) int {
	return c.caculate(a, b)
}
