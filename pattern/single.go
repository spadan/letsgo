// 单例模式
package pattern

import (
	"fmt"
	"sync"
)

type single struct {
}

var instance *single
var lock sync.Mutex

// double check
func getInstance() *single {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			fmt.Println("Creating single instance now.")
			instance = &single{}
		} else {
			fmt.Println("Single instance already created")
		}
	} else {
		fmt.Println("Single instance already created")
	}
	return instance
}

var once sync.Once

func getInstance2() *single {
	once.Do(func() {
		if instance == nil {
			instance = &single{}
		}
	})
	return instance
}
