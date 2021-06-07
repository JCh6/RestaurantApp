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
                <v-col md="3" align="end">
                    <v-btn
                        :disabled="page == 1 || reveal"
                        icon
                        x-large
                        outlined
                        @click="page--"
                    >
                        <v-icon>chevron_left</v-icon>
                    </v-btn>
                </v-col>
                <v-col md="6">
                    <v-card align="center" outlined :loading="cardLoading">
                        <v-card-text>
                            <span>{{ `# ${page}` }}</span> <br />
                            <div>Transaction ID</div>
                            <p class="text-h4 text--primary">
                                {{ trx.transactionId }}
                            </p>
                            <span>{{ trx.device }}</span> <br />
                            <span>{{ trx.ip }}</span>
                        </v-card-text>
                        <v-chip class="ma-2" color="indigo" text-color="white">
                            <v-avatar left class="indigo darken-4">
                                {{ trx.quantity }}
                            </v-avatar>
                            Products
                        </v-chip>
                        <v-chip class="ma-2" color="indigo" text-color="white">
                            <v-avatar left class="indigo darken-4">
                                <v-icon>credit_card</v-icon>
                            </v-avatar>
                            {{ formatPrice(trx.totalPrice) }}
                        </v-chip>
                        <v-card-actions>
                            <v-btn
                                :disabled="numMaxCards == 0"
                                icon
                                @click="reveal = !reveal"
                            >
                                <v-icon>{{
                                    reveal ? "expand_less" : "expand_more"
                                }}</v-icon>
                            </v-btn>
                        </v-card-actions>

                        <v-expand-transition>
                            <div v-show="showMoreInfo" align="start">
                                <v-divider />
                                <v-subheader>Products</v-subheader>
                                <v-list-item
                                    two-line
                                    v-for="(product, index) in trx.products"
                                    v-bind:key="index"
                                >
                                    <v-list-item-content>
                                        <v-list-item-title
                                            v-text="product.name"
                                        ></v-list-item-title>
                                        <v-list-item-subtitle
                                            v-text="formatPrice(product.price)"
                                        ></v-list-item-subtitle>
                                    </v-list-item-content>
                                </v-list-item>
                                <v-divider />
                                <v-subheader
                                    >Buyers using the same IP</v-subheader
                                >
                                <v-list-item
                                    v-for="(buyer, index) in trx.otherBuyers"
                                    v-bind:key="buyer + index"
                                >
                                    <v-list-item-content>
                                        <v-list-item-title
                                            v-text="buyer"
                                        ></v-list-item-title>
                                    </v-list-item-content>
                                </v-list-item>
                                <v-divider />
                                <v-subheader
                                    >Other people also bought</v-subheader
                                >
                                <v-list-item
                                    v-for="(product, index) in trx.recommended"
                                    v-bind:key="product + index"
                                >
                                    <v-list-item-content>
                                        <v-chip label v-text="product" />
                                    </v-list-item-content>
                                </v-list-item>
                            </div>
                        </v-expand-transition>
                    </v-card>
                </v-col>
                <v-col md="3" align="start">
                    <v-btn
                        :disabled="page >= numMaxCards || reveal"
                        icon
                        x-large
                        outlined
                        @click="page++"
                    >
                        <v-icon>chevron_right</v-icon>
                    </v-btn>
                </v-col>
            </v-row>
        </div>
    </v-container>
</template>

<script>
import { mapState, mapMutations, mapActions } from "vuex";
import BackButton from "@/components/BackButton.vue";

export default {
    name: "Transactions",
    data() {
        return {
            showCard: true,
            reveal: false,
            cardLoading: false,
            showMoreInfo: false,
            page: 1,
            pageSize: 1,
            recommendedMaxSize: 5,
            buyerId: "",
            trx: {
                transactionId: "Not found",
                device: "",
                ip: "",
                quantity: 0,
                totalPrice: 0,
                products: [],
                otherBuyers: [],
                recommended: [],
            },
            numMaxCards: 0,
        };
    },
    components: {
        BackButton,
    },
    methods: {
        ...mapMutations(["setActiveViewName", "showAlert"]),
        ...mapActions(["getJSON"]),
        formatPrice(price) {
            let formatter = new Intl.NumberFormat("en-US", {
                style: "currency",
                currency: "USD",
            });
            return formatter.format(price);
        },
        async getTransactions() {
            this.cardLoading = true;
            if (this.showCard) {
                try {
                    let resp = null;
                    const req = {
                        endpoint: "transactions",
                        params: {
                            buyer: this.buyerId,
                            page: this.page,
                            limit: this.pageSize,
                        },
                    };

                    resp = await this.getJSON(req);

                    if (resp.data && resp.data.code === 200) {
                        let { aggregateTransaction, queryTransaction } =
                            resp.data.body;
                        let trx = queryTransaction[0];
                        this.numMaxCards = aggregateTransaction.count;
                        if (aggregateTransaction.count > 0) {
                            this.trx.transactionId = trx.id;
                            this.trx.ip = trx.ip;
                            this.trx.device = trx.device;
                            this.trx.quantity = trx.productsAggregate.count;
                            this.trx.totalPrice =
                                trx.productsAggregate.priceSum;
                        }
                    } else {
                        this.showAlert({
                            type: "error",
                            text: resp.data.errorMessage,
                        });
                    }
                } catch (err) {
                    this.showAlert({
                        type: "error",
                        text: err,
                    });
                } finally {
                    this.cardLoading = false;
                }
            }
        },
        async getMoreInfo() {
            this.cardLoading = true;
            try {
                let resp = null;
                const req = {
                    endpoint: "transactions",
                    params: {
                        id: this.trx.transactionId,
                        ip: this.trx.ip,
                    },
                };

                resp = await this.getJSON(req);

                if (resp.data && resp.data.code === 200) {
                    let { buyerInfo, getTransaction, other } = resp.data.body;
                    this.trx.products = getTransaction.products;
                    this.setOtherBuyers(buyerInfo.queryBuyer);
                    this.setRecommended(other);
                    this.showMoreInfo = true;
                } else {
                    this.showAlert({
                        type: "error",
                        text: resp.data.errorMessage,
                    });
                }
            } catch (err) {
                this.showAlert({
                    type: "error",
                    text: err,
                });
            } finally {
                this.cardLoading = false;
            }
        },
        setOtherBuyers(buyers) {
            let names = [];
            for (let buyer of buyers) {
                names.push(buyer.name);
            }
            this.trx.otherBuyers = names;
        },
        setRecommended(otherProducts) {
            let recProducts = new Set();

            for (let buyer of otherProducts) {
                for (let product of buyer.products) {
                    recProducts.add(product.name);
                }
            }

            for (let product of this.trx.products) {
                recProducts.delete(product.name);
            }

            this.trx.recommended = 
                Array.from(recProducts).slice(0, this.recommendedMaxSize);
        },
    },
    computed: {
        ...mapState(["infoLoaded"]),
    },
    watch: {
        page() {
            this.getTransactions();
        },
        reveal() {
            if (this.reveal) {
                this.getMoreInfo();
            } else this.showMoreInfo = false;
        },
    },
    created() {
        this.buyerId = this.$route.params.buyerId;
        this.setActiveViewName(`Buyer: (${this.buyerId})`);

        if (this.infoLoaded.numBuyers == 0) {
            /*this.showCard = false;
            this.showAlert({
                type: "warning",
                text: "There is no data loaded. Please select a new date.",
            });*/
        }
    },
    mounted() {
        this.getTransactions();
    },
};
</script>