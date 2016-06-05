// +build l4ka
package runtime;

import (
	"unsafe"
)

const (
	_NSIG = 1
)
type mOS struct {}

func sigpanic() {}

func getRandomData(r []byte) {}

func gogetenv(key string) string {
	return ""
}


func semacreate(mp *m) {}

func osyield() {}

func semasleep(ns int64) int32 {
	return 0
}

func semawakeup(mp *m) {}

func signame(sig uint32) string {
	return ""
}

func crash() {}

func goenvs(){}

func mpreinit(mp *m) {}

func msigsave(mp *m) {}

func msigrestore(sigmask sigset) {}

func sigblock() {}

func minit() {}

func unminit() {}

func initsig(preinit bool) {}

func newosproc(mp *m, stk unsafe.Pointer) {}

func resetcpuprofiler(hz int32) {
}

func sigenable(sig uint32) {}

func sigdisable(sig uint32) {}

func sigignore(sig uint32) {}


func read(fd int32, p unsafe.Pointer, n int32) int32 {
	return 0
}
func closefd(fd int32) int32 {
	return 0
}

func exit(code int32) {}
func nanotime() int64 {
	return 0
}
func usleep(usec uint32) {

}

func munmap(addr unsafe.Pointer, n uintptr) {}

func write(fd uintptr, p unsafe.Pointer, n int32) int32 {
	return -1
}

func open(name *byte, mode, perm int32) int32 {
	return -1
}

func madvise(addr unsafe.Pointer, n uintptr, flags int32){}
