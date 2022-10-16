


# For creating article
echo "Test Post request to get response using streaming"
curl http://localhost:8080/article -X POST -d '{"title":"test","contents":"test","userName":"bob"}'

echo "Create Table"
mysql -h 127.0.0.1 -u docker sampledb -p < database/createTable.sql

echo "Populate Data"
mysql -h 127.0.0.1 -u docker sampledb -p < database/init.sql