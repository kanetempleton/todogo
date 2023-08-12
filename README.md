# RESTful Task Manager API

Backend: Go

# Featues
- RESTful API for task manager, test via curl or main_test.go (details below)
- Test file to test the functionality of all REST methods

# TODO
- implement API keys
- user-specific tasks
- frontend component (javascript)


## How to run

1. Clone the repository:
```
git clone https://github.com/yourusername/todogo.git
cd todogo
```

2a. Build + test + run by running the run.sh script:
```
./run.sh
```

2b. Otherwise, build, test, and run individually:
```
go build
go test
./todogo
```

3. (optional) Manual Test with curl:
Run the following script while the server is running and inspect the output to make sure the data is behaving properly:
```
./curl_test.sh
```



## API Endpoints

- `GET /tasks` - Retrieve all tasks.
- `POST /tasks` - Create a new task.
- `GET /tasks/{id}` - Retrieve a task by its ID.
- `PUT /tasks/{id}` - Update a task by its ID.
- `DELETE /tasks/{id}` - Delete a task by its ID.