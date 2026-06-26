এই function **`entity.go`-তে রাখা উচিত নয়**, কারণ এতে `event` package → `dto` package dependency তৈরি হয়। পরে যদি `dto` আবার `event` import করে, তাহলে import cycle হওয়ার সম্ভাবনা থাকে।

# What is import Cycle

**Import Cycle** মানে দুই বা তার বেশি Go package একে অপরকে **ঘুরে ফিরে (circular)** import করছে।

## Example 1 (Import Cycle)

```text
event
   │
   ▼
dto
   │
   ▼
event
```

`event/service.go`

```go
package event

import "myapp/internal/event/dto"
```

`dto/request.go`

```go
package dto

import "myapp/internal/event"
```

এখন dependency হলো:

```text
event → dto → event
```

Go Compiler বুঝতে পারে না কোন package আগে compile করবে, তাই error দেয়:

```text
import cycle not allowed
```

---

## Example 2 (৩টি Package)

```text
handler
   │
   ▼
service
   │
   ▼
repository
   │
   ▼
handler
```

এটাও import cycle।

---

## Example 3 (No Cycle ✅)

```text
handler
   │
   ▼
service
   │
   ▼
repository
```

এখানে dependency একদিকে যাচ্ছে, তাই কোনো সমস্যা নেই।

---

# কেন Go Import Cycle Allow করে না?

ধরো,

```text
Package A
    ↓
Package B
    ↓
Package A
```

এখন প্রশ্ন হলো:

- আগে `A` compile হবে?
- নাকি `B`?

`A`-কে compile করতে `B` দরকার।
আবার `B`-কে compile করতে `A` দরকার।

এই infinite dependency-এর কারণে Go compiler এটি নিষিদ্ধ করেছে।

---

## Production Rule

Dependency সবসময় **একদিকে** যাবে।

```text
Handler
   ↓
Service
   ↓
Repository
   ↓
Database
```

❌ এমন হবে না:

```text
Handler
   ↓
Service
   ↓
Repository
   ↓
Handler
```

---

## তোমার Project-এ

যদি থাকে:

```text
event
   ↓
dto
```

তাহলে `dto` কখনো `event` import করবে না।

অর্থাৎ:

```text
event  ─────► dto
```

✅ ঠিক

কিন্তু

```text
event ─────► dto
   ▲           │
   └───────────┘
```

❌ Import Cycle (Error)

### সহজভাবে মনে রাখো

> **যখন একটি package-এ ফিরে আসার জন্য import-এর একটি বৃত্ত (loop) তৈরি হয়, তখন সেটাই Import Cycle।** Go-তে dependency graph অবশ্যই **acyclic (loop-free)** হতে হবে।

Production-level architecture-এ **mapper/converter** আলাদা রাখা হয়।

### Recommended Structure

```text
internal/
└── event/
    ├── dto/
    │   ├── request.go
    │   └── response.go
    ├── mapper/
    │   └── event_mapper.go
    ├── entity.go
    ├── repository.go
    ├── service.go
    ├── handler.go
    └── routes.go
```

### `mapper/event_mapper.go`

```go
package mapper

import (
	"myapp/internal/event"
	"myapp/internal/event/dto"
)

func ToEventResponse(e *event.Event) dto.EventResponse {
	return dto.EventResponse{
		ID:               e.ID,
		Title:            e.Title,
		Description:      e.Description,
		Location:         e.Location,
		StartAt:          e.StartAt,
		EndAt:            e.EndAt,
		TotalTickets:     e.TotalTickets,
		AvailableTickets: e.AvailableTickets,
		BookedTickets:    e.TotalTickets - e.AvailableTickets,
		CreatedAt:        e.CreatedAt,
		UpdatedAt:        e.UpdatedAt,
	}
}
```

Service-এ ব্যবহার করবে:

```go
import "myapp/internal/event/mapper"

response := mapper.ToEventResponse(event)
```

## যদি আলাদা `mapper` package না রাখতে চাও

তাহলে `mapper.go` file একই `event` package-এর ভিতরেও রাখতে পারো:

```text
event/
├── entity.go
├── mapper.go
├── service.go
```

```go
package event

import "myapp/internal/event/dto"

func ToEventResponse(e *Event) dto.EventResponse {
    ...
}
```

এটি তখনই ঠিক থাকবে, **যদি `dto` package কখনো `event` package import না করে**।

### আমার Recommendation

Production-level Clean Architecture-এর জন্য:

- ✅ `entity` → শুধু Entity
- ✅ `dto` → শুধু Request/Response struct
- ✅ `mapper` → Entity ↔ DTO conversion
- ✅ `service` → Business logic
- ✅ `repository` → Database access

এতে import cycle এড়ানো সহজ হয় এবং code maintain করাও সহজ হয়।
