
"""
TypedKeys class
"""
class Input:

  """
  Constructor
  """
  def __init__(self):
    self.text = ""

  """
  Handles key cleanup when it's pressed
  """
  def press(self, key):

    # Remove a character if it's a backspace
    if key == 'Key.backspace':
      self.text = self.text[:-1]

    # Add space if space key was pressed
    elif key == 'Key.space':
      self.text += ' '

    # Add [shift] if shift key was pressed
    elif key == 'Key.shift':
      self.text += '[shift]'

    # Add [capslock] if caps_lock key was pressed
    elif key == 'Key.caps_lock':
      self.text += '[capslock]'

    # Add newline if enter was pressed
    elif key == 'Key.enter':
      self.text += '\n'

    # Otherwise, parse the key
    else:

      # Parsed key variable
      parsed = None

      #'\x01' is unicode so we remove it
      if (len(key) and (key[0:3] != 'Key') and (key[0:3] != '\'\\x')):
        parsed = key[1:-1]

      # Append parsed key to text
      if parsed is not None:

        # Append presed key to text
        self.text += parsed

        return True

    return False

  # Get typed keys
  def getText(self):
    return self.text

  # Clear text
  def clear(self):
    self.text = ""
