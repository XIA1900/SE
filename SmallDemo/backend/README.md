#### Required Packages
- gin (web engine)
- gorm (SQL database)
  - gorm mysql driver
- go-yaml/yaml (config)
- log
  - go.uber.org/zap (log)
  - gopkg.in/natefinch/lumberjack.v2 (log files cutting or backup etc.)
- JWT



#### Small Demo

- Test API:

  - Login POST:

    - http://localhost:10010/user/login

    - request parameters

      ```JSON
      {
          username: "",
          password: ""
      }
      ```

      Fixed Data(in code) Test

      - username: jake16
      - password: 12345

  - Register POST

    - http://localhost:10010/user/register

    - request parameters

      ```JSON
      {
          username: "jake21",
          password: "12345"
      }
      ```

      Insert Data to DB

- Project Architecture

  ```
  ./SmallDemo/backend
  ├── README.md
  ├── cache
  │   └── InitRedis.go
  ├── component
  │   └── jwt
  │       ├── GenerateToken.go
  │       ├── InitJWT.go
  │       ├── RefreshToken.go
  │       └── VerifyToken.go
  ├── config
  │   ├── LoadAppConfig.go
  │   └── application-demo.yaml
  ├── go.mod
  ├── go.sum
  ├── log
  │   └── InitLog.go
  ├── main.go
  ├── model
  │   ├── InitDB.go
  │   └── UserDao.go
  ├── router
  │   ├── InitRouter.go
  │   └── UserController.go
  └── service
      └── UserService.go
  ```

  