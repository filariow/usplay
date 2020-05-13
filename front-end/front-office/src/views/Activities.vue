<template>
  <div class="Activities">
    <Activity />
  </div>
</template>

<script lang="ts">
import Activity from "@/components/Activity.vue";
import {
  ActivityTypeSvcClient,
  ServiceError
} from "../gen/activitytypecomm/activitytype_pb_service";
import {
  ListActivityTypesRequest,
  ListActivityTypesReply
} from "../gen/activitytypecomm/activitytype_pb";

function listActivityTypes() {
  console.log("hello");

  const acttypeHost = process.env.VUE_APP_US_ACTTYPE_HOST;
  if (acttypeHost == undefined) {
    console.log("Environment variable VUE_APP_US_ACTTYPE_HOST is not set");
    return;
  }
  console.log(acttypeHost);

  const client = new ActivityTypeSvcClient(acttypeHost);
  const request = new ListActivityTypesRequest();

  console.log(request);

  client.list(
    request,
    (err: ServiceError | null, message: ListActivityTypesReply | null) => {
      console.log("Client list ");
      if (err != null) {
        console.log(`Error ${err.code}: ${err.message}`);
      } else {
        console.log("Error is null");
      }
      console.log(message?.toObject());
    }
  );

  console.log("Request sent");
}

export default {
  name: "Activities",

  components: {
    Activity
  },

  methods: {},

  mounted() {
    listActivityTypes();
  },

  data() {
    return {};
  }
};
</script>
