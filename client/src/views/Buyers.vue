<template>
    <v-container>
        <BackButton
            :justify="justifyBackBtn"
            icon="date_range"
            msg="Back to calendar"
            color="info"
            route="/"
        />
        <div v-if="showTable">
            <DataTable
                :header="headers"
                :data="buyers"
                :handleClick="handleClick"
            />
            <div class="text-center pt-2">
                <v-pagination
                    color="green lighten-1"
                    v-model="page"
                    :length="pageCount"
                    circle
                ></v-pagination>
            </div>
        </div>
    </v-container>
</template>

<script>
import { mapState, mapMutations, mapActions } from "vuex";
import BackButton from "@/components/BackButton.vue";
import DataTable from "@/components/DataTable.vue";

export default {
    name: "Buyers",
    components: {
        BackButton,
        DataTable,
    },
    data() {
        return {
            showTable: true,
            justifyBackBtn: "start",
            page: 1,
            pageSize: 10,
            pageCount: 0,
            headers: [],
            buyers: [],
        };
    },
    methods: {
        ...mapMutations(["setActiveViewName", "showAlert"]),
        ...mapActions(["getJSON"]),
        async getBuyersData() {
            if (this.showTable) {
                try {
                    let resp = null;
                    const req = {
                        endpoint: "buyers",
                        params: {
                            page: this.page,
                            limit: this.pageSize,
                        },
                    };

                    resp = await this.getJSON(req);

                    if (resp.data && resp.data.code === 200) {
                        let body = resp.data.body;
                        this.setHeader();
                        this.buyers = body.queryBuyer;
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
                }
            }
        },
        setHeader() {
            this.headers = [
                { text: "ID", value: "id" },
                { text: "Name", value: "name" },
                { text: "Age", value: "age" },
            ];
        },
        handleClick(row) {
            this.$router.push(`/buyers/${row.id}/transactions`);
        },
    },
    computed: {
        ...mapState(["infoLoaded"]),
    },
    watch: {
        page() {
            this.getBuyersData();
        },
    },
    created() {
        this.setActiveViewName("Buyers");
        this.pageCount = Math.round(this.infoLoaded.numBuyers / this.pageSize);

        if (this.infoLoaded.numBuyers == 0) {
            this.showTable = false;
            this.justifyBackBtn = "center";
            this.showAlert({
                type: "warning",
                text: "There is no data loaded. Please select a new date.",
            });
        }
    },
    mounted() {
        this.getBuyersData();
    },
};
</script>