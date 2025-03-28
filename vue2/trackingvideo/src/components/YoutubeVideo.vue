<template>
    <div>
        <youtube v-if="!show" :player-vars="{
            autoplay: 0
        }" :id="videoID" ref="videoRef" :video-id="id" @playing="playing" @end="ended" @paused="paused"
            @ready="onPlayerReady" />

        <v-skeleton v-else width="100%" height="319px"></v-skeleton>
    </div>
</template>

<script lang="ts">



import { ControlVideo } from '@/common/control';
import { Component, Prop, Vue } from 'vue-property-decorator';
@Component({
    name: "youtube-video",
})
export default class YoutubeVideo extends Vue {
    @Prop({
        required: true
    }) private ctrl !: ControlVideo
    @Prop() private onChangeState !: (value: string) => void
    @Prop({
        required: true
    }) private id !: string
    @Prop({
        required: true
    }) private userid !: string
    @Prop({
        required: true
    }) private extendsid !: string

    public intervalPauseNumber = 0
    public intervalReadyNumber = 0
    public intervalPlayingNumber = 0

    public timeVideo = 0
    public state: "playing" | "pause" | "ended" | "ready" | "none" = "none"
    public interval = 0
    public show = false
    public videoRef: any
    public videoID = `${this.$props.userid}${this.$props.id}${this.$props.extendsid}`

    async created() {
        this.videoID = `${this.$props.userid}${this.$props.id}${this.$props.extendsid}`
        await this.ctrl.loadtimevideoServer(this.videoID)
        //this.videoRef = this.$refs.videoRef;
        this.videoRef = (this.$refs.videoRef as any)?.player;
    }

    public onloadVideoTime = () => {
        //await this.ctrl.loadtimevideoServer(this.videoID);
        this.timeVideo = parseFloat(localStorage.getItem(this.videoID) ?? "0");
        this.show = true;
        this.seekTo(this.timeVideo);
    }
    public seekTo(value: number) {
        const player = (this.$refs.videoRef as any).player;
        player?.seekTo(value, true)
    }

    private intervalReady = async () => {
        console.log(this.state);
        if (this.state === "ready") {
            const player = (this.$refs.videoRef as any).player;
            const currentTime = await player?.getCurrentTime();

            if (currentTime > this.timeVideo) {
                player?.seekTo(this.timeVideo, true)
            }
        }
    }
    private intervalPause = async () => {
        console.log(this.state);
        if (this.state === "pause") {
            const player = (this.$refs.videoRef as any).player;
            const currentTime = await player?.getCurrentTime();

            if (currentTime > this.timeVideo) {
                player?.seekTo(this.timeVideo, true)
            }
        }
    }

    public onPlayerReady = () => {
        if (this.state !== "ready") {
            if (this.intervalPlayingNumber) clearInterval(this.intervalPlayingNumber);
            if (this.intervalPauseNumber) clearInterval(this.intervalPauseNumber);

            this.seekTo(this.timeVideo);
            if (!this.intervalReadyNumber) this.intervalReadyNumber = setInterval(this.intervalReady.bind(this), 500);
        }
        this.state = "ready";
        // this.onChangeState(this.state)
    }
    public ended = async () => {
        if(this.state !== "ended"){
            const player = (this.$refs.videoRef as any).player;
            const duration = await player?.getDuration();
            this.timeVideo = parseFloat(localStorage.getItem(this.videoID) ?? "0");
            if(duration > this.timeVideo - 15){
                this.seekTo(this.timeVideo);
            }
            else{
                this.ctrl.timeupdate(this.videoID,duration);
            }
        }
        this.state = "ended";
        // this.onChangeState(this.state)
    }
    public paused = async () => {
        if (this.state !== "pause") {
            if (this.intervalPlayingNumber) clearInterval(this.intervalPlayingNumber);
            if (this.intervalReadyNumber) clearInterval(this.intervalReadyNumber);
            // this.onChangeState(this.state)
            if (!this.intervalPauseNumber) this.intervalPauseNumber = setInterval(this.intervalPause.bind(this), 500);
        }
        this.state = "pause";

    }

    private updateTime = async () => {
        console.log(this.state);
        if (this.state === "playing") {
            const player = (this.$refs.videoRef as any).player;
            const time = await player?.getCurrentTime();
            this.timeVideo = parseFloat(localStorage.getItem(this.videoID) ?? "0");
            if (time > this.timeVideo) {
                this.timeVideo = time;
                this.$props.ctrl.timeupdate(this.videoID, time)
                localStorage.setItem(this.videoID, time);
            }
        }
    }

    public playing = async () => {
        if (this.state !== "playing") {
            const player = (this.$refs.videoRef as any).player;
            const time = await player?.getCurrentTime();
            const duration = await player?.getDuration();
            this.timeVideo = parseFloat(localStorage.getItem(this.videoID) ?? "0");

            if (duration > this.timeVideo) {
                if (time > this.timeVideo) {
                    this.seekTo(this.timeVideo);
                }
            }
        }
        this.state = "playing";
        if (this.intervalReadyNumber) clearInterval(this.intervalReadyNumber);
        if (this.intervalPauseNumber) clearInterval(this.intervalPauseNumber);
        // this.onChangeState(this.state);
        if (!this.intervalPlayingNumber) this.intervalPlayingNumber = setInterval(this.updateTime.bind(this), 10000);
    }
    public visibilityChange = () => {
        if (document.hidden) {
            if (this.intervalReadyNumber) clearInterval(this.intervalReadyNumber);
            if (this.intervalPauseNumber) clearInterval(this.intervalPauseNumber);
            if (this.intervalPlayingNumber) clearInterval(this.intervalPlayingNumber);
            (this.$refs.videoRef as any).player?.pauseVideo();
        }
    }
    mounted() {
        this.onloadVideoTime();
        document.addEventListener("visibilitychange", this.visibilityChange);
    }
}
</script>