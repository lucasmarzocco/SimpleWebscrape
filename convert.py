import json

#Mini file to convert a file of URLs to a list to pass into Postman
def convertToArray():

	arr = []

	for line in open("urls_100.txt"):

		arr.append(line.rstrip())
		
	print(json.dumps(arr))

convertToArray()


