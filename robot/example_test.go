package robot_test

import (
	"github.com/golang/protobuf/proto"
	"github.com/thkhxm/tgf/robot"
	"github.com/thkhxm/tgf/tgf-tutorial/common/api"
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

	rb.RegisterCallbackMessage(api.HelloWorld.MessageType, func(i robot.IRobot, bytes []byte) {
		resp := &pb.HelloWorldResponse{}
		proto.Unmarshal(bytes, resp)
		t.Log(resp.Message)
	}).RegisterCallbackMessage(api.Login.MessageType, func(i robot.IRobot, bytes []byte) {
		resp := &pb.LoginResponse{}
		proto.Unmarshal(bytes, resp)
		if resp.Success {
			i.Send(api.HelloWorld.MessageType, &pb.HelloWorldRequest{Name: "tgf"})
		} else {
			t.Log("login fail")
		}
	})
	//

	rb.Send(api.Login.MessageType, &pb.LoginRequest{
		Account:  "admin",
		Password: "123",
	})

	select {}
}

func TestRPCRobot(t *testing.T) {
	rb := CreateRobot()

	rb.RegisterCallbackMessage(api.Login.MessageType, func(i robot.IRobot, bytes []byte) {
		resp := &pb.LoginResponse{}
		proto.Unmarshal(bytes, resp)
		if resp.Success {
			i.Send(api.LoadUserData.MessageType, &pb.LoadUserDataRequest{})
		} else {
			t.Log("login fail")
		}
	}).RegisterCallbackMessage(api.LoadUserData.MessageType, func(i robot.IRobot, bytes []byte) {
		resp := &pb.LoadUserDataResponse{}
		proto.Unmarshal(bytes, resp)
		t.Log(resp.Name, resp.PropCount)
	})

	rb.Send(api.Login.MessageType, &pb.LoginRequest{
		Account:  "admin",
		Password: "123",
	})

	select {}
}
