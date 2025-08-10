# Invitation

## Create Invitation
  - **POST** `/api/invitation/create_invitation`

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
  "name": "string"
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

## Add Data Invitation
  - **POST** `/api/invitation/add_data_invitation/:id_invitation`

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
  "data_invitation": {}
}
```

### Response Success
```json
{
  "message": "Undangan siap dibagikan"
}
```

### Response Validation Fail
```json
{
  "validation_errors": {}
}
```

## Add Background Image
  - **POST** `/api/invitation/add_background_image/:id_invitation`

### Request
- Headers
```json
{
  "Authorization": "Bearer token"
}
```

- Body
Berupa form-data
```
bg_image: file
```

### Response Success
```json
{
  "message": "Latar belakang berhasil ditambahkan"
}
```

### Response Validation Fail
```json
{
  "validation_errors": {}
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

## Delete Invitation
  - **POST** `/api/invitation/delete/:id_invitation`

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
  "message": "Undangan berhasil dihapus",
}
```

### Response Validation Fail
```json
{
  "message_error": ""
}
```

## Guest View
  - **GET** `/api/invitation/guest_view`

### Request
- Body
```json
{
  "id_invitation_link": "string",
  "ip_address": "string",
  "user_agent": "string"
}
```

### Response Success
```json
{
  "message": "string",
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
