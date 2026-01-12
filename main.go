package main

import (
	"machine"
	"machine/usb/hid/keyboard"
	"time"
)

// keys - キー配列を指定する変数
var keys = [7]keyboard.Keycode{
	keyboard.KeyA, keyboard.KeyB, keyboard.KeyC, // 上列
	keyboard.KeyD, keyboard.KeyE, keyboard.KeyF, keyboard.KeyG, // 下列
}

// 例 (Aseprite用)
// var keys = [7]keyboard.Keycode{
//	keyboard.KeyB, keyboard.KeyE, keyboard.KeyG, // Pencil Tool, Eraser Tool, Paint Bucket Tool
//	keyboard.KeyL, keyboard.KeyU, keyboard.KeyX, keyboard.KeyY, // Line Tool, Rectangle Tool, ...
//}

// main - ピンへの入力をキーボードの入力として変換する
func main() {
	// input Config
	d0 := machine.D0
	d1 := machine.D1
	d2 := machine.D2
	d3 := machine.D3
	d4 := machine.D4
	d5 := machine.D5
	d6 := machine.D6
	pins := [7]machine.Pin{d0, d1, d2, d3, d4, d5, d6}
	for _, pin := range pins {
		pin.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	}

	kb := keyboard.Port()

	for {
		for i, pin := range pins {
			if pin.Get() {
				kb.Down(keys[i])
				time.Sleep(20 * time.Millisecond)
				kb.Up(keys[i])
			}
		}
		time.Sleep(100 * time.Millisecond)
	}
}
