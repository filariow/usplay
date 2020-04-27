<template>
  <div class="Activities">
    <Activity />
  </div>
</template>

<script>
// @ is an alias to /src
import Activity from "@/components/Activity.vue";

export default {
  name: "Activities",

  components: {
    Activity,
  },
};
</script>
