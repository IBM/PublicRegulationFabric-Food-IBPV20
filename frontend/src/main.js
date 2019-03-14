import Vue from 'vue'
import App from './App.vue'
import VFC from 'vfc'
import DataTable from 'v-data-table'
import VModal from 'vue-js-modal'
// import Modal from './components/Modal.vue'

import vSelect from 'vue-select'

import 'vfc/dist/vfc.css'
import './dist/json-tree.css'


import TreeView from "vue-json-tree-view"

Vue.config.productionTip = false

Vue.use(DataTable)
Vue.use(VFC)
Vue.use(VModal)
Vue.use(TreeView)
Vue.component('v-select', vSelect)
// Vue.use(VModal, { componentName: "modal" })


// Vue.use(Vuetable)
// Vue.component('modal', {
//   template: '#modal-template'
// })

new Vue({
  render: h => h(App),
  data: function() {
    return {
      showModal: false
    }
  }
}).$mount('#app')
