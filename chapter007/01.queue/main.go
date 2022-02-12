package main

import (
	"fmt"
	"sync"
)

type Queue struct {
	sync.Mutex
	data []interface{}
}

func (q *Queue) Push(data interface{}) {
	q.Lock()
	defer q.Unlock()
	q.data = append(q.data, data)

}

func (q *Queue) Pop() (interface{}, bool) {
	/*
	(q *Queue)不是参数列表，函数参数列表就是
	函数名后面第一个括号的，如果有第二个括号是返回值列表，
	函数名前面的(q *Queue)指的是结构体的方法
	结构体实例化后才能使用。
	*/
	q.Lock()
	defer q.Unlock()
	if len(q.data) > 0 {
		o := q.data[0]
		q.data = q.data[1:]
		return o, true

	}
	return nil, false

}

func main() {
	q := &Queue{}
	q.Push(333)
	q.Push(222)
	q.Push(111)
	q.Push(nil)

	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())

}
/*
package main

import (
	"fmt"

	"sync"
)

func main() {
	p := &Person{}
	p.Man(1)
	p.Man(2)
	p.Man(3)
	p.Man(nil)

	fmt.Println(p.Girl())
	fmt.Println(p.Girl())
	fmt.Println(p.Girl())
	fmt.Println(p.Girl())

}

type Person struct {
	sync.Mutex
	tall   float64
	weight int
	bmi    []interface{}
}

func (p *Person) Man(bmi interface{}) {
	p.Lock()
	defer p.Unlock()
	p.bmi = append(p.bmi, bmi)
}

func (p *Person) Girl() (interface{}, bool) {
	p.Lock()
	defer p.Unlock()
	if len(p.bmi) > 1 {
		o := p.bmi[1]     //切片索引下标为1,的值
		p.bmi = p.bmi[2:] //方括号下标索引
		//[2:]切片从索引2(包含)到默认上线为len(s)
		return o, true
	}
	return nil, false

}

*/