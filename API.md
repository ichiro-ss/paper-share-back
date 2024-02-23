---
title: 個人プロジェクト
language_tabs:
  - shell: Shell
  - http: HTTP
  - javascript: JavaScript
  - ruby: Ruby
  - python: Python
  - php: PHP
  - java: Java
  - go: Go
toc_footers: []
includes: []
search: true
code_clipboard: true
highlight_theme: darkula
headingLevel: 2
generator: "@tarslib/widdershins v4.0.20"

---

# 個人プロジェクト

Base URLs:

* <a href="https://prod.your-api-server.com">本番環境: https://prod.your-api-server.com</a>

# Authentication

- HTTP Authentication, scheme: bearer

# user

## POST ユーザ認証

POST /login

> Body Parameters

```json
{
  "loginId": "string",
  "password": "string"
}
```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|object| no |none|
|» loginId|body|string| yes |none|
|» password|body|string| yes |none|

> Response Examples

> 200 Response

```json
{
  "token": "string"
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### Responses Data Schema

HTTP Status Code **200**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|» token|string|true|none||認証トークン|

## POST ユーザ情報作成

POST /users

ユーザ情報を作成します。
ユーザの名前情報、メールアドレス、パスワードをリクエストで受け取り、ユーザIDと認証用のトークンを生成しデータベースへ保存します。
生成された認証用のトークンがレスポンスとして返されます。

> Body Parameters

```json
{
  "loginId": "string",
  "password": "string"
}
```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|object| no |none|
|» loginId|body|string| yes |none|
|» password|body|string| yes |none|

> Response Examples

> 200 Response

```json
{
  "token": "string"
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### Responses Data Schema

HTTP Status Code **200**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|» token|string|true|none||認証トークン|

## GET ユーザ情報取得

GET /users

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|Authorization|header|string| yes |認証トークン(JWT)|

> Response Examples

> 200 Response

```json
{
  "name": "string"
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### Responses Data Schema

HTTP Status Code **200**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|» name|string|true|none||none|

## PUT ユーザ情報更新

PUT /users

> Body Parameters

```json
{
  "name": "string"
}
```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|Authorization|header|string| no |認証トークン(JWT)|
|body|body|object| no |none|
|» name|body|string| yes |none|

> Response Examples

> 200 Response

```json
{
  "name": "string"
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### Responses Data Schema

HTTP Status Code **200**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|» name|string|true|none||none|

# summary

## GET 要約一覧取得

GET /summaries

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|Authorization|header|string| no |認証トークン(JWT)|

> Response Examples

> 200 Response

```json
{
  "id": 0,
  "userId": 0,
  "title": "string",
  "markdown": "string",
  "isMine": true
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### Responses Data Schema

HTTP Status Code **200**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|» id|integer|true|none||none|
|» userId|integer|true|none||id of user who wrote the summary|
|» title|string|true|none||paper's title|
|» markdown|string|true|none||paper's summary|
|» isMine|boolean|true|none||none|

## POST 要約投稿

POST /summaries

> Body Parameters

```json
{
  "title": "string",
  "markdown": "string"
}
```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|Authorization|header|string| no |認証トークン(JWT)|
|body|body|object| no |none|
|» title|body|string| yes |none|
|» markdown|body|string| yes |none|

> Response Examples

> 200 Response

```json
{}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### Responses Data Schema

## GET 要約取得

GET /summaries/{id}

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|id|path|string| yes |none|
|Authorization|header|string| no |認証トークン(JWT)|

> Response Examples

> 200 Response

```json
{
  "id": 0,
  "userId": 0,
  "title": "string",
  "markdown": "string",
  "isMine": true
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### Responses Data Schema

HTTP Status Code **200**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|» id|integer|true|none||none|
|» userId|integer|true|none||id of user who wrote the summary|
|» title|string|true|none||paper's title|
|» markdown|string|true|none||paper's summary|
|» isMine|boolean|true|none||none|

## PUT 要約更新

PUT /summaries/{id}

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|id|path|string| yes |none|
|Authorization|header|string| no |認証トークン(JWT)|

> Response Examples

> 200 Response

```json
{
  "id": 0,
  "userId": 0,
  "title": "string",
  "markdown": "string",
  "isMine": true
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### Responses Data Schema

HTTP Status Code **200**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|» id|integer|true|none||none|
|» userId|integer|true|none||id of user who wrote the summary|
|» title|string|true|none||paper's title|
|» markdown|string|true|none||paper's summary|
|» isMine|boolean|true|none||none|

## DELETE 要約削除

DELETE /summaries/{id}

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|id|path|string| yes |none|
|Authorization|header|string| no |認証トークン(JWT)|

> Response Examples

> 200 Response

```json
{}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### Responses Data Schema

# Data Schema

<h2 id="tocS_Authors">Authors</h2>

<a id="schemaauthors"></a>
<a id="schema_Authors"></a>
<a id="tocSauthors"></a>
<a id="tocsauthors"></a>

```json
{
  "id": 0,
  "name": "string"
}

```

### Attribute

|Name|Type|Required|Restrictions|Title|Description|
|---|---|---|---|---|---|
|id|integer|true|none||none|
|name|string|true|none||author's name|

<h2 id="tocS_Paper_authors">Paper_authors</h2>

<a id="schemapaper_authors"></a>
<a id="schema_Paper_authors"></a>
<a id="tocSpaper_authors"></a>
<a id="tocspaper_authors"></a>

```json
{
  "id": 0,
  "paperId": 0,
  "userId": 0
}

```

### Attribute

|Name|Type|Required|Restrictions|Title|Description|
|---|---|---|---|---|---|
|id|integer|true|none||none|
|paperId|integer|true|none||none|
|userId|integer|true|none||none|

<h2 id="tocS_Summaries">Summaries</h2>

<a id="schemasummaries"></a>
<a id="schema_Summaries"></a>
<a id="tocSsummaries"></a>
<a id="tocssummaries"></a>

```json
{
  "id": 0,
  "userId": 0,
  "title": "string",
  "markdown": "string"
}

```

### Attribute

|Name|Type|Required|Restrictions|Title|Description|
|---|---|---|---|---|---|
|id|integer|true|none||none|
|userId|integer|true|none||id of user who wrote the summary|
|title|string|true|none||paper's title|
|markdown|string|true|none||paper's summary|

<h2 id="tocS_Authorizations">Authorizations</h2>

<a id="schemaauthorizations"></a>
<a id="schema_Authorizations"></a>
<a id="tocSauthorizations"></a>
<a id="tocsauthorizations"></a>

```json
{
  "loginId": "string",
  "userId": 0,
  "password": "string",
  "createdAt": "string",
  "updatedAt": "string"
}

```

### Attribute

|Name|Type|Required|Restrictions|Title|Description|
|---|---|---|---|---|---|
|loginId|string|true|none||none|
|userId|integer|true|none||none|
|password|string|true|none||none|
|createdAt|string|true|none||none|
|updatedAt|string|true|none||none|

<h2 id="tocS_Tag">Tag</h2>

<a id="schematag"></a>
<a id="schema_Tag"></a>
<a id="tocStag"></a>
<a id="tocstag"></a>

```json
{
  "id": 1,
  "name": "string"
}

```

### Attribute

|Name|Type|Required|Restrictions|Title|Description|
|---|---|---|---|---|---|
|id|integer(int64)|false|none||タグID番号|
|name|string|false|none||タグ名|

<h2 id="tocS_Users">Users</h2>

<a id="schemausers"></a>
<a id="schema_Users"></a>
<a id="tocSusers"></a>
<a id="tocsusers"></a>

```json
{
  "id": 0,
  "name": "string",
  "createdAt": "string",
  "updatedAt": "string",
  "deletedAt": "string"
}

```

### Attribute

|Name|Type|Required|Restrictions|Title|Description|
|---|---|---|---|---|---|
|id|integer|true|none||none|
|name|string¦null|true|none||none|
|createdAt|string|true|none||none|
|updatedAt|string|true|none||none|
|deletedAt|string|true|none||none|

<h2 id="tocS_Category">Category</h2>

<a id="schemacategory"></a>
<a id="schema_Category"></a>
<a id="tocScategory"></a>
<a id="tocscategory"></a>

```json
{
  "id": 1,
  "name": "string"
}

```

### Attribute

|Name|Type|Required|Restrictions|Title|Description|
|---|---|---|---|---|---|
|id|integer(int64)|false|none||カテゴリーID番号|
|name|string|false|none||カテゴリー名|

<h2 id="tocS_Pet">Pet</h2>

<a id="schemapet"></a>
<a id="schema_Pet"></a>
<a id="tocSpet"></a>
<a id="tocspet"></a>

```json
{
  "id": 1,
  "category": {
    "id": 1,
    "name": "string"
  },
  "name": "doggie",
  "photoUrls": [
    "string"
  ],
  "tags": [
    {
      "id": 1,
      "name": "string"
    }
  ],
  "status": "available"
}

```

### Attribute

|Name|Type|Required|Restrictions|Title|Description|
|---|---|---|---|---|---|
|id|integer(int64)|true|none||ペットID番号|
|category|[Category](#schemacategory)|true|none||カテゴリー|
|name|string|true|none||名前|
|photoUrls|[string]|true|none||写真URL|
|tags|[[Tag](#schematag)]|true|none||タグ|
|status|string|true|none||ペット販売状況|

#### Enum

|Name|Value|
|---|---|
|status|available|
|status|pending|
|status|sold|

