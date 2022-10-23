# golang-graphql

a graphql implementation written in golang

1. Run server.go `go run server.go`
2. Open `http://localhost:8080/`

3. To add user

```graph
mutation addUser {
  createUser(input: {name: "your name", class: 2}){
    _id
    name
    class
  }
}
```

4. To get all users

```graph
query getAll {
  users {
    _id
    name
    class
  }
}
```

5. To get user by id

```graph
query getUser {
  user(_id: "id_to_find"){
    _id
    name
    class
  }
}
```

Note: Make sure you have installed MongoDB and run it
The installation steps of MongoDB can be seen here: https://www.mongodb.com/docs/manual/tutorial/
