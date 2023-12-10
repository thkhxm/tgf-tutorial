package robot_test

import (
	"github.com/golang/protobuf/proto"
	"github.com/thkhxm/tgf/robot"
	"github.com/thkhxm/tgf/tgf-tutorial/common/pb"
	"testing"
)

//***************************************************
//@Link  https://github.com/thkhxm/tgf
//@Link  https://gitee.com/timgame/tgf
//@QQ群 7400585
//author tim.huang<thkhxm@gmail.com>
//@Description
//2023/12/10
//***************************************************

func CreateRobot() robot.IRobot {
	rb := robot.NewRobotWs("/tgf")
	rb.Connect("127.0.0.1:8443")
	return rb
}

func TestHelloWorld(t *testing.T) {
	rb := CreateRobot()
	req := &pb.HelloWorldRequest{Name: "tgf"}

	rb.RegisterCallbackMessage("hall.HelloWorld", func(i robot.IRobot, bytes []byte) {
		resp := &pb.HelloWorldResponse{}
		proto.Unmarshal(bytes, resp)
		t.Log(resp.Message)
	})
	//

	rb.SendMessage("hall", "HelloWorld", req)

	select {}
}
