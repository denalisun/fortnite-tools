package utilities

import (
	"syscall"
	"unsafe"
)

func GetFortnitePID() (uint32, error) {
	hSnapshot, err := syscall.CreateToolhelp32Snapshot(0x00000002, 0)
	if err != nil {
		return 0, err
	}
	defer syscall.CloseHandle(hSnapshot)

	var hFortniteHandle uint32 = 0

	if hSnapshot != 0 {
		var pe32 syscall.ProcessEntry32
		pe32.Size = uint32(unsafe.Sizeof(pe32))

		if err = syscall.Process32First(hSnapshot, &pe32); err == nil {
			for {
				if exeName := syscall.UTF16ToString(pe32.ExeFile[:]); exeName == "FortniteClient-Win64-Shipping.exe" {
					hFortniteHandle = pe32.ProcessID
				}

				if err = syscall.Process32Next(hSnapshot, &pe32); err != nil {
					break
				}
			}
		}
	}

	return hFortniteHandle, nil
}
