<template>
    <div class="skeleton" :style="{
        width: width,
        height: height
    }">
    </div>
</template>
<style>
.skeleton {
    position: relative;
    overflow: hidden;

    background: #ccc;
    border-radius: 3px;
}

.skeleton:after {
    content: "";
    position: absolute;

    height: 100%;
    width: 100px;
    left: -100px;
    top: 0;

    background: linear-gradient(to right, transparent, #ffffff70, transparent);
    animation: reflect 800ms ease-out infinite;
}

@keyframes reflect {
    to {
        left: calc(100% + 100px);
    }
}
</style>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';

@Component({
    name: "v-skeleton"
})
export default class VSkeleton extends Vue {
    @Prop() private width!: number
    @Prop() private height!: number
}
</script>