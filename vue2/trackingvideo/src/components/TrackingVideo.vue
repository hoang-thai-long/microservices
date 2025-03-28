<template>
    <!-- <youtube-video id="test"></youtube-video> -->
    <div>
        <youtube-video v-if="show" :id="videoid" :userid="userid" :extendsid="classid" :onChangeState="onChangeState" :ctrl="funCtr(url, userid)" />
        <v-skeleton v-else width="100%" height="100%"></v-skeleton>
    </div>
</template>
<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';
import YoutubeVideo from './YoutubeVideo.vue';
import controlVideoByServer from '../common/control';
import VSkeleton from './VSkeleton.vue';
@Component({
    name: 'tracking-video',
    components: {
        YoutubeVideo,
        VSkeleton
    }
})
export default class TrackingVideo extends Vue {
    @Prop({
        required: true
    }) private classid!: string;
    @Prop({
        required: true
    }) private userid!: string;
    @Prop({
        required: true
    }) private videoid!: string;
    @Prop({
        required: true
    }) private type !: string;
    @Prop({
        required: true
    }) private url !: string;

    public show = false;
    public funCtr = controlVideoByServer;
    public onChangeState = (state: string) => {
        console.log(state);
    }
    mounted() {
        this.show = true;
    }
}
</script>