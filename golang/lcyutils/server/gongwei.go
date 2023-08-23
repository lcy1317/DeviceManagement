package server

import (
	"code.corp.bcollie.net/whistle/lcyutils/models"
	"code.corp.bcollie.net/whistle/lcyutils/restapi/operations/shebei"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/xuri/excelize/v2"
	"os"
	"strconv"
)

func getDeviceError(message string, statusCode int64) middleware.Responder {
	response := &models.APIResponse{
		Code:    statusCode,
		Message: message,
	}
	return shebei.NewGetGetDeviceInternalServerError().WithPayload(response)
}

func modifyDeviceOK() middleware.Responder {
	response := &models.APIResponse{
		Code:    200,
		Message: "修改成功",
	}
	return shebei.NewPostModifyDeviceOK().WithPayload(response)
}
func modifyDeviceError(message string, statusCode int64) middleware.Responder {
	response := &models.APIResponse{
		Code:    statusCode,
		Message: message,
	}
	return shebei.NewPostModifyDeviceInternalServerError().WithPayload(response)
}

var Device_Info = "./device.xlsx"
var verifyCode = "jhwisbest"

func (s *Server) handleModifyDevice(params shebei.PostModifyDeviceParams) middleware.Responder {
	fmt.Println(params.Verify)
	if params.Verify != verifyCode {
		return modifyDeviceError("授权码错误", 500)
	}
	f, err := excelize.OpenFile(Device_Info)
	//fmt.Println(params)
	//fmt.Println(params.ID)
	if err != nil {
		return modifyDeviceError("Modify DeviceInfo Error!"+err.Error(), 500)
	}
	iID, err := strconv.Atoi(params.ID)
	if err != nil {
		return modifyDeviceError("Modify DeviceInfo Error! ID convert error!"+err.Error(), 500)
	}
	params.ID = strconv.Itoa(iID)

	err = f.SetCellValue("Sheet1", "A"+params.ID, params.ID)
	if err != nil {
		return modifyDeviceError("Modify DeviceInfo Error!"+err.Error(), 500)
	}

	err = f.SetCellValue("Sheet1", "B"+params.ID, params.Categorie)
	if err != nil {
		return modifyDeviceError("Modify DeviceInfo Error!"+err.Error(), 500)
	}
	err = f.SetCellValue("Sheet1", "C"+params.ID, params.Dtype)
	if err != nil {
		return modifyDeviceError("Modify DeviceInfo Error!"+err.Error(), 500)
	}
	err = f.SetCellValue("Sheet1", "D"+params.ID, params.Time)
	if err != nil {
		return modifyDeviceError("Modify DeviceInfo Error!"+err.Error(), 500)
	}
	err = f.SetCellValue("Sheet1", "E"+params.ID, params.Belong)
	if err != nil {
		return modifyDeviceError("Modify DeviceInfo Error!"+err.Error(), 500)
	}
	err = f.SetCellValue("Sheet1", "F"+params.ID, params.Ifmark)
	if err != nil {
		return modifyDeviceError("Modify DeviceInfo Error!"+err.Error(), 500)
	}
	err = f.SetCellValue("Sheet1", "G"+params.ID, params.Status)
	if err != nil {
		return modifyDeviceError("Modify DeviceInfo Error!"+err.Error(), 500)
	}
	err = f.SetCellValue("Sheet1", "H"+params.ID, params.Username)
	if err != nil {
		return modifyDeviceError("Modify DeviceInfo Error!"+err.Error(), 500)
	}
	err = f.SetCellValue("Sheet1", "I"+params.ID, params.Yikatong)
	if err != nil {
		return modifyDeviceError("Modify DeviceInfo Error!"+err.Error(), 500)
	}
	err = f.SetCellValue("Sheet1", "J"+params.ID, params.Position)
	if err != nil {
		return modifyDeviceError("Modify DeviceInfo Error!"+err.Error(), 500)
	}
	err = f.SetCellValue("Sheet1", "K"+params.ID, params.GdzcNumber)
	if err != nil {
		return modifyDeviceError("Modify DeviceInfo Error!"+err.Error(), 500)
	}
	if err = f.SaveAs(Device_Info); err != nil {
		return modifyDeviceError("Modify DeviceInfo Error!"+err.Error(), 500)
	}
	return modifyDeviceOK()
}

