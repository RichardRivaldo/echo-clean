# echo-clean

Echo Clean Architecture

# Description

Straightforward Online Shop backend application with simplified Clean Architecture. Created with Go Echo - MySQL - GORM.

# Setup

## Dependencies

1. Add environment variables to `.env` file as usual. Below are example of the configurations.

```
DB_USER=<db_username>
DB_PASS=<db_password>
DB_HOST=<db_host>
DB_PORT=<db_port>
DB_NAME=<db_name>
DB_TYPE=<db_type>

PORT=<server_port>
```

2. You can migrate the testing database with the provided `src/schema/migrate.sql` script or initialize an empty database with defined tables through `src/schema/database.sql` script. This project is tested with MariaDB-flavoured MySQL (v8.0.27).

3. Simply execute `go build` or `go run .` to install all dependencies from `go mod` and run the server.

# API Documentation

`root`: `/api/v1`

## Members

`root`: `/members`

### Creating New Member

```
POST /
{
    "username": string,
    "gender": enum("Male", "Female"),
    "skin_type": enum("Normal", "Sensitive", "Dry", "Oily"),
    "skin_color": enum("White", "Yellow", "Black")
}
```

### Update Member Data

```
PUT /:member_id
{
    "username": string,
    "gender": enum("Male", "Female"),
    "skin_type": enum("Normal", "Sensitive", "Dry", "Oily"),
    "skin_color": enum("White", "Yellow", "Black")
}
```

### Delete Member Data

```
DELETE /:member_id
```

### Get All Members

```
GET /
```

## Product Review

`root`: `/reviews`

### Get Product Reviews

```
GET /:product_id
```

## Review Interaction

`root`: `/like_reviews`

### Like a Review

```
POST /like
{
    "review_id": string,
    "member_id": string
}
```

### Unlike a Review

```
POST /unlike
{
    "review_id": string,
    "member_id": string
}
```

# Things to Improve

1. After trying it, `GORM`, in my opinion, is not really a good ORM library for Go. Will need to try another ORM libraries.
2. Several APIs might lack clarities in the response due to the ORM.
