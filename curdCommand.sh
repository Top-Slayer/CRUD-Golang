if [ $1 = "-p" ]; then
	curl -X POST http://localhost:8080/person -H "Content-Type: application/json" -d '{
    "id": "123",
    "name": "John Doe",
    "gmail": "johndoe@gmail.com",
    "phoneNumber": 1234567890
  }'
elif [ $1 = "-g" ]; then
	curl http://localhost:8080/person
else
	echo "No command used"
fi
