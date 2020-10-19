<template>
<b-table-simple table-variant="light" responsive="sm" caption-top>
  <caption class="text-center" style="text-transform:uppercase; font-weight:bold;">Top Five Item Types</caption>
  <b-thead>
    <b-tr style="border-bottom: 2px solid #dddddd;">
      <b-th>Item Type</b-th>
      <b-th>Total Profit</b-th>
      <b-th>Percentage</b-th>
    </b-tr>
  </b-thead>
  <b-tbody>
    <b-tr v-for="(item,index) in items" :key="index" style="border-bottom: 2px solid #dddddd;">
      <b-td>
        {{item.ItemType}}
      </b-td>
      <b-td>
        {{item.Total}}
      </b-td>
      <b-td style="width:50%;">
        <b-progress :value="item.Total" :max="totalProfit" show-progress animated variant="info"></b-progress>
      </b-td>
    </b-tr>
  </b-tbody>
</b-table-simple>
</template>
<script lang="ts">
import { Component, Mixins, Watch } from 'vue-property-decorator'
import axios from 'axios'
import StoreMixin from '@/mixins/StoreMixin'
interface Result {
  ItemType: string;
  Total: number;
}
@Component
export default class StoreItems extends Mixins(StoreMixin) {
  items: Result[] = []
  get totalProfit (): number {
    let total = 0
    this.items.map(el => {
      total += el.Total
    })
    return total
  }

  mounted () {
    this.getItems()
  }

  @Watch('dateCommits')
  public getItems (): void{
    axios.get('http://localhost:3000/sales/item_types', {
      params: {
        startDate: this.startDate,
        endDate: this.endDate
      }
    }).then(res => {
      const items: Result[] = res.data
      this.items = items
    })
  }
}
</script>
