package main

//執行:
//swag init -g swagger.go
//or
//swag init

// notes:
// swag init -g swagger.go (-g 指定用哪個file抓swagger general API Info,沒指定的話預設抓main同一層的。下cmd的目錄需為main目錄)
// 如果main在內層其他swagger註解在外層會抓不到!
// docs要和執行的main放在同一層，不然啟動會error
// swag init -g swagger.go  -o ./cmd/servicemanager/docs  [-g] 指定main註釋在swagger.go [-o] 產生出來的docs放指定位置

// --------swagger general API Info--------

// @title Account Detail
// @version 2.0
// @description Account Detail API

// -------APIs--------

/*
sample:
 @status  statuscode 物件     物件名        描述
 @Success 200        {object} provisionreq "{"code":0,"message":"OK","data":{"username":"kong"}}"
*/

// @Tags Customer API
// @Summary customer get vip account detail
// @Description 客戶查詢帳號資訊
// @Accept  json
// @Produce  json
// @Param username query string true "Account Email"
// @Success 200 {object} models.AccountIdDetail
// @Failure 400 {object} models.Err "{"error":"account not exist"}"
// @Router /customer/account [get]
func queryAccount_customer() {

}

// @Tags API
// @Summary get all vip accounts detail
// @Description 查詢所有帳號資訊
// @Accept  json
// @Produce  json
// @Success 200 {array} models.AccountIdDetail
// @Failure 400 {object} models.Err
// @Router /accounts [get]
func queryAccounts() {

}

// @Tags API
// @Summary get vip account detail
// @Description 查詢帳號資訊
// @Accept  json
// @Produce  json
// @Param id path string true "Account ID"
// @Success 200 {object} models.AccountIdDetail
// @Failure 400 {object} models.Err "{"error":"account not exist"}"
// @Router /account/{id} [get]
func queryAccount() {

}

// @Tags API
// @Summary add vip account detail
// @Description 新增帳號資訊
// @Accept  json
// @Produce  json
// @param detail body models.AccountIdDetail true "Account Detail"
// @Success 200 string string "Account Id"
// @Failure 400 {object} models.Err
// @Router /account [post]
func addAccount() {

}

// @Tags API
// @Summary edit vip account detail
// @Description 修改帳號資訊
// @Accept  json
// @Produce  json
// @Param id path string true "Account ID"
// @param detail body models.AccountIdDetail true "Account Detail"
// @Success 200 string string
// @Failure 400 {object} models.Err
// @Router /account/{id} [put]
func editAccount() {

}

// @Tags API
// @Summary delete vip account detail
// @Description 刪除帳號資訊
// @Accept  json
// @Produce  json
// @Param id path string true "Account ID"
// @Success 200 string string
// @Failure 400 {object} models.Err
// @Router /account/{id} [delete]
func deleteAccount() {

}