func (s *Server) handleGetDevice(params shebei.GetGetDeviceParams) middleware.Responder {

	if len(params.ID) >= 4 {
		return getDeviceError("Get DeviceInfo Error!", 500)
	}

	com := &models.DeviceInfo{
		Categories: "Undefined",
		Dtype:      "Undefined",
		Time:       "Undefined",
		Belong:     "Undefined",
		Ifmark:     "Undefined",
		Status:     "Undefined",
		Username:   "Undefined",
		Yikatong:   "Undefined",
		Position:   "Undefined",
		GdzcNumber: "Undefined",
	}
	f, err := excelize.OpenFile(Device_Info)
	if err != nil {
		return getDeviceError("Get DeviceInfo Error!"+err.Error(), 500)
	}
	com.Categories, _ = f.GetCellValue("Sheet1", "B"+params.ID)
	com.Dtype, _ = f.GetCellValue("Sheet1", "C"+params.ID)
	com.Time, _ = f.GetCellValue("Sheet1", "D"+params.ID)
	com.Belong, _ = f.GetCellValue("Sheet1", "E"+params.ID)
	com.Ifmark, _ = f.GetCellValue("Sheet1", "F"+params.ID)
	com.Status, _ = f.GetCellValue("Sheet1", "G"+params.ID)
	com.Username, _ = f.GetCellValue("Sheet1", "H"+params.ID)
	com.Yikatong, _ = f.GetCellValue("Sheet1", "I"+params.ID)
	com.Position, _ = f.GetCellValue("Sheet1", "J"+params.ID)
	com.GdzcNumber, _ = f.GetCellValue("Sheet1", "K"+params.ID)

	//fmt.Println(com)

	return shebei.NewGetGetDeviceOK().WithPayload(com)
	//return getDeviceError("No Device!", 500)
}

func getDeviceListError(message string, statusCode int64) middleware.Responder {
	response := &models.APIResponse{
		Code:    statusCode,
		Message: message,
	}
	return shebei.NewGetGetDeviceListInternalServerError().WithPayload(response)
}

func (s *Server) handleGetDeviceList(params shebei.GetGetDeviceListParams) middleware.Responder {

	f, err := excelize.OpenFile(Device_Info)

	if err != nil {
		return getDeviceListError("Get DeviceListInfo Error!"+err.Error(), 500)
	}

	var deviceList []*models.DeviceListInfoDataItems0
	for i := 1; i <= 512; i++ {
		id, err := f.GetCellValue("Sheet1", "A"+strconv.Itoa(i))
		if err != nil {
			return getDeviceListError("modify DeviceInfo Html Error!"+err.Error(), 500)
		}
		user, err := f.GetCellValue("Sheet1", "H"+strconv.Itoa(i))
		if err != nil {
			return getDeviceListError("modify DeviceInfo Html Error!"+err.Error(), 500)
		}
		if len(user) < 2 {
			continue
		}
		dtype, err := f.GetCellValue("Sheet1", "C"+strconv.Itoa(i))
		if err != nil {
			return getDeviceListError("modify DeviceInfo Html Error!"+err.Error(), 500)
		}
		if len(dtype) < 2 {
			continue
		}
		status, err := f.GetCellValue("Sheet1", "G"+strconv.Itoa(i))
		if err != nil {
			return getDeviceListError("modify DeviceInfo Html Error!"+err.Error(), 500)
		}

		gdzc_number, err := f.GetCellValue("Sheet1", "K"+strconv.Itoa(i))
		if err != nil {
			return getDeviceListError("modify DeviceInfo Html Error!"+err.Error(), 500)
		}

		scode := id
		for len(scode) < 5 {
			scode = "0" + scode
		}
		deviceItem := &models.DeviceListInfoDataItems0{
			ID:         id,
			SnCode:     "JH-W-" + scode,
			Name:       user,
			Type:       dtype,
			Status:     status,
			GdzcNumber: gdzc_number,
		}
		deviceList = append(deviceList, deviceItem)
	}

	response := &models.DeviceListInfo{
		Code: 200,
		Data: deviceList,
	}
	return shebei.NewGetGetDeviceListOK().WithPayload(response)
}

