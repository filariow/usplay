<template>
  <div class="ActivityTypes">
    <ActivityTypes :types="activitytypes" />
  </div>
</template>

<script lang="ts">
import ActivityTypes from "@/components/ActivityTypes.vue";
import Vue from "vue";
import {
  ActivityTypeSvcClient,
  ServiceError,
} from "../gen/bookmastergrpc/activitytype_pb_service";
import {
  ListActivityTypesRequest,
  ListActivityTypesReply,
} from "../gen/bookmastergrpc/activitytype_pb";
import Component from "vue-class-component";

export class ActivityTypeViewModel {
  Name: string;
  Code: number;
  Description: string;

  /**
   *
   */
  constructor(name: string, code: number, description: string) {
    this.Name = name;
    this.Code = code;
    this.Description = description;
  }
}

@Component({
  components: { ActivityTypes }
})
export default class ActivityTypesList extends Vue {
  private activitytypes = new Array<ActivityTypeViewModel>();

  fetchData() {
    const acttypeHost = process.env.VUE_APP_US_ACTTYPE_HOST;
    if (acttypeHost == undefined) {
      console.log("Environment variable VUE_APP_US_ACTTYPE_HOST is not set");
      return;
    }

    const client = new ActivityTypeSvcClient(acttypeHost);
    const request = new ListActivityTypesRequest();

    const call = client.list(
      request,
      (err: ServiceError | null, message: ListActivityTypesReply | null) => {
        if (err != null) {
          console.log(`Error ${err.code}: ${err.message}`);
          return;
        }

        const atList = message?.toObject()?.activitytypesList;
        atList?.forEach((at, _) => {
          const avm = new ActivityTypeViewModel(
            at.name,
            at.code,
            at.description
          );
          this.activitytypes.push(avm);
        });
      }
    );
  }

  created() {
    this.fetchData();
  }
}
</script>
