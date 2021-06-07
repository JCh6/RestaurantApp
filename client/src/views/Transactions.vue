<template>
    <v-container>
        <BackButton
            v-if="!showCard"
            justify="center"
            icon="date_range"
            msg="Back to calendar"
            color="info"
            route="/"
        />
        <div v-if="showCard">
            <BackButton
                justify="start"
                icon="people"
                msg="Back to buyers"
                route="/buyers"
                color="info"
            />
            <v-row align="center">
                <v-col md="2" align="end">
                    <v-btn icon x-large outlined>
                        <v-icon>chevron_left</v-icon>
                    </v-btn>
                </v-col>
                <v-col md="8">
                    <v-card elevation="2" shaped>
                        <v-list-item three-line>
                            <v-list-item-content>
                                <div class="text-overline mb-4">
                                    Transaction ID: {{ transactionId }}
                                </div>
                                <v-list-item-title>
                                    <v-chip
                                        class="ma-2"
                                        color="green"
                                        text-color="white"
                                    >
                                        <v-avatar left class="green darken-4">
                                            {{ quantity }}
                                        </v-avatar>
                                        Products
                                    </v-chip>
                                    <v-chip
                                        class="ma-2"
                                        color="indigo"
                                        text-color="white"
                                    >
                                        <v-avatar left class="indigo darken-4">
                                            <v-icon>credit_card</v-icon>
                                        </v-avatar>
                                        {{ formatPrice() }}
                                    </v-chip>
                                </v-list-item-title>
                            </v-list-item-content>
                            <v-list-item-action>
                                <v-list-item-action-text v-text="ip" />
                                <v-list-item-action-text v-text="device" />
                            </v-list-item-action>
                        </v-list-item>
                        <v-card-actions>
                            <v-btn icon rounded>
                                <v-icon>expand_more</v-icon>
                            </v-btn>
                        </v-card-actions>
                    </v-card>
                </v-col>
                <v-col md="2" align="start">
                    <v-btn icon x-large outlined>
                        <v-icon>chevron_right</v-icon>
                    </v-btn>
                </v-col>
            </v-row>
        </div>
    </v-container>
</template>

<script>
import { mapState, mapMutations } from "vuex";
import BackButton from "@/components/BackButton.vue";

export default {
    name: "Transactions",
    data() {
        return {
            showCard: true,
            transactionId: "000060bd6bb7",
            device: "linux",
            ip: "11.224.157.163",
            quantity: 5,
            totalPrice: 28754,
        };
    },
    components: {
        BackButton,
    },
    methods: {
        ...mapMutations(["setActiveViewName", "showAlert"]),
        formatPrice() {
            let formatter = new Intl.NumberFormat("en-US", {
                style: "currency",
                currency: "USD",
            });
            return formatter.format(this.totalPrice);
        },
    },
    computed: {
        ...mapState(["infoLoaded"]),
    },
    created() {
        let id = this.$route.params.buyerId;
        this.setActiveViewName(`Buyer (${id}) Transactions`);

        if (this.infoLoaded.numBuyers == 0) {
            /*this.showCard = false;
            this.showAlert({
                type: "warning",
                text: "There is no data loaded. Please select a new date.",
            });*/
        }
    },
};
</script>
