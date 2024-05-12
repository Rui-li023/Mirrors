<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule"
        @keyup.enter="onSubmit">

        <el-form-item label="容器ID" prop="containerId">
          <el-input v-model="searchInfo.containerId" placeholder="容器ID" />
        </el-form-item>
        <el-form-item label="镜像名称" prop="imageName">
          <el-input v-model="searchInfo.imageName" placeholder="模糊查询" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
      </div>
      <el-table ref="multipleTable" style="width: 100%" tooltip-effect="dark" :data="tableData" row-key="ID"
        v-fit-columns>
        <el-table-column type="index" width="55" />

        <!-- <el-table-column align="left" label="日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column> -->

        <el-table-column align="left" label="容器ID" prop="containerId" width="120" />
        <el-table-column align="left" label="容器名字" prop="name" min-width="80" />
        <el-table-column align="left" label="镜像名称" prop="imageName" min-width="180" />
        <el-table-column align="left" label="启动命令" prop="command" min-width="150" />
        <el-table-column align="left" label="运行状态" prop="status" width="200" />
        <el-table-column align="left" label="端口" prop="ports" min-width="300" />
        <el-table-column align="left" label="操作" fixed="right" min-width="160">
          <template #default="scope">
            <el-button-group>
              <el-dropdown trigger="click" @command="handleCommand(scope.row, $event)">
                <el-button type="primary" link icon="more">操作</el-button>
                <template #dropdown>
                  <el-dropdown-menu>
                    <el-dropdown-item v-if="scope.row.status.indexOf('Up')" type="primary" link icon=""
                      @click="startRow(scope.row)">启动</el-dropdown-item>
                    <el-dropdown-item v-if="scope.row.status.indexOf('Exited')" type="primary" link icon=""
                      @click="stopRow(scope.row)">停止</el-dropdown-item>
                    <el-dropdown-item v-if="scope.row.status.indexOf('Exited')" type="primary" link icon=""
                      @click="restartRow(scope.row)">重启</el-dropdown-item>
                  </el-dropdown-menu>
                </template>
              </el-dropdown>

              <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
            </el-button-group>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination layout="total, sizes, prev, pager, next, jumper" :current-page="page" :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]" :total="total" @current-change="handleCurrentChange"
          @size-change="handleSizeChange" />
      </div>
    </div>
    <el-drawer destroy-on-close size="800" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">{{ type === 'create' ? '添加' : '修改' }}</span>
          <div>
            <el-button type="primary" @click="enterDialog">确 定</el-button>
            <el-button @click="closeDialog">取 消</el-button>
          </div>
        </div>
      </template>

      <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
        <!-- <el-form-item label="容器ID:" prop="containerId">
          <el-input v-model="formData.containerId" :clearable="true" placeholder="请输入容器ID" />
        </el-form-item> -->
        <el-form-item label="镜像名称:" prop="imageName">
          <el-select
            v-model="formData.imageName"
            placeholder="请输入镜像名称"
          >
            <el-option
              v-for="item in options1"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>

        <el-form-item label="容器名称:" prop="containerName">
          <el-input v-model="formData.containerName" :clearable="true" placeholder="请输入容器名称" />
        </el-form-item>

  

        <el-form-item label="启动命令:" prop="command">
          <el-button type="primary" @click="addInput">+</el-button>
          <el-button type="primary" @click="removeInput">-</el-button>
          <div v-for="(input, index1) in formData.cmds" :key="index1" style="width: 100%; margin-top: 10px;">
            <el-input v-model="formData.cmds[index1]"  placeholder="请输入启动命令"></el-input>
          </div>
        </el-form-item>
        <!-- <el-form-item label="创建时间:" prop="create">
          <el-input v-model="formData.create" :clearable="true" placeholder="请输入创建时间" />
        </el-form-item>
        <el-form-item label="运行状态:" prop="status">
          <el-input v-model="formData.status" :clearable="true" placeholder="请输入运行状态" />
        </el-form-item> -->
        <!-- <el-form-item label="端口:" prop="ports">
          <el-input v-model="formData.ports" :clearable="true" placeholder="请输入端口" />
        </el-form-item> -->

        <el-form-item label="端口:" prop="ports">
          <el-button type="primary" @click="addPortMapping">+</el-button>
          <el-button type="danger" @click="removePortMapping(index)">-</el-button>

          <div v-for="(portMapping, index) in formData.portMappings" :key="index" style="width: 100%; margin-top: 10px;">
            <el-input v-model="portMapping.containerPort" placeholder="容器端口" style="width: 33%;"></el-input>
            <!-- <el-input v-model="portMapping.containerType" placeholder="传输类型" style="width: 33%;"></el-input> -->
            
            <el-select
              v-model="portMapping.containerType"
              placeholder="传输类型"
              style="width: 33%;"
            >
              <el-option
                :key="'TCP'"
                :label="'TCP'"
                :value="'tcp'"
              >TCP</el-option>
              <el-option
                :key="'UDP'"
                :label="'UDP'"
                :value="'udp'"
              >UDP</el-option>
            </el-select>

            <el-input v-model="portMapping.hostPort" placeholder="主机端口" style="width: 33%;"></el-input>
          
          </div>
        </el-form-item>

      </el-form>
    </el-drawer>
  </div>
</template>

<script setup>
import {
  createContainer,
  deleteContainer,
  updateContainer,
  findContainer,
  startContainer,
  stopContainer,
  restartContainer,
  getContainerList,
} from '@/api/environment/container'

import {
  getImagesList
} from '@/api/environment/envImages'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, filterDataSource, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
  name: 'Container'
})

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  imageName: '',
  containerName:'',
  cmds:[],
  ports: '',
  portMappings: []
})


// 验证规则
const rule = reactive({
})

const searchRule = reactive({
  createdAt: [
    {
      validator: (rule, value, callback) => {
        if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
          callback(new Error('请填写结束日期'))
        } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
          callback(new Error('请填写开始日期'))
        } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() || searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
          callback(new Error('开始日期应当早于结束日期'))
        } else {
          callback()
        }
      }, trigger: 'change'
    }
  ],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async (valid) => {
    if (!valid) return
    page.value = 1
    pageSize.value = 10
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async () => {
  const table = await getContainerList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () => {
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    deleteContainerFunc(row)
  })
}

const startRow = (row) => {
  let message = '确定要启动 ' + row.imageName + ' 吗?'
  ElMessageBox.confirm(message, '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    startContainerFunc(row)
  })
}

const stopRow = (row) => {
  let message = '确定要停止 ' + row.imageName + ' 吗?'
  ElMessageBox.confirm(message, '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    stopContainerFunc(row)
  })
}

const restartRow = (row) => {
  let message = '确定要重启 ' + row.imageName + ' 吗?'
  ElMessageBox.confirm(message, '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    restartContainerFunc(row)
  })
}
// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')



// 删除行
const deleteContainerFunc = async (row) => {
  const res = await deleteContainer({ ID: row.containerId })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--
    }
    getTableData()
  }
}

const startContainerFunc = async (row) => {
  const res = await startContainer({ ID: row.containerId })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '启动成功'
    })
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--
    }
    getTableData()
  }
}
const stopContainerFunc = async (row) => {
  const res = await stopContainer({ ID: row.containerId })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '停止成功'
    })
    getTableData()
  }
}

const restartContainerFunc = async (row) => {
  const res = await restartContainer({ ID: row.containerId })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '重启成功'
    })
    getTableData()
  }
}


// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
  type.value = 'create'
  dialogFormVisible.value = true
  getImagesNameList()
}


// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    imageName: '',
    containerName:'',
    cmds:[],
    portMappings: [],
  }
}

const getImagesNameList = async () =>{
  options1.value=[]
  let res = await getImagesList()
  res.data.list.forEach(element => {
    options1.value.push({'label': element['repository'],'value': element['repository']})
  });
}

const addInput = ()=> {
  // console.log(formData)
  formData.value.cmds.push('')
}

const removeInput = ()=> {
  // console.log(formData)
  formData.value.cmds.pop()
}

const addPortMapping = () => {
      formData.value.portMappings.push({ containerPort: '', hostPort: '', containerType:'tcp' })
    }
const removePortMapping = (index) => {
      formData.value.portMappings.splice(index, 1)
    }
// 弹窗确定
const enterDialog = async () => {
  elFormRef.value?.validate(async (valid) => {
    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createContainer(formData.value)
        break
      case 'update':
        res = await updateContainer(formData.value)
        break
      default:
        res = await createContainer(formData.value)
        break
    }
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '创建/更改成功'
      })
      closeDialog()
      getTableData()
    }
  })
}

// 更新行
// const updateContainerFunc = async (row) => {
//   const res = await findContainer({ ID: row.ID })
//   type.value = 'update'
//   if (res.code === 0) {
//     formData.value = res.data.recontainer
//     dialogFormVisible.value = true
//   }
// }
const options1 = ref([
  {
    value: 'nginx',
    label: 'nginx',
  },
])

</script>

<style></style>