গো (Go) ল্যাঙ্গুয়েজের ফোল্ডার স্ট্রাকচার, প্রজেক্ট লেআউট কিংবা `internal/` ফোল্ডার নিয়ে যদি আপনার মনে কোনো কনফিউশন থেকে থাকে, তবে অভিনন্দন — আপনি এই বিষয়ের সবচেয়ে সহজ এবং সাবলীল একটি গাইডে এসে পৌঁছেছেন।

এই টিউটোরিয়ালে আমরা শিখবো:

- কীভাবে একজন প্রফেশনালের মতো আসল গো প্রজেক্ট সাজাতে হয়।
- কেন গো ডেভেলপার হিসেবে আমি এই সাধারণ ফোল্ডার লেআউটটি পছন্দ করি (যদিও এটি নিয়ে অনেক বিতর্ক ও ভিন্ন ভিন্ন মত রয়েছে)।
- অন্যান্য আর্কিটেকচার যেমন: ক্লিন আর্কিটেকচার (Clean Architecture), হেক্সাগোনাল (Hexagonal), এবং রিপোজিটরি প্যাটার্ন (Repository Pattern)।
- চমৎকার কিছু ফোল্ডার ডায়াগ্রাম।
- এবং অবশ্যই: গো ল্যাঙ্গুয়েজটি কেন এত অসাধারণ!

---

## শুরুতেই জানা যাক: কেন গো (Go) বেছে নেবেন?

ফোল্ডার নিয়ে কথা বলার আগে, চলুন ল্যাঙ্গুয়েজটির কিছু দুর্দান্ত দিক দেখে নেওয়া যাক। গো আপনাকে দিচ্ছে:

1. **ফাস্ট রানটাইম (Fast runtime):** এটি সরাসরি মেশিন কোডে কম্পাইল হয়। কোনো JVM বা পাইথন ইন্টারপ্রেটারের ঝামেলা নেই। শুধু বিল্ড করুন, রান করুন এবং রকেটের গতিতে চালান। এপিআই (API), মাইক্রোসার্ভিস এবং সিএলআই (CLI) তৈরির জন্য এটি একদম পারফেক্ট।
2. **স্ট্যাটিকালি টাইপড + সহজে পাঠযোগ্য (Statically typed + readable):** গো মূলত এমন একটি ল্যাঙ্গুয়েজ যা পাইথনের মতো সহজ, আবার সি (C) ল্যাঙ্গুয়েজের মতো নিরাপদ ও দ্রুতগতির।
3. **খুব দ্রুত কম্পাইল হয় (Compiled but really fast):** এর কম্পাইল টাইম এতটাই কম যে আপনি বুঝতেই পারবেন না কখন বিল্ড হয়ে গেছে।
4. **বিল্ট-ইন কনকারেন্সি (Built-in concurrency):** গো-রুটিন (Goroutines) + চ্যানেলস (Channels) = একেবারে ম্যাজিক!
   `go doSomething()` লিখলেই ব্যস, ব্যাকগ্রাউন্ডে অ্যাসিনক্রোনাস কাজ শুরু।
5. **অটোমেটিক গার্বেজ কালেকশন (Automatic garbage collection):** মেমোরি ম্যানেজমেন্ট নিয়ে আপনাকে নিজে থেকে মাথা ঘামাতে হবে না, গো নিজেই মেমোরি সেফটি নিশ্চিত করে।
6. **সহজে শিপ বা ডেপ্লয় করা যায় (Easy to ship):** একটি গো বাইনারি মানে মাত্র একটি সিঙ্গেল ফাইল। ফাইলটি যেকোনো জায়গায় নিয়ে যান, বন্ধুকে পাঠান কিংবা সার্ভারে ডেপ্লয় করুন — কোনো এক্সটার্নাল ডিপেন্ডেন্সির ঝামেলা নেই। একারণেই বড় বড় কোম্পানিগুলো গো ব্যবহার করে।

---