// 获取完整的所有设备信息的列表。
func getDeviceListFullError(message string, statusCode int64) middleware.Responder {
	response := &models.APIResponse{
		Code:    statusCode,
		Message: message,
	}
	return shebei.NewGetGetDeviceListFullInternalServerError().WithPayload(response)
}

func (s *Server) handleGetDeviceListFull(params shebei.GetGetDeviceListFullParams) middleware.Responder {

	f, err := excelize.OpenFile(Device_Info)

	if err != nil {
		return getDeviceListFullError("Get DeviceListInfo Error!"+err.Error(), 500)
	}

	var deviceList []*models.DeviceListInfoFullDataItems0
	for i := 1; i <= 512; i++ {
		id, err := f.GetCellValue("Sheet1", "A"+strconv.Itoa(i)) // 设备的ID
		if err != nil {
			return getDeviceListFullError("modify DeviceInfo Html Error!"+err.Error(), 500)
		}
		categories, err := f.GetCellValue("Sheet1", "B"+strconv.Itoa(i)) // 设备的分类
		if err != nil {
			return getDeviceListFullError("modify DeviceInfo Html Error!"+err.Error(), 500)
		}
		dtype, err := f.GetCellValue("Sheet1", "C"+strconv.Itoa(i))
		if err != nil {
			return getDeviceListFullError("modify DeviceInfo Html Error!"+err.Error(), 500)
		}
		if len(dtype) < 2 {
			continue
		}
		buytime, err := f.GetCellValue("Sheet1", "D"+strconv.Itoa(i)) // 入库时间
		if err != nil {
			return getDeviceListFullError("modify DeviceInfo Html Error!"+err.Error(), 500)
		}
		belong, err := f.GetCellValue("Sheet1", "E"+strconv.Itoa(i)) // 归属
		if err != nil {
			return getDeviceListFullError("modify DeviceInfo Html Error!"+err.Error(), 500)
		}
		if len(belong) < 2 {
			continue
		}
		ifmark, err := f.GetCellValue("Sheet1", "F"+strconv.Itoa(i)) // 是否有资产标
		if err != nil {
			return getDeviceListFullError("modify DeviceInfo Html Error!"+err.Error(), 500)
		}

		user, err := f.GetCellValue("Sheet1", "H"+strconv.Itoa(i))
		if err != nil {
			return getDeviceListFullError("modify DeviceInfo Html Error!"+err.Error(), 500)
		}
		if len(user) < 2 {
			continue
		}
		status, err := f.GetCellValue("Sheet1", "G"+strconv.Itoa(i))
		if err != nil {
			return getDeviceListFullError("modify DeviceInfo Html Error!"+err.Error(), 500)
		}
		yikatong, err := f.GetCellValue("Sheet1", "I"+strconv.Itoa(i))
		if err != nil {
			return getDeviceListFullError("modify DeviceInfo Html Error!"+err.Error(), 500)
		}

		position, err := f.GetCellValue("Sheet1", "J"+strconv.Itoa(i))
		if err != nil {
			return getDeviceListFullError("modify DeviceInfo Html Error!"+err.Error(), 500)
		}

		gdzc_number, err := f.GetCellValue("Sheet1", "K"+strconv.Itoa(i))
		if err != nil {
			return getDeviceListFullError("modify DeviceInfo Html Error!"+err.Error(), 500)
		}

		scode := id
		for len(scode) < 5 {
			scode = "0" + scode
		}
		deviceItem := &models.DeviceListInfoFullDataItems0{
			ID:         id,
			Categories: categories,
			Buytime:    buytime,
			Belong:     belong,
			Ifmark:     ifmark,
			SnCode:     "JH-W-" + scode,
			Name:       user,
			Type:       dtype,
			Status:     status,
			Yikatong:   yikatong,
			Position:   position,
			GdzcNumber: gdzc_number,
		}
		deviceList = append(deviceList, deviceItem)
	}

	response := &models.DeviceListInfoFull{
		Code: 200,
		Data: deviceList,
	}
	return shebei.NewGetGetDeviceListFullOK().WithPayload(response)
}

