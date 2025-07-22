package order

import (
	"context"
	"fmt"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/house"
	"github.com/flipped-aurora/gin-vue-admin/server/model/order"
	orderReq "github.com/flipped-aurora/gin-vue-admin/server/model/order/request"
	orderResp "github.com/flipped-aurora/gin-vue-admin/server/model/order/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/user"
)

type RentalOrderService struct{}

// CreateRentalOrder 创建rentalOrder表记录
// Author [yourname](https://github.com/yourname)
func (rentalOrderService *RentalOrderService) CreateRentalOrder(ctx context.Context, rentalOrder *order.RentalOrder) (err error) {
	err = global.GVA_DB.Create(rentalOrder).Error
	return err
}

// DeleteRentalOrder 删除rentalOrder表记录
// Author [yourname](https://github.com/yourname)
func (rentalOrderService *RentalOrderService) DeleteRentalOrder(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&order.RentalOrder{}, "id = ?", ID).Error
	return err
}

// DeleteRentalOrderByIds 批量删除rentalOrder表记录
// Author [yourname](https://github.com/yourname)
func (rentalOrderService *RentalOrderService) DeleteRentalOrderByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]order.RentalOrder{}, "id in ?", IDs).Error
	return err
}

// UpdateRentalOrder 更新rentalOrder表记录
// Author [yourname](https://github.com/yourname)
func (rentalOrderService *RentalOrderService) UpdateRentalOrder(ctx context.Context, rentalOrder order.RentalOrder) (err error) {
	err = global.GVA_DB.Model(&order.RentalOrder{}).Where("id = ?", rentalOrder.ID).Updates(&rentalOrder).Error
	return err
}

// GetRentalOrder 根据ID获取rentalOrder表记录
// Author [yourname](https://github.com/yourname)
func (rentalOrderService *RentalOrderService) GetRentalOrder(ctx context.Context, ID string) (rentalOrder order.RentalOrder, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&rentalOrder).Error
	return
}

// GetRentalOrderInfoList 分页获取rentalOrder表记录
// Author [yourname](https://github.com/yourname)
func (rentalOrderService *RentalOrderService) GetRentalOrderInfoList(ctx context.Context, info orderReq.RentalOrderSearch) (list []orderResp.RentalOrderListItem, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Debug().Model(&order.RentalOrder{})
	var rentalOrders []order.RentalOrder
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}

	if info.TenantId != nil {
		db = db.Where("tenant_id = ?", *info.TenantId)
	}
	if info.LandlordId != nil {
		db = db.Where("landlord_id = ?", *info.LandlordId)
	}
	// 优化：只JOIN一次每个表，并在JOIN中加LIKE条件，避免多次JOIN导致SQL错乱
	// 只有有名称筛选时才JOIN
	if info.HouseTitle != "" {
		db = db.Joins("LEFT JOIN house ON rental_order.house_id = house.house_id AND house.title LIKE ?", "%"+info.HouseTitle+"%")
	}
	if info.TenantName != "" {
		db = db.Joins("LEFT JOIN user_base AS tenant ON rental_order.tenant_id = tenant.user_id AND tenant.name = ?", info.TenantName)
	}
	if info.LandlordName != "" {
		db = db.Joins("LEFT JOIN user_base AS landlord ON rental_order.landlord_id = landlord.user_id AND landlord.name = ?", info.LandlordName)
	}
	// 新增：订单状态精确查询
	if info.Status != "" {
		db = db.Where("rental_order.status = ?", info.Status)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	// 只查需要的字段
	err = db.Select("rental_order.id, rental_order.house_id, rental_order.tenant_id, rental_order.landlord_id, rental_order.rent_start, rental_order.rent_end, rental_order.rent_amount, rental_order.deposit, rental_order.status, rental_order.signed_at, rental_order.created_at").Order("rental_order.created_at DESC").Find(&rentalOrders).Error
	if err != nil {
		return
	}

	// 批量收集所有 houseId、tenantId、landlordId
	houseIdSet := make(map[int]struct{})
	userIdSet := make(map[int]struct{})
	for _, ro := range rentalOrders {
		if ro.HouseId != nil {
			houseIdSet[*ro.HouseId] = struct{}{}
		}
		if ro.TenantId != nil {
			userIdSet[*ro.TenantId] = struct{}{}
		}
		if ro.LandlordId != nil {
			userIdSet[*ro.LandlordId] = struct{}{}
		}
	}
	// 转为切片
	houseIds := make([]int, 0, len(houseIdSet))
	for id := range houseIdSet {
		houseIds = append(houseIds, id)
	}
	userIds := make([]int, 0, len(userIdSet))
	for id := range userIdSet {
		userIds = append(userIds, id)
	}

	// 批量查房源
	var houses []house.House
	houseMap := make(map[int]string)
	if len(houseIds) > 0 {
		err = global.GVA_DB.Model(&house.House{}).Select("house_id, title").Where("house_id in ?", houseIds).Find(&houses).Error
		if err != nil {
			return
		}
		for _, h := range houses {
			houseMap[int(h.HouseID)] = h.Title
		}
	}
	// 批量查用户
	var users []user.UserBase
	userMap := make(map[int]string)
	if len(userIds) > 0 {
		err = global.GVA_DB.Model(&user.UserBase{}).Select("user_id, name").Where("user_id in ?", userIds).Find(&users).Error
		if err != nil {
			return
		}
		for _, u := range users {
			userMap[int(u.UserID)] = u.Name
		}
	}

	// 组装返回
	var result []orderResp.RentalOrderListItem
	for _, ro := range rentalOrders {
		houseTitle := ""
		if ro.HouseId != nil {
			houseTitle = houseMap[*ro.HouseId]
		}
		tenantName := ""
		if ro.TenantId != nil {
			tenantName = userMap[*ro.TenantId]
		}
		landlordName := ""
		if ro.LandlordId != nil {
			landlordName = userMap[*ro.LandlordId]
		}
		result = append(result, orderResp.RentalOrderListItem{
			ID:           int(ro.ID),
			HouseTitle:   houseTitle,
			TenantName:   tenantName,
			LandlordName: landlordName,
			RentStart:    derefDate(ro.RentStart),
			RentEnd:      derefDate(ro.RentEnd),
			RentAmount:   derefFloat(ro.RentAmount),
			Deposit:      derefFloat(ro.Deposit),
			Status:       derefString(ro.Status),
			SignedAt:     derefTime(ro.SignedAt),
			CreatedAt:    derefTime(&ro.CreatedAt),
		})
	}
	return result, total, nil
}
func (rentalOrderService *RentalOrderService) GetRentalOrderPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

// GetRentalOrderDetail 根据ID获取订单及其合同信息
func (rentalOrderService *RentalOrderService) GetRentalOrderDetail(ctx context.Context, ID string) (*orderResp.RentalOrderDetail, error) {
	if ID == "" || ID == "0" {
		return nil, fmt.Errorf("订单ID不能为空或为0")
	}
	var rentalOrder order.RentalOrder
	if err := global.GVA_DB.Where("id = ?", ID).First(&rentalOrder).Error; err != nil {
		return nil, err
	}
	var contract order.Contract
	var contractInfo *orderResp.ContractInfo
	if err := global.GVA_DB.Where("order_id = ?", rentalOrder.ID).First(&contract).Error; err == nil {
		contractInfo = &orderResp.ContractInfo{
			ContractURL: contract.ContractURL,
			Status:      contract.Status,
			SignTime:    contract.SignTime.Format("2006-01-02 15:04:05"),
		}
	}
	// 查询房源名称
	var houseInfo house.House
	houseTitle := ""
	houseImages := []string{}
	if rentalOrder.HouseId != nil && *rentalOrder.HouseId != 0 {
		if err := global.GVA_DB.Select("title").Where("house_id = ?", *rentalOrder.HouseId).First(&houseInfo).Error; err == nil {
			houseTitle = houseInfo.Title
		}
		// 查询房源图片
		var images []house.HouseImage
		if err := global.GVA_DB.Select("image_url").Where("house_id = ?", *rentalOrder.HouseId).Order("sort_order").Find(&images).Error; err == nil {
			for _, img := range images {
				houseImages = append(houseImages, img.ImageURL)
			}
		}
	}
	// 查询租客名称
	var tenant user.UserBase
	tenantName := ""
	if rentalOrder.TenantId != nil && *rentalOrder.TenantId != 0 {
		if err := global.GVA_DB.Select("name").Where("user_id = ?", *rentalOrder.TenantId).First(&tenant).Error; err == nil {
			tenantName = tenant.Name
		}
	}
	// 查询房东名称
	var landlord user.UserBase
	landlordName := ""
	if rentalOrder.LandlordId != nil && *rentalOrder.LandlordId != 0 {
		if err := global.GVA_DB.Select("name").Where("user_id = ?", *rentalOrder.LandlordId).First(&landlord).Error; err == nil {
			landlordName = landlord.Name
		}
	}
	detail := &orderResp.RentalOrderDetail{
		HouseTitle:   houseTitle,
		HouseImages:  houseImages,
		TenantName:   tenantName,
		LandlordName: landlordName,
		RentAmount:   derefFloat(rentalOrder.RentAmount),
		Status:       derefString(rentalOrder.Status),
		SignedAt:     derefTime(rentalOrder.SignedAt),
		Contract:     contractInfo,
	}
	return detail, nil
}

func derefInt(ptr *int) int {
	if ptr != nil {
		return *ptr
	}
	return 0
}
func derefFloat(ptr *float64) float64 {
	if ptr != nil {
		return *ptr
	}
	return 0
}
func derefString(ptr *string) string {
	if ptr != nil {
		return *ptr
	}
	return ""
}

func derefTime(ptr *time.Time) string {
	if ptr != nil {
		return ptr.Format("2006-01-02 15:04:05")
	}
	return ""
}

func derefDate(ptr *time.Time) string {
	if ptr != nil {
		return ptr.Format("2006-01-02")
	}
	return ""
}
