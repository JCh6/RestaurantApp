<template>
    <v-row class="my-3" :justify="justify">
        <v-btn :color="color" elevation="2" @click="$router.push(route)">
            <v-icon class="mr-2">{{ icon }}</v-icon>
            {{ msg }}
        </v-btn>
    </v-row>
</template>

<script>
export default {
    name: "BackButton",
    props: {
        icon: String,
        msg: String,
        justify: String,
        route: String,
        color: String,
    },
};
</script>