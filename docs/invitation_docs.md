# Invitation

## Create
  - **POST** `/api/invitation/create`

### Request
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

## Get Invitations
  - **GET** `/api/invitation`

### Request
- Headers
```json
{
  "Authorization": "Bearer token"
}
```

### Response Success
```json
{
  "invitations": [],
}
```

## Generate Link Invitation
  - **POST** `/api/invitation/generate_link`

### Request
- Headers
```json
{
  "Authorization": "Bearer token"
}
```

- Body
```json
{
  "id_invitation": "string"
}
```

### Response Success
```json
{
  "link": "string",
}
```

## Share Socia Media
  - **POST** `/api/invitation/share_social_media`

### Request
- Headers
```json
{
  "Authorization": "Bearer token"
}
```

- Body
```json
{
  "id_invitation": "string",
  "name_platform": "string"
}
```

### Response Success
```json
{
  "link": "string",
}
```