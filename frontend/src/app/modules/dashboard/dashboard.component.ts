import { AfterViewInit, Component, OnInit } from '@angular/core';
import { Chart, registerables } from 'chart.js';

Chart.register(...registerables);

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.scss']
})
export class DashboardComponent implements OnInit, AfterViewInit {
  constructor() { }

  ngOnInit(): void { }

  ngAfterViewInit(): void {
    this.createBarChart();
    this.createProductEditionChart();
    this.createProductActionPieChart();
  }

  createBarChart() {
    new Chart('barChart', {
      type: 'bar',
      data: {
        labels: ['Jan', 'Feb', 'Mar', 'Apr'],
        datasets: [{
          label: 'Vendas',
          data: [30, 45, 60, 70],
          backgroundColor: 'rgba(75, 192, 192, 0.2)',
          borderColor: 'rgba(75, 192, 192, 1)',
          borderWidth: 1
        }]
      },
      options: {
        responsive: true,
        scales: {
          y: {
            beginAtZero: true,
          }
        }
      }
    });
  }

  createProductEditionChart() {
    new Chart('editionChart', {
      type: 'line',
      data: {
        labels: ['Jan', 'Feb', 'Mar', 'Apr', 'May'],
        datasets: [{
          label: 'Produtos Editados',
          data: [50, 70, 90, 120, 150],
          fill: false,
          borderColor: 'rgba(255, 159, 64, 1)',
          tension: 0.1
        }]
      },
      options: {
        responsive: true,
        scales: {
          y: {
            beginAtZero: true
          }
        }
      }
    });
  }

  createProductActionPieChart() {
    new Chart('actionPieChart', {
      type: 'pie',
      data: {
        labels: ['Registros', 'Edições', 'Remoções'],
        datasets: [{
          data: [500, 300, 150],
          backgroundColor: ['#42A5F5', '#FFCA28', '#FF7043'],
          hoverOffset: 4
        }]
      }
    });
  }
}
