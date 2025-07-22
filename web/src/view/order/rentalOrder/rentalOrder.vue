
<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
        <el-form-item label="创建日期" prop="createdAtRange">
          <template #label>
            <span>
              创建日期
              <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
                <el-icon><QuestionFilled /></el-icon>
              </el-tooltip>
            </span>
          </template>
          <el-date-picker
            v-model="searchInfo.createdAtRange"
            class="w-[380px]"
            type="datetimerange"
            range-separator="至"
            start-placeholder="开始时间"
            end-placeholder="结束时间"
          />
        </el-form-item>
        <el-form-item label="房源名称" prop="houseTitle">
          <el-input v-model="searchInfo.houseTitle" placeholder="请输入房源名称" clearable />
        </el-form-item>
        <el-form-item label="租客名称" prop="tenantName">
          <el-input v-model="searchInfo.tenantName" placeholder="请输入租客名称" clearable />
        </el-form-item>
        <el-form-item label="房东名称" prop="landlordName">
          <el-input v-model="searchInfo.landlordName" placeholder="请输入房东名称" clearable />
        </el-form-item>
        <el-form-item label="订单状态" prop="status">
          <el-select v-model="searchInfo.status" placeholder="请选择订单状态" clearable style="width: 140px">
            <el-option v-for="item in orderStatusOptions" :key="item.value" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <template v-if="showAllQuery">
          <!-- 可扩展更多查询条件 -->
        </template>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
          <el-button link type="primary" icon="arrow-down" @click="showAllQuery=true" v-if="!showAllQuery">展开</el-button>
          <el-button link type="primary" icon="arrow-up" @click="showAllQuery=false" v-else>收起</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button v-auth="btnAuth.batchDelete" icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
      </div>
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column sortable align="left" label="创建时间" prop="createdAt" width="180">
          <template #default="scope">{{ formatDate(scope.row.createdAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="房源名称" prop="houseTitle" width="180" />
        <el-table-column align="left" label="租客名称" prop="tenantName" width="120" />
        <el-table-column align="left" label="房东名称" prop="landlordName" width="120" />
        <el-table-column align="left" label="租期开始" prop="rentStart" width="120" />
        <el-table-column align="left" label="租期结束" prop="rentEnd" width="120" />
        <el-table-column align="left" label="租金" prop="rentAmount" width="120" />
        <el-table-column align="left" label="押金" prop="deposit" width="120" />
        <el-table-column
          align="left"
          label="订单状态（待处理、活动、已完成、已取消）"
          prop="status"
          width="120"
        >
          <template #default="scope">
            {{ filterDict(scope.row.status, orderStatusOptions) }}
          </template>
        </el-table-column>
        <el-table-column align="left" label="签约时间" prop="signedAt" width="180">
          <template #default="scope">{{ formatDate(scope.row.signedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="操作" fixed="right" :min-width="appStore.operateMinWith">
          <template #default="scope">
            <el-button v-auth="btnAuth.info" type="primary" link class="table-button" @click="goDetail(scope.row)">
              <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看
            </el-button>
            <el-button v-auth="btnAuth.delete" type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          layout="total, sizes, prev, pager, next, jumper"
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>
  </div>
</template>

<script setup>
import { getDictFunc, formatDate, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref } from 'vue'
import { useBtnAuth } from '@/utils/btnAuth'
import { useAppStore } from "@/pinia"
import { useRouter } from 'vue-router'
import {
  deleteRentalOrder,
  deleteRentalOrderByIds,
  findRentalOrder,
  getRentalOrderList
} from '@/api/order/rentalOrder'

const btnAuth = useBtnAuth()
const appStore = useAppStore()
const router = useRouter()

const showAllQuery = ref(false)
const orderStatusOptions = ref([])
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
const multipleSelection = ref([])

const onReset = () => {
  searchInfo.value = {}
  getTableData()
}
const onSubmit = () => {
  // 去除前后空白
  if (searchInfo.value.houseTitle) searchInfo.value.houseTitle = searchInfo.value.houseTitle.trim()
  if (searchInfo.value.tenantName) searchInfo.value.tenantName = searchInfo.value.tenantName.trim()
  if (searchInfo.value.landlordName) searchInfo.value.landlordName = searchInfo.value.landlordName.trim()
  page.value = 1
  getTableData()
}
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}
const getTableData = async() => {
  const table = await getRentalOrderList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    console.log('订单列表返回的原始数据：', table.data.list)
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}
getTableData()
const setOptions = async () =>{
  orderStatusOptions.value = await getDictFunc('orderStatus')
}
setOptions()
const goDetail = (row) => {
  console.log('跳转详情页，行数据：', row)
  console.log('订单ID：', row.id, '类型：', typeof row.id)
  if (!row.id || row.id === 0) {
    ElMessage({
      type: 'warning',
      message: '订单ID无效'
    })
    return
  }
  router.push({ path: '/order/rentalOrder/detail', query: { id: row.id } })
}
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    deleteRentalOrderFunc(row)
  })
}
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
    const IDs = []
    if (multipleSelection.value.length === 0) {
      ElMessage({
        type: 'warning',
        message: '请选择要删除的数据'
      })
      return
    }
    multipleSelection.value &&
      multipleSelection.value.map(item => {
        IDs.push(item.ID)
      })
    const res = await deleteRentalOrderByIds({ IDs })
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '删除成功'
      })
      if (tableData.value.length === IDs.length && page.value > 1) {
        page.value--
      }
      getTableData()
    }
  })
}
const deleteRentalOrderFunc = async (row) => {
  const res = await deleteRentalOrder({ ID: row.ID })
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
</script>

<style>

</style>
