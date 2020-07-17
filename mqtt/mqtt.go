/*
 * @Description: 备注
 * @Author: wangcong0918
 * @Date: 2019-11-29 18:08:04
 * @LastEditTime: 2020-07-16 15:38:53
 * @LastEditors: wangcong0918
 */
// mqtt client service
package mqtt

import (
	"fmt"

	"github.com/wangcong0918/sunrise/surgemq/surgemq/service"

	"github.com/wangcong0918/sunrise/surgemq/message"
)

type MQTT interface {
	GetMqttClient() (*service.Client, error)
}

type mqtt struct {
	ConnectUrl   string
	WillQos      int // 如果遗嘱标志被设置为1，遗嘱QoS的值可以等于0(0x00)，1(0x01)，2(0x02)
	Version      int // 0x03 MQIsdp 0x04 MQTT
	CleanSession bool
	KeepAlive    uint16 // 链接存活时间
	WillTopic    string // 遗嘱标志
	WillMessage  string // 遗嘱信息
	Username     string // 用户名
	Password     string // 密码
	Port         int    // 端口
	ClientId     string // 链接ID
}

func NewMqtt(m mqtt) mqtt {
	return m
}

func (m mqtt) GetMqttClient() (*service.Client, error) {
	conn := &service.Client{}
	msg := message.NewConnectMessage()

	msg.SetWillQos(byte(m.WillQos)) // 如果遗嘱标志被设置为1，遗嘱QoS的值可以等于0(0x00)，1(0x01)，2(0x02)
	msg.SetVersion(byte(m.Version))
	msg.SetCleanSession(m.CleanSession)
	msg.SetClientId([]byte(m.ClientId))
	if m.KeepAlive != 0 {
		msg.SetKeepAlive(m.KeepAlive)
	}
	msg.SetWillTopic([]byte(m.WillTopic)) // 遗嘱标志
	msg.SetWillMessage([]byte(m.WillMessage))
	msg.SetUsername([]byte(m.Username))
	msg.SetPassword([]byte(m.Password))
	connUrl := fmt.Sprintf("tcp://%s:%d", m.ConnectUrl, m.Port)
	err := conn.Connect(connUrl, msg)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
