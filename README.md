# call api

=====Đăng nhập====================================================================================================
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

4. API Lấy Danh Sách Users đã được tạo
Method: GET
URL: /admin/users
Authentication: JWT token required
Authorization: Chỉ admin role mới có quyền truy cập

Reponse:
{
    "status_code": 200,
    "message": "Get users successfully",
    "data": [
        {
            "id": 2,
            "username": "hoangthai",
            "user_id": "e17138e6-980a-4382-a8fc-875d1a4b2d46",
            "role": "USER"
        }
    ]
}


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

========infra-components=================================================================================

Bước 1: Đăng nhập để lấy token (với user có role admin)
curl -X POST http://localhost:3000/admin/login \
  -H "Content-Type: application/json" \
  -d '{"email":"admin@example.com","password":"admin123"}'

Bước 2: Sử dụng token để gọi API lấy infra components
curl -X GET http://localhost:3000/admin/infra-components \
  -H "Authorization: Bearer <JWT_TOKEN_FROM_STEP_1>" \
  -H "Content-Type: application/json"

# 1. 	api.Echo.GET("/admin/infra-components")
Lấy tất cả với phân trang: ?page=1&limit=10

curl -X GET http://localhost:3000/admin/infra-components \
  -H "Authorization: Bearer <JWT_TOKEN_FROM_STEP_1>" \
  -H "Content-Type: application/json"


# 2.	api.Echo.GET("/admin/infra-components/all")
Lấy tất cả không phân trang. Trả về toàn bộ dữ liệu

curl -X GET http://localhost:3000/admin/infra-components/all \
  -H "Authorization: Bearer <JWT_TOKEN_FROM_STEP_1>" \
  -H "Content-Type: application/json"

# 3.	api.Echo.GET("/admin/infra-components/pending")
Chỉ lấy các bản ghi có status = "Đang chờ". Không phân trang, trả về tất cả bản ghi phù hợp

curl -X GET http://localhost:3000/admin/infra-components/pending \
  -H "Authorization: Bearer <JWT_TOKEN_FROM_STEP_1>" \
  -H "Content-Type: application/json"

# 4.    api.Echo.PUT("/admin/infra-components/status")
User phải có role = "admin". Cả ID và hostname phải khớp với bản ghi trong database

curl -X PUT "http://localhost:3000/admin/infra-components/status" \
  -H "Authorization: Bearer <JWT_TOKEN>" \
  -H "Content-Type: application/json" \
  -d '{
    "id": 15,
    "hostname": "server15",
    "new_status": "Hoàn thành"
  }'

# 5.    api.Echo.PUT("/admin/infra-components")
Cập Nhật Thông Tin Infra Component. User phải có role = "admin". ID phải tồn tại trong database
Không được sửa id, status và create_by
created_at sẽ được cập nhật thành thời gian hiện tại (thời gian sửa)

curl -X PUT "http://localhost:3000/admin/infra-components" \
  -H "Authorization: Bearer <JWT_TOKEN>" \
  -H "Content-Type: application/json" \
  -d '{
    "id": 15,
    "hostname": "server15-updated",
    "dns": "server15-updated.example.com",
    "description": "Updated web server description",
    "public_internet": "Yes",
    "class": "Production",
    "ipaddress": "192.168.1.15",
    "subnet": "192.168.1.0/24",
    "site": "HN",
    "it_component_type": "Web Server",
    "request_type": "Update",
    "appid": "APP015",
    "vlan": "VLAN100",
    "app_name": "Updated Web Application",
    "app_owner": "Updated IT Team",
    "level": "L1",
    "ci_owners": "Updated Admin",
    "im_cm": "CM015"
  }'


# 6. api.Echo.POST("/admin/infra-components")
API Tạo Mới Infra-Components
Method: POST
URL: /admin/infra-components
Authentication: JWT token required

body:
{
    "hostname": "server01.example.com",
    "dns": "server01.dns.example.com", 
    "description": "Web server for production",
    "public_internet": "Yes",
    "class": "Production",
    "ipaddress": "192.168.1.100",
    "subnet": "192.168.1.0/24",
    "site": "Data Center A",
    "it_component_type": "Server",
    "request_type": "New",
    "appid": "APP001",
    "vlan": "VLAN100",
    "app_name": "Web Application",
    "app_owner": "IT Team",
    "level": "Critical",
    "ci_owners": "Infrastructure Team",
    "im_cm": "ServiceNow"
}

respon: 
{
    "status_code": 201,
    "message": "Create infra component successfully",
    "data": {
        "id": 15,
        "hostname": "server01.example.com",
        "dns": "server01.dns.example.com",
        "description": "Web server for production",
        "public_internet": "Yes",
        "class": "Production",
        "ipaddress": "192.168.1.100",
        "subnet": "192.168.1.0/24", 
        "site": "Data Center A",
        "it_component_type": "Server",
        "request_type": "New",
        "appid": "APP001",
        "vlan": "VLAN100",
        "app_name": "Web Application",
        "app_owner": "IT Team",
        "level": "Critical",
        "ci_owners": "Infrastructure Team",
        "im_cm": "ServiceNow",
        "status": "Đang chờ",
        "created_at": "2025-09-19 15:30:45",
        "create_by": "admin@example.com"
    }
}


Error Responses:
401 Unauthorized: JWT token không hợp lệ
400 Bad Request: Validation errors
500 Internal Server Error: Database errors hoặc user không tồn tại
