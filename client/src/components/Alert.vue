<template>
    <v-alert
        v-if="alertInfo.show"
        elevation="1"
        :type="alertInfo.type"
        dense
        text
        dismissible
        @click="hideAlert"
    >
        {{ alertInfo.text }}
    </v-alert>
</template>

<script>
import { mapState, mapMutations } from "vuex";

export default {
    name: "Alert",
    methods: {
        ...mapMutations(["hideAlert"]),
    },
    computed: {
        ...mapState(["alertInfo"]),
    },
};
</script>
