## 概述：

後端專案主要用於提供佈告欄系統，使用者可以透過 API 呼叫新增、修改、刪除佈告欄，並透過 API 取得目前所有佈告欄的列表。

後端架構使用 Gin 作為 HTTP 伺服器，並使用 Swagger 來產生 API 文件。
## 啟動方式

git pull repostory

go mod init app

go mod tidy


API 文件在產生時引入包的版本有誤:
需要引入穩定版本，添加這段代碼到go.sum內 
require(
github.com/swaggo/gin-swagger v1.4.3
	github.com/swaggo/swag v1.8.1)
    
go mod tidy

swag init
將文件
docs/docs.go init() 改成大寫 Init()
go run main.go

swagger文件路徑: http://localhost:3000/swagger/index.html#/
前端靜態文件: http://localhost:3000/

  
  API 使用文件：

      

## 新增佈告欄

URL 網址 : /announcements
API 呼叫方式 : POST  Json  Body
| Body 參數 |  格式  |  必填  | 說明  |
|----------|----------|----------|----------|
| Title    | string   | Required |佈告欄名稱 |
| Content  | string   | Required |佈告欄內容 |

回傳格式 :

| Body 參數 |  格式     | 說明     |
|---------- |----------|----------|
|Bulletin[] |struct[]  |----------|
| id        | int      |佈告欄編號 |
| title     | string   |佈告欄名稱 |
| content   | string   |佈告欄內容 |

## 刪除佈告欄

URL 網址 : /announcements/:id
API 呼叫方式 : DELETE path
| Body 參數 |  格式  |  必填  | 說明  |
|----------|----------|----------|----------|
| ID       | int      | Required |佈告欄編號 |
"announcements/123"或"announcements/abc"。路徑指向一組特定ID的佈告欄料，用於在"announcements"資源中刪除資料。


回傳格式 :
| Body 參數 |  格式     | 說明     |
|---------- |----------|----------|
| "Success"   | string   |成功     |

## 修改佈告欄
URL 網址 : /announcements/:id
API 呼叫方式 : PUT  Json  Body
| Body 參數 |  格式  |  必填  | 說明  |
|----------|----------|----------|----------|
| ID       | int      | Required |佈告欄編號 |
| Title    | string   | Required |佈告欄名稱 |
| Content  | string   | Required |佈告欄內容 |


回傳格式 :
| Body 參數 |  格式     | 說明     |
|---------- |----------|----------|
| Bulletin[] |struct[]  |----------|
| id        | int      |佈告欄編號 |
| title     | string   |佈告欄名稱 |
| content   | string   |佈告欄內容 |

## 取得所有佈告欄
URL 網址 : /announcements
API 呼叫方式 : GET 

回傳格式 :
| Body 參數 |  格式     | 說明     |
|---------- |----------|----------|
|Bulletin[] |struct[]  |----------|
| id        | int      |佈告欄編號 |
| title     | string   |佈告欄名稱 |
| content   | string   |佈告欄內容 |

