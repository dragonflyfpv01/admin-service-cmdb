package errors

import "errors"

var (
	UserConflict = errors.New("Người dùng đã tồn tại")
	SignUpFail   = errors.New("Đăng ký thất bại")
	UserNotFound = errors.New("Tài khoản không tồn tại")
)
