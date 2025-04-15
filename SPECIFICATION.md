# Техническое задание

____

## Сводное HTTP API

API системы доступен по пути верхнего уровня `/api`.

### Система работы с пользователем должна предоставлять следующие HTTP-хендлеры:

Система работы с пользователем доступна по пути среднего уровня `/consumers`.

- **POST** `/api/consumers/registration` — регистрация пользователя.
- **POST** `/api/consumers/login` — аутентификация пользователя.
- **DELETE** `/api/consumers/delete` - удаление пользователя.
- **GET** `/api/consumers/get_me` — получение профиля аутентифицированного пользователя.
- **PUT** `/api/consumers/update_password` - изменение пароля пользователя.
- **POST** `/api/consumers/refresh-token` — перевыпуск пары авторизационных токенов.

#### Предполагается:

- **GET** `/api/consumers/all` — получение списка пользователей;

### Система работы с тестами должна предоставлять следующие HTTP-хендлеры:

Система работы с тестами доступна по пути среднего уровня `/tests`.

- **GET** `/api/test/all` - получение всех тестов, которые имеются в системе.
- **GET** `/api/test/:test_id` - получение теста по его идентификатору.

### Система работы с решенными тестами должна предоставлять следующие HTTP-хендлеры:

Система работы с тестами доступна по пути среднего уровня `/resolved`.

### Система работы с результатами тестирования должна предоставлять следующие HTTP-хендлеры:

Система работы с результатами тестирования доступна по пути среднего уровня `/results`.

- **POST** `/api/results/send` - сохранение результата тестирования.
- **GET** `/api/results/:id&:test_id` - получение результата тестирования по идентификатору пользователя и
  идентификатору теста.
- **GET** `/api/results/:result_id` - получение результатов тестирования по идентификатору теста.
- **GET** `/api/results/all/:id` - получение всех результатов тестирования по идентификатору пользователя.

### Описание методов системы для работы с пользователем.

- #### Регистрация пользователя
##### Хендлер: **POST** `/api/consumers/registration`.

Формат запроса:

```json lines
POST /api/consumers/registration HTTP/1.1
Content-Type application/json
...
{
  "login": "<login>",
  "password": "<password>"
} 
```

Формат ответа:

```json lines
HTTP/1.1
Content-Type application/json
...

{
  "access_token": "<access_token>",
  "refresh_token": "<refresh_token>"
}
```

- #### Аутентификация пользователя
##### Хендлер: **POST** `/api/consumers/login`.

Формат запроса:

```json lines
POST /api/consumers/login HTTP/1.1
Content-Type application/json
...

{
  "login": "<login>",
  "password": "<password>"
}
```

Формат ответа:

```json lines
HTTP/1.1
Content-Type application/json
...

{
  "access_token": "<access_token>",
  "refresh_token": "<refresh_token>"
}
```

- #### Удаление пользователя
##### Хендлер: **DELETE** `/api/consumers/delete`.

Формат запроса:

```json lines
DELETE /api/consumers/delete HTTP/1.1
Content-Type application/json
Authorization Bearer <token>
...
```

Формат ответа:

```json lines
HTTP/1.1
Content-Type application/json
...
```

- #### Получение профиля аутентифицированного пользователя.
##### Хендлер: **GET** `/api/consumers/get_me`.

Формат запроса:

```json lines
GET /api/consumers/get_me HTTP/1.1
Content-Type application/json
Authorization Bearer <token>
...
```

Формат ответа:

```json lines
HTTP/1.1
Content-Type application/json
...

{
  "id": "<consumer_id>",
  "login": "<login>",
  "created_at": "<timestamp>",
}
```

- #### Изменение пароля пользователя.
##### Хендлер: **PUT** `/api/consumers/update_password`.

Формат запроса:

```json lines
PUT /api/consumers/update_password HTTP/1.1
Content-Type application/json
Authorization Bearer <token>
...

{
  "old_password": "<old_password>",
  "new_password": "<new_password>"
}
```

Формат ответа:

```json lines
HTTP/1.1
Content-Type application/json
...
```

- #### Перевыпуск пары авторизационных токенов.
##### Хендлер: **POST** `/api/consumers/refresh-token`.

Формат запроса:

```json lines
POST /api/consumers/refresh-token HTTP/1.1
Content-Type application/json
Authorization Bearer <token>
...
```

Формат ответа:

```json lines
HTTP/1.1
Content-Type application/json
...

{
  "access_token": "<access_token>",
  "refresh_token": "<refresh_token>"
}
```

### Описание методов системы для работы с тестами.

- #### Получение всех тестов, которые имеются в системе.
##### Хендлер: **GET** `/api/test/all`.

Формат запроса:

```json lines
GET /api/test/all HTTP/1.1
Content-Type application/json
Authorization Bearer <token>
...
```

Формат ответа:

```json lines
HTTP/1.1
Content-Type application/json
...

{
  "tests": [
    {
      "id": "<id>",
      "name": "<name>",
      "description": "<description",
      "questions": [
        {
          "order": "<order>",
          "question": "<question>"
        },
        ...
      ]
    },
    ...
  ]
}
```

- #### Получение теста по его идентификатору.
##### Хендлер: **GET** `/api/test/:test_id`.

Формат запроса:

```json lines
GET /api/test/{test_id:"<test_id>"} HTTP/1.1
Content-Type application/json
Authorization Bearer <token>
...
```

Формат ответа:

```json lines
HTTP/1.1
Content-Type application/json
...

{
  "id": "<id>",
  "name": "<name>",
  "description": "<description",
  "questions": [
    {
      "order": "<order>",
      "question": "<question>"
    },
    ...
  ]
}
```

### Описание методов системы для работы с результатами тестирования.

- #### Сохранение результата тестирования.
##### Хендлер: **POST** `/api/results/send`.

Формат запроса:

```json lines
POST /api/results/send HTTP/1.1
Content-Type application/json
Authorization Bearer <token>
...
```

Формат ответа:

```json lines
HTTP/1.1
Content-Type application/json
...

{
  "tests": [
    {
      "id": "<id>",
      "name": "<name>",
      "questions": [
        {
          "order": "<order>",
          "question": "<question>"
        },
        ...
      ]
    },
    ...
  ]
}
```

- **GET** `/api/results/:id&:test_id` - получение результата тестирования по идентификатору пользователя и
  идентификатору теста.
- **GET** `/api/results/:result_id` - получение результатов тестирования по идентификатору теста.
- **GET** `/api/results/all/:id` - получение всех результатов тестирования по идентификатору пользователя.