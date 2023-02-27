# golang-oauth2-example

An example project how a jwt token can be validated according to the oAuth2 standard.

## TODO's

- create login form with redirect

### Folder Structure:

```bash
.
├── cmd
│   └── main.go
├── go.mod
├── go.sum
└── pkg
    └── auth
        ├── middleware.go
        ├── models.go
        └── token.go
```

### Configure Microsoft Azure

Go to Microsoft Azure Portal: https://portal.azure.com/

1. `App-Registration`
2. select our app or creat a new one
3. `ApplicationID_URI`
4. Add Scope with name `default` -> this will be used [here](https://github.com/fullstack-devops/golang-oauth2-example/blob/main/pkg/auth/middleware.go#L56)