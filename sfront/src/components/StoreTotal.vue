<template>
<b-card-group deck>
  <b-card bg-variant="white" header="Total Profits" class="text-center" header-bg-variant="light">
    <b-card-text>
      <span class="total-text">{{result.Total.toLocaleString()}}</span>
    </b-card-text>
  </b-card>
  <b-card bg-variant="white" header="Profit Summary" class="text-left">
    <b-card-text>
      Duration : {{duration}} days
    </b-card-text>
    <b-card-text>
      Daily Average: <span class="text-success font-weight-bold">{{dailyAverage}}</span>
    </b-card-text>
  </b-card>
</b-card-group>

</template>
<style scoped>
.total-text{
  font-size:2em;
}
</style>
<script lang="ts">
import { Component, Watch, Mixins } from 'vue-property-decorator'
import axios from 'axios'
import moment from 'moment'
import StoreMixin from '@/mixins/StoreMixin'
interface Result {
  Total: number;
}

@Component
export default class StoreTotal extends Mixins(StoreMixin) {
  public result: Result= { Total: 0 }

  mounted () {
    this.getTotal()
  }

  get duration (): number {
    const sdate = moment(this.startDate)
    const edate = moment(this.endDate)
    const duration = moment.duration(edate.diff(sdate))
    return Math.round(duration.asDays())
  }

  get dailyAverage (): number {
    return Math.round(this.result.Total / this.duration)
  }

  @Watch('dateCommits')
  public getTotal (): void {
    axios.get('http://localhost:3000/sales/profit_n_period', {
      params: {
        startDate: this.startDate,
        endDate: this.endDate
      }
    }).then(res => {
      const result: Result = res.data
      this.result = result
    })
  }
}
</script>
