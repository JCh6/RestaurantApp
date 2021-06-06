<template>
    <v-container fill-height>
        <v-row>
            <v-col align="center">
                <v-date-picker
                    elevation="5"
                    v-model="date"
                    :max="maxDate"
                    show-adjacent-months
                    :landscape="true"
                    color="green lighten-1"
                />
                <v-col align="center">
                    <v-btn
                        @click="load(date)"
                        color="blue-grey"
                        :loading="loading"
                        elevation="2"
                        rounded
                        dark
                        >Load Transactions
                        <v-icon class="ml-2">cloud_download</v-icon>
                    </v-btn>
                </v-col>
            </v-col>
        </v-row>
    </v-container>
</template>

<script>
import { mapMutations, mapActions } from "vuex";

export default {
    name: "Home",
    data: () => {
        return {
            date: new Date().toISOString().substr(0, 10),
            maxDate: new Date().toISOString().substr(0, 10),
            loading: false,
        };
    },
    methods: {
        ...mapMutations(["setActiveViewName", "showAlert", "setInfoLoaded"]),
        ...mapActions(["getJSON"]),
        async load(selectedDate) {
            this.loading = true;
            const newAlert = { type: "info" };
            try {
                const req = {
                    endpoint: "load",
                    params: { date: new Date(selectedDate).getTime() },
                };
                let resp = await this.getJSON(req);

                if (resp.data && resp.data.code === 200) {
                    let body = resp.data.body;
                    this.setInfoLoaded({
                        numProducts: body.aggregateProduct.count,
                        numBuyers: body.aggregateBuyer.count,
                        numTransactions: body.aggregateTransaction.count,
                        lastDate: selectedDate,
                    });
                    newAlert.text = "Data successfully loaded";
                } else {
                    newAlert.type = "error";
                    newAlert.text = resp.data.errorMessage;
                }
            } catch (e) {
                newAlert.text = e;
                newAlert.type = "error";
            } finally {
                this.loading = false;
                this.showAlert(newAlert);
            }
        },
    },
    mounted() {
        this.setActiveViewName("Calendar");
    },
};
</script>
