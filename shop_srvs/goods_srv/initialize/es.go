package initialize

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"go.uber.org/zap"
	"log"
	"os"
	"shop_srvs/goods_srv/global"
	"shop_srvs/goods_srv/model"
	"strconv"
)

func InitEs() {
	host := fmt.Sprintf("http://%s:%d", global.ServerConfig.EsInfo.Host, global.ServerConfig.EsInfo.Port)
	logger := log.New(os.Stdout, "shop", log.LstdFlags)
	var err error
	global.EsClient, err = elastic.NewClient(elastic.SetURL(host), elastic.SetSniff(false), elastic.SetTraceLog(logger))
	if err != nil {
		zap.S().Panic(err)
	}

	//新建mapping和index
	exists, err := global.EsClient.IndexExists(model.EsGoods{}.GetIndexName()).Do(context.Background())
	if err != nil {
		zap.S().Panic(err)
	}
	if !exists {
		_, err = global.EsClient.CreateIndex(model.EsGoods{}.GetIndexName()).BodyString(model.EsGoods{}.GetMapping()).Do(context.Background())
		if err != nil {
			zap.S().Panic(err)
		}
	}

	// 将mysql数据导入es
	var goods []model.Goods
	global.DB.Find(&goods)
	for _, g := range goods {
		esModel := model.EsGoods{
			ID:          g.ID,
			CategoryID:  g.CategoryID,
			BrandsID:    g.BrandsID,
			OnSale:      g.OnSale,
			ShipFree:    g.ShipFree,
			IsNew:       g.IsNew,
			IsHot:       g.IsHot,
			Name:        g.Name,
			ClickNum:    g.ClickNum,
			SoldNum:     g.SoldNum,
			FavNum:      g.FavNum,
			MarketPrice: g.MarketPrice,
			GoodsBrief:  g.GoodsBrief,
			ShopPrice:   g.ShopPrice,
		}
		_, err := global.EsClient.Index().Index(esModel.GetIndexName()).BodyJson(esModel).Id(strconv.Itoa(int(g.ID))).Do(context.Background())
		if err != nil {
			zap.S().Panic(err)
		}
		//强调一下 一定要将docker启动es的java_ops的内存设置大一些 否则运行过程中会出现 bad request错误
	}
}
