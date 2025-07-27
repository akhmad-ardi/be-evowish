# Auth

## Register
  - **POST** `/api/user/register`

### Request Body
```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "password123",
  "confirm_password": "password123"
}
```

### Response Success
```json
{
  "message": "Registration successful",
}
```

### Response Validation Fail
```json
{
  "validation_error": {
    "Email": "must be a valid email",
    "Password": "minimum length is 6",
    "ConfirmPassword": "ConfirmPassword must match Password"
  }
}
```

## Login
  - **POST** `/api/user/login`

### Request Body
```json
{
  "email": "john@example.com",
  "password": "password123",
}
```

### Response Success
```json
{
  "message": "Login successful",
}
```

### Response Validation Fail
```json
{
  "validation_error": {
    "Email": "must be a valid email",
    "Password": "minimum length is 6",
  }
}
```
