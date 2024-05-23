<template>
  <div class="grid grid-cols-1 md:grid-cols-6 lg:grid-cols-6 p-2 gap-4 md:gap-2 gva-container2">
    <gva-card custom-class="col-span-3 lg:col-span-2 h-80">
      <gva-chart :type="0" title="来源IP TOP10" :data="data" v-if="dataLoaded"/>
    </gva-card>
    <gva-card custom-class="col-span-3 lg:col-span-2 h-80">
      <gva-chart :type="4" title="来源端口 TOP10" :data="data" v-if="dataLoaded"/>
    </gva-card>
    <gva-card custom-class="col-span-3 lg:col-span-2 h-80">
      <gva-chart :type="3" title="被攻击蜜罐 TOP10" :data="data" v-if="dataLoaded"/>
    </gva-card>
    <gva-card custom-class="col-span-3 lg:col-span-2 h-80">
      <gva-chart :type="2" title="被攻击IP TOP10" :data="data" v-if="dataLoaded"/>
    </gva-card>
    <gva-card custom-class="col-span-3 lg:col-span-2 h-80">
      <gva-chart :type="5" title="被攻击端口 TOP10" :data="data" v-if="dataLoaded"/>
    </gva-card>
    <gva-card custom-class="col-span-3 lg:col-span-2 h-80">
      <gva-chart :type="1" title="网络协议组成" :data="data" v-if="dataLoaded"/>
    </gva-card>
    <gva-card title="攻击次数" custom-class="col-span-1 md:col-span-2 md:row-start-3 lg:col-span-6 col-start-1 row-span-2">
      <gva-chart :type="6" />
    </gva-card>
    <!-- <gva-card without-padding custom-class="overflow-hidden h-40 md:row-start-8 md:col-start-3 lg:row-start-3 lg:col-start-7" >
      <gva-banner />
    </gva-card> -->

    <!-- <gva-card title="快捷功能" show-action custom-class="col-start-1 md:col-start-3 lg:col-start-7 row-span-2 h-38">
      <gva-quick-link />
    </gva-card> -->
    <!-- <gva-card title="线上热门内容" custom-class="col-span-1 md:col-span-3 row-span-2">
      <gva-table />
    </gva-card>
    <gva-card title="线上热门内容" custom-class="col-span-1 md:col-span-3 row-span-2">
      <gva-table />
    </gva-card> -->
<!-- 
    <gva-card title="公告" show-action custom-class="col-span-1 lg:col-start-7">
      <gva-notice />
    </gva-card> -->

    <!-- <gva-card title="文档" show-action custom-class="col-span-1 md:col-start-2 md:col-span-1 lg:col-start-7">
      <gva-wiki />
    </gva-card> -->
  </div>
</template>

<script setup>
import { GvaTable, GvaChart , GvaNotice , GvaQuickLink , GvaCard , GvaBanner } from "./componenst"
import {
  getScanlogTop10
} from '@/api/environment/envImages'
import {ref} from "vue"

const dataLoaded = ref(false)

var data = []

getScanlogTop10().then((res) => {
  console.log(res)
  data = res.data.results
  dataLoaded.value = true
  console.log(data)
})

</script>

<style lang="scss" scoped>
</style>
