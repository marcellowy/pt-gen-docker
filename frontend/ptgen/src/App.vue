<script setup lang="ts">
import axios from 'axios'
import {ref} from "vue";
import {ElMessage} from "element-plus";

let douBanId = ref(34429795)
let result = ref("")
let okLoading=  ref(false)

function getData(){
  result.value = ""
  okLoading.value = true
  axios.post( import.meta.env.VITE_APP_PATH+"/api/v1/ptgen?sid="+douBanId.value, {}).then(function (response) {
    okLoading.value = false
    if (response.status !== 200 || !response.data.success) {
      // 查询失败
      ElMessage.error("查询失败")
      return
    }
    ElMessage.success("查询成功")
    result.value = response.data.format;
  })
}

</script>

<template>
  <el-row>
    <el-col :span="3"></el-col>
    <el-col :span="12">
      <el-form :inline="true">
        <el-form-item label="输入豆瓣ID">
          <el-input placeholder="数字ID" v-model="douBanId" ></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="getData" :loading="okLoading" >查询</el-button>
        </el-form-item>
      </el-form>
      <el-input type="textarea" :rows="30" v-model="result"/>
    </el-col>
  </el-row>
</template>

<style scoped>

</style>
