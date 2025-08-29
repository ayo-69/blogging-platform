# Blogging Platform

A simple blogging platform built with Go.

## Getting Started

### Prerequisites

- Go
- Git

### Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/ayo-69/blogging-platform.git
   ```
2. Navigate to the project directory:
   ```sh
   cd blogging-platform
   ```
3. Install dependencies:
   ```sh
   go mod tidy
   ```

## Usage

Run the application with the following command:

```sh
go run main.go
```

## API Endpoints

The following API endpoints are available:

### Authentication
- `POST /auth/register` => Register a new user (Profile is also created)
- `POST /auth/login` => Login a user

### Posts
- `POST /post` => Create post
- `GET /post` => Get posts
- `GET /post/:id` => Get a post by ID
- `GET /post/user/:id` => Get posts by user ID
- `PUT /post/:id` => Update post by ID
- `DELETE /post/:id` => Delete post by ID

### Comments
- `POST /comments` => Create comment
- `GET /comments/:id` => Get a comment by ID
- `GET /comments/post/:post_id` => Get comments by post ID
- `GET /comments/user/:user_id` => Get comments by user ID
- `PUT /comments/:id` => Update comment by ID
- `DELETE /comments/:id` => Delete comment by ID

### Profile
- `GET /profile/:id` => Get profile by ID
- `GET /profile/user/:user_id` => Get profiles by user ID
- `PUT /profile/:id` => Update profile by ID
- `DELETE /profile/:id` => Delete profile by ID

## Note
More documentation on /

## Folder Structure

```
.
├── config
├── controllers
├── database
├── middleware
├── models
├── routes
└── utils
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

[MIT](https.choosealicense.com/licenses/mit/)
