package controller

import (
	"context"
	"errors"
	"github.com/astaxie/beego/logs"
	"test/common/pb"
	"test/service/mvp/dao"
	"test/service/mvp/model"
	"time"
)

type MvpServer struct {
	pb.UnimplementedMvpServer
}

func (MvpServer) AddGood(ctx context.Context, req *pb.AddGoodReq) (*pb.AddGoodRes, error) {
	logs.Info("AddGood", ctx, req)

	var res pb.AddGoodRes

	// 1. 判断商品是否存在
	good, err := dao.GoodDao.GetByName(req.GoodName)
	if err != nil {
		logs.Error(err)
		return nil, errors.New("Fail to finish GoodDao.GetByName")
	}
	if good != nil {
		return nil, errors.New("Good has existed")
	}

	// 2. 创建商品
	if err := dao.GoodDao.Create(req.GoodName, float64(req.Price)); err != nil {
		logs.Error(err)
		return nil, err
	}
	return &res, nil
}

func (MvpServer) SellGood(ctx context.Context, req *pb.SellGoodReq) (*pb.SellGoodRes, error) {
	logs.Info("SellGood", ctx, req)

	var res pb.SellGoodRes

	// 1. 判断商品是否存在
	good, err := dao.GoodDao.GetByName(req.GoodName)
	if err != nil {
		logs.Error(err)
		return nil, errors.New("Fail to finish GoodDao.GetByName")
	}
	if good == nil {
		return nil, errors.New("Good not existed")
	}

	// 2. 创建销售记录
	if err := dao.SellRecordDao.Create(int(good.ID), float64(req.SellPrice)); err != nil {
		logs.Error(err)
		return nil, err
	}
	return &res, nil
}

func (MvpServer) AddBilliardDesk(ctx context.Context, req *pb.AddBilliardDeskReq) (*pb.AddBilliardDeskRes, error) {
	logs.Info("AddBilliardDesk", ctx, req)

	var res pb.AddBilliardDeskRes

	// 1. 判断桌名是否重复
	desk, err := dao.BilliardDeskDao.GetByName(req.BilliardDeskName)
	if err != nil {
		logs.Error(err)
		return nil, errors.New("Fail to finish BilliardDeskDao.GetByName")
	}
	if desk != nil {
		return nil, errors.New("Desk name dumplicate")
	}

	// 2. 创建台球桌
	if err := dao.BilliardDeskDao.Create(req.BilliardDeskName); err != nil {
		logs.Error(err)
		return nil, err
	}
	return &res, nil
}

func (MvpServer) BeginPlayBilliard(ctx context.Context, req *pb.BeginPlayBilliardReq) (*pb.BeginPlayBilliardRes, error) {
	logs.Info("BeginPlayBilliard", ctx, req)

	var res pb.BeginPlayBilliardRes

	// 1. 判断桌名是否存在
	desk, err := dao.BilliardDeskDao.GetByName(req.BilliardDeskName)
	if err != nil {
		logs.Error(err)
		return nil, errors.New("Fail to finish BilliardDeskDao.GetByName")
	}
	if desk == nil {
		return nil, errors.New("Desk not exist")
	}

	// 2. 判断记录是否存在
	record, err := dao.PlayRecordDao.Get(int(desk.ID), time.Unix(req.BeginPlayTimestamp, 0))
	if err != nil {
		logs.Error(err)
		return nil, errors.New("Fail to finish PlayRecordDao.Get")
	}
	if record != nil {
		return nil, errors.New("Record has exist")
	}

	// 3. 创建玩台球的记录
	if err := dao.PlayRecordDao.Create(
		int(desk.ID),
		time.Unix(req.BeginPlayTimestamp, 0),
	); err != nil {
		logs.Error(err)
		return nil, err
	}
	return &res, nil
}

func (MvpServer) StopPlayBilliard(ctx context.Context, req *pb.StopPlayBilliardReq) (*pb.StopPlayBilliardRes, error) {
	logs.Info("StopPlayBilliard", ctx, req)

	var res pb.StopPlayBilliardRes

	// 1. 判断桌名是否存在
	desk, err := dao.BilliardDeskDao.GetByName(req.BilliardDeskName)
	if err != nil {
		logs.Error(err)
		return nil, errors.New("Fail to finish BilliardDeskDao.GetByName")
	}
	if desk == nil {
		return nil, errors.New("Desk not exist")
	}

	// 2. 更新玩台球的记录
	if err := dao.PlayRecordDao.UpdateStopPlayTime(
		int(desk.ID),
		time.Unix(req.BeginPlayTimestamp, 0),
		time.Unix(req.StopPlayTimestamp, 0),
	); err != nil {
		logs.Error(err)
		return nil, err
	}
	return &res, nil
}

func (MvpServer) Order(ctx context.Context, req *pb.OrderReq) (*pb.OrderRes, error) {
	logs.Info("Order", ctx, req)

	var res pb.OrderRes
	orderID := getOrderID()

	// 1. 处理原商品
	for _, pbGood := range req.Goods {
		// 1.1 判断原商品是否存在
		sGood, err := dao.GoodDao.GetByName(pbGood.Name)
		if err != nil {
			logs.Error(err)
			return nil, errors.New("Fail to finish GoodDao.GetByName")
		}
		if sGood == nil {
			return nil, errors.New("Good not existed")
		}

		// 1.2 插入订单记录
		if err := dao.OrderRecordDao.Create(orderID, int(sGood.ID), model.FlagOfNotAttachGood); err != nil {
			return nil, errors.New("Fail to finish OrderRecordDao.Create")
		}

		thingID := getThingID()
		// 1.3 判断附属商品是否存在
		for _, pbAttachGood := range pbGood.AttachGoods {
			// 1.3.1 判断附属商品是否存在
			good, err := dao.GoodDao.GetByName(pbAttachGood.Name)
			if err != nil {
				logs.Error(err)
				return nil, errors.New("Fail to finish GoodDao.GetByName")
			}
			if good == nil {
				logs.Error("Good not existed")
				return nil, errors.New("Good not existed")
			}
			// 1.3.2 插入附属记录
			if err := dao.AttachRecordDao.Create(orderID, thingID, int(sGood.ID), int(good.ID)); err != nil {
				logs.Error(err)
				return nil, errors.New("Fail to finish AttachRecordDao.Create")
			}

			// 1.3.3 插入订单记录
			if err := dao.OrderRecordDao.Create(orderID, int(good.ID), model.FlagOfAttachGood); err != nil {
				return nil, errors.New("Fail to finish OrderRecordDao.Create")
			}

			// 1.3.4 形成下一个物品号
			nextThingID()
		}

	}

	// 2. 形成下一个订单号
	nextOrderID()

	res.OrderID = int64(orderID)
	return &res, nil
}

func (MvpServer) GetOrderGoods(ctx context.Context, req *pb.GetOrderGoodsReq) (*pb.GetOrderGoodsRes, error) {
	logs.Info("GetOrderGoods", ctx, req)

	var res pb.GetOrderGoodsRes

	// 1. 获取订单的商品
	orderRecords, err := dao.OrderRecordDao.GetByOrderID(int(req.OrderID))
	if err != nil {
		logs.Error(err)
		return nil, errors.New("Fail to finish OrderRecordDao.GetByOrderID")
	}

	// 2. 形成显示商品
	var pbGoods []*pb.Good
	for _, orderRecord := range orderRecords {
		if orderRecord.IsAttachGood == model.FlagOfNotAttachGood {
			// 2.1 获得商品名
			good, err := dao.GoodDao.Get(orderRecord.GoodID)
			if err != nil {
				logs.Error(err)
				return nil, errors.New("Fail to finish GoodDao.Get")
			}
			if good == nil {
				logs.Error("Good not existed")
				return nil, errors.New("Good not existed")
			}

			// 2.2 形成顶层显示商品
			pbGoods = append(pbGoods, &pb.Good{
				Name: good.Name,
			})
		} else {

			// 2.3 判断商品记录是否正确
			if len(pbGoods) == 0 {
				logs.Error("Order record error")
				return nil, errors.New("Order record error")
			}

			// 2.4 获得附属商品名
			good, err := dao.GoodDao.Get(orderRecord.GoodID)
			if err != nil {
				logs.Error(err)
				return nil, errors.New("Fail to finish GoodDao.Get")
			}
			if good == nil {
				logs.Error("Good not existed")
				return nil, errors.New("Good not existed")
			}

			// 2.5 形成附属显示商品
			pbGoods[len(pbGoods)-1].AttachGoods = append(
				pbGoods[len(pbGoods)-1].AttachGoods,
				&pb.AttachGood{
					Name: good.Name,
				},
			)
		}
	}

	// 3. 返回
	res.Goods = pbGoods
	return &res, nil
}

var curOrderID int // 这里不考虑高并发

func getOrderID() int {
	return curOrderID
}

func nextOrderID() {
	curOrderID++
}

var curThingID int // 这里不考虑高并发

func getThingID() int {
	return curThingID
}

func nextThingID() {
	curThingID++
}
