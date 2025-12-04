curl -X POST http://localhost:8081/books \
  -H "Content-Type: application/json" \
  -d '{
    "title": "<Book Title>",
    "author": "<Autor Name>",
    "ownerId": "<ownerId>",
    "status": "available"
  }'
