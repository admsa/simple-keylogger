import threading, time

DELAY=1
DEFAULT_INTERVAL=3

"""
Thread class
"""
class Thread(threading.Thread):

  """
  Constructor
  """
  def __init__(self, callback=None):

    # Call main thread constructor
    threading.Thread.__init__(self, daemon=True)

    # Initialize values
    self.callback = callback
    self.interval = DEFAULT_INTERVAL

  """
  Run thread
  """
  def run(self):

    # Loop through given interval
    while self.interval:
      self.interval -= 1
      time.sleep(DELAY)

    # Call callback if there's any
    if self.callback is not None:
      self.callback()

  """
  Reset thread interval
  """
  def resetInterval(self):
    self.interval = DEFAULT_INTERVAL
