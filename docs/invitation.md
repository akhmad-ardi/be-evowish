# Invitation

## Create
  - **POST** `/api/invitation/create`

### Request Body
- Headers
```json
{
  "Authorization": "Bearer token"
}
```

- Body
```json
{
  "id_template": "string",
  "title": "string",
  "date": "2006-01-02",
  "time": "15:04:05",
  "location": "string",
  "description": "string",
  "primary_color": "",
  "secondary_color": ""
}
```

### Response Success
```json
{
  "message": "Undangan berhasil dibuat",
}
```

### Response Validation Fail
```json
{
  "validation_error": {
    "title": "Judul acara wajib diisi",
    "date": "Tanggal wajib diisi",
  }
}
```