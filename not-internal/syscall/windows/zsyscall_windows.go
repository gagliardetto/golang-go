// Code generated by 'go generate'; DO NOT EDIT.

package windows

import (
	"github.com/gagliardetto/golang-go/not-internal/syscall/windows/sysdll"
	"syscall"
	"unsafe"
)

var _ unsafe.Pointer

// Do the interface allocations only once for common
// Errno values.
const (
	errnoERROR_IO_PENDING = 997
)

var (
	errERROR_IO_PENDING error = syscall.Errno(errnoERROR_IO_PENDING)
)

// errnoErr returns common boxed Errno values, to prevent
// allocations at runtime.
func errnoErr(e syscall.Errno) error {
	switch e {
	case 0:
		return nil
	case errnoERROR_IO_PENDING:
		return errERROR_IO_PENDING
	}
	// TODO: add more here, after collecting data on the common
	// error values see on Windows. (perhaps when running
	// all.bat?)
	return e
}

var (
	modiphlpapi = syscall.NewLazyDLL(sysdll.Add("iphlpapi.dll"))
	modkernel32 = syscall.NewLazyDLL(sysdll.Add("kernel32.dll"))
	modws2_32   = syscall.NewLazyDLL(sysdll.Add("ws2_32.dll"))
	modnetapi32 = syscall.NewLazyDLL(sysdll.Add("netapi32.dll"))
	modadvapi32 = syscall.NewLazyDLL(sysdll.Add("advapi32.dll"))
	moduserenv  = syscall.NewLazyDLL(sysdll.Add("userenv.dll"))
	modpsapi    = syscall.NewLazyDLL(sysdll.Add("psapi.dll"))

	procGetAdaptersAddresses         = modiphlpapi.NewProc("GetAdaptersAddresses")
	procGetComputerNameExW           = modkernel32.NewProc("GetComputerNameExW")
	procMoveFileExW                  = modkernel32.NewProc("MoveFileExW")
	procGetModuleFileNameW           = modkernel32.NewProc("GetModuleFileNameW")
	procWSASocketW                   = modws2_32.NewProc("WSASocketW")
	procLockFileEx                   = modkernel32.NewProc("LockFileEx")
	procUnlockFileEx                 = modkernel32.NewProc("UnlockFileEx")
	procGetACP                       = modkernel32.NewProc("GetACP")
	procGetConsoleCP                 = modkernel32.NewProc("GetConsoleCP")
	procMultiByteToWideChar          = modkernel32.NewProc("MultiByteToWideChar")
	procGetCurrentThread             = modkernel32.NewProc("GetCurrentThread")
	procNetShareAdd                  = modnetapi32.NewProc("NetShareAdd")
	procNetShareDel                  = modnetapi32.NewProc("NetShareDel")
	procGetFinalPathNameByHandleW    = modkernel32.NewProc("GetFinalPathNameByHandleW")
	procCreateEnvironmentBlock       = moduserenv.NewProc("CreateEnvironmentBlock")
	procDestroyEnvironmentBlock      = moduserenv.NewProc("DestroyEnvironmentBlock")
	procImpersonateSelf              = modadvapi32.NewProc("ImpersonateSelf")
	procRevertToSelf                 = modadvapi32.NewProc("RevertToSelf")
	procOpenThreadToken              = modadvapi32.NewProc("OpenThreadToken")
	procLookupPrivilegeValueW        = modadvapi32.NewProc("LookupPrivilegeValueW")
	procAdjustTokenPrivileges        = modadvapi32.NewProc("AdjustTokenPrivileges")
	procDuplicateTokenEx             = modadvapi32.NewProc("DuplicateTokenEx")
	procSetTokenInformation          = modadvapi32.NewProc("SetTokenInformation")
	procGetProfilesDirectoryW        = moduserenv.NewProc("GetProfilesDirectoryW")
	procNetUserGetLocalGroups        = modnetapi32.NewProc("NetUserGetLocalGroups")
	procGetProcessMemoryInfo         = modpsapi.NewProc("GetProcessMemoryInfo")
	procGetFileInformationByHandleEx = modkernel32.NewProc("GetFileInformationByHandleEx")
)

func GetAdaptersAddresses(family uint32, flags uint32, reserved uintptr, adapterAddresses *IpAdapterAddresses, sizePointer *uint32) (errcode error) {
	r0, _, _ := syscall.Syscall6(procGetAdaptersAddresses.Addr(), 5, uintptr(family), uintptr(flags), uintptr(reserved), uintptr(unsafe.Pointer(adapterAddresses)), uintptr(unsafe.Pointer(sizePointer)), 0)
	if r0 != 0 {
		errcode = syscall.Errno(r0)
	}
	return
}

