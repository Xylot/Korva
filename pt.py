import subprocess
import json
import os
import time
from subprocess import Popen, PIPE, STDOUT
from pprint import pprint

args = ['rephead-windows.exe', 'resources/replays/2016_09.replay']
#output = Popen(args, stdout=PIPE).communicate()[0].decode()
#pprint(output)



def main():
	count = 0
	dirs = getReplayFileNames('resources/replays/')
	for file in dirs:
		epoch = int(time.time())
		# pprint(epoch)
		# pprint(count)
		print('\n\n' + file + '\n')
		#pprint('resources/replays/' + file)
		#args = ['parser.exe', 'resources/replays/' + file]
		args = ['./korva', 'resources/replays/' + file]
		#output = Popen(args, stdout=PIPE).communicate()[0].strip()
		output = Popen(args, stdout=PIPE).communicate()[0].decode()
		#_json = json.loads(output)
		pprint(output)
		logParseErrors(output, file)
		#pprint(_json)
		
		count = count + 1

def getReplayFileNames(replaysDirectory):
	return os.listdir(replaysDirectory)

def logParseErrors(output, file): 
	if len(output) < 200:
		pprint(file + ' breaks everything')

def runMac():
	pass


main()

