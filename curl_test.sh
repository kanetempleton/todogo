#!/bin/bash

# Test GET /tasks
echo "Testing GET /tasks"
curl http://localhost:8080/tasks
echo ""

# Test POST /tasks
echo "Testing POST /tasks"
curl -X POST -H "Content-Type: application/json" -d '{"title": "New Task", "completed": false}' http://localhost:8080/tasks
echo ""

# Test GET /tasks
echo "Testing GET /tasks"
curl http://localhost:8080/tasks
echo ""

# Test GET /tasks/{id}
echo "Testing GET /tasks/2"
curl http://localhost:8080/tasks/2
echo ""

# Test PUT /tasks/{id}
echo "Testing PUT /tasks/2"
curl -X PUT -H "Content-Type: application/json" -d '{"title": "Updated Task", "completed": true}' http://localhost:8080/tasks/2
echo ""

# Test DELETE /tasks/{id}
echo "Testing DELETE /tasks/{id}"
curl -X DELETE http://localhost:8080/tasks/1
echo ""

# Test GET /tasks
echo "Testing GET /tasks"
curl http://localhost:8080/tasks
echo ""