<template>
  <div class="system">
    <el-form ref="form" :model="config" label-width="240px">
      <!--  System start  -->
      <el-tabs v-model="activeNames">
        <el-tab-pane label="业务设置" name="1" class="mt-3.5">
          <el-form-item label="文件上传最大值">
            <el-input
              v-model.number="config.business['文件上传最大值']"
              placeholder="请输入文件上传最大值"
            />
          </el-form-item>
        </el-tab-pane>

        <!-- 系统配置 -->
        <el-tab-pane label="系统配置" name="2" class="mt-3.5">
          <el-form-item label="端口值">
            <el-input-number v-model="config.system.addr" placeholder="请输入端口值" />
          </el-form-item>
          <el-form-item label="数据库类型">
            <el-select v-model="config.system['db-type']" class="w-full">
              <el-option value="mysql" />
              <el-option value="pgsql" />
              <el-option value="mssql" />
              <el-option value="sqlite" />
              <el-option value="oracle" />
            </el-select>
          </el-form-item>
          <el-form-item label="Oss类型">
            <el-select v-model="config.system['oss-type']" class="w-full">
              <el-option value="local">本地</el-option>
              <el-option value="qiniu">七牛</el-option>
              <el-option value="tencent-cos">腾讯云COS</el-option>
              <el-option value="aliyun-oss">阿里云OSS</el-option>
              <el-option value="huawei-obs">华为云OBS</el-option>
              <el-option value="cloudflare-r2">cloudflare R2</el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="多点登录拦截">
            <el-switch v-model="config.system['use-multipoint']" />
          </el-form-item>
          <el-form-item label="开启redis">
            <el-switch v-model="config.system['use-redis']" />
          </el-form-item>
          <el-form-item label="开启Mongo">
            <el-switch v-model="config.system['use-mongo']" />
          </el-form-item>
          <el-form-item label="严格角色模式">
            <el-switch v-model="config.system['use-strict-auth']" />
          </el-form-item>
          <el-form-item label="限流次数">
            <el-input-number v-model.number="config.system['iplimit-count']" />
          </el-form-item>
          <el-form-item label="限流时间">
            <el-input-number v-model.number="config.system['iplimit-time']" />
          </el-form-item>
          <el-tooltip
            content="请修改完成后，注意一并修改前端env环境下的VITE_BASE_PATH"
            placement="top-start"
          >
            <el-form-item label="全局路由前缀">
              <el-input
                v-model.trim="config.system['router-prefix']"
                placeholder="请输入全局路由前缀"
              />
            </el-form-item>
          </el-tooltip>
        </el-tab-pane>
      </el-tabs>
    </el-form>
    <div class="mt-4">
      <el-button type="primary" @click="save">保存</el-button>
      <el-button type="primary" @click="refeshConfigCacheToDBFunc">刷新缓存存储到数据库</el-button>
      <!-- <el-button type="primary" @click="refeshConfigCacheToEtcdFunc">刷新缓存存储到etcd</el-button> -->
    </div>
  </div>
</template>

<script setup>
  import {
    getBusinessConfig,
    setBusinessConfig,
    refeshConfigCacheToDB,
    refeshConfigCacheToEtcd,
  } from '@/plugin/BusinessConfig/api/businessConfig'
  import { ref } from 'vue'
  import { ElMessage } from 'element-plus'
  import { number } from 'echarts'
  defineOptions({
    name: 'BusinessConfig',
  })

  const activeNames = ref('1')
  const config = ref({
    business: {},
    system: {
      'iplimit-count': 0,
      'iplimit-time': 0,
    },
    jwt: {},
    mysql: {},
    mssql: {},
    sqlite: {},
    pgsql: {},
    oracle: {},
    excel: {},
    autocode: {},
    redis: {},
    mongo: {
      coll: '',
      options: '',
      database: '',
      username: '',
      password: '',
      'min-pool-size': '',
      'max-pool-size': '',
      'socket-timeout-ms': '',
      'connect-timeout-ms': '',
      'is-zap': false,
      hosts: [
        {
          host: '',
          port: '',
        },
      ],
    },
    qiniu: {},
    'tencent-cos': {},
    'aliyun-oss': {},
    'hua-wei-obs': {},
    'cloudflare-r2': {},
    captcha: {},
    zap: {},
    local: {},
    email: {},
    timer: {
      detail: {},
    },
  })

  const initForm = async () => {
    const res = await getBusinessConfig()
    if (res.code === 0) {
      config.value.business = res.data
    }
  }

  const save = async () => {
    const res = await setBusinessConfig(config.value.business)
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '保存成功',
      })
    }
  }

  const refeshConfigCacheToDBFunc = async () => {
    const res = await refeshConfigCacheToDB()
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '缓存存储成功',
      })
    }
  }
  const refeshConfigCacheToEtcdFunc = async () => {
    const res = await refeshConfigCacheToEtcd()
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '缓存存储成功',
      })
    }
  }
  initForm()
</script>

<style lang="scss" scoped>
  .system {
    @apply bg-white p-9 rounded dark:bg-slate-900;
    h2 {
      @apply p-2.5 my-2.5 text-lg shadow;
    }
  }
</style>
