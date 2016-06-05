// +build l4ka

package runtime

import (
	"unsafe"
)

func sysAlloc(n uintptr, sysStat *uint64) unsafe.Pointer {
	return nil
}

func sysFree(v unsafe.Pointer, n uintptr, sysStat *uint64) {
}


func sysReserve(v unsafe.Pointer, n uintptr, reserved *bool) unsafe.Pointer {
	return nil
}

func sysMap(v unsafe.Pointer, n uintptr, reserved bool, sysStat *uint64) {}

func sysFault(v unsafe.Pointer, n uintptr) {}

func sysUnused(v unsafe.Pointer, n uintptr) {}

func sysUsed(v unsafe.Pointer, n uintptr) {}

