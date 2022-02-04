In this file, we will briefly talk about what we've accomplished.


# Frontend
 1. We have found a library(@ant-design-pro) which we think is modifiable and can be used by our Gator Forum.
 2. We have set English as default language.
 3. After login, user profile page is set as default page at present. But later, we will change this default page as the main page of our forum.
 4. Loading page and login page have been modified.






# Backend

## Related Functions Decision

- User Authentication & Authorization
- Save & Search Articles
- Private cloud space management
- Cache Information (Related User Information)

## Database Design

- Discuss Database Schema
- Tables Definition in backend branch "tables.sql"

## Components Combination

Components Combination in "SmalleDemo" (in backend branch) for Sprint1

- **gin** for web server
- **gorm** from **MySQL** database operation
- **Redis** for cache
- uber/**zap** for log
- **JWT** for user authentication
- Configuration Information stored in yaml file (Load when Server Starting)

## Implemented Some Functions

- Load Configuration for different components
  - load configuration information from yaml file when starting server

- CRUD User Table in DB

- User Authentication 

  - Generation, Verification, Refreshing  of JWT token

- User Login/Register Request/Response

  - with data wrote into database
  - related API in GitHub Wiki

- Logging Component

  - different log level

  
