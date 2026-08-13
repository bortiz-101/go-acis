import { Component } from '@angular/core';
import { bootstrapApplication } from '@angular/platform-browser';

@Component({
  selector: 'app-root',
  template: `
    <button (click)="checkHealth()">Health</button>

    @if (response) {
      <p>{{ response }}</p>
    }
  `,
  styles: `
    button {
      margin: 2rem;
    }

    p {
      margin: 0 2rem;
    }
  `,
})
class App {
  response = '';

  async checkHealth() {
    try {
      const response = await fetch('/health');
      this.response = await response.text();
    } catch {
      this.response = 'Could not reach the API';
    }
  }
}

bootstrapApplication(App).catch((error) => console.error(error));
