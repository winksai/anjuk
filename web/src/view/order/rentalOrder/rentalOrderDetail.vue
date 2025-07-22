<template>
  <el-card>
    <el-descriptions :column="1" border>
      <el-descriptions-item label="房源名称">{{ detail.houseTitle || '-' }}</el-descriptions-item>
      <el-descriptions-item label="房源图片" v-if="detail.houseImages && detail.houseImages.length > 0">
        <div style="display: flex; gap: 10px; flex-wrap: wrap;">
          <el-image 
            v-for="(img, index) in detail.houseImages" 
            :key="index"
            :src="img" 
            style="width: 100px; height: 100px; object-fit: cover;"
            :preview-src-list="detail.houseImages"
            fit="cover"
          />
        </div>
      </el-descriptions-item>
      <el-descriptions-item label="租客名称">{{ detail.tenantName || '-' }}</el-descriptions-item>
      <el-descriptions-item label="房东名称">{{ detail.landlordName || '-' }}</el-descriptions-item>
      <el-descriptions-item label="租金">{{ detail.rentAmount !== undefined && detail.rentAmount !== null ? detail.rentAmount : '-' }}</el-descriptions-item>
      <el-descriptions-item label="订单状态">{{ filterDict(detail.status, orderStatusOptions) || '-' }}</el-descriptions-item>
      <el-descriptions-item label="签约时间">{{ detail.signedAt || '-' }}</el-descriptions-item>
      <el-descriptions-item label="合同URL" v-if="detail.contract">
        <a v-if="detail.contract.contractUrl" :href="detail.contract.contractUrl" target="_blank">{{ detail.contract.contractUrl }}</a>
        <span v-else>-</span>
      </el-descriptions-item>
      <el-descriptions-item label="合同状态" v-if="detail.contract">{{ contractStatusMap[detail.contract.status] || detail.contract.status || '-' }}</el-descriptions-item>
      <el-descriptions-item label="合同签约时间" v-if="detail.contract">{{ detail.contract.signTime || '-' }}</el-descriptions-item>
      <el-descriptions-item label="合同内容" v-if="detail.contract">
        <a :href="detail.contract.contractUrl" target="_blank" style="margin-right: 16px;">
          在线预览合同
        </a>
        <el-button
          type="primary"
          size="small"
          @click="downloadContractByApi"
          :icon="Download"
          plain
        >
          下载合同
        </el-button>
      </el-descriptions-item>
    </el-descriptions>
  </el-card>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { findRentalOrder, downloadContract } from '@/api/order/rentalOrder'
import { getDictFunc, filterDict } from '@/utils/format'
import { Download } from '@element-plus/icons-vue'
import { ElMessage } from 'element-plus'

const route = useRoute()
const detail = ref({})
const orderStatusOptions = ref([])

const contractStatusMap = {
  signed: '已签署',
  unsigned: '未签署',
  expired: '已过期'
}

const isImage = (url) => {
  return /\.(png|jpe?g|gif|bmp|webp)$/i.test(url || '')
}

const downloadContractByApi = () => {
  const id = route.query.id
  if (!id) return
  downloadContract(id).then(res => {
    const blob = new Blob([res.data], { type: 'application/pdf' })
    const url = window.URL.createObjectURL(blob)
    const link = document.createElement('a')
    link.href = url
    link.setAttribute('download', '合同.pdf')
    document.body.appendChild(link)
    link.click()
    link.remove()
    window.URL.revokeObjectURL(url)
    ElMessage.success('合同下载成功')
  }).catch(() => {
    ElMessage.error('合同下载失败')
  })
}

onMounted(async () => {
  orderStatusOptions.value = await getDictFunc('orderStatus')
  const id = route.query.id
  console.log('订单详情页接收到的ID参数：', id, '类型：', typeof id)
  if (!id || id === '0') {
    console.error('订单ID参数为空或为0')
    return
  }
  const res = await findRentalOrder({ ID: id })
  if (res.code === 0) {
    detail.value = res.data
  }
})
</script> 