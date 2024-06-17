//go:build windows

package playwright

import "syscall"

var defaultSysProcAttr = &syscall.SysProcAttr{CreationFlags: syscall.CREATE_NEW_PROCESS_GROUP, HideWindow: true}

// for WritableStream.Copy
const defaultCopyBufSize = 64 * 1024
