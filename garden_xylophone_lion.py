#imports
import sys
import os

# Define variables
days_since_refresh = 0
refresh_needed = False

# Functions 
def has_refresh_been_done():
  global days_since_refresh
  if days_since_refresh == 0:
    return False
  else:
    return True

def check_for_refresh_needed():
  global days_since_refresh
  global refresh_needed
  if days_since_refresh >= 7:
    refresh_needed = True

def creative_refresh():
  global refresh_needed
  global days_since_refresh
  if refresh_needed:
    # Perform refresh activities
    os.system('ls') # debugging only
    days_since_refresh = 0
    refresh_needed = False

# Program logic
def main():
  if has_refresh_been_done():
    days_since_refresh += 1
  check_for_refresh_needed()
  creative_refresh()

# Main program
if __name__ == '__main__':
  main()
  sys.exit(0)