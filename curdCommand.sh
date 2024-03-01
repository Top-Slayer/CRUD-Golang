if [ -z "$1" ]; then
	echo "No command used"
elif [ $1 = "-p" ]; then
	curl -X POST http://localhost:8080/person -H "Content-Type: application/json" -d '{
    "id": "123",
    "name": "John Doe",
    "gmail": "johndoe@gmail.com",
    "phoneNumber": 1234567890
  }'
elif [ $1 = "-g" ]; then
	if [ -z "$2" ]; then
		curl -X GET http://localhost:8080/person -H "Content-Type: application/json"
	else
		curl -X GET http://localhost:8080/person/$2 -H "Content-Type: application/json"
	fi
fi
