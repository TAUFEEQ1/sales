import { Component, Vue } from 'vue-property-decorator'
@Component
class StoreMixin extends Vue {
  get startDate (): string {
    return this.$store.state.sDate
  }

  get endDate (): string {
    return this.$store.state.eDate
  }

  get dateCommits (): number {
    return this.$store.state.dComm
  }
}
export default StoreMixin
