# SmartCampus Workbench Backend API Documentation

## Authentication
- `POST /api/auth/login` — User login
  - Request: `{ "username": string, "password": string }`
  - Response: `{ "token": string, "user": object }`

## Users
- `GET /api/users` — List all users
- `GET /api/users/:id` — Get user by ID
- `POST /api/users` — Create user
- `PUT /api/users/:id` — Update user
- `DELETE /api/users/:id` — Delete user

## Courses
- `GET /api/courses` — List all courses
- `GET /api/courses/:id` — Get course by ID
- `POST /api/courses` — Create course
- `PUT /api/courses/:id` — Update course
- `DELETE /api/courses/:id` — Delete course
- `GET /api/courses/:id/students` — List students in a course

## Assignments
- `GET /api/assignments` — List all assignments
- `GET /api/assignments/:id` — Get assignment by ID
- `POST /api/assignments` — Create assignment
- `PUT /api/assignments/:id` — Update assignment
- `DELETE /api/assignments/:id` — Delete assignment

## Submissions
- `GET /api/submissions` — List all submissions
- `GET /api/submissions/:id` — Get submission by ID
- `POST /api/submissions` — Create submission
- `PUT /api/submissions/:id` — Update submission
- `DELETE /api/submissions/:id` — Delete submission

## Messages
- `GET /api/messages` — List all messages
- `GET /api/messages/:id` — Get message by ID
- `POST /api/messages` — Create message
- `PUT /api/messages/:id` — Update message
- `DELETE /api/messages/:id` — Delete message

## Files
- `POST /api/files/upload` — Upload file
- `GET /api/files/:id` — Get file by ID
- `DELETE /api/files/:id` — Delete file

---

### Common Response Format
```
{
  "success": true/false,
  "data": object, // on success
  "error": string // on error
}
```

### Auth Required
Most endpoints require JWT authentication in the `Authorization` header.

### Notes
- All requests and responses use JSON.
- Pagination, filtering, and sorting can be added as query parameters where applicable.
- Error codes: 400 (Bad Request), 401 (Unauthorized), 403 (Forbidden), 500 (Server Error)
