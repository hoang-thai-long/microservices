<template>
  <div class="hello">
    <tracking-video :userid="'longht2'" :videoid="'Hlz93KRJv00'" type="youtube" :classid="'classid'" url="https://report.eduso.vn/api/video"></tracking-video>
  </div>
</template>

<script lang="ts">

import { Component, Prop, Vue } from 'vue-property-decorator';
import TrackingVideo from './TrackingVideo.vue';
@Component({
  components: {
    TrackingVideo
  }
})
export default class HelloWorld extends Vue {
  @Prop() private msg!: string;
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}

ul {
  list-style-type: none;
  padding: 0;
}

li {
  display: inline-block;
  margin: 0 10px;
}

a {
  color: #42b983;
}
</style>
