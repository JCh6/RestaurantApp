import Vue from 'vue';
import Vuex from 'vuex';
import axios from "axios";

Vue.use(Vuex)

export default new Vuex.Store({
    state: {
        activeViewName: "",
        alertInfo: {
            show: false,
            type: "info",
            text: ""
        },
        infoLoaded: {
            numProducts: 0,
            numBuyers: 0,
            numTransactions: 0,
            date: ""
        }
    },
    mutations: {
        setActiveViewName(state, viewName) {
            state.activeViewName = viewName;
        },
        showAlert(state, info) {
            state.alertInfo.show = true;
            state.alertInfo.type = info.type;
            state.alertInfo.text = info.text;
        },
        hideAlert(state) {
            state.alertInfo.show = false;
        },
        setInfoLoaded(state, newInfo) {
            state.infoLoaded.numProducts = newInfo.numProducts;
            state.infoLoaded.numBuyers = newInfo.numBuyers;
            state.infoLoaded.numTransactions = newInfo.numTransactions;
            state.infoLoaded.date = newInfo.lastDate;
        },
    },
    actions: {
        getJSON(state, req) {
            let url = process.env.VUE_APP_MY_ENDPOINT;
            return axios.get(url + req.endpoint, { params: req.params });
        }
    },
    modules: {
    }
})
