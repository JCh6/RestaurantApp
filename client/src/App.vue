<template>
    <v-app>
        <Header app />
        <v-main>
            <Alert />
            <router-view />
        </v-main>
        <Footer app />
    </v-app>
</template>

<script>
import Header from "./components/Header.vue";
import Alert from "./components/Alert.vue";
import Footer from "./components/Footer.vue";

export default {
    name: "App",
    components: {
        Header,
        Alert,
        Footer,
    },
};
</script>