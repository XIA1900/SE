### Roles

There are three roles in this system, **admin**, **regular** and **visitor**. Only authorization data about **admin** and **regular** will be recorded in authorization management table. After users register, they will be assigned **admin** or **regular** role. **Visitor** can be considered as users without login.

### Roles Authorization

Base Path: `/gf/api`

|           | admin | regular            | visitor            |
| --------- | ----- | ------------------ | ------------------ |
| /user     |       |                    |                    |
| /register |       |                    | :heavy_check_mark: |
| /login    |       |                    | :heavy_check_mark: |
| /logout   |       | :heavy_check_mark: |                    |
| /password |       | :heavy_check_mark: |                    |

- If **visitor** has :heavy_check_mark:, meaning this request has no interceptor.
- If **regular** has :heavy_check_mark:, meaning this request only **admin** and **regular** role can request.
- If **admin** has :heavy_check_mark:, meaning this request only **admin **role can request.