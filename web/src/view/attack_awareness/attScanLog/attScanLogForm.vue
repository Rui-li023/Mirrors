<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="源IP:" prop="source_ip">
          <el-input v-model="formData.source_ip" :clearable="true"  placeholder="请输入源IP" />
       </el-form-item>
        <el-form-item label="源端口:" prop="source_port">
          <el-input v-model.number="formData.source_port" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="协议:" prop="protocol">
           <el-select v-model="formData.protocol" placeholder="请选择协议" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in protocolOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="目的IP:" prop="dest_ip">
          <el-input v-model="formData.dest_ip" :clearable="true"  placeholder="请输入目的IP" />
       </el-form-item>
        <el-form-item label="目的端口:" prop="dest_port">
          <el-input v-model.number="formData.dest_port" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="详细内容:" prop="detail">
          <RichEdit v-model="formData.detail"/>
       </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createScanlog,
  updateScanlog,
  findScanlog
} from '@/api/attack_awareness/attScanLog'

defineOptions({
    name: 'ScanlogForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
// 富文本组件
import RichEdit from '@/components/richtext/rich-edit.vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const protocolOptions = ref([])
const formData = ref({
            source_ip: '',
            source_port: 0,
            protocol: '',
            dest_ip: '',
            dest_port: 0,
            detail: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findScanlog({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rescanlog
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    protocolOptions.value = await getDictFunc('protocol')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createScanlog(formData.value)
               break
             case 'update':
               res = await updateScanlog(formData.value)
               break
             default:
               res = await createScanlog(formData.value)
               break
           }
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
