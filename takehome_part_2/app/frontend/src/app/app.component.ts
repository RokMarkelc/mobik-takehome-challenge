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