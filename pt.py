import subprocess
import json
import os
import time
from subprocess import Popen, PIPE, STDOUT
from pprint import pprint
from tqdm import tqdm

demoPath = 'C:/Users/Joseph/Documents/My Games/Rocket League/TAGame/Demos/Demos/PreOctober2018/'
#demoPath = 'resources/replays/set of 100/'
#output = Popen(args, stdout=PIPE).communicate()[0].decode()
#pprint(output)



def main():
	count = 0
	#dirs = getReplayFileNames('resources/replays/set of 100/')
	dirs = getReplayFileNames(demoPath)
	epoch = int(time.time())
	for file in tqdm(dirs):
		# pprint(epoch)
		# pprint(count)
		#print('\n\n' + file + '\n')
		#pprint('resources/replays/' + file)
		#args = ['Korva.exe', 'resources/replays/set of 100/' + file]
		rhargs = ['rephead-windows.exe', demoPath + file]
		args = ['Korva.exe', demoPath + file]
		#args = ['./korva', 'resources/replays/' + file]
		#output = Popen(args, stdout=PIPE).communicate()[0].strip()
		output = Popen(args, stdout=PIPE).communicate()[0].decode()
		#_json = json.loads(output)
		#pprint(output)
		#logParseErrors(output, file)
		#pprint(_json)
		
		count = count + 1
	totalTime = time.time() - epoch
	pprint(totalTime)

def getReplayFileNames(replaysDirectory):
	return os.listdir(replaysDirectory)

def logParseErrors(output, file): 
	if len(output) < 200:
		pprint(file + ' breaks everything')

def runMac():
	pass


main()

