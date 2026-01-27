SHELL := /bin/bash

# Load .env variables if present
-include .env
export

DC := docker compose -f docker-compose.dev.yml

.PHONY: help init up down logs ps clean \
        frontend-init frontend-wire backend-tidy \
        migrate-up migrate-down migrate-create psql

help:
	@echo "Common commands:"
	@echo "  make init             - Bootstrap everything (Angular app, DB up, migrations, services)"
	@echo "  make up               - Start db, backend, frontend in dev mode"
	@echo "  make down             - Stop all services"
	@echo "  make logs             - Tail logs from all services"
	@echo "  make migrate-create n=name - Create new timestamped migration"
	@echo "  make migrate-up       - Apply all migrations"
	@echo "  make migrate-down     - Rollback one migration step"
	@echo "  make psql             - Open psql shell into db"
	@echo "  make clean            - Stop & remove volumes (destroys db data)"

init: .env frontend-init frontend-wire up migrate-up
	@echo "✅ Init complete. Visit http://localhost:4200"

.env:
	cp .env.example .env

up:
	$(DC) up -d --build db
	sleep 2
	$(DC) up -d --build backend frontend
	$(DC) ps

down:
	$(DC) down

logs:
	$(DC) logs -f

ps:
	$(DC) ps

clean:
	$(DC) down -v

backend-tidy:
	$(DC) exec backend bash -lc "go mod tidy && go fmt ./..."

# --- FRONTEND (Angular) ---

frontend-init:
	@mkdir -p frontend
	# Generate Angular app (module-based; standalone disabled)
	docker run --rm -v $$PWD/frontend:/app -w /app node:20 bash -lc "\
		npm i -g @angular/cli@18 && \
		ng new frontend --directory . --skip-git --style=scss --routing=true --package-manager npm --standalone=false \
	"

frontend-wire:
	# Ensure dev server uses proxy + binds 0.0.0.0 (run edit inside container)
	docker run --rm -v $$PWD/frontend:/app -w /app node:20 node -e '\
	const fs=require("fs");\
	const p=JSON.parse(fs.readFileSync("package.json"));\
	p.scripts = p.scripts || {};\
	p.scripts.start = "ng serve --host 0.0.0.0 --port 4200 --poll 2000 --proxy-config proxy.conf.json";\
	fs.writeFileSync("package.json", JSON.stringify(p, null, 2));\
	'

	# Proxy config to route /api to backend service in Docker network
	@cat > frontend/proxy.conf.json <<'EOF'
{
  "/api": {
    "target": "http://backend:8080",
    "secure": false,
    "changeOrigin": true,
    "logLevel": "debug"
  }
}
EOF

	# ApiService
	@cat > frontend/src/app/api.service.ts <<'EOF'
import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

export interface Todo {
  id: number;
  title: string;
  done: boolean;
  created_at: string;
}

@Injectable({ providedIn: 'root' })
export class ApiService {
  constructor(private http: HttpClient) {}

  getTodos() {
    return this.http.get<Todo[]>('/api/todos');
  }
  createTodo(title: string) {
    return this.http.post<Todo>('/api/todos', { title });
  }
  health() {
    return this.http.get<{status: string}>('/api/health');
  }
}
EOF

	# app.module.ts (imports HttpClientModule + FormsModule)
	@cat > frontend/src/app/app.module.ts <<'EOF'
import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http';
import { FormsModule } from '@angular/forms';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';

@NgModule({
  declarations: [AppComponent],
  imports: [BrowserModule, AppRoutingModule, HttpClientModule, FormsModule],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule {}
EOF

	# app.component.ts
	@cat > frontend/src/app/app.component.ts <<'EOF'
import { Component, OnInit } from '@angular/core';
import { ApiService, Todo } from './api.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss'],
})
export class AppComponent implements OnInit {
  title = 'Three-tier Angular + Go + Postgres';
  todos: Todo[] = [];
  newTitle = '';
  health = 'checking...';

  constructor(private api: ApiService) {}

  ngOnInit(): void {
    this.api.health().subscribe({
      next: (h) => (this.health = h.status),
      error: () => (this.health = 'down'),
    });
    this.load();
  }

  load() {
    this.api.getTodos().subscribe((list) => (this.todos = list));
  }

  add() {
    const t = this.newTitle.trim();
    if (!t) return;
    this.api.createTodo(t).subscribe(() => {
      this.newTitle = '';
      this.load();
    });
  }
}
EOF

	# app.component.html
	@cat > frontend/src/app/app.component.html <<'EOF'
<div style="max-width:800px;margin:2rem auto;font-family:system-ui,Arial;">
  <h1>{{ title }}</h1>
  <p>API health: <strong>{{ health }}</strong></p>

  <div style="margin:1rem 0;">
    <input [(ngModel)]="newTitle" placeholder="New todo title" style="padding:.5rem;width:70%;" />
    <button (click)="add()" style="padding:.55rem 1rem;margin-left:.5rem;">Add</button>
  </div>

  <ul>
    <li *ngFor="let t of todos">
      {{ t.title }} <small style="opacity:.6;">(#{{t.id}})</small>
    </li>
  </ul>
</div>
EOF

# --- MIGRATIONS ---

migrate-create:
	@if [ -z "$(n)" ]; then echo "Usage: make migrate-create n=name_of_migration"; exit 1; fi
	docker run --rm -v $$PWD/migrations:/migrations migrate/migrate:4 create -ext sql -dir /migrations -seq $(n)

migrate-up:
	$(DC) run --rm migrate -path=/migrations -database "postgres://$(DB_USER):$(DB_PASSWORD)@db:$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)" up

migrate-down:
	$(DC) run --rm migrate -path=/migrations -database "postgres://$(DB_USER):$(DB_PASSWORD)@db:$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)" down 1

psql:
	$(DC) exec db psql -U $(POSTGRES_USER) -d $(POSTGRES_DB)
