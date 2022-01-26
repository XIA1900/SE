#### Required Packages
- gin (web engine)
- gorm (SQL database)
  - gorm mysql driver
- go-yaml/yaml (config)
- log
  - go.uber.org/zap (log)
  - gopkg.in/natefinch/lumberjack.v2 (log files cutting or backup etc.)



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
  │  go.mod
  │  go.sum
  │  main.go // program enter point
  │  README.md
  ├─bin
  │      Small_Demo_Backend.exe
  ├─config
  │      application.yaml // application config file
  │      LoadAppConfig.go // load config parameters
  │
  ├─model
  │      InitDB.go // data source
  │      UserDao.go // user table data access operation
  │
  ├─router
  │      InitRouter.go // router initializer
  │      UserController.go // controller for user operation
  │
  └─service
          UserService.go // user s
  ```

  