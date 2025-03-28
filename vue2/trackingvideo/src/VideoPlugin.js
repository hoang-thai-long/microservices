import TrackingVideo from "./components/TrackingVideo.vue";
import VueYouTube from "vue-youtube";
const VideoPlugin = {
  install(Vue, options) {
    console.log("tracking-video");
    Vue.component("tracking-video", TrackingVideo);
  },
};

if (typeof window !== 'undefined' && window.Vue) {
  // window.VideoPlugin = VideoPlugin;
  // window.TrackingVideo = TrackingVideo;
  // window.Vue.use(VueYouTube);
  window.Vue.use(VideoPlugin);
}

export default VideoPlugin;
