package main

import "fmt"

/**
事件的发布
*/
var eventMap = make(map[string][]func(interface{}))

func RegisterEvent(name string, callback func(interface{})) {
	list := eventMap[name]
	list = append(list, callback)
	eventMap[name] = list
}

func CallEvent(name string, params interface{}) {
	list := eventMap[name]
	for _, callback := range list {
		callback(params)
	}
}

type Player struct {
	hp  int
	atk int
	def int
}

func (p Player) heal(params interface{}) {
	amount := params.(int)
	p.hp = p.hp + amount
	println("受到了治疗，当前血量为:", p.hp)
}

func GlobalHeal(params interface{}) {
	fmt.Printf("全局治疗:%d", params)
}

func main() {
	player := &Player{}
	RegisterEvent("heal", player.heal)
	RegisterEvent("heal", GlobalHeal)
	CallEvent("heal", 1000)
}
