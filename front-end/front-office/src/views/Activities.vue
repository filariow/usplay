<template>
  <div class="Activities">
    <Activity />
  </div>
</template>

<script lang="ts">
import Activity from "@/components/Activity.vue";
import { ActivityTypeSvcClient } from "@/gen/comm/activitytypecomm/ActivitytypeServiceClientPb";
import {
  ListActivityTypesRequest,
  ListActivityTypesReply,
} from "@/gen/comm/activitytypecomm/activitytype_pb";

export default {
  name: "Activities",

  components: {
    Activity,
  },

  listActivityTypes() {
    console.log("hello");

    let client = new ActivityTypeSvcClient(process.env.US_ACTTYPE_HOST);
    var request = new ListActivityTypesRequest();

    client.list(request, {}, (err, response) => {
      console.log(response);
    });
  },

  mounted() {
    this.listActivityTypes();
  },

  data() {
    return {};
  },
};
</script>
./src/gen/comm/activitytypecomm/ActivitytypeServiceClientPb.ts