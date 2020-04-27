<template>
  <v-container>
    <v-row class="text-center">
      <v-col>
        <v-data-table
          :headers="headers"
          :items="activities"
          :items-per-page="5"
          class="elevation-1"
        ></v-data-table>
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts">
import Vue from "vue";

export default Vue.extend({
  name: "Activity",

  data: () => ({
    headers: [
      {
        text: "name",
        align: "start",
        sortable: false,
        value: "name",
      },
      {
        text: "code",
        align: "start",
        sortable: false,
        value: "code",
      },
    ],
    activities: [
      { name: "zero", code: 0 },
      { name: "one", code: 1 },
      { name: "two", code: 2 },
      { name: "three", code: 3 },
      { name: "four", code: 4 },
      { name: "five", code: 5 },
      { name: "six", code: 6 },
      { name: "seven", code: 7 },
      { name: "eigth", code: 8 },
      { name: "nine", code: 9 },
    ],
  }),
});
</script>
