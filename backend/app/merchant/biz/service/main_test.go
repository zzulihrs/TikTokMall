package service

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/tiktokmall/backend/app/merchant/biz/dal"
	"github.com/tiktokmall/backend/app/merchant/biz/dal/mysql"
)

func TestMain(m *testing.M) {
	os.Chdir("../../")
	_ = godotenv.Load()
	dal.Init()
	os.Exit(m.Run())
}

func TestDB(t *testing.T) {
	log.Printf("%v", mysql.DB)
}
