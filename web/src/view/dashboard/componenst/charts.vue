<!--
    @auther: bypanghu<bypanghu@163.com>
    @date: 2024/5/8
!-->

<template>
  <div class="">
    <div class="flex items-center justify-between mb-2">
      <div v-if="title" class="font-bold">
        {{ title }}
      </div>
      <slot v-else name="title" />
    </div>
    <div class=" w-full relative h-full">
      <div class="top-1 w-[100%] h-80" v-if="dataLoaded">
        <charts-people-number v-if="type === 0" :data="data[0]" height="100%" />
        <charts-people-number v-if="type === 1" :data="data[1]" height="100%" />
        <charts-people-number v-if="type === 2" :data="data[2]" height="100%" />
        <charts-people-number v-if="type === 3" :data="data[3]" height="100%" />
        <charts-people-number v-if="type === 4" :data="data[4]" height="100%" />
        <charts-people-number v-if="type === 5" :data="data[5]" height="100%" />
        <charts-content-number v-if="type === 6" />
      </div>

    </div>
  </div>
</template>

<script setup>

import {ref} from "vue"
import chartsPeopleNumber from "./charts-people-numbers.vue"
import chartsContentNumber from "./charts-content-numbers.vue"
import {
  getScanlogTop10
} from '@/api/environment/envImages'


defineProps({
  type :{
    type : Number,
    default : 1
  },
  title : {
    type : String,
    default : ""
  },
  data :{   
    type : Array,
    default : () => []
  }
})

const dataLoaded = ref(true)

// var data = []

// getScanlogTop10().then((res) => {
//   console.log(res)

//   data = res.data.results
//   dataLoaded.value = true
//   console.log(data)
// })

// const data = [
//     [
//           { value: 1048, name: 'Search Engine' },
//           { value: 735, name: 'Direct' },
//           { value: 580, name: 'Email' },
//           { value: 484, name: 'Union Ads' },
//           { value: 300, name: 'Video Ads' }
//     ],
//     [
//           { value: 1048, name: 'Search Engine' },
//           { value: 735, name: 'Direct' },
//           { value: 580, name: 'Email' },
//           { value: 484, name: 'Union Ads' },
//           { value: 300, name: 'Video Ads' }
//     ],
//     [
//           { value: 1048, name: 'Search Engine' },
//           { value: 735, name: 'Direct' },
//           { value: 580, name: 'Email' },
//           { value: 484, name: 'Union Ads' },
//           { value: 300, name: 'Video Ads' }
//     ],
//     [
//           { value: 1048, name: 'Search Engine' },
//           { value: 735, name: 'Direct' },
//           { value: 580, name: 'Email' },
//           { value: 484, name: 'Union Ads' },
//           { value: 300, name: 'Video Ads' }
//     ],
// ]
</script>

<style scoped lang="scss">

</style>