func modifyDeviceListOK() middleware.Responder {
	response := &models.APIResponse{
		Code:    200,
		Message: "设备列表更新成功",
	}
	return shebei.NewPostModifyDeviceListOK().WithPayload(response)
}
func modifyDeviceListError(message string, statusCode int64) middleware.Responder {
	response := &models.APIResponse{
		Code:    statusCode,
		Message: message,
	}
	return shebei.NewPostModifyDeviceListInternalServerError().WithPayload(response)
}
func (s *Server) handleModifyDeviceList(params shebei.PostModifyDeviceListParams) middleware.Responder {

	if params.Verify != verifyCode {
		return modifyDeviceError("授权码错误", 500)
	}

	before := "<!doctype html>\n<html>\n<head>\n\t<meta charset='UTF-8'><meta name='viewport' content='width=device-width initial-scale=1'>\n\n\t<link href='https://fonts.loli.net/css?family=PT+Serif:400,400italic,700,700italic&subset=latin,cyrillic-ext,cyrillic,latin-ext' rel='stylesheet' type='text/css' />\n\n\t<link href=\"middleware/index.css\" rel=\"stylesheet\" type=\"text/css\" />\n\n\t</style><title>王家恒课题组设备列表</title>\n\n\t<script src=\"middleware/taidu_common.js\"></script>\n\t<script src=\"middleware/taidu_smartcommunity.js\"></script>\n\t<script src=\"middleware/jq-3.6.0.js\"></script>\n\n</head>\n<body class='typora-export os-windows'><div class='typora-export-content'>\n<div id='write'  class=''><h1 id='王家恒课题组设备登记'><span>王家恒课题组设备登记</span></h1>\n</div>\n<div id='write'  class=''>\n\t<figure>\n\t\t<table><thead><tr><th><span>设备ID</span></th><th><span>设备使用者</span></th><th><span>设备类型</span></th></tr></thead>\n"
	after := "\t\t</table></figure><p>&nbsp;</p></div>\n\t\n</div></div>\n</body>\n</html>"

	f, err := excelize.OpenFile(Device_Info)

	if err != nil {
		return modifyDeviceListError("Modify DeviceInfo Html Error!"+err.Error(), 500)
	}
	deviceInformation := ""
	for i := 1; i <= 108; i++ {
		id, err := f.GetCellValue("Sheet1", "A"+strconv.Itoa(i))
		if err != nil {
			return modifyDeviceListError("modify DeviceInfo Html Error!"+err.Error(), 500)
		}
		user, err := f.GetCellValue("Sheet1", "H"+strconv.Itoa(i))
		if err != nil {
			return modifyDeviceListError("modify DeviceInfo Html Error!"+err.Error(), 500)
		}
		if len(user) < 2 {
			continue
		}
		dtype, err := f.GetCellValue("Sheet1", "C"+strconv.Itoa(i))
		if err != nil {
			return modifyDeviceListError("modify DeviceInfo Html Error!"+err.Error(), 500)
		}
		if len(dtype) < 2 {
			continue
		}

		gdzc_number, err := f.GetCellValue("Sheet1", "K"+strconv.Itoa(i))
		if err != nil {
			return modifyDeviceListError("modify DeviceInfo Html Error!"+err.Error(), 500)
		}
		if len(dtype) < 2 {
			continue
		}
		temp := "<tbody><tr><td><span><a href = \"http://www.seujyh.cn/device?id="
		temp = temp + id
		temp = temp + "\"> JH-W-"
		temp = temp + id
		temp = temp + "</a></span> </td><td><span>"
		temp = temp + user
		temp = temp + "</span></td><td><span>"
		temp = temp + dtype
		temp = temp + "</span></td></tr></tbody>"
		temp = temp + gdzc_number
		temp = temp + "</span></td></tr></tbody>"
		deviceInformation = deviceInformation + temp + "\n"
	}

	htmlPage := before + deviceInformation + after
	fileName := "deviceIndex.html"
	htmlFile, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return modifyDeviceListError("文件读取出错 ", 500)
	}
	defer htmlFile.Close()

	htmlFile.WriteString(htmlPage + "\n")

	return modifyDeviceListOK()
}
