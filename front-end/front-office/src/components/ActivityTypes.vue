<template>
  <v-container>
    <v-row class="text-center">
      <v-col>
        <v-data-table
          :headers="headers"
          :items="types"
          :items-per-page="10"
          class="elevation-1"
        ></v-data-table>
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";

import { Prop } from "vue-property-decorator";
import ActivityTypeViewModel from "@/views/ActivityTypeList.vue";

@Component
export default class ActivityTypes extends Vue {
  @Prop({ default: "Types" }) types!: Array<ActivityTypeViewModel>;
 
  headers = [
    { text: "name", align: "start", sortable: true, value: "Name" },
    { text: "code", align: "start", sortable: true, value: "Code" },
    { text: "description", align: "start", sortable: false, value: "Description" }
  ];
}
</script>
