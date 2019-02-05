import subprocess
import json
import os
from subprocess import Popen, PIPE, STDOUT
from pprint import pprint

args = ['rephead-windows.exe', 'resources/replays/2016_09.replay']
#output = Popen(args, stdout=PIPE).communicate()[0].decode()
#pprint(output)

dirs = os.listdir('resources/replays/')
for file in dirs:
	#print('\n\n' + file + '\n')
	#pprint('resources/replays/' + file)
	args = ['korva.exe', 'resources/replays/' + file]
	#output = Popen(args, stdout=PIPE).communicate()[0]
	output = Popen(args, stdout=PIPE).communicate()[0].decode()
	#_json = json.loads(output)

	#pprint(_json)

	if len(output) < 200:
		pprint(file + ' breaks everything')

	print(output)



