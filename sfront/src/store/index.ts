import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    sDate: '2012-01-01',
    eDate: '2015-12-12',
    dComm: 0
  },
  mutations: {
    updateStart (state, sdate) {
      state.sDate = sdate
    },
    updateEnd (state, edate) {
      state.eDate = edate
    },
    saveDates (state) {
      state.dComm = state.dComm + 1
    }
  },
  actions: {
  },
  modules: {
  }
})
