# Blackheaven admin panel/manager
This project uses a modern web stack combining Go with lightweight frontend tooling:

### Backend
- **Go** – Main language powering the backend.
- **[Chi](https://github.com/go-chi/chi)** – Lightweight and idiomatic router for HTTP things.
- **Stdlib** – For performance and simplicity.

### Frontend
- **[Templ](https://templ.guide/)** – Type-safe HTML templating in Go
- **[TemplUI](https://templui.io/)** - Component library built on top of Templ
- **[HTMX](https://htmx.org/)** – Enables dynamic, AJAX-like behavior without JS
- **[Alpine.js](https://alpinejs.dev/)** – For handling minimal frontend interactivity
- **[Tailwind CSS](https://tailwindcss.com/)** – CSS framework for UI styling

### Project Structure & Approach
```go
.
├── internal // where most of what's considered "backend" logic exists
│   ├── database // database functionalities, schema, queries, etc.
│   ├── handlers // http handlers functions
│   ├── models   // structs that represent domain objects (and often map to database tables)
│   └── services // functions that are called by http handlers (usually upon successful requests) 
├── server // setup for the web server (routing, middleware, etc.)
├── utils  // general project wide utilities
└── views  // the frontend UI the user ends up seeing
    ├── assets     // frontend assets like css, images, js
    ├── components // UI components imported by TemplUI
    ├── layouts    // boilerplate that define an HTML doc that pages can use
    ├── modules    // custom project components (can also be wrappers over TemplUI components) 
    └── pages      // the pages the user ends up seeing (usually tied to a route)
```
