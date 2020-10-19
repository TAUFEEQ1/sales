<template>
  <div id="app" class="container-fluid">
    <div class="row">
      <div class="pt-2" style="width:54%">
        <h4 class="text-center text-uppercase app-title my-2">Store Sales Manager</h4>
        <store-index :saleCount="saleCount"></store-index>
        <store-sales></store-sales>
      </div>
      <div style="width:45%" class="pt-3">
        <store-dates></store-dates>
        <div id="affected-mems" class="mt-1">
          <div class="stotal pl-3 pr-3 mb-2">
            <h6 class="text-center font-weight-bold mt-1 mb-3" style="text-transform:uppercase;"> PROFITS IN RANGE </h6>
            <store-total></store-total>
          </div>
          <div class="pl-3 pr-3">
            <store-items></store-items>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import StoreIndex from '@/components/StoreIndex.vue'
import StoreDates from '@/components/StoreDates.vue'
import StoreTotal from '@/components/StoreTotal.vue'
import StoreItems from '@/components/StoreItems.vue'
import StoreSales from '@/components/StoreSales.vue'
import axios from 'axios'
interface SaleCount {
  Total: number;
}
export default Vue.extend({
  name: 'App',
  components: {
    StoreIndex,
    StoreDates,
    StoreTotal,
    StoreItems,
    StoreSales
  },
  data () {
    return {
      saleCount: 0
    }
  },
  mounted () {
    this.CountSales()
  },
  methods: {
    CountSales (): void {
      axios.get('http://localhost:3000/sales/count').then(res => {
        const saleCount: SaleCount = res.data
        this.saleCount = saleCount.Total
      })
    }
  }
})
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #606060;
  background-color:#F2EBC2;
  height:100%;
  padding-top:15px;
  padding-bottom: 18px;
}
.app-title{
  font-weight:600;
}
</style>
<style scoped>
#affected-mems{
 display: flex;
 flex-direction: column;
}
.stotal{
  max-width: 800px;
}
</style>
