package custom_pool

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

type Conn struct {
	alive         bool
	isExipre      bool
	maxConn       int                     // 最大连接数
	maxIdle       int                     // 最大可用连接数
	freeConn      int                     // 线程池空闲连接数
	connPool      []int                   // 连接池
	openCount     int                     // 已经打开的连接数
	waitConn      map[int]chan Permission // 排队等待的连接队列
	waitCount     int                     // 等待个数
	lock          sync.Mutex              // 锁
	nextConnIndex int
	freeConns     map[int]Permission // 连接池的连接
	expiredCh     chan string
}

type Config struct {
	MaxConn int
	MaxIdle int
}

type Permission struct {
	NextConnIndex
	Content     string
	CreatedAt   time.Time
	MaxLifeTime time.Duration
}

type NextConnIndex struct {
	Index int
}

var nowFunc = time.Now

// Ping Ping是一个获取连接并释放连接的过程
func Ping() {

}

// TODO 1、使用channel归还连接 2、连接添加超时时间，如果超过超时时间没有拿到连接则退出
func Prepare(ctx context.Context, config *Config) (conn *Conn) {
	go func() {
		for {
			conn.expiredCh = make(chan string, len(conn.freeConns))
			for _, value := range conn.freeConns {
				if value.CreatedAt.Add(value.MaxLifeTime).Before(nowFunc()) {
					conn.expiredCh <- "CLOSE"
				}
			}
			//select {
			//case <-conn.waitConn:
			//	fmt.Println("receive")
			//	conn.lock.Lock()
			//	if conn.waitCount > 1 {
			//		conn.waitCount--
			//	}
			//	conn.lock.Unlock()
			//}
		}

	}()
	return &Conn{
		alive:     false,
		isExipre:  false,
		maxConn:   config.MaxConn,
		maxIdle:   config.MaxIdle,
		openCount: 0,
		freeConn:  0,
		connPool:  []int{},
		waitConn:  make(map[int]chan Permission),
		waitCount: 0,
		freeConns: make(map[int]Permission),
	}
}

// 创建连接
func (conn *Conn) New(ctx context.Context) (permission Permission, err error) {
	/**
	1、如果当前连接池已满，即freeConn=0
	2、判定openConn是否大于maxConn，如果大于，则丢弃获取加入队列进行等待
	3、如果小于，则考虑创建新连接
	注意设置超时时间以及用完连接的回收

	任务满了就加入等待队列，什么时候监听消息
	*/
	conn.lock.Lock()

	select {
	default:
	case <-ctx.Done():
		nextConnIndex := getNextConnIndex(conn)
		conn.lock.Unlock()

		return Permission{NextConnIndex: NextConnIndex{nextConnIndex},
				Content: "PASSED", CreatedAt: nowFunc(), MaxLifeTime: time.Second * 5},
			errors.New("new conn failed, context cancelled!")
	}

	//if conn.freeConn > 0 {
	//	conn.connPool = conn.connPool[1:] // 取出一个连接
	//	conn.freeConn--
	//	fmt.Println("openCount: ", conn.openCount, " freeConn: ", conn.freeConn, " connPool: ", conn.connPool)
	//	conn.lock.Unlock()
	//	return true, nil
	//}

	if conn.openCount > conn.maxConn { // 当前连接数大于上限，则加入等待队列
		//conn.lock.Lock()
		//conn.waitConn <- 1
		nextConnIndex := getNextConnIndex(conn)

		req := make(chan Permission, 1)
		conn.waitConn[nextConnIndex] = req
		//conn.waitConn = append(conn.waitConn, 1) // TODO 等待队列的处理
		conn.waitCount++
		conn.lock.Unlock()

		select {
		case <-conn.expiredCh:
			return Permission{}, errors.New("new conn failed, timeout")
		case ret, ok := <-req: // 有放回的连接, 直接拿来用
			fmt.Println("received release conn !!!")
			if !ok {
				return Permission{}, errors.New("new conn failed, no available conn release")
			}
			conn.lock.Lock()
			//delete(conn.waitConn, )
			conn.openCount++
			conn.lock.Unlock()
			return ret, nil
		}
		return Permission{}, errors.New("new conn failed")
	}

	// 新建连接
	conn.connPool = append(conn.connPool, 1)
	conn.openCount++
	conn.freeConn++
	conn.lock.Unlock()
	nextConnIndex := getNextConnIndex(conn)
	permission = Permission{NextConnIndex: NextConnIndex{nextConnIndex},
		Content: "PASSED", CreatedAt: nowFunc(), MaxLifeTime: time.Second * 5}
	conn.freeConns[nextConnIndex] = permission
	fmt.Println("openCount: ", conn.openCount, " freeConn: ", conn.freeConn, " connPool: ", conn.connPool)
	return permission, nil
}

func getNextConnIndex(conn *Conn) int {
	currentIndex := conn.nextConnIndex
	//conn.lock.Lock()
	conn.nextConnIndex = currentIndex + 1
	fmt.Println("connIndex: ", conn.nextConnIndex)
	//conn.lock.Lock()
	return conn.nextConnIndex
}

// 释放连接
func (conn *Conn) Release(ctx context.Context) (result bool, err error) {
	conn.lock.Lock()
	conn.openCount--
	conn.connPool = append(conn.connPool, 1)
	conn.freeConn++
	conn.lock.Unlock()
	fmt.Println("openCount: ", conn.openCount, " freeConn: ", conn.freeConn, " connPool: ", conn.connPool)

	if len(conn.waitConn) > 0 {
		var req chan Permission
		var reqKey int
		for reqKey, req = range conn.waitConn {
			break
		}

		permission := Permission{NextConnIndex: NextConnIndex{reqKey},
			Content: "PASSED", CreatedAt: nowFunc(), MaxLifeTime: time.Second * 5}
		req <- permission
		conn.freeConns[reqKey] = permission
		conn.lock.Lock()
		conn.waitCount--
		delete(conn.waitConn, reqKey)
		conn.lock.Unlock()
	}
	return
}
