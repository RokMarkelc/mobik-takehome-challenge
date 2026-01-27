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