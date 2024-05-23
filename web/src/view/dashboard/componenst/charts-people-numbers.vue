<!--
    本组件参考 arco-pro 的实现 将 ts 改为 js 写法
    https://github.com/arco-design/arco-design-pro-vue/blob/main/arco-design-pro-vite/src/views/dashboard/workplace/components/content-chart.vue
    @auther: bypanghu<bypanghu@163.com>
    @date: 2024/5/8
    @desc: 人数统计图表
!-->

<template>
  <Chart :height="height" :option="chartOption" />
</template>

<script setup>
import Chart from "@/components/charts/index.vue";
import useChartOption from '@/hooks/charts';
import { graphic } from 'echarts'
import { ref } from 'vue';
import { storeToRefs } from "pinia"
import { useAppStore } from '@/pinia'
const appStore = useAppStore()
const { config } = storeToRefs(appStore)

const  prop = defineProps({
  height: {
    type: String,
    default: '128px',
  },
  data : {
    type : Array,
    default : () => []
  }
})
const  graphicFactory = (side) => {
  return {
      type: 'text',
      bottom: '8',
      ...side,
      style: {
        text: '',
        textAlign: 'center',
        fill: '#4E5969',
        fontSize: 12,
      },
    };
  }
  const graphicElements = ref([
    graphicFactory({ left: '5%' }),
    graphicFactory({ right: 0 }),
  ]);
  const { chartOption } = useChartOption(() => {
    console.log(prop.data)
    return {
    tooltip: {
      trigger: 'item'
    },
    legend: {
      orient: 'vertical',
      left: 'left'
    },
    zoom: 2,
    series: [
      {
        name: 'Access From',
        type: 'pie',
        radius: '50%',
        center: ['65%', '50%'],
        data: prop.data,
        emphasis: {
          itemStyle: {
            shadowBlur: 10,
            shadowOffsetX: 0,
            shadowColor: 'rgba(0, 0, 0, 0.5)'
          }
        }
      }
    ]
  }
});
</script>

<style scoped lang="scss">

</style>
