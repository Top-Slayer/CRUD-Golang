if [ -z "$1" ]; then
	echo "Don't used any command"
elif [ $1 = "-p" ]; then
	curl -X POST http://localhost:8080/ -H "Content-Type: application/json" -d @post.json
elif [ $1 = "-g" ]; then
	if [ -z "$2" ]; then
		curl -X GET http://localhost:8080/ -H "Content-Type: application/json"
	else
		curl -X GET http://localhost:8080/$2 -H "Content-Type: application/json"
	fi
elif [ $1 = "-d" ]; then
  curl -X DELETE http://localhost:8080/$2 -H "Content-Type: application/json"
elif [ $1 = "-h" ]; then
  echo -e "1. -p for POST and you can edit data in post.json file\n2. -g for GET all informations and by id\n\tEx: bash curdCommand -g || bash curdCommand -g {id}\n3. -d for DELETE by id\n\tEx: bash curdCommand -d {id}\n4. -h for help it's display help guide"
fi
