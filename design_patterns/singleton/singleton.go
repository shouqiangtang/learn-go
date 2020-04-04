package singleton

import "sync"

// Single : 单例结构体，该结构体只能被初始化一次
type Single struct {
	Name string
}

// single : 包全局变量，该对象分配在堆里面，在整个进程生命周期内有效
// 单例模式保证single只能被赋值一次
//
// 可以参考net/http包，var DefaultClient = &Client{}，在包中定一个默认的结构体变量；同时提供NewClient方法。
// 包使用者可以灵活选择使用默认的DefaultClient还是使用NewClient新建变量。
var single *Single

// 全局锁
var mutex sync.Mutex

// GetInstance1 : 懒汉模式
// 懒汉模式在高并发时存在数据安全问题，single变量可能被赋值多次
func GetInstance1(name string) *Single {
	if single == nil {
		single = &Single{
			Name: name,
		}
	}
	return single
}

// GetInstance2 : 饿汉模式
// 饿汉模式有个缺点每次请求该方法时都得加锁，降低性能
// 数据争用(data race) 和竞态条件(race condition)
func GetInstance2(name string) *Single {
	// http://ifeve.com/race-conditions-and-critical-sections/
	mutex.Lock()
	defer mutex.Unlock()
	if single == nil {
		single = &Single{
			Name: name,
		}
	}
	return single
}

// GetInstance3 : 双重检测模式
// 只有在未初始化时，才会进行加锁/解锁
// TODO 此种方法使用go run -race检测是存在data races的
func GetInstance3(name string) *Single {
	if single == nil {
		mutex.Lock()
		defer mutex.Unlock()
		if single == nil {
			single = &Single{
				Name: name,
			}
		}
	}
	return single
}

// sync.Once的Do方法可以实现在程序运行过程中只运行一次其中的回调
var once sync.Once

// GetInstance4 : 获取实例
func GetInstance4(name string) *Single {
	once.Do(func() {
		single = &Single{
			Name: name,
		}
	})
	return single
}
