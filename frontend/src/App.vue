<template>
  <div id="app">

    <!-- <demo-login-modal/> -->

    <!-- <img alt="Vue logo" src="./assets/logo.png"> -->
    <!-- <HelloWorld msg="Welcome to Your Vue.js App"/> -->

    <!-- <button id="show-modal" @click="showModal = true">Show Modal</button> -->
    <!-- use the modal component, pass in the prop -->
    <!-- <modal v-if="showModal" @close="showModal = false"> -->
      <!-- <h3 slot="header">custom header</h3> -->
    <!-- </modal> -->
    <!-- <input v-model="message" placeholder="edit me">
    <p>Message is: {{ message }}</p> -->

    <!-- <vue-form v-model="form" > -->

    <!-- <vue-form v-model="form" >        -->
    <div>
      <vue-button type="default" v-on:click="showInvokeModal({'function': 'init_product_listing', 'fields': ['Product Listing Id', 'Supplier ID', 'Product ID'], 'title': 'Create Product Listing'})">Create Product Listing</vue-button>
      <vue-button type="default" v-on:click="showInvokeModal({'function': 'init_product', 'fields': ['Product Id', 'Quantity', 'CountryId'], 'title': 'Create Product'})">Create Product</vue-button>
      <vue-button type="default" v-on:click="showInvokeModal({'function': 'init_user', 'fields': ['ID'], 'title': 'Create User'})">Create User</vue-button>
      <vue-button type="default" v-on:click="showInvokeModal({'function': 'init_regulator', 'fields': ['ID'], 'title': 'Create Regulator'})">Create Regulator</vue-button>
    </div>
    </br>
    <div>

      <vue-button type="default" v-on:click="showInvokeModal({'function': 'transfer_product_listing', 'fields': ['Product Listing Id', 'New Owner ID'], 'title': 'Transfer Product Listing'})">Transfer Product Listing</vue-button>
      <vue-button type="default" v-on:click="showInvokeModal({'function': 'check_products', 'fields': ['Product Listing ID', 'Regulator ID'], 'title': 'Check Products'})">Check Products</vue-button>


      <vue-button type="default" v-on:click="getLedger">Refresh Ledger</vue-button>
    </div>

    <!-- <button
      class="btn green"
      @click="$modal.show('demo-login')">
      Demo: Login
    </button> -->
    <!-- <p>Ledger State -->
      <!-- <pre>
      {{ ledgerState | pretty }}
      </pre> -->
    <!-- </p> -->

    <!-- <data-table :data="gridData">
    </data-table> -->
    <!-- <v-card> -->
    <div>


      <modal name="invoke-modal" height="auto" >
        <h2 align="center"> {{title}} </h2>
        <vue-form
          id="chaincode-form"
          :model="form">

          <template v-for="(field, idx) in fields">
            <vue-form-item style="width:500px;" align=center>
              <vue-input
                :placeholder=field
                v-model=input[idx]>
              </vue-input>
            </vue-form-item>
          </template>
          <template v-if="func == 'init_user'">
            <vue-form-item >
              <v-select @change="addFields" style="width:400px;margin-left:100px" id="user_type" v-model="user_type" placeholder="User Type" :options="['retailer','importer', 'supplier']"></v-select>
            </vue-form-item>
          </template>
          <template v-for="(field, idx) in user_fields">
            <vue-form-item>
              <vue-input
                :placeholder=field
                v-model=user_input[idx]>
              </vue-input>
            </vue-form-item>
          </template>
          <vue-form-item style="margin-left:100px">
            <vue-button type="default" v-on:click="hideModal">Cancel</vue-button>
            <vue-button type="success" v-on:click="invoke" >Submit</vue-button>
          </vue-form-item>
        </vue-form>
      </modal>


      <!-- <modal name="invoke-modal">
        <vue-form
          id="chaincode-form"
          :model="form"
          style="width: 75%; position: fixed; left: 50%; margin-left: -37.5%;">
          <h2 style="float:center"> Invoke Chaincode </h2>
          <vue-form-item label="Function">
            <vue-input
              placeholder="Function"
              v-model="form.function"
              style="width: 100%">
            </vue-input>
          </vue-form-item>

          <vue-form-item label="Arguments">
            <vue-input
              placeholder="Arguments"
              v-model="form.args"
              style="width: 100%">
            </vue-input>
          </vue-form-item>
          <vue-form-item>
            <vue-button type="default" v-on:click="isHidden = true">Cancel</vue-button>
            <vue-button type="success" v-on:click="invoke" >Submit</vue-button>
          </vue-form-item>
        </vue-form>
      </modal> -->


      <!-- <template v-if="Object.keys(ledgerState).length"> -->


        <div>

          <template v-if="ledgerState.products">
          <data-table :data=ledgerState.products height="auto" width="auto">
            <template slot="caption">Products</template>
          </data-table>
          </template>

          <template v-if="ledgerState.products">
          <data-table :data="ledgerState.retailers.map( s =>  ({ Id: s.Id, Products: s.products }) )">
              <template slot="caption">Retailers</template>
          </data-table>
          </template>

          <template v-if="ledgerState.suppliers">
          <data-table :data="ledgerState.suppliers.map( s =>  ({ Id: s.Id, CountryId: s.countryId, OrgId: s.orgId }) )" >
          <!-- <data-table :data="ledgerState.suppliers" > -->
            <template slot="caption">Suppliers</template>
          </data-table>
          </template>
        </div>
          <!-- </template> -->
      <div style="clear: both;margin:auto;margin-top:50px">

        <template v-if="ledgerState.ProductListingContracts">
        <data-table style="margin:auto" :data=ledgerState.ProductListingContracts>
          <template slot="caption">Product Listing Contracts</template>
        </data-table>
        </template>
      </div>
      <div style="margin:auto">
        <template v-if="ledgerState.regulators">

        <data-table  :data=ledgerState.regulators>
          <template slot="caption">Regulators</template>
        </data-table>
        </template>
        <template v-if="ledgerState.importers">
        <data-table  :data="ledgerState.importers.map( s =>  ({ Id: s.Id }) )">
        <!-- <data-table  :data="ledgerState.importers"> -->
          <template slot="caption">Importers</template>
        </data-table>
        </template>

        <template v-if="ledgerState.retailers">
        <data-table  :data="ledgerState.retailers.map( s =>  ({ Id: s.Id, Products: s.products }) )">
        <!-- <data-table  :data="ledgerState.importers"> -->
          <template slot="caption">Retailers</template>
        </data-table>
        </template>
      </div>
      <!-- </template> -->
    </div>
    <!-- </v-card> -->
    <!-- <p>Ledger State</p> -->
    <div class="col-xs-12">
      <div class="well">
        <tree-view :data="ledgerState" :options="{maxDepth: 1, rootObjectKey: 'ledgerState'}"></tree-view>
      </div>
    </div>
    </br></br></br></br>

    <!--  TODO, this is being hidden by datatables since it's fixed -->
    <!-- <div v-if="!isHidden" style="z-index:9000">
      <vue-form
        id="chaincode-form"
        :model="form"
        style="width: 75%; position: fixed; left: 50%; margin-left: -37.5%;">
        <h2 style="float:center"> Invoke Chaincode </h2>
        <vue-form-item label="Function">
          <vue-input
            placeholder="Function"
            v-model="form.function"
            style="width: 100%">
          </vue-input>
        </vue-form-item>

        <vue-form-item label="Arguments">
          <vue-input
            placeholder="Arguments"
            v-model="form.args"
            style="width: 100%">
          </vue-input>
        </vue-form-item>
        <vue-form-item>
          <vue-button type="default" v-on:click="isHidden = true">Cancel</vue-button>
          <vue-button type="success" v-on:click="invoke" >Submit</vue-button>
        </vue-form-item>
      </vue-form>
    </div> -->
      <!-- <vue-form-item> item 1 </vue-form-item>
      <vue-form-item> item 2 </vue-form-item> -->
      <!-- <vue-input placeholder="Please input"></vue-input>
      <vue-input placeholder="Please input"></vue-input> -->



  </div>


