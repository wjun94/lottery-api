package app

/**
 * 文档作者: wjun94
 * 创建时间：2019年09月22日
 * 修改时间：2019年10月10日
 * 描述信息：路由
 */
import (
	"context"
	"encoding/json"
	"mango-api/config"
	"mango-api/db"
	"mango-api/handler"
	"mango-api/response"
	"mango-api/utils"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/rs/cors"
)

// InitApp 初始化
func InitApp() {
	println("version: 0.1.6")
	router := mux.NewRouter()
	r := router.PathPrefix(config.Config.Route.PathPrefix).Subrouter()
	r.HandleFunc("/wxLogin", handler.LoginWxHandler).Methods("POST")
	r.HandleFunc("/createWxUser", handler.CreateWxUser).Methods("POST")
	r.HandleFunc("/getOpenID", handler.GetOpenID).Methods("POST")
	r.HandleFunc("/getAllCommodity", handler.GetAllCommodity).Methods("GET")
	r.HandleFunc("/getCommodityList", handler.GetCommodityList).Methods("GET")
	r.HandleFunc("/getRelatedCommodity", handler.GetRelatedCommodity).Methods("GET")
	r.HandleFunc("/getCouponByID", handler.GetCouponByID).Methods("GET")
	r.HandleFunc("/loginAdmin", handler.LoginAdminHandler).Methods("POST")
	r.HandleFunc("/getCommodityAd", handler.GetCommodityAdByID).Methods("GET")
	r.HandleFunc("/getPromote", handler.GetPromote).Methods("GET")
	r.HandleFunc("/getComment", handler.GetComment).Methods("GET")
	r.HandleFunc("/getPromoteByLinkId", handler.GetPromoteByLinkID).Methods("GET")

	admin := r.PathPrefix("").Subrouter()
	wx := r.PathPrefix("").Subrouter()
	{
		wx.HandleFunc("/addAddress", handler.AddAddress).Methods("POST")
		wx.HandleFunc("/getAllAddress", handler.GetAllAddress).Methods("GET")
		wx.HandleFunc("/updateAddress", handler.UpdateAddress).Methods("POST")
		wx.HandleFunc("/deleteAddress", handler.DeleteAddress).Methods("POST")
		wx.HandleFunc("/updateOrderAddress", handler.UpdateOrderAddress).Methods("POST")
	}
	{
		wx.HandleFunc("/createOrder", handler.CreateOrder).Methods("POST")
		wx.HandleFunc("/getWeappOrderList", handler.GetWeappOrderList).Methods("GET")
		wx.HandleFunc("/getOrderDetails", handler.GetOrderDetails).Methods("GET")
		wx.HandleFunc("/updateOrder", handler.UpdateOrder).Methods("POST")
		wx.HandleFunc("/wxPay", handler.WxPay).Methods("POST")
	}

	{
		wx.HandleFunc("/deleteCoupon", handler.DeleteCoupon).Methods("POST")
		wx.HandleFunc("/getAllUserCoupon", handler.GetAllUserCoupon).Methods("GET")
		wx.HandleFunc("/createUserCoupon", handler.CreateUserCoupon).Methods("POST")
		wx.HandleFunc("/getCommodityUserCouponById", handler.GetCommodityUserCouponByID).Methods("GET")
	}
	{
		wx.HandleFunc("/createComment", handler.CreateComment).Methods("POST")
		r.HandleFunc("/getCommodityComments", handler.GetCommodityComments).Methods("GET")
	}
	{
		admin.HandleFunc("/getOrder", handler.GetOrder).Methods("GET")
	}
	{
		admin.HandleFunc("/getAllCoupon", handler.GetAllCoupon).Methods("GET")
		admin.HandleFunc("/addCoupon", handler.AddCoupon).Methods("POST")
		admin.HandleFunc("/updateCoupon", handler.UpdateCoupon).Methods("POST")
		admin.HandleFunc("/deleteCoupon", handler.DeleteCoupon).Methods("POST")
	}
	{
		admin.HandleFunc("/createLink", handler.CreateLink).Methods("POST")
		admin.HandleFunc("/getAllLink", handler.GetAllLink).Methods("GET")
		admin.HandleFunc("/getLink", handler.SelectLink).Methods("GET")
		admin.HandleFunc("/deleteLink", handler.DeleteLink).Methods("POST")
		admin.HandleFunc("/updateLink", handler.UpdateLink).Methods("POST")
	}
	{
		admin.HandleFunc("/createPromote", handler.CreatePromote).Methods("POST")
		admin.HandleFunc("/updatePromote", handler.UpdatePromote).Methods("POST")
		admin.HandleFunc("/deletePromote", handler.DeletePromote).Methods("POST")
	}
	{
		admin.HandleFunc("/addCommodity", handler.AddCommodity).Methods("POST")
		admin.HandleFunc("/updateCommodity", handler.UpdateCommodity).Methods("POST")
		admin.HandleFunc("/deleteCommodity", handler.DeleteCommodity).Methods("POST")
	}
	{
		admin.HandleFunc("/getAllAccount", handler.GetAllAccount).Methods("GET")
		admin.HandleFunc("/updateAccount", handler.UpdateAccount).Methods("POST")
		admin.HandleFunc("/addAccount", handler.AddAccount).Methods("POST")
		admin.HandleFunc("/deleteAccount", handler.DeleteAccount).Methods("POST")
	}
	wx.HandleFunc("/upload", handler.UploadFile).Methods("POST") // 上传图片
	wx.HandleFunc("/deleteUserCoupon", handler.DeleteUserCoupon).Methods("POST")
	admin.HandleFunc("/getUserInfo", handler.GetUserInfo).Methods("GET")
	admin.HandleFunc("/getAuth", handler.GetAuth).Methods("GET")

	admin.Use(adminMiddleware)
	wx.Use(wxMiddleware)
	// handler := cors.Default().Handler(r)
	// handler := cors.New(cors.Options{
	// 	AllowedOrigins:   []string{"*"},
	// 	AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	// 	AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	// 	ExposedHeaders:   []string{"Link"},
	// 	AllowCredentials: true,
	// 	MaxAge:           300, // Maximum value not ignored by any of major browsers
	// }).Handler(r)
	http.ListenAndServe(":8080", r)
}

// adminMiddleware 后台中间件
func adminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 检验token
		tokenStr := r.Header.Get("authorization")
		loginUserID, level := utils.GetUserID(tokenStr)
		rd, _ := db.RedisInit().Get(loginUserID).Result()
		if tokenStr == "" || loginUserID == "" {
			response, _ := json.Marshal(response.JSONErrorCode(304))
			w.Write(response)
		} else if rd != tokenStr {
			response, _ := json.Marshal(response.JSONErrorCode(304))
			w.Write(response)
		} else {
			ctx := context.WithValue(r.Context(), "info", map[string]interface{}{"Type": level, "UserID": loginUserID})
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		}
	})
}

// wxMiddleware 微信中间件
func wxMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 检验token
		tokenStr := r.Header.Get("authorization")
		id := utils.GetWxID(tokenStr)
		rd, _ := db.RedisInit().Get(id).Result()
		if tokenStr == "" || id == "" {
			response, _ := json.Marshal(response.JSONErrorCode(304))
			w.Write(response)
		} else if rd != tokenStr {
			response, _ := json.Marshal(response.JSONErrorCode(304))
			w.Write(response)
		} else {
			ctx := context.WithValue(r.Context(), "info", map[string]string{"UserID": id})
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		}
	})
}
