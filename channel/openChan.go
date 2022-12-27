package channel

import (
	"github.com/pkg/errors"
	"reflect"
	"sync"
	"unsafe"
)

// 管道关闭后状态 1
// 理论上修改关闭后的状态即可
// 关闭 closed == 1.开启 closed = 0

func lock(l *sync.Mutex) {}

func unlock(l *sync.Mutex) {}

func Open(c interface{}) error {
	v := reflect.ValueOf(c)
	if v.Type().Kind() != reflect.Chan {
		return errors.New("type must be channel")
	}
	i := (*[2]uintptr)(unsafe.Pointer(&c)) //定位c所在数据空间，这里的c是个指针所以要进行一步取值
	var closedOffset, lockOffset uintptr = 28, 88
	closed := (*uint32)(unsafe.Pointer(i[1] + closedOffset)) //指向closed的地址
	if *closed == 1 {
		lockPtr := (*sync.Mutex)(unsafe.Pointer(i[1] + lockOffset)) //指向lock地址
		lock(lockPtr)                                               //上锁
		if *closed == 1 {
			*closed = 0 //直接修改值
		}
		unlock(lockPtr) //解锁
	}
	return nil
}
