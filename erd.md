# ER diagram

```mermaid
---
title: "PAPER_SHARE"
---
erDiagram
    Users ||--|| Authorization : ""
    Users }|--o{ Summaries: ""
    Paper_Authors }|--|{ Summaries: ""
    Paper_Authors }|--|{ Authors: ""

    Users {
        bigint id PK "user id"
        varchar name "name"
        timestamp created_at "created datetime"
        timestamp updated_at "updated datetime"
        timestamp deleted_at "deleted datetime"
    }
    Authorization {
        varchar login_id PK "login id"
        bigint user_id FK "users.id"
        varchar password "password"
        timestamp created_at "created datetime"
        timestamp updated_at "updated datetime"
    }
    Summaries {
        bigint id PK "summary id"
        bigint user_id FK "users.id"
        varchar title "paper title"
        text summary_md "summary markdown"
    }
    Paper_Authors {
        bigint id PK "paper-author id"
        bigint paper_id FK "paper(summary) id"
        bigint author_id FK "author id"
    }
    Authors {
        bigint id PK "author id"
        varchar name "author name"
    }
```
