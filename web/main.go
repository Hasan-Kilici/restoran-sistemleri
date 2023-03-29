package main

import (
	"fmt"
	"Kawethra/utils"
	"Kawethra/dataset"
	"net/http"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"Kawethra/api"
	"os"
	"io"
)

func main() {
	userdata , err := dataset.UserCSV()
	if err != nil {
		fmt.Println("Kullanıcı Bilgileri Okunamadı, Bir yanlışlık var.")
	}
	permdata , err := dataset.PermCSV()
	if err != nil {
		fmt.Println("Yetki Bilgileri Okunamadı, Bir yanlışlık var.")
	}
	tablesdata , err := dataset.TableCSV()
	if err != nil {
		fmt.Println("Masa Bilgileri Okunamadı, Bir yanlışlık var.")
	}

	fmt.Println(userdata, permdata, tablesdata)

	r := gin.Default()

	r.LoadHTMLGlob("src/*.tmpl")
	r.Static("/static", "./static/")
	r.Static("/components", "./components/")
	r.Static("/uploads", "./uploads/")

	r.GET("/", func(ctx *gin.Context){
		token , err := ctx.Cookie("token")
		if err != nil {
			ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
				"title":      "Anasayfa",
				"userStatus": "false",
			})
			return
		}
		
		user, err := dataset.FindUserByToken(token)
		if err != nil {
			ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
				"title":      "Anasayfa",
				"userStatus": "false",
			})
			return
		}

		ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title":      "Anasayfa",
			"userStatus": "true",
			"userId":     user.ID,
			"username":   user.Username,
			"email":      user.Email,
			"perms":	  strings.Split(user.Perms, ">"),
		})
	})

	r.POST("/add/table/to-company",func(ctx *gin.Context){
		if !utils.HasRequiredPerms(ctx, []int{3}) {
			ctx.Redirect(http.StatusFound, "/")
			return
		}
		dataset.AddTableToCSV("./data/tables.csv");
		ctx.Redirect(http.StatusFound, "/#/yonetici/#/paneli")
	})

	r.POST("/add/worker",func(ctx *gin.Context){
		if !utils.HasRequiredPerms(ctx, []int{3}) {
			ctx.Redirect(http.StatusFound, "/")
			return
		}
		isim := ctx.PostForm("isim")
		email := ctx.PostForm("email")
		sifre := ctx.PostForm("sifre")
		yetki := ctx.PostForm("yetki")

		dataset.AddUserToCSV("./data/users.csv", isim,sifre,email,yetki)
		ctx.Redirect(http.StatusFound, "/#/yonetici/#/paneli?param=yetkili")
	})

	r.POST("/delete/worker/:id", func(ctx *gin.Context){
		if !utils.HasRequiredPerms(ctx, []int{3}) {
			ctx.Redirect(http.StatusFound, "/")
			return
		}
	})

	r.POST("/add/food", func(ctx *gin.Context){
		if !utils.HasRequiredPerms(ctx, []int{3}) {
			ctx.Redirect(http.StatusFound, "/")
			return
		}	

		name := ctx.PostForm("name")
		price, err := strconv.Atoi(ctx.PostForm("price"))
		if err != nil {
			ctx.Redirect(http.StatusFound, "/")
			return	
		}

		file, header , err := ctx.Request.FormFile("file")
        filename := header.Filename
        fmt.Println(header.Filename)
		dosya := "./uploads/"+filename
        out, err := os.Create("./uploads/"+filename)
        if err != nil {
            ctx.Redirect(http.StatusFound, "/#/yonetici/#/paneli?param=yemekler")
			return
        }
        defer out.Close()
        _, err = io.Copy(out, file)
        if err != nil {
            ctx.Redirect(http.StatusFound, "/#/yonetici/#/paneli?param=yemekler")
			return
        }   

		dataset.AddFoodToCSV("./data/foods.csv", name, price, dosya)
		ctx.Redirect(http.StatusFound, "/#/yonetici/#/paneli?param=yemekler")
	})

	r.POST("/delete/food/:id", func(ctx *gin.Context){
		if !utils.HasRequiredPerms(ctx, []int{3}) {
			ctx.Redirect(http.StatusFound, "/")
			return
		}	

		id := ctx.Param("id")
		dataset.DeleteFood(id)
	})

	r.POST("/delete/order/:orderid",func(ctx *gin.Context){
		if !utils.HasRequiredPerms(ctx, []int{3}) || !utils.HasRequiredPerms(ctx, []int{2})  || !utils.HasRequiredPerms(ctx, []int{1}) {
			ctx.Redirect(http.StatusFound, "/")
			return
		}
		orderid := ctx.Param("orderid")
		dataset.DeleteFoodOrder(orderid)
		ctx.Redirect(http.StatusFound, "/#/siparisler")
	})

	r.POST("/add/food/for-table/:tableid/:foodname/:price",func(ctx *gin.Context){
		if !utils.HasRequiredPerms(ctx, []int{3}) || !utils.HasRequiredPerms(ctx, []int{2})  || !utils.HasRequiredPerms(ctx, []int{1}) {
			ctx.Redirect(http.StatusFound, "/")
			return
		}	

		tableid, _ := strconv.Atoi(ctx.Param("tableid"))
		foodname := utils.ReplaceSpacesWithDash(ctx.Param("foodname"))
		price , _ := strconv.Atoi(ctx.Param("price"))
		urlparam1 := strconv.Itoa(tableid)
		urlparam2 := "/#/masa?param="
		url := urlparam2+urlparam1
		dataset.AddFoodToTable(tableid, foodname, price)
		dataset.AddFoodToOrder(tableid, foodname, price)
		ctx.Redirect(http.StatusFound, url)
	})

	r.POST("/remove/food/for-table/:orderid/:tableid",func(ctx *gin.Context){
		if !utils.HasRequiredPerms(ctx, []int{3}) || !utils.HasRequiredPerms(ctx, []int{2})  || !utils.HasRequiredPerms(ctx, []int{1}) {
			ctx.Redirect(http.StatusFound, "/")
			return
		}	
		order := ctx.Param("orderid")
		tableid := ctx.Param("tableid")
		url := "/#/masa?param="+tableid
		dataset.DeleteOrder(order)
		ctx.Redirect(http.StatusFound, url)
		
	})

	r.POST("/reset/table/:tableid/:count",func(ctx *gin.Context){
		if !utils.HasRequiredPerms(ctx, []int{3}) || !utils.HasRequiredPerms(ctx, []int{2})  || !utils.HasRequiredPerms(ctx, []int{1}) {
			ctx.Redirect(http.StatusFound, "/")
			return
		}	

		tableid := ctx.Param("tableid")
		count, _ := strconv.Atoi(ctx.Param("count"))
		for i := 0;i<count;i++{
			dataset.ResetTable(tableid);
			url := "/#/masa?param="+tableid
			ctx.Redirect(http.StatusFound,url)
		}
	})

	r.POST("/login", func(ctx *gin.Context) {
		email := ctx.PostForm("email")
		password := ctx.PostForm("password")
	
		token, err := utils.LoginUser(email, password)
		if err != nil {
			ctx.String(http.StatusUnauthorized, "Email ya da şifre yanlış!")
			return
		}
	
		ctx.SetCookie("token", token, 36000, "/", "", false, true)
		ctx.Redirect(http.StatusFound, "/")
	})

	r.POST("/create/perm",func(ctx *gin.Context){
		if !utils.HasRequiredPerms(ctx, []int{3}) {
			ctx.Redirect(http.StatusFound, "/")
			return
		}	

		isim := ctx.PostForm("name")
		girissaat := ctx.PostForm("girissaat")
		cikissaat := ctx.PostForm("cikissaat")
		izinVerilenSaat := girissaat+">"+cikissaat
		dataset.AddPermToCSV("./data/perms.csv",izinVerilenSaat, isim)
		ctx.Redirect(http.StatusFound, "/#/yonetici/#/paneli?param=calisanlar")
	})
	//Api
	r.POST("/users/api", func(ctx *gin.Context){
		users , err := api.ReadUsers("./data/users.csv")
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"Hata": "CSV Dosyası okunamadı",
			})
			return
		}
		ctx.JSON(http.StatusOK, users)
	})

	r.POST("/foods/api", func(ctx *gin.Context){
		foods , err := api.ReadFoods("./data/foods.csv")
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"Hata": "CSV Dosyası okunamadı",
			})
			return
		}
		ctx.JSON(http.StatusOK, foods)
	})

	r.POST("/tables/api", func(ctx *gin.Context){
		tables , err := api.ReadTables("./data/tables.csv")
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"Hata": "CSV Dosyası okunamadı",
			})
			return
		}
		ctx.JSON(http.StatusOK, tables)
	})

	r.POST("/orders/api", func(ctx *gin.Context){
		orders , err := api.ReadOrders("./data/orders.csv")
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"Hata": "CSV Dosyası okunamadı",
			})
			return
		}
		ctx.JSON(http.StatusOK, orders)
	})

	r.POST("/perms/api",func(ctx *gin.Context){
		perms , err := api.ReadPerms("./data/perms.csv")
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"Hata": "CSV Dosyası okunamadı",
			})
			return
		}
		ctx.JSON(http.StatusOK, perms)
	})

	r.POST("/orders/api/for-cheff",func(ctx *gin.Context){
		orders , err := api.ReadOrders("./data/foodlist.csv")
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"Hata": "CSV Dosyası okunamadı",
			})
			return
		}
		ctx.JSON(http.StatusOK, orders)
	})
	r.Run(":5000")
}
