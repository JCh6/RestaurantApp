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
            </v-col>
        </v-row>
    </v-container>
</template>

<script>
export default {
    name: "Home",
    data: () => {
        return {
            date: new Date().toISOString().substr(0, 10),
            maxDate: new Date().toISOString().substr(0, 10),
        };
    },
};
</script>