## গো প্রজেক্ট স্ট্রাকচার (Go Project Structure)

```text
go-level-structure/
│
├── cmd/
│   └── main.go                  # অ্যাপ্লিকেশনের এন্ট্রি পয়েন্ট (বুটস্ট্র্যাপ, সার্ভার স্টার্ট)
│
├── internal/
│   ├── handlers/                # ইনকামিং REST রিকোয়েস্টের জন্য HTTP হ্যান্ডলার
│   │   └── todo.go              # টুডু HTTP হ্যান্ডলার (Gin লজিক, JSON বাইন্ডিং)
│   │
│   ├── store/                   # ডাটা পারসিস্টেন্স লেয়ার (Repository pattern)
│   │   └── memory.go            # ইন-মেমোরি স্টোর (ডেভেলপমেন্ট/টেস্টিংয়ের জন্য দরকারী)
│   │
│   ├── routes/                  # অ্যাপ রাউটিং সেটআপ
│   │   └── routes.go            # এন্ডপয়েন্ট রেজিস্টার + মিডলওয়্যার বাইন্ডিং
│   │
│   ├── services/                # বিজনেস লজিক লেয়ার (যোগ করার পরামর্শ দেওয়া হয়)
│   │   └── todo_service.go      # স্টোর ও এক্সটার্নাল ক্লায়েন্ট ব্যবহার করে টুডু বিজনেস লজিক
│   │
│   ├── clients/                 # এক্সটার্নাল REST API ক্লায়েন্ট (অন্য সার্ভিসে HTTP কল করার জন্য)
│   │   └── weather_client.go    # উদাহরণ: এক্সটার্নাল ওয়েদার API কল করা
│   │
│   ├── dto/                     # রিকোয়েস্ট/রেসপন্স DTOs (ঐচ্ছিক)
│   │   └── todo_dto.go
│   │
│   └── config/                  # ইন্টারনাল কনফিগ হেল্পার (যদি pkg/config ব্যবহার না করেন)
│       └── loader.go
│
├── pkg/                         # পুনরায় ব্যবহারযোগ্য হেল্পার লাইব্রেরি (কোনো বিজনেস লজিক থাকবে না)
│   ├── config/
│   │   └── config.go            # এনভায়রনমেন্ট ভেরিয়েবল লোড, কনফিগ স্ট্রাক্ট, Viper ইত্যাদি
│   │
│   ├── logger/
│   │   └── logger.go            # প্রজেক্ট-ওয়াইড লগার (zerolog/logrus/zap র‍্যাপার)
│   │
│   ├── middleware/
│   │   ├── cors.go              # CORS সেটআপ
│   │   └── requestid.go         # ট্রেসিংয়ের জন্য X-Request-ID ইনজেক্ট করা
│   │
│   ├── response/
│   │   └── json.go              # ইউনিফাইড JSON সাকসেস/এরর রেসপন্স হেল্পার
│   │
│   ├── utils/
│   │   ├── strings.go           # স্ট্রিং ইউটিলিটি (slugify, split, trim ইত্যাদি)
│   │   └── conv.go              # টাইপ কনভার্সন ইউটিলিটি (Atoi, float ইত্যাদি)
│   │
│   └── security/
│       ├── hash.go              # পাসওয়ার্ড হ্যাশিং (bcrypt/argon2)
│       └── jwt.go               # JWT সাইন/ভেরিফাই ইউটিলিটি
│
├── .air.toml                    # লাইভ রিলোড কনফিগ (Live reload config)
├── .env                         # লোকাল এনভায়রনমেন্ট ভেরিয়েবল
├── go.mod
└── Makefile                     # বিল্ড, রান, লিন্ট ও টেস্ট অটোমেশন

```

---

## প্রতিটি ফোল্ডারের বিস্তারিত আলোচনা

### ১. cmd/ ফোল্ডার

এটি আপনার অ্যাপ্লিকেশনের শুরু বা প্রবেশদ্বার (Entry point)। `cmd/main.go` ফাইল থেকেই পুরো অ্যাপ্লিকেশন স্টার্ট হয়। একটি বড় গো অ্যাপ্লিকেশনে একাধিক এন্ট্রি পয়েন্ট থাকতে পারে, যেমন:

```text
cmd/
  api/
    main.go
  worker/
    main.go
  migrate/
    main.go

```

### ২. internal/ ফোল্ডার

এখানেই আসল ম্যাজিক ঘটে। `internal/` ফোল্ডারটি গো ল্যাঙ্গুয়েজ নিজেই প্রটেক্ট বা সুরক্ষিত রাখে। অন্য কোনো মডিউল যদি আপনার `internal/` ফোল্ডারের কোড ইমপোর্ট করতে চায়, তবে গো সরাসরি বলবে: _“Nope. That’s illegal. Move along.”_ এর মাধ্যমে আপনি আপনার কোডের ভেতরের ইমপ্লিমেন্টেশন লুকিয়ে রাখতে পারেন।

`internal/` এর ভেতরে সাধারণত যা থাকে:

- **internal/handlers:** এগুলো হলো আপনার HTTP হ্যান্ডলার — যা ইনকামিং API কলের রেসপন্স দেয়। (যেমন: `func (h *UserHandler) CreateUser(c *gin.Context) { ... }`)। এখানে মূলত Request এবং Response-এর লজিক থাকে।

  > **মনে রাখবেন:** হ্যান্ডলারের কাজ শুধু HTTP রিকোয়েস্ট ও বিজনেস লজিক হ্যান্ডেল করা। এদের ডাটাবেজ লজিক বা রাউটিং পাথ জানার কোনো প্রয়োজন নেই।

- **internal/store:** এখানে ডাটা এক্সেসের কাজ করা হয়। এই লেয়ারটি সরাসরি ডাটাবেজের সাথে কথা বলে। (যেমন: `func (s *Store) CreateUser(ctx context.Context, user User) error`)। হ্যান্ডলার কল করে স্টোরকে, আর স্টোর কল করে ডাটাবেজকে। একদম পরিষ্কার বিভাজন!
- **internal/routes:** এখানে সব এন্ডপয়েন্ট রেজিস্টার করা হয়। (যেমন: `users.POST("/", h.CreateUser)`)। এর ফলে হ্যান্ডলারকে রাউটিং পাথ নিয়ে ভাবতে হয় না, আবার সার্ভারকেও হ্যান্ডলারের ভেতরের কোড নিয়ে মাথা ঘামাতে হয় না।

### ৩. pkg/ ফোল্ডার

এই ফোল্ডারটি পাবলিক। এর ভেতরের যেকোনো কোড বাইরের অন্য যেকোনো মডিউল ইমপোর্ট করে ব্যবহার করতে পারবে। এখানে সাধারণত জেনেরিক বা পুনরায় ব্যবহারযোগ্য কোড রাখা হয় (যেমন: `validator`, `jwt`, `logger`, `config`, `utils`, `middleware`)।

> **সহজ নিয়ম:** কোড যদি জেনেরিক বা সবার জন্য উন্মুক্ত হয় তবে `pkg/`-এ রাখুন। আর যদি বিজনেস-স্পেসিফিক বা গোপন রাখতে চান, তবে `internal/`-এ রাখুন।

### ৪. Makefile

একটি `Makefile` আপনার প্রজেক্টের প্রয়োজনীয় কমান্ডগুলোকে অটোমেট বা সহজ করে দেয়। যেমন আপনি ছোট ছোট কমান্ড দিয়ে বড় কাজ করতে পারেন:

- `make run` (প্রজেক্ট রান করা)
- `make test` (টেস্ট চালানো)
- `make build` (কোড বিল্ড করা)

### ৫. .air.toml (লাইভ রিলোডার)

`air` হলো গো-র একটি হট-রিলোড টুল। আপনি কোড পরিবর্তন করার সাথে সাথে এটি অটোমেটিকভাবে সার্ভার রিবিল্ড ও রিলোড করে দেয়। লোকাল ডেভেলপমেন্টের সুবিধার জন্য `Air` এবং প্রোডাকশন বা সিআই/সিডি অটোমেশনের জন্য `Makefile` — দুটোই দারুণ কার্যকর।

---

## বড় প্রজেক্টের ক্ষেত্রে যেভাবে স্কেল করবেন

প্রজেক্ট বড় হলে আপনি ডোমেইন বা ফিচার ভিত্তিক ফোল্ডার তৈরি করতে পারেন:

```text
internal/
├── users/
│   ├── handler.go
│   ├── store.go
│   ├── routes.go
│   └── model.go
├── products/
│   └── ...
└── auth/
    └── ...

```

---

## বিকল্প আর্কিটেকচার এবং কখন কোনটি ব্যবহার করবেন

### ১. ক্লিন আর্কিটেকচার (Clean Architecture)

বড় টিম এবং দীর্ঘমেয়াদী প্রজেক্টের জন্য এটি দারুণ, তবে ছোট প্রজেক্টের জন্য কিছুটা জটিল মনে হতে পারে। এর লেয়ারগুলো হলো:

```text
project-name/
│
├── cmd/
│   └── app/
│       └── main.go
│
├── internal/
│   ├── domain/               ← এন্টিটি এবং ডোমেইন লজিক (Core Business Logic)
│   │   └── todo/
│   │       ├── entity.go
│   │       └── repository.go
│   │
│   ├── usecase/              ← অ্যাপ্লিকেশন সার্ভিস / বিজনেস রুলস
│   │   └── todo/
│   │       └── create.go
│   │
│   ├── interface/            ← অ্যাডাপ্টারস (REST হ্যান্ডলার, gRPC, CLI)
│   │   └── rest/
│   │       └── todo_handler.go
│   │
│   ├── infrastructure/       ← ফ্রেমওয়ার্ক এবং ড্রাইভারস (DB, HTTP Router)
│   │   ├── persistence/
│   │   └── http/
│   │
│   └── bootstrap/            ← ডিপেন্ডেন্সি ইনজেকশন ও অ্যাপ স্টার্টআপ
│       └── initialize.go
...

```

### ২. হেক্সাগোনাল আর্কিটেকচার (Hexagonal Architecture - Ports & Adapters)

বিজনেস লজিককে বাইরের যেকোনো সিস্টেম (যেমন ডাটাবেজ বা থার্ড-পার্টি API) থেকে সম্পূর্ণ আলাদা রাখার জন্য এটি ব্যবহার করা হয়। আধুনিক গো মাইক্রোসার্ভিসে এটি প্রচুর ব্যবহৃত হয় কারণ এটি ফ্রেমওয়ার্কের ওপর নির্ভরশীল নয় এবং খুব সহজে টেস্ট করা যায়।

```text
project/
│
├── cmd/
│   └── app/
│       └── main.go
│
├── internal/
│   ├── domain/                 ← কোর বিজনেস লজিক
│   │   └── todo/
│   │       ├── entity.go
│   │       └── repository.go   (port)
│   │
│   ├── application/            ← ইউজ কেস (Input Ports)
│   │
│   ├── adapters/               ← OUTBOUND অ্যাডাপ্টারস (ডাটাবেজ বা এক্সটার্নাল সার্ভিস ইমপ্লিমেন্টেশন)
│   │   ├── persistence/
│   │   └── external/
│   │
│   ├── ports/                  ← INBOUND ও OUTBOUND ইন্টারফেস (Ports)
│   │   ├── inbound/
│   │   └── outbound/
│   │
│   ├── transport/              ← INBOUND অ্যাডাপ্টারস (REST, gRPC, CLI)
│   │   ├── rest/
│   │   └── grpc/
│   │
│   └── bootstrap/              ← সবকিছু একসাথে যুক্ত করা (Wiring)
...

```

### ৩. রিপোজিটরি প্যাটার্ন (Repository Pattern)

ক্লিন বা হেক্সাগোনাল আর্কিটেকচারের চেয়ে এটি অনেক সহজ, কিন্তু অত্যন্ত স্কেলেবল এবং টেস্টেবল। বাস্তব জীবনের গো প্রজেক্টে এটি সবচেয়ে বেশি জনপ্রিয়। এর মূল লক্ষ্য হলো ডাটা সেভ বা রিড করার লজিককে একটি ইন্টারফেসের (Interface) আড়ালে আলাদা করে ফেলা।

```text
project/
│
├── cmd/
│   └── app/
│       └── main.go
│
├── internal/
│   ├── handlers/              ← HTTP হ্যান্ডলারস (Gin, Fiber, Chi ইত্যাদি)
│   │   └── todo_handler.go
│   │
│   ├── repositories/          ← ইন্টারফেস এবং ডাটাবেজ ইমপ্লিমেন্টেশন
│   │   ├── todo_repository.go         (interface)
│   │   ├── todo_postgres.go          (postgres impl)
│   │   └── mocks/                    (টেস্টিংয়ের জন্য মক রিপোজিটরি)
│   │
│   ├── services/              ← রিপোজিটরি ব্যবহার করে বিজনেস লজিক চালানো
│   │   └── todo_service.go
│   │
│   ├── models/                ← ডোমেইন এন্টিটি/মডেল
│   │   └── todo.go
│   │
│   ├── routes/
│   └── config/
...

```

---

## সারসংক্ষেপ

এখন আপনি জানেন:

1. গো ল্যাঙ্গুয়েজের মূল দর্শন কী।
2. `internal/` এবং `pkg/` ফোল্ডারের আসল পার্থক্য এবং উদ্দেশ্য।
3. হ্যান্ডলার, স্টোর এবং রাউট কীভাবে একে অপরের সাথে কানেক্ট হয়।
4. প্রজেক্ট বড় হলে কীভাবে একাধিক ডোমেইনে কোড ভাগ করতে হয়।
5. বিভিন্ন অ্যাডভান্সড আর্কিটেকচার এবং আপনার প্রজেক্টের জন্য কোনটি সঠিক।

---

# English Version

---

If you’re learning Go and confused about folder structure, project layout, or why everyone has strong opinions about `internal` — congrats, you’ve just opened one of the cleanest and easiest-to-follow guides on the topic.

This tutorial will walk you through:

- How to structure real Go projects like a pro.
- Why, as a Go developer, I love this simple folder layout even though it is highly debatable and opinionated.
- Other architectures: Clean Architecture, Hexagonal, and the Repository Pattern.
- Beautiful folder diagrams.
- And of course: why Go itself is such a fantastic language.

---

## Why Choose Go in the First Place?

Before we talk folders, let’s appreciate the language itself. Go gives you:

1. **Fast runtime:** Compiled straight to machine code. No JVM. No Python interpreter. Just: build → run → fly. It is perfect for APIs, microservices, and CLIs.
2. **Statically typed + readable:** Go basically said: _“What if we made a language as simple as Python, but as safe and fast as C?”_ And then they actually did it.
3. **Compiled (but really fast):** Compile times are so fast that sometimes you barely notice the build happened.
4. **Built-in concurrency:** Goroutines + channels = absolute magic. Write `go doSomething()` and boom—you’re doing async tasks like a wizard.
5. **Automatic garbage collection:** You get memory safety without babysitting your RAM manually.
6. **Easy to ship:** A Go binary is just one file. One. Single. File. Move it anywhere, send it to your friend, or deploy it to a server without worrying about missing dependencies. This is why Go is used by major global brands.

---

## Go Project Structure

```text
go-level-structure/
│
├── cmd/
│   └── main.go                  # Application entrypoint (bootstrap, start server)
│
├── internal/
│   ├── handlers/                # HTTP handlers for incoming REST requests
│   │   └── todo.go              # Todo HTTP handler (Gin logic, JSON binding)
│   │
│   ├── store/                   # Data persistence layer (Repository pattern)
│   │   └── memory.go            # In-memory store (useful for dev/testing)
│   │
│   ├── routes/                  # App routing setup
│   │   └── routes.go            # Registers endpoints + middleware binding
│   │
│   ├── services/                # Business logic layer (recommended addition)
│   │   └── todo_service.go      # Todo business logic using stores + external clients
│   │
│   ├── clients/                 # External REST API clients (HTTP calls to other services)
│   │   └── weather_client.go    # Example: calling external weather API
│   │
│   ├── dto/                     # Request/response DTOs (optional)
│   │   └── todo_dto.go
│   │
│   └── config/                  # Internal config helpers (if not using pkg/config)
│       └── loader.go
│
├── pkg/                         # Reusable helper libraries (no business logic)
│   ├── config/
│   │   └── config.go            # Load env vars, config structs, Viper, etc.
│   │
│   ├── logger/
│   │   └── logger.go            # Project-wide logger (zerolog/logrus/zap wrapper)
│   │
│   ├── middleware/
│   │   ├── cors.go              # CORS setup
│   │   └── requestid.go         # Injects X-Request-ID for tracing
│   │
│   ├── response/
│   │   └── json.go              # Unified JSON success/error response helpers
│   │
│   ├── utils/
│   │   ├── strings.go           # String utilities (slugify, split, trim, etc.)
│   │   └── conv.go              # Type conversion utilities (Atoi, float, etc.)
│   │
│   └── security/
│       ├── hash.go              # Password hashing (bcrypt/argon2)
│       └── jwt.go               # JWT sign/verify utilities
│
├── .air.toml                    # Live reload config
├── .env                         # Local environment variables
├── go.mod
└── Makefile                     # Build, run, lint, test automation

```

---

## Breaking Down Each Part

### 1. The cmd/ Folder

The entry point of your application. `cmd/main.go` is the exact place where your entire application boots up. In a large-scale Go application, you might have multiple application targets:

```text
cmd/
  api/
    main.go
  worker/
    main.go
  migrate/
    main.go

```

### 2. The internal/ Folder

This is where the fun happens. `internal/` is a special directory protected by Go's build tools. If another code module outside your project tries to import something from your `internal/` folder, the Go compiler will explicitly block it. This gives you a powerful way to encapsulate and hide your private implementation details.

Inside `internal/`, you cleanly split responsibilities:

- **internal/handlers:** These are your HTTP handlers—the functions that capture incoming API requests and map out responses (e.g., `func (h *UserHandler) CreateUser(c *gin.Context) { ... }`).

  > **Rule of thumb:** Handlers should only handle HTTP logic. They shouldn't write SQL queries directly or register routing paths.

- **internal/store:** This is your data access layer. The store interacts directly with your database (e.g., `func (s *Store) CreateUser(ctx context.Context, user User) error`). Handlers call services, services call the store, and the store executes database actions.
- **internal/routes:** This is where endpoints get registered (e.g., `users.POST("/", h.CreateUser)`). It acts as the perfect, decoupled bridge between the underlying server engine and your request handlers.

### 3. The pkg/ Folder

This folder is entirely public. Anything placed in `pkg/` can be freely imported by external applications or packages. It acts as an internal repository for generic utility configurations like loggers, custom encryption helpers (`jwt`), text manipulation scripts (`utils`), or framework-agnostic middlewares. If a piece of code is tied to your core business logic, it stays in `internal/`; if it's generic and re-usable, it goes into `pkg/`.

### 4. Makefile

A `Makefile` automates command-line steps to dramatically speed up workflows. Instead of writing long commands, you can instantly run tests, compile code, or handle database migrations using clean wrappers like:

- `make run`
- `make test`
- `make build`

### 5. .air.toml (Your Live Reloader)

`air` is an incredible daemon tool that watches your project directories for file modifications and triggers instantaneous hot-reloads. While it shares execution jobs with your `Makefile`, it occupies a distinct lane: `Air` is dedicated to saving development time locally, while your `Makefile` acts as the single source of truth for automation, CI/CD pipelines, and production tasks.

---

## Scaling Your Project Structure

As your product grows, you can easily scale horizontally by organizing your `internal/` codebase by domain/feature rather than technical layer:

```text
internal/
├── users/
│   ├── handler.go
│   ├── store.go
│   ├── routes.go
│   └── model.go
├── products/
│   └── ...
└── auth/
    └── ...

```

---

## Alternative Architectures & When to Use Them

### 1. Clean Architecture

This is incredibly powerful for massive development teams working on enterprise applications, though it introduces boilerplate overhead for lean teams. It separates code into strict layers:

```text
project-name/
│
├── cmd/
│   └── app/
│       └── main.go
│
├── internal/
│   ├── domain/               ← Core domain models, entities, and interfaces
│   │   └── todo/
│   │       ├── entity.go
│   │       └── repository.go
│   │
│   ├── usecase/              ← Application specific business logic rules
│   │   └── todo/
│   │       └── create.go
│   │
│   ├── interface/            ← Interface adapters (REST handlers, gRPC, CLI)
│   │   └── rest/
│   │       └── todo_handler.go
│   │
│   ├── infrastructure/       ← Concrete drivers (DB connections, HTTP routers)
│   │   ├── persistence/
│   │   └── http/
│   │
│   └── bootstrap/            ← Core dependency wiring & application initialization
...

```

### 2. Hexagonal Architecture (Ports & Adapters)

Hexagonal architecture decouples core domain algorithms entirely from external dependencies like databases, third-party messaging queues, or user interfaces. Highly favored in microservice patterns, it provides excellent unit testability and keeps code framework-independent.

```text
project/
│
├── cmd/
│   └── app/
│       └── main.go
│
├── internal/
│   ├── domain/                 ← Pure application business kernel
│   │   └── todo/
│   │       ├── entity.go
│   │       └── repository.go   (port)
│   │
│   ├── application/            ← Core use-cases (Input Ports)
│   │
│   ├── adapters/               ← Outbound adapters (Database/External implementations)
│   │   ├── persistence/
│   │   └── external/
│   │
│   ├── ports/                  ← Boundary interfaces defining communication contracts
│   │   ├── inbound/
│   │   └── outbound/
│   │
│   ├── transport/              ← Inbound adapters (REST endpoints, gRPC streams)
│   │   ├── rest/
│   │   └── grpc/
│   │
│   └── bootstrap/              ← Manual dependency injection setups
...

```

### 3. Repository Pattern

This layout is more accessible and practical than Clean or Hexagonal architecture while retaining great horizontal scalability. It concentrates heavily on segregating state persistence from business functionality using clean Go interfaces.

```text
project/
│
├── cmd/
│   └── app/
│       └── main.go
│
├── internal/
│   ├── handlers/              ← Web traffic handlers (Gin, Chi, Fiber, etc.)
│   │   └── todo_handler.go
│   │
│   ├── repositories/          ← Storage layer interfaces and variations
│   │   ├── todo_repository.go         (interface)
│   │   ├── todo_postgres.go          (PostgreSQL implementation)
│   │   └── mocks/                    (Mock databases generated for automated testing)
│   │
│   ├── services/              ← Intermediary layers linking routes to storage access
│   │   └── todo_service.go
│   │
│   ├── models/                ← Central domain models
│   │   └── todo.go
│   │
│   ├── routes/
│   └── config/
...

```

---

## Summary

You now have a solid understanding of:

1. Go’s design philosophies.
2. The exact boundary definitions separating `internal/` from `pkg/`.
3. How handlers, service stores, and system routing pathways interconnect seamlessly.
4. How to scale your repository layers across independent functional domains.
5. The various layout alternative frameworks and how to confidently select the right tool for your next project.
