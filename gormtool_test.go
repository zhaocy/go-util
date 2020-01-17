package go_util

import (
	"fmt"
	"testing"
)

func TestAddGameCfgTag(t *testing.T) {
	rs := AddJsonFormGormTag(`
	type GameCfg struct{
		Id        uint32
		Icon string
		Name  string
		Flag int32
	}
	`)
	fmt.Println(rs)
}

func TestAddMatchUserTag(t *testing.T) {
	rs := AddJsonFormGormTag(`
		type MatchUser struct {
			MatchId int64
			Uid     int32
			GameId  int32
			Score   int64
			Data    string
		}
	`)
	fmt.Println(rs)
}

func TestAddMatchCfgTag(t *testing.T) {
	rs := AddJsonFormGormTag(`
		type MatchCfg struct {
			Id        int32
			GameId    int32
			MatchType int32
			MatchName string
			MatchDesc string
			FeeType   byte
			Fee       int64
			WinFee    int64
			WinExp    int32
			LoseExP   int32
			ticketz   int32
			UserNum   int16
			PerTime   int32
			Medal     int32
		}
	`)
	fmt.Println(rs)
}

func TestAddMatchParamTag(t *testing.T) {
	rs := AddJsonFormGormTag(`
		type MatchParam struct{
			GameId int32 
			CfgId  int32 
			ParamKey    string
			ParamValue  string
		}
	`)
	fmt.Println(rs)
}

func TestAddUserGameTag(t *testing.T) {
	rs := AddJsonFormGormTag(`
		type UserGame struct {
			Id int32
			UserId int32
			GameId int32
			Zmoney int64
			Exp int32
			Mscore int64
			WinProb int32
			Lv int32
		}
	`)
	fmt.Println(rs)
}

func TestAddMatchTag(t *testing.T) {
	rs := AddJsonFormGormTag(`
		type Match struct{
			Id int32
			GameId int32
			CfgId int32
			Status int8
			St time.Time
			Et time.Time
			Rd int16
			WinId int32
			Uid1 int32
			Status1 int8
			Score1 int64
			St1 time.Time
			Et1 time.Time
			Uid2 int32
			Status2 int8
			Score2 int64
			St2 time.Time
			Et2 time.Time
			Flag int8
			Seed int32
		}
    `)
	fmt.Println(rs)
}

func TestAddScoreLogTag(t *testing.T) {
	rs := AddJsonFormGormTag(`
		type ScoreLog struct {
			Id        int32
			GameId    int32
			UserId    int32
			MatchType int32
			Score     int64
			Zscore    int64
		}
	`)
	fmt.Println(rs)
}

func TestAddGameExpLvTag(t *testing.T) {
	rs := AddJsonFormGormTag(`
		type GameExpLv struct {
			Id     int32
			GameId int32
			Lv     int32
			Exp    int32
		}
	`)
	fmt.Println(rs)
}

func TestAddTransTag(t *testing.T) {
	rs := AddJsonFormGormTag(`
		type Trans struct {
			Id int32
			UserId int32
			GameId int32
			EvName string
			MatchId int64
			Tm int64
			FeeType int8
			Fee int64
			Balance int64
		}
	`)
	fmt.Println(rs)
}

func TestAddPlatformExpLvTag(t *testing.T) {
	rs := AddJsonFormGormTag(`
		type PlatformExpLv struct {
			Id  int32
			Lv  int32
			Exp int32
		}
	`)
	fmt.Println(rs)
}

func TestAddGameScoreLogTag(t *testing.T) {
	rs := AddJsonFormGormTag(`
		type GameScoreLog struct {
			Id int32
			GameId int32
			UserId int32
			FeeType int8
			Score int64
			Zscore int64
			tm int64
		}
	`)
	fmt.Println(rs)
}

func TestAddBonusCfgTag(t *testing.T) {
	rs := AddJsonFormGormTag(`
		type BonusCfg struct {
			Id int32
			GameId int32
			Zmoney int32
			per int32
		}
	`)
	fmt.Println(rs)
}

func TestAddCashCfgTag(t *testing.T) {
	rs := AddJsonFormGormTag(`
		type CashCfg struct {
			Id int32
			CoinType int8
			Amount int32 
			Added int32
			Sale int8
			Status int8
		}
	`)
	fmt.Println(rs)
}

func TestAddBonusLogTag(t *testing.T) {
	rs := AddJsonFormGormTag(`
		type BonusLog struct {
			Id int32
			GameId int32
			UserId int32
			St int64
			Et int64
			Status int8
		}
	`)
	fmt.Println(rs)
}

func TestAddUserCashTag(t *testing.T) {
	rs := AddJsonFormGormTag(`
		type UserCash struct {
			Id int32
			UserId int32
			Cash int32
			Reward int32
			Ut int64
		}
	`)
	fmt.Println(rs)
}

func TestAddSloganCfgTag(t *testing.T) {
	rs := AddJsonFormGormTag(`
		type SloganCfg struct {
			Id int32
			Slogan string
			zmoney int64
		}
	`)
	fmt.Println(rs)
}

func TestAddUserSloganTag(t *testing.T) {
	rs := AddJsonFormGormTag(`
		type UserSlogan struct {
			Id int32
			GameId int32
			UserId int32
			SloganId int32
			Status int8
		}
	`)
	fmt.Println(rs)
}

func TestAddTaskTag(t *testing.T) {
	rs := AddJsonFormGormTag(`
		type Task struct {
			Id int32
			Icon string
			Title string
			Desc string
			Steps int32
			Zmoney int32
			Cash int32
		}
	`)
	fmt.Println(rs)
}

func TestAddUserTaskTag(t *testing.T) {
	rs := AddJsonFormGormTag(`
		type UserTask struct {
			Id int32
			GameId int32
			TaskId int32
			Step int32
			Tm int64
		}
	`)
	fmt.Println(rs)
}

func TestAddPrizeTag(t *testing.T) {
	rs := AddJsonFormGormTag(`
		type Prize struct {
			Id int32
			ImageUrl string
			Title string
			Desc string
			Ticketz int64
			Status int32
		}
	`)
	fmt.Println(rs)
}

func TestAddUserPrizeTag(t *testing.T) {
	rs := AddJsonFormGormTag(`
		type UserPrize struct {
			Id int32
			PrizeId int32
			PrizeType int32
			ExchangeNum string
			ExchangePrize string
			Tm int64
		}
	`)
	fmt.Println(rs)
}