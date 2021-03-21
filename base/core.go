package base

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	fmt.Println("core init")
	testFor()
	testPoint()
	testSyntax()
	testDefer()
	testGoMuti()
	testContext()
}

func testFor() {
	strings := []string{"google", "runoob"}
	for i, s := range strings {
		fmt.Println(i, s)
	}
}
func testPoint() {
	fmt.Println("Hello, World!")
	var age = -1
	var name = "abc"
	fstr := &name
	fmt.Println(age, name, fstr)
	fmt.Println(*fstr)
}

func testGoMuti() {
	go func() {
		ch := make(chan bool) //阻塞的channel
		go func() {
			rand.Seed(time.Now().UnixNano())
			for a := rand.Intn(2000); ; a = rand.Intn(2000) {
				if a%2 == 0 {
					fmt.Println("gorutine", a)
					ch <- true
					return
				}
			}
		}()
		/**
		运行到这里该取值，实际还没有放进去，所以等待放入之后再执行取的操作
		ch是阻塞，里面的东西不被读掉，程序无法往下执行
		*/
		<-ch
		fmt.Println("阻塞的 gorutine finish")
	}()
}

func testContext() {
	var ctx, cancel = context.WithCancel(context.Background())
	defer cancel()
	go func(ctx context.Context) {
		rand.Seed(time.Now().UnixNano())
		for {
			a := rand.Int()
			if a%9 == 0 {
				fmt.Println("context gorutine 处理完成", a)
				<-ctx.Done()
				return
			} else {
				fmt.Println("监控 goroutine 中...", a)
			}
		}
	}(ctx)

	fmt.Println("主进程运行中。")
}

func testSyntax() {
	var a interface{}
	a = "abc"
	switch t := a.(type) {
	case string:
		fmt.Printf("string %s\n", t)
	case int:
		fmt.Printf("int %d\n", t)
	default:
		fmt.Printf("unexpected type %T", t)
	}
	//a := sync.Pool{}
}

/*
内部 defer 标记的函数按照 先进后出 执行
*/
func testDefer() (n int) {
	defer func() {
		n++
		fmt.Println("2st:", n)
	}()

	defer func() {
		n++
		fmt.Println("1st:", n)
	}()

	return n
}

func TestPanicRecover() {
	defer func() {
		if info := recover(); info != nil {
			fmt.Println("抓住异常->", info)
			//debug.PrintStack()

		} else {
			fmt.Println("hello")
		}
	}()

	defer func() {
		fmt.Println("最后的defer!")
	}()

	panic("刚抛出的异常")

	/**
	异常后的代码不执行
	*/
	fmt.Println("3")
	defer func() {
		fmt.Println("4")
	}()
}