func GetComputerNameEx(nameformat uint32, buf *uint16, n *uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procGetComputerNameExW.Addr(), 3, uintptr(nameformat), uintptr(unsafe.Pointer(buf)), uintptr(unsafe.Pointer(n)))
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func MoveFileEx(from *uint16, to *uint16, flags uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procMoveFileExW.Addr(), 3, uintptr(unsafe.Pointer(from)), uintptr(unsafe.Pointer(to)), uintptr(flags))
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func GetModuleFileName(module syscall.Handle, fn *uint16, len uint32) (n uint32, err error) {
	r0, _, e1 := syscall.Syscall(procGetModuleFileNameW.Addr(), 3, uintptr(module), uintptr(unsafe.Pointer(fn)), uintptr(len))
	n = uint32(r0)
	if n == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func WSASocket(af int32, typ int32, protocol int32, protinfo *syscall.WSAProtocolInfo, group uint32, flags uint32) (handle syscall.Handle, err error) {
	r0, _, e1 := syscall.Syscall6(procWSASocketW.Addr(), 6, uintptr(af), uintptr(typ), uintptr(protocol), uintptr(unsafe.Pointer(protinfo)), uintptr(group), uintptr(flags))
	handle = syscall.Handle(r0)
	if handle == syscall.InvalidHandle {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func LockFileEx(file syscall.Handle, flags uint32, reserved uint32, bytesLow uint32, bytesHigh uint32, overlapped *syscall.Overlapped) (err error) {
	r1, _, e1 := syscall.Syscall6(procLockFileEx.Addr(), 6, uintptr(file), uintptr(flags), uintptr(reserved), uintptr(bytesLow), uintptr(bytesHigh), uintptr(unsafe.Pointer(overlapped)))
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func UnlockFileEx(file syscall.Handle, reserved uint32, bytesLow uint32, bytesHigh uint32, overlapped *syscall.Overlapped) (err error) {
	r1, _, e1 := syscall.Syscall6(procUnlockFileEx.Addr(), 5, uintptr(file), uintptr(reserved), uintptr(bytesLow), uintptr(bytesHigh), uintptr(unsafe.Pointer(overlapped)), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func GetACP() (acp uint32) {
	r0, _, _ := syscall.Syscall(procGetACP.Addr(), 0, 0, 0, 0)
	acp = uint32(r0)
	return
}

func GetConsoleCP() (ccp uint32) {
	r0, _, _ := syscall.Syscall(procGetConsoleCP.Addr(), 0, 0, 0, 0)
	ccp = uint32(r0)
	return
}

func MultiByteToWideChar(codePage uint32, dwFlags uint32, str *byte, nstr int32, wchar *uint16, nwchar int32) (nwrite int32, err error) {
	r0, _, e1 := syscall.Syscall6(procMultiByteToWideChar.Addr(), 6, uintptr(codePage), uintptr(dwFlags), uintptr(unsafe.Pointer(str)), uintptr(nstr), uintptr(unsafe.Pointer(wchar)), uintptr(nwchar))
	nwrite = int32(r0)
	if nwrite == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func GetCurrentThread() (pseudoHandle syscall.Handle, err error) {
	r0, _, e1 := syscall.Syscall(procGetCurrentThread.Addr(), 0, 0, 0, 0)
	pseudoHandle = syscall.Handle(r0)
	if pseudoHandle == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func NetShareAdd(serverName *uint16, level uint32, buf *byte, parmErr *uint16) (neterr error) {
	r0, _, _ := syscall.Syscall6(procNetShareAdd.Addr(), 4, uintptr(unsafe.Pointer(serverName)), uintptr(level), uintptr(unsafe.Pointer(buf)), uintptr(unsafe.Pointer(parmErr)), 0, 0)
	if r0 != 0 {
		neterr = syscall.Errno(r0)
	}
	return
}

func NetShareDel(serverName *uint16, netName *uint16, reserved uint32) (neterr error) {
	r0, _, _ := syscall.Syscall(procNetShareDel.Addr(), 3, uintptr(unsafe.Pointer(serverName)), uintptr(unsafe.Pointer(netName)), uintptr(reserved))
	if r0 != 0 {
		neterr = syscall.Errno(r0)
	}
	return
}

func GetFinalPathNameByHandle(file syscall.Handle, filePath *uint16, filePathSize uint32, flags uint32) (n uint32, err error) {
	r0, _, e1 := syscall.Syscall6(procGetFinalPathNameByHandleW.Addr(), 4, uintptr(file), uintptr(unsafe.Pointer(filePath)), uintptr(filePathSize), uintptr(flags), 0, 0)
	n = uint32(r0)
	if n == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func CreateEnvironmentBlock(block **uint16, token syscall.Token, inheritExisting bool) (err error) {
	var _p0 uint32
	if inheritExisting {
		_p0 = 1
	} else {
		_p0 = 0
	}
	r1, _, e1 := syscall.Syscall(procCreateEnvironmentBlock.Addr(), 3, uintptr(unsafe.Pointer(block)), uintptr(token), uintptr(_p0))
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func DestroyEnvironmentBlock(block *uint16) (err error) {
	r1, _, e1 := syscall.Syscall(procDestroyEnvironmentBlock.Addr(), 1, uintptr(unsafe.Pointer(block)), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func ImpersonateSelf(impersonationlevel uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procImpersonateSelf.Addr(), 1, uintptr(impersonationlevel), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func RevertToSelf() (err error) {
	r1, _, e1 := syscall.Syscall(procRevertToSelf.Addr(), 0, 0, 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func OpenThreadToken(h syscall.Handle, access uint32, openasself bool, token *syscall.Token) (err error) {
	var _p0 uint32
	if openasself {
		_p0 = 1
	} else {
		_p0 = 0
	}
	r1, _, e1 := syscall.Syscall6(procOpenThreadToken.Addr(), 4, uintptr(h), uintptr(access), uintptr(_p0), uintptr(unsafe.Pointer(token)), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func LookupPrivilegeValue(systemname *uint16, name *uint16, luid *LUID) (err error) {
	r1, _, e1 := syscall.Syscall(procLookupPrivilegeValueW.Addr(), 3, uintptr(unsafe.Pointer(systemname)), uintptr(unsafe.Pointer(name)), uintptr(unsafe.Pointer(luid)))
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func adjustTokenPrivileges(token syscall.Token, disableAllPrivileges bool, newstate *TOKEN_PRIVILEGES, buflen uint32, prevstate *TOKEN_PRIVILEGES, returnlen *uint32) (ret uint32, err error) {
	var _p0 uint32
	if disableAllPrivileges {
		_p0 = 1
	} else {
		_p0 = 0
	}
	r0, _, e1 := syscall.Syscall6(procAdjustTokenPrivileges.Addr(), 6, uintptr(token), uintptr(_p0), uintptr(unsafe.Pointer(newstate)), uintptr(buflen), uintptr(unsafe.Pointer(prevstate)), uintptr(unsafe.Pointer(returnlen)))
	ret = uint32(r0)
	if true {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func DuplicateTokenEx(hExistingToken syscall.Token, dwDesiredAccess uint32, lpTokenAttributes *syscall.SecurityAttributes, impersonationLevel uint32, tokenType TokenType, phNewToken *syscall.Token) (err error) {
	r1, _, e1 := syscall.Syscall6(procDuplicateTokenEx.Addr(), 6, uintptr(hExistingToken), uintptr(dwDesiredAccess), uintptr(unsafe.Pointer(lpTokenAttributes)), uintptr(impersonationLevel), uintptr(tokenType), uintptr(unsafe.Pointer(phNewToken)))
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func SetTokenInformation(tokenHandle syscall.Token, tokenInformationClass uint32, tokenInformation uintptr, tokenInformationLength uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procSetTokenInformation.Addr(), 4, uintptr(tokenHandle), uintptr(tokenInformationClass), uintptr(tokenInformation), uintptr(tokenInformationLength), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func GetProfilesDirectory(dir *uint16, dirLen *uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procGetProfilesDirectoryW.Addr(), 2, uintptr(unsafe.Pointer(dir)), uintptr(unsafe.Pointer(dirLen)), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func NetUserGetLocalGroups(serverName *uint16, userName *uint16, level uint32, flags uint32, buf **byte, prefMaxLen uint32, entriesRead *uint32, totalEntries *uint32) (neterr error) {
	r0, _, _ := syscall.Syscall9(procNetUserGetLocalGroups.Addr(), 8, uintptr(unsafe.Pointer(serverName)), uintptr(unsafe.Pointer(userName)), uintptr(level), uintptr(flags), uintptr(unsafe.Pointer(buf)), uintptr(prefMaxLen), uintptr(unsafe.Pointer(entriesRead)), uintptr(unsafe.Pointer(totalEntries)), 0)
	if r0 != 0 {
		neterr = syscall.Errno(r0)
	}
	return
}

func GetProcessMemoryInfo(handle syscall.Handle, memCounters *PROCESS_MEMORY_COUNTERS, cb uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procGetProcessMemoryInfo.Addr(), 3, uintptr(handle), uintptr(unsafe.Pointer(memCounters)), uintptr(cb))
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func GetFileInformationByHandleEx(handle syscall.Handle, class uint32, info *byte, bufsize uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procGetFileInformationByHandleEx.Addr(), 4, uintptr(handle), uintptr(class), uintptr(unsafe.Pointer(info)), uintptr(bufsize), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}
