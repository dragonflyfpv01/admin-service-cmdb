# call api

=====Đăng nhập=====
1. Đăng ký tài khoản (Sign Up)
Phương thức: POST
URL: /admin/sign-up
Body (JSON):
{
  "email": "your_email@example.com",
  "password": "your_password",
  "full_name": "Your Name"
}

2. Đăng nhập (Sign In)
Phương thức: POST
URL: /admin/sign-in
Body (JSON):
{
  "email": "your_email@example.com",
  "password": "your_password"
}
Kết quả trả về: Nếu thành công sẽ nhận được token JWT trong trường token của user.

3. Lấy thông tin profile (Yêu cầu xác thực JWT)
Phương thức: GET
URL: /admin/profile
Header: Authorization: Bearer <token>. (Token lấy từ kết quả đăng nhập)

Ví dụ gọi bằng curl:
# 1 Đăng ký
curl -X POST http://localhost:3000/admin/sign-up \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"123456","full_name":"Test User"}'

# 2 Đăng nhập
curl -X POST http://localhost:3000/admin/sign-in \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"123456"}'

# 3 Lấy profile (thay <token> bằng JWT nhận được)
curl -X GET http://localhost:3000/admin/profile \
  -H "Authorization: Bearer <token>"
