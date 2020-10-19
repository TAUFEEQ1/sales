<template>
  <b-card class="ml-4 mr-auto" style="width:500px;" >
    <h6 class="font-weight-bold text-center">REPORT UPLOAD</h6>
    <b-row>
      <b-col style="flex:5">
        <b-form-file accept=".csv"  @change="storeSales($event)"></b-form-file>
      </b-col>
      <b-col>
        <b-spinner v-if="spshow" :variant="loadstate?'success':'warning'"></b-spinner>
        <span v-show="spshow">{{loadstate?'Processing':'Uploading'}}</span>
      </b-col>
    </b-row>
  </b-card>
</template>
<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import axios from 'axios'
@Component
export default class StoreSales extends Vue {
  spshow: boolean;
  loadstate: boolean;

  constructor () {
    super()
    this.spshow = false
    this.loadstate = false
  }

  checkProgress (progressEvent: { loaded: number; total: number }) {
    const percentCompleted = progressEvent.loaded / progressEvent.total
    if (Math.round(percentCompleted) > 0.99) {
      this.loadstate = true
    }
  }

  storeSales (event: { target: { files: (string|Blob)[] } }): void {
    const uform = new FormData()
    const file = event.target.files[0]
    uform.append('sales_file', file)
    this.spshow = true
    axios.post('http://localhost:3000/sales/store', uform, {
      headers: {
        'Content-Type': 'multipart/form-data'
      },
      onUploadProgress: this.checkProgress
    }).then(() => {
      this.spshow = false
      alert('Operation successful')
    }).catch(err => {
      this.spshow = false
      this.loadstate = false
      if (err.response.status === 409) {
        alert(err.response.data.message)
      } else {
        alert('Unkown Error')
      }
    })
  }
}
</script>
