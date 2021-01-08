package models

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"mango-api/config"
	"mango-api/response"

	"github.com/iGoogle-ink/gopay"
	"github.com/iGoogle-ink/gopay/wechat"
)

/**
 * 文档作者: wjun94
 * 创建时间：2019年10月18日
 * 修改时间：2019年10月18日
 * 描述信息：用户信息数据库操作
 */

// OrderAddress 订单地址数据
type OrderAddress struct {
	ID       string `gorm:"primary_key" json:"id"`    // id
	OrderID  string `gorm:"not null" json:"-"`        // 订单id
	Name     string `gorm:"not null" json:"name"`     // 头像链接
	Province string `gorm:"not null" json:"province"` // 省
	City     string `gorm:"not null" json:"city"`     // 城市
	Area     string `gorm:"not null" json:"area"`     // 区
	Address  string `gorm:"not null" json:"address"`  // 详细地址
	Phone    string `gorm:"not null" json:"phone"`    // 手机
}

// OrderCommodity 订单的相关商品
type OrderCommodity struct {
	CommodityID string  `gorm:"not null" json:"id"`                              // 商品id
	OrderID     string  `gorm:"not null" json:"-"`                               // 订单id
	Name        string  `gorm:"not null" json:"name"`                            // 商品名
	Img         string  `gorm:"not null;type:longtext" json:"img"`               // 商品主图
	Unit        string  `gorm:"not null" json:"unit"`                            // 单位
	Price       float64 `gorm:"not null" gorm:"type:decimal(10,2)" json:"price"` // 价格
	Count       int     `gorm:"not null" json:"count"`                           // 数量
	Code        string  `json:"code"`                                            // 商品编码
	Photos      string  `gorm:"not null;type:longtext" json:"photos"`            // 要打印的照片
	Options     byte    `gorm:"not null" json:"options"`                         // 附加选项
	Extend1     byte    `gorm:"not null" json:"extend1,omitempty"`               // 扩展字段1
	Extend2     string  `gorm:"not null" json:"extend2,omitempty"`               // 扩展字段2
}

// Order 订单表数据结构
type Order struct {
	ID         string           `gorm:"primary_key" json:"id"`                                // 用户唯一标识
	UserID     string           `gorm:"not null" json:"userId"`                               // 用户id
	YzyID      string           `gorm:"not null" json:"yzyId"`                                // yzy平台id
	Platform   byte             `gorm:"not null" json:"platform"`                             // 哪个平台下的单(1: h5 2:小程序)[不要删]
	CreateTime string           `gorm:"not null" json:"createTime"`                           // 创建时间
	Method     string           `gorm:"not null" json:"method"`                               // 支付方式
	Price      float64          `gorm:"not null" gorm:"type:decimal(10,2)" json:"price"`      // 付款金额
	Offer      int              `gorm:"not null" json:"offer"`                                // 优惠
	Type       byte             `gorm:"not null" json:"type"`                                 // 优惠券类型
	Unit       string           `gorm:"not null" json:"unit"`                                 // 优惠券单位
	TotalPrice float64          `gorm:"not null" gorm:"type:decimal(10,2)" json:"totalPrice"` // 总价
	Progress   byte             `gorm:"not null" json:"progress"`                             // 邮费
	Status     byte             `gorm:"not null" json:"status"`                               // 订单状态
	ShipName   string           `gorm:"not null" json:"shipName"`                             // 快递
	ShipSn     string           `gorm:"not null" json:"shipSn"`                               // 快递单号
	Success    *bool            `json:"success"`
	Commodity  []OrderCommodity `json:"commodity"` // 订单相关商品
	Address    OrderAddress     `json:"address"`   // 订单地址
}

// 微信订单列表
type weappOrderList struct {
	ID          string   `json:"id"`          // 订单id
	CommodityID string   `json:"commodityId"` // 商品id
	CreateTime  string   `json:"createTime"`  // 创建时间
	Name        string   `json:"name"`        // 商品名
	Status      byte     `json:"status"`      // 订单状态  1.待付款|2.已取消|3.已付款|4.配送中|5.已签收|6.已评价|7.已退款|8.付款删除|9.未付款删除
	Count       int      `json:"count"`       // 数量
	Imgs        []string `json:"imgs"`        // 商品图片
	TotalPrice  float64  `json:"totalPrice"`  // 实付金额
	ShipName    string   `json:"shipName"`    // 快递
	ShipSn      string   `json:"shipSn"`      // 快递单号
	Success     *bool    `json:"success"`     // 是否付款成功
}

type orderList struct {
	Order
	Count  int      `json:"count"`  // 数量
	Photos []string `json:"photos"` // 商品图片
	User   User     `json:"-"`      // 用户
}

// InsertOrder 创建订单
func InsertOrder(order Order) response.Response {
	order.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	order.ID = strings.ToUpper(gopay.GetRandomString(16))
	order.Address.ID = gopay.GetRandomString(16)
	var result response.Response
	if dbc := Db().Create(&order); dbc.Error != nil {
		result = response.JSONError(dbc.Error.Error())
	} else {
		result = response.JSON(order.ID)
	}
	return result
}

type Photos struct {
	Img   string
	Count int
}

// GetAllOrder 分页查询订单
func GetAllOrder(current int, pageSize int, orderID string, status []string, start string, end string, links []string) response.OrderResponse {
	var (
		order []Order
		total int // 订单量
		count float64
		inc   float64 // 总收入
	)
	model := Db().Model(&Order{}).Select("orders.*").Order("orders.create_time desc").Joins("left join users on orders.user_id = users.id").Where("orders.id LIKE ? AND orders.create_time BETWEEN DATE_SUB(?,INTERVAL 1 DAY) AND DATE_ADD(?,INTERVAL 1 DAY) AND orders.status IN (?)", "%"+orderID+"%", start, end, status)
	if len(links) > 0 {
		model = model.Where("users.link_id IN (?)", links)
	}
	model.Count(&total).Limit(pageSize).Offset((current - 1) * pageSize).Preload("Commodity").Preload("Address").Find(&order)
	row := model.Select("sum(orders.total_price)").Row()
	row.Scan(&count)
	row1 := model.Where("orders.status IN (?)", []string{"3", "4", "5", "6", "7", "8"}).Select("sum(orders.total_price)").Row()
	row1.Scan(&inc)
	return response.OrderJSON(order, total, count, inc)
}

// WeappOrderList 微信订单列表
func WeappOrderList(current int, pageSize int, id string) response.Response {
	var (
		order []Order
		list  []weappOrderList
		count int
	)
	model := Db().Model(&Order{}).Order("create_time desc").Where("user_id = ?", id).Not("status", []int{8, 9})
	model.Count(&count).Limit(pageSize).Offset((current - 1) * pageSize).Find(&order)
	bt, _ := json.Marshal(order)
	json.Unmarshal(bt, &list)
	for k, v := range order {
		Db().Model(&v).Related(&v.Commodity).Related(&v.Address)
		if v.YzyID != "" && v.Status == 3 {
			shipSn, shipName, sta := SelectYzy(v.ID, v.YzyID)
			if sta == 20 {
				// 已发货
				Db().Model(&Order{}).Where("id=?", v.ID).Updates(map[string]interface{}{"ship_name": shipName, "ship_sn": shipSn, "status": 4})
			}
		}
		list[k].Count = 0
		list[k].ShipSn = v.ShipSn
		list[k].ShipName = v.ShipName
		for _, y := range v.Commodity {
			list[k].Imgs = append(list[k].Imgs, y.Img)
			list[k].Count += y.Count
			list[k].Name = y.Name
			list[k].CommodityID = y.CommodityID
		}
	}
	return response.PageJSON(list, count)
}

// UpdateOrder 更新数据
func UpdateOrder(w http.ResponseWriter, order Order, openid string, transactionID string) {
	var result response.Response
	if order.Status == 3 {
		var Config = config.Config.App
		client := wechat.NewClient(Config.AppID, Config.MchID, Config.Key, true)
		// 初始化参数结构体
		bm := make(gopay.BodyMap)
		bm.Set("out_trade_no", transactionID)
		bm.Set("nonce_str", gopay.GetRandomString(32))
		bm.Set("sign_type", wechat.SignType_MD5)
		// bm.Set("transaction_id", transactionID)

		// 请求订单查询，成功后得到结果
		wxRsp, _ := client.QueryOrder(bm)
		if wxRsp.ReturnCode == "SUCCESS" && wxRsp.TradeStateDesc == "支付成功" {
			if dbc := Db().Model(&Order{}).Where("id = ?", order.ID).Updates(map[string]interface{}{"status": 3, "success": true}); dbc.RowsAffected == 0 {
				result = response.JSONError("更新失败,没有该订单号!")
			} else {
				wg := new(sync.WaitGroup)
				wg.Add(2)
				go func() {
					defer wg.Done()
					Db().Where("id = ?", order.ID).Find(&order)
					Db().Model(&order).Related(&order.Commodity).Related(&order.Address)

					ints := make([]Photos, 0)
					json.Unmarshal([]byte(order.Commodity[0].Photos), &ints)
					addr := order.Address.Province + " " + order.Address.City + " " + order.Address.Area + " " + order.Address.Address
					str := CreateYzyOrder(order.ID, order.Address.Name, addr, order.Address.Phone, order.Commodity[0].Code, order.Commodity[0].Count, ints, order.Commodity[0].Options)
					Db().Model(&Order{}).Where("id = ?", order.ID).Update("yzy_id", str)
				}()
				go func() {
					defer wg.Done()
					result = response.JSONSuccess()
					response.ResultSuccess(w, result)
				}()
				wg.Wait()
				return
			}
		} else {
			result = response.JSONError("支付失败")
			fmt.Println(wxRsp)
		}
	} else {
		if dbc := Db().Model(&Order{}).Where("id = ?", order.ID).Update("status", order.Status); dbc.Error != nil {
			result = response.JSONError(dbc.Error.Error())
		} else {
			result = response.JSONSuccess()
		}
	}
	response.ResultSuccess(w, result)
}

// OrderDetails 小程序获取详情
func OrderDetails(id string) response.Response {
	var (
		order Order
	)
	Db().Where("id = ?", id).Find(&order)
	Db().Model(&order).Related(&order.Commodity).Related(&order.Address)
	return response.JSON(order)
}

// UpdateOrderAddress 更新订单地址
func UpdateOrderAddress(orderAddress OrderAddress) response.Response {
	var result response.Response
	if dbc := Db().Model(&OrderAddress{}).Where("id = ?", orderAddress.ID).Updates(&orderAddress); dbc.Error != nil {
		result = response.JSONError(dbc.Error.Error())
	} else {
		result = response.JSONSuccess()
	}
	return result
}
