<template>
  <div class="pl-4">
    <b-table :items="items" head-variant="light" :fields="fields" tbody-class="bg-light" bordered responsive primary-key="id" ></b-table>
    <div id="load_sign" v-if="loading">
      <b-spinner variant="info" style="width:10em; height:10em;" class="text-center mx-auto"></b-spinner>
    </div>
    <ul class="list-inline list-unstyled" style="margin-top:-10px;">
      <li class="list-inline-item">
        <b-btn variant="link" @click="decrement" :disabled="pageNo == 1"  class="text-info font-weight-bold">
          <span>PREVIOUS</span>
        </b-btn>
      </li>
      <li class="list-inline-item">
        {{pageNo}}
      </li>
      <li class="list-inline-item">
        <b-btn variant="link" @click="increment" class="text-info font-weight-bold" :disabled="pageNo == pageMax">
          NEXT
        </b-btn>
      </li>
      <li class="list-inline-item">
        <i>Showing {{pageNo}} of {{Math.ceil(saleCount/pageSize)}}</i>
      </li>
    </ul>
  </div>
</template>
<style scoped>
#load_sign{
  position: absolute;
  top:33%;
  right:67%;
  z-index: 200;
}
</style>
<script lang="ts">
import Vue from 'vue'
import axios from 'axios'
import moment from 'moment'
interface Order {
  id: number;
  orderDate: string;
  orderPriority: string;
  unitsSold: number;
  unitPrice: number;
  totalCost: number;
  totalRevenue: number;
  itemType: string;
}
class OrderModel {
  private orderDate: string;
  public id: number;
  OrderPriority: string;
  UnitsSold: number;
  UnitPrice: number;
  TotalCost: number;
  TotalRevenue: number;
  ItemType: string;
  constructor (orderMd: Order) {
    this.id = orderMd.id
    this.orderDate = orderMd.orderDate
    this.OrderPriority = orderMd.orderPriority
    this.UnitsSold = orderMd.unitsSold
    this.UnitPrice = orderMd.unitPrice
    this.TotalCost = orderMd.totalCost
    this.TotalRevenue = orderMd.totalRevenue
    this.ItemType = orderMd.itemType
  }

  get OrderDate (): string {
    return moment(this.orderDate).format('YYYY-MM-DD')
  }
}
export default Vue.extend({
  data () {
    const fields = [
      {
        key: 'OrderDate',
        label: 'Order Date'
      },
      {
        key: 'OrderPriority',
        label: 'Priority'
      },
      {
        key: 'UnitsSold',
        label: 'U.Sold'
      },
      {
        key: 'UnitPrice',
        label: 'U.Price'
      },
      {
        key: 'TotalCost',
        label: 'T.Cost'
      },
      {
        key: 'TotalRevenue',
        sortable: true,
        label: 'T.Revenue'
      },
      {
        key: 'ItemType',
        label: 'Item Type'
      }
    ]
    return {
      loading: true,
      pageNo: 1,
      pageSize: 8,
      fields: fields,
      items: new Array<OrderModel>()
    }
  },
  props: ['saleCount'],
  computed: {
    pageMax (): number {
      return Math.ceil(this.saleCount / this.pageSize)
    }
  },
  mounted () {
    this.getSales()
  },
  watch: {
    pageNo () {
      this.getSales()
    }
  },
  methods: {
    getSales (): void {
      this.loading = true
      axios.get('http://localhost:3000/sales/index', {
        params: {
          pageNo: this.pageNo,
          pageSize: this.pageSize
        }
      }).then(res => {
        const data: Order[] = res.data
        this.items = data.map(el => {
          return new OrderModel(el)
        })
        this.loading = false
      })
    },
    increment (): void {
      this.pageNo = this.pageNo + 1
    },
    decrement (): void {
      this.pageNo = this.pageNo - 1
    }
  }
})
</script>
