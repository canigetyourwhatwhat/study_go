
# These are the set of Shell commands to test APIs

# For creating article
echo "Test Post request to get response using streaming"
curl http://localhost:8080/article -X POST -d '{"title":"test","contents":"test","userName":"bob"}'

echo "Dump data"
mysql -h 127.0.0.1 -u docker sampledb -p < database/init.sql

curl http://localhost:8080/article/2

curl http://localhost:8080/article/nice -X POST -d '{"articleId": 1}'

curl http://localhost:8080/article/all

curl http://localhost:8080/article/comment -X POST -d '{"articleId":2,"message":"it is great"}'