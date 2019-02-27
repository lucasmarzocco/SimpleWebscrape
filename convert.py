import json

def convertToArray():

	arr = []

	for line in open("urls_100.txt"):

		arr.append(line.rstrip())
		
	print(json.dumps(arr))

convertToArray()


