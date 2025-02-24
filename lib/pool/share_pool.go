package pool

//
//import (
//	"context"
//	"github.com/panjf2000/ants/v2"
//	"runtime"
//	"sync"
//	"time"
//)
//
//var (
//	// _pool协程池
//	_pool *sharePool
//	// 后台任务池，用于rainbow/loopload/singleflight-cache等
//	_taskPool, _ = newSharePool(3000)
//	// PanicBufLen panic调用栈日志buffer大小，默认2048
//	PanicBufLen = 2048
//
//	// GetPool TODO
//	GetPool = func() Pool {
//		// 如果不进行判断，会报空指针异常，排查问题更麻烦
//		if _pool == nil {
//			panic("未进行share.InitPool初始化操作!!!!, 另外需要注意code.oa和woa")
//		}
//		return _pool
//	}
//
//	// GetTaskPool 获取taskPool
//	GetTaskPool = func() Pool {
//		return _taskPool
//	}
//)
//
//// CustomLogger 用户日志
//type CustomLogger struct{}
//
//// Printf 打印自定义日志
//func (c CustomLogger) Printf(format string, args ...interface{}) {
//	log.Errorf(format, args)
//}
//
//// InitPool 初始化协程池, 需要在main中调用
//func InitPool(size int, opts ...ants.Option) error {
//	var err error
//
//	// 初始化
//	antsPool, err := ants.NewPool(size, opts...)
//	if err != nil {
//		return err
//	}
//
//	_pool = &sharePool{
//		antsPool,
//	}
//
//	return err
//}
//
//// HandlerFunc ....
//type HandlerFunc func() error
//
//type ContextHandlerFunc func(ctx context.Context) error
//
//// Pool 协程池
////
////go:generate mockgen -destination pool_mock.go -package share -source pool.go
//type Pool interface {
//	// Submit 提交任务，提交失败返回error
//	Submit(func()) error
//	// SubmitAndWait 批量提交任务，并等待所有任务结束。
//	// 如果在提交或者执行过程中发生错误, 则记下第一个error，并返回
//	// 能处理过程中的panic
//	SubmitAndWait(...HandlerFunc) error
//	// TuneCapacity 调整容量
//	TuneCapacity(size int)
//	// ContextSubmit 对比 Submit 会自动帮你clone context
//	ContextSubmit(context.Context, func(context.Context)) error
//	// ContextSubmitAndWait 对比 SubmitAndWait 会自动帮你clone context
//	ContextSubmitAndWait(context.Context, ...ContextHandlerFunc) error
//}
//
//type sharePool struct {
//	*ants.Pool
//}
//
//func (s *sharePool) ContextSubmit(ctx context.Context, f func(context.Context)) error {
//	cloneContext := trpc.CloneContext(ctx)
//	return s.Pool.Submit(func() {
//		f(cloneContext)
//	})
//}
//
//func (s *sharePool) ContextSubmitAndWait(ctx context.Context, handlerFunc ...ContextHandlerFunc) error {
//	var funcList = make([]HandlerFunc, 0, len(handlerFunc))
//
//	for _, fc := range handlerFunc {
//		fc := fc
//		// 这里不需要clone ctx，因为 Wait，应该是需要等待所有执行完成的。
//		funcList = append(funcList, func() error {
//			return fc(ctx)
//		})
//	}
//
//	return s.SubmitAndWait(funcList...)
//}
//
//// 创建sharePool
//func newSharePool(size int, opts ...ants.Option) (*sharePool, error) {
//	p, err := ants.NewPool(size, opts...)
//	if err != nil {
//		return nil, err
//	}
//
//	return &sharePool{
//		p,
//	}, nil
//}
//
//// Submit 无返回任务，直接包装
//func (s *sharePool) Submit(f func()) error {
//	return s.Pool.Submit(f)
//}
//
//// TuneCapacity 调整容量
//func (s *sharePool) TuneCapacity(size int) {
//	if s.Pool != nil {
//		s.Pool.Tune(size)
//	}
//}
//
//// SubmitAndWait , 批量提交任务 参考trpc-go中的实现, 只要有一个error出现，则返回err，且为第一个
//func (s *sharePool) SubmitAndWait(handlers ...HandlerFunc) error {
//	once := &sync.Once{}
//	var err error
//	// setErr 设置错误, 只执行一次
//	setErr := func(e error) {
//		once.Do(func() {
//			err = e
//		})
//	}
//
//	wg := &sync.WaitGroup{}
//	for _, f := range handlers {
//		wg.Add(1)
//		tempFunc := f // 闭包问题
//		if submitErr := s.Pool.Submit(func() {
//			defer func() {
//				// 发生panic,打印错误，并记录
//				if e := recover(); e != nil {
//					buf := make([]byte, PanicBufLen)
//					buf = buf[:runtime.Stack(buf, false)]
//					log.Errorf("[PANIC]%v\n%s\n", e, buf)
//					metric.ServerPanicTotal.WithLabelValues("trpc").Inc()
//					setErr(errs.New(errs.RetServerSystemErr, "panic found when call handlers"))
//				}
//
//				// 完成
//				wg.Done()
//			}()
//
//			// 执行函数
//			if e := tempFunc(); e != nil {
//				// 失败，记录原因
//				setErr(e)
//			}
//		}); submitErr != nil {
//			setErr(submitErr)
//			wg.Done() // 提交失败
//		}
//	}
//
//	wg.Wait()
//	return err
//}
//
//// PanicHandler4Pool 协程池panic上报
//func PanicHandler4Pool(interface{}) {
//	buf := make([]byte, PanicBufLen)
//	buf = buf[:runtime.Stack(buf, false)]
//	log.Errorf("[PANIC]\n%s\n", buf)
//	metric.ServerPanicTotal.WithLabelValues("trpc").Inc()
//}
//
//// RetryWrap 重试n次，间隔为interval * 2 ^ n
//func RetryWrap(handlers []HandlerFunc, retry int, interval time.Duration) []HandlerFunc {
//	res := make([]HandlerFunc, 0, len(handlers))
//
//	for _, temp := range handlers {
//		h := temp // 闭包
//		res = append(res, func() error {
//			var err error
//			for i := 0; i < retry; i++ {
//				err = h() // 执行
//				if err == nil {
//					return nil // 成功则返回
//				}
//
//				dur := interval * (1 << i) // 指数退避
//				time.Sleep(dur)
//			}
//
//			return err // 最后一次尝试的Err
//		})
//	}
//
//	return res
//}
