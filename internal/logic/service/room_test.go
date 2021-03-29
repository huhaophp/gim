package service

import (
	"fmt"
	"gim/internal/logic/cache"
	"gim/pkg/db"
	"gim/pkg/pb"
	"gim/pkg/util"
	"testing"
	"time"
)

func Test_roomService_DelExpireMessage(t *testing.T) {
	db.InitByTest()
	err := RoomService.DelExpireMessage(1)
	fmt.Println(err)
}

func Test_roomService_List(t *testing.T) {
	db.InitByTest()
	msgs, err := cache.RoomMessageCache.List(1, 1)
	fmt.Println(err)
	fmt.Println(msgs)
}

func Test_roomService_AddMessage(t *testing.T) {
	db.InitByTest()

	for i := 1; i <= 20; i++ {
		RoomService.AddMessage(1, pb.Message{
			Seq:      int64(i),
			SendTime: util.UnixMilliTime(time.Now()),
		})
		time.Sleep(time.Second)
		fmt.Println(i)
	}

	err := RoomService.DelExpireMessage(1)
	fmt.Println(err)
}