</template>

<script>
  import 'vfc/dist/vfc.css'
  import './dist/json-tree.css'

  import {
    Input
  } from 'vfc'
  import {
    Form
  } from 'vfc'
  import {
    FormItem
  } from 'vfc'
  import {
    Button
  } from 'vfc'
  // import DemoLoginModal       from './components/modals/DemoLoginModal.vue'

  // import 'vue-js-modal'
  // import { Card } from 'v-card'
  // import { DataTable } from 'v-data-table'
  // import { Button } from 'vfc'



  export default {
    name: 'app',
    data() {
      return {
        isHidden: false,
        form: {
          function: '',
          args: ''
        },
        args: [],
        ledgerState: {},
        products: [],
        fields: [],
        user_fields: [],
        user_type: '',
        user_input: [],
        input: [],
        func: '',
        title: ''
        // gridData: [{
        //   name: 'Chuck Norris',
        //   power: Infinity
        // }, {
        //   name: 'Bruce Lee',
        //   power: 9000
        // }, {
        //   name: 'Jackie Chan',
        //   power: 7000
        // }, {
        //   name: 'Jet Li',
        //   power: 8000
        // }]
        // gridData: [{"id":"product1","quantity":"300","countryId":"US"},{"id":"product2","quantity":"350","countryId":"US"}]
      }
    },
    components: {
      Form,
      FormItem,
      // DemoLoginModal,
      [Input.name]: Input,
      [Button.name]: Button
    },
    methods: {
      invoke() {
        // console.log(this.$data.input)
        // console.log(this.$data.user_input)
        console.log(this.$data.input.concat(this.$data.user_input))
        var options = {
          method: "POST",
          body: JSON.stringify({
            method: "invoke",
            params: {
              ctorMsg: {
                function: this.$data.func, //this.$data.input[0],
                args: this.$data.input.concat(this.$data.user_input)
                // function: this.$data.form.function,
                // args: this.$data.form.args.split(',') // ["retailer1", "retailer"]
              }
            }
          }),
          headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
          }
        }
        fetch("http://localhost:30001/api/chaincode", options).then((response) => {
          console.log("api call complete")
          // this.$data.isHidden = true
          this.$modal.hide('invoke-modal');
          this.$data.input = []
          this.$data.user_input = []
        })


      },
      getLedger() {
        console.log(this.$data.form.function)
        console.log("publishing")

        var options = {
          method: "POST",
          body: JSON.stringify({
            method: "query",
            params: {
              ctorMsg: {
                function: "read_everything",
                args: []
              }
            }
          }),
          headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
          }
        }
        fetch("http://localhost:30001/api/chaincode", options).then((response) => {
          console.log("ledger state retrieved")
          response.json().then((json) => {
            console.log(json)
            var result = JSON.parse(json)
            // for (reg in json.regulators) {
            //   delete reg.id
            // }
            // for (reg in json.retailers) {
            //   delete reg.id
            // }

            this.$data.ledgerState = result
            // this.$data.products = JSON.parse(json.products)
            // console.log("json.products")
            // console.log(json.products)
          })

        })
      },
      showInvokeModal(config) {
        console.log("opening modal")
        // console.log(fields)
        this.$data.input = []
        console.log(config.function)
        console.log(config.fields)
        this.$data.func = config.function
        this.$data.fields = config.fields
        this.$data.title = config.title
        this.$data.user_fields = []
        this.$data.user_type = ''
        this.$modal.show('invoke-modal', { "fields": config.fields });
      },
      hideModal() {
        this.$modal.hide('invoke-modal');
        console.log("hiding modal")
        this.$data.user_fields = []
        this.$data.user_type = ''
      },
      getFormValues () {
        console.log("getting form vals")
        console.log(this.$data.input)
        // this.output = this.$refs.my_input.value
        // console.log(this.$refs.my_input.value)
      },
      addFields () {
        this.$data.input.push(this.$data.user_type)
        if (this.user_type == "supplier") {
          this.$data.user_fields.push('Country Id')
          this.$data.user_fields.push('Org Id')
        } else {
          this.$data.user_fields = []
        }
      }

    },
    filters: {
      pretty: function(value) {
        return JSON.stringify(value, null, 2);
      }
    }


  }
</script>

<!-- TODO, finish modal -->
<!-- <script type="text/x-template" id="modal-template">
  <transition name="modal">
    <div class="modal-mask">
      <div class="modal-wrapper">
        <div class="modal-container">

          <div class="modal-header">
            <slot name="header">
              default header
            </slot>
          </div>

          <div class="modal-body">
            <slot name="body">
              default body
            </slot>
          </div>

          <div class="modal-footer">
            <slot name="footer">
              default footer
              <button class="modal-default-button" @click="$emit('close')">
                OK
              </button>
            </slot>
          </div>
        </div>
      </div>
    </div>
  </transition>
</script> -->


<style>
  #app {
    font-family: 'Avenir', Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    text-align: center;
    color: #2c3e50;
    margin-top: 60px;
  }
</style>
