

```markdown
# API Documentation

## Register

**Endpoint:** `/api/register`

**Method:** POST

**Description:** Registers a new user.

**Request Body:**
```json
{
    "username": "your_username",
    "password": "your_password",
    "email": "your_email"
}
```

**Response:**
```json
{
    "message": "User registered successfully"
}
```

---

## Login

**Endpoint:** `/api/login`

**Method:** POST

**Description:** Logs in a user.

**Request Body:**
```json
{
    "username": "your_username",
    "password": "your_password"
}
```

**Response:**
```json
{
    "token": "your_access_token"
}
```

---

## Create Checklist

**Endpoint:** `/api/checklist`

**Method:** POST

**Description:** Creates a new checklist.

**Request Body:**
```json
{
    "name": "checklist_name",
    "description": "checklist_description"
}
```

**Response:**
```json
{
    "message": "Checklist created successfully"
}
```

---

## Get All Checklists

**Endpoint:** `/api/checklist`

**Method:** GET

**Description:** Retrieves all checklists.

**Response:**
```json
{
    "checklists": [
        {
            "id": 1,
            "name": "checklist_name",
            "description": "checklist_description"
        },
        ...
    ]
}
```

---

## Delete Checklist

**Endpoint:** `/api/checklist/{checklistId}`

**Method:** DELETE

**Description:** Deletes a checklist by ID.

**Response:**
```json
{
    "message": "Checklist deleted successfully"
}
```

---

## Create Checklist Item

**Endpoint:** `/api/checklist/{checklistId}/item`

**Method:** POST

**Description:** Creates a new checklist item in a checklist.

**Request Body:**
```json
{
    "itemName": "item_name"
}
```

**Response:**
```json
{
    "message": "Checklist item created successfully"
}
```

---

## Get All Checklist Items

**Endpoint:** `/api/checklist/{checklistId}/item`

**Method:** GET

**Description:** Retrieves all checklist items in a checklist.

**Response:**
```json
{
    "items": [
        {
            "id": 1,
            "checklist_id": 1,
            "item_name": "item_name",
            "status": "status"
        },
        ...
    ]
}
```

---

## Get Checklist Item

**Endpoint:** `/api/checklist/{checklistId}/item/{checklistItemId}`

**Method:** GET

**Description:** Retrieves a checklist item by ID.

**Response:**
```json
{
    "id": 1,
    "checklist_id": 1,
    "item_name": "item_name",
    "status": "status"
}
```

---

## Update Checklist Item Status

**Endpoint:** `/api/checklist/{checklistId}/item/{checklistItemId}`

**Method:** PUT

**Description:** Updates the status of a checklist item by ID.

**Request Body:**
```json
{
    "status": "new_status"
}
```

**Response:**
```json
{
    "message": "Checklist item status updated successfully"
}
```

---

## Rename Checklist Item

**Endpoint:** `/api/checklist/{checklistId}/item/rename/{checklistItemId}`

**Method:** PUT

**Description:** Renames a checklist item by ID.

**Request Body:**
```json
{
    "itemName": "new_item_name"
}
```

**Response:**
```json
{
    "message": "Checklist item renamed successfully"
}
```
