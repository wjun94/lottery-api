### 数据库

** user表 **
level 等级划分
1:管理员
2:业务员

## GET请求
```go
type SelectLotList struct {
	Name   string `form:"name"` // 需要form
}

var selectLotList model.SelectLotList
c.ShouldBind(&selectLotList)
```