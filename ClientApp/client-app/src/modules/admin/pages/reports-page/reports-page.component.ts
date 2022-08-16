import { Component, OnInit } from '@angular/core';
import { ChartData, ChartType } from 'chart.js';
import { Report } from '../../models/Report';
import { ReportsService } from '../../services/reports.service';

@Component({
  selector: 'app-reports-page',
  templateUrl: './reports-page.component.html',
  styleUrls: ['./reports-page.component.scss']
})
export class ReportsPageComponent implements OnInit {

  doughnutChartLabels: string[] = [];
  doughnutChartData: ChartData<'doughnut'> = {
    datasets: []
  }
  doughnutChartType: ChartType = 'doughnut';
  
  doughnutChartLabels2: string[] =  [];
  doughnutChartData2: ChartData<'doughnut'> = {
    datasets: [],
  }
  constructor(private reportsService: ReportsService) { }

  ngOnInit(): void {
    this.reportsService.getReports()
    .subscribe((response) => {
      var report = response.body as Report;
      const map = new Map(Object.entries(report.map_restaurants));

      let temp = []

      for (let [key, value] of map) {
        this.doughnutChartLabels.push(key)
        temp.push(value)
      }

      this.doughnutChartData = {
        labels: this.doughnutChartLabels,
        datasets: [ {data:[]} ]
      }
      this.doughnutChartData.datasets[0].data = temp;

      const map2 = new Map(Object.entries(report.map_articles));

      let temp1 = []
      let temp2 = []
      for (let [key, value] of map2) {
        this.doughnutChartLabels2.push(key)
        temp1.push(value.quantity)
        temp2.push(value.total_price)
      }

      this.doughnutChartData2 = {
        labels: this.doughnutChartLabels2,
        datasets: [ {data:[]}, {data:[]}]
      }
      this.doughnutChartData2.datasets[0].data = temp1;
      this.doughnutChartData2.datasets[1].data = temp2;
    })
  }

}
