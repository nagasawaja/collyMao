package Models

import "collyMao/orm"

type TableMaoGames struct {
	Id     int64
	GameId int64 `gorm:"game_id"`
	Title  string
	Url    string
}

func (h *TableMaoGames) TableName() string {
	return "mao_games"
}

func (mg *TableMaoGames) Create() int64 {
	var dbResult = orm.Gorm.Create(mg)
	return dbResult.RowsAffected
}

type TableMaoGamesGoodsCount struct {
	Id         int64  `gorm:"id"`
	GameId     int64  `gorm:"game_id"`
	CreateDate string `gorm:"create_date"`
	GoodsCount int64  `gorm:"goods_count"`
}

func (t *TableMaoGamesGoodsCount) TableName() string {
	return "mao_games_goods_count"
}

func (mggc *TableMaoGamesGoodsCount) Create() int64 {
	var dbResult = orm.Gorm.Create(mggc)
	return dbResult.RowsAffected
}

type TableMaoGamesGoods struct {
	Id         int64
	GameId     int64  `gorm:"game_id"`
	GoodsId    int64  `gorm:"goods_id"`
	Url        string `gorm:"url"`
	SellerName string
	SellerType int64
}

func (t *TableMaoGamesGoods) TableName() string {
	return "mao_games_goods"
}

type TableMaoGamesGoodsDetail struct {
	Id             int64
	CreateDatetime string  `gorm:"create_datetime"`
	Price          float64 `gorm:"price"`
	GoodsCount     int64   `gorm:"goods_count"`
	Title          string  `gorm:"title"`
	GoodsId        int64   `gorm:"goods_id"`
}

func (t *TableMaoGamesGoodsDetail) TableName() string {
	return "mao_games_goods_detail"
}

type TableMaoGamesStc struct {
	Id              int64
	SaleCount       int64
	GoodsTotalCount int64
	Stc             float64
	Title           string
	Url             string
	GameId          int64
	CreateDatetime  string `gorm:"create_datetime"`
}

func (t *TableMaoGamesStc) TableName() string {
	return "mao_games_stc"
}
