<template>
    <v-app-bar app dense>
        <div class="text-md-h5 font-weight-light">Vue {{ activeViewName }}</div>
        <v-spacer />
        <v-tooltip bottom>
            <template v-slot:activator="{ on, attrs }">
                <v-icon color="green lighten-1" dark v-bind="attrs" v-on="on">
                    info
                </v-icon>
            </template>
            <span>Products: {{ infoLoaded.numProducts }}</span><br />
            <span>Buyers: {{ infoLoaded.numBuyers }}</span><br />
            <span>Transactions: {{ infoLoaded.numTransactions }}</span><br />
            <span v-if="infoLoaded.date != ''">Date: {{ infoLoaded.date }}</span> 
        </v-tooltip>
    </v-app-bar>
</template>

<script>
import { mapState } from "vuex";

export default {
    name: "Header",
    computed: {
        ...mapState(["activeViewName", "infoLoaded"]),
    },
};
</script>
