<template>
<nav class="pl-3">
    <b-form inline @submit.prevent="saveDate">
        <b-form-group>
            <label for="example-datepicker-1" class="font-weight-bold">Choose  Start Date</label>
            <b-form-datepicker id="example-datepicker-1" v-model="startDate" class="mb-1" size="sm"></b-form-datepicker>
        </b-form-group>
        <b-form-group class="ml-1">
            <label for="example-datepicker" class="font-weight-bold">Choose  End date</label>
            <b-form-datepicker id="example-datepicker" v-model="endDate" class="mb-1" size="sm"></b-form-datepicker>
        </b-form-group>
        <b-form-group class="mt-3 ml-1">
        <b-btn type="submit" variant="info" class="text-light font-weight-bold mt-1" size="sm">
          FILTER
        </b-btn>
        </b-form-group>

    </b-form>
</nav>
</template>

<script lang="ts">
import Vue from 'vue'
export default Vue.extend({
  computed: {
    startDate: {
      set (ndate: string) {
        this.$store.commit('updateStart', ndate)
      },

      get (): string {
        return this.$store.state.sDate
      }
    },
    endDate: {
      set (ndate: string) {
        this.$store.commit('updateEnd', ndate)
      },

      get (): string {
        return this.$store.state.eDate
      }
    }
  },
  methods: {
    saveDate (): void {
      this.$store.commit('saveDates')
    }
  }
})
</script>
