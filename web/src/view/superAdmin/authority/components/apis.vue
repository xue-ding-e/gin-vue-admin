<template>
  <div>
    <div class="sticky top-0.5 z-10 flex flex-wrap space-x-2">
      <el-input v-model="filterTextName" class="flex-1" placeholder="筛选名字" />
      <el-input v-model="filterTextPath" class="flex-1" placeholder="筛选路径" />
      <el-button type="primary" @click="toggleSelectAll">
        {{ isAllSelected ? '取消全选' : '一键全选' }}
      </el-button>
      <el-button class="float-right" type="primary" @click="authApiEnter">确 定</el-button>
    </div>

    <!-- x-role筛选区域 -->
    <div class="mb-2 mt-2 flex items-center space-x-2">
      <span>按x-role筛选:</span>
      <el-select
          v-model="selectedRole"
          placeholder="选择角色"
          clearable
          style="width: 200px"
      >
        <el-option
            v-for="role in availableRoles"
            :key="role"
            :label="role"
            :value="role"
        />
      </el-select>
      <el-button
          type="primary"
          size="small"
          :disabled="!selectedRole"
          @click="selectApisByRole(selectedRole)"
      >
        添加该角色所有API
      </el-button>
      <el-tooltip content="在现有选择的基础上添加该角色的API权限">
        <i class="el-icon-question" />
      </el-tooltip>
    </div>

    <div class="tree-content">
      <el-scrollbar>
        <el-tree
          ref="apiTree"
          :data="apiTreeData"
          :default-checked-keys="apiTreeIds"
          :props="apiDefaultProps"
          default-expand-all
          highlight-current
          node-key="onlyId"
          show-checkbox
          :filter-node-method="filterNode"
          @check="nodeChange"
        >
          <template #default="{ _, data }">
            <div class="flex items-center justify-between w-full pr-1">
              <span>{{ data.description }}</span>
              <div class="flex items-center">
                <el-tag v-if="data.roles && data.roles.length" size="small" type="success" class="mr-2">
                  {{ data.roles.join(', ') }}
                </el-tag>
                <el-tooltip :content="data.path">
                  <span class="max-w-[240px] break-all overflow-ellipsis overflow-hidden">
                    {{ data.path }}
                  </span>
                </el-tooltip>
              </div>
            </div>
          </template>
        </el-tree>
      </el-scrollbar>
    </div>
  </div>
</template>

<script setup>
  import { getAllApis, getSwaggerDoc } from '@/api/api'
  import { UpdateCasbin, getPolicyPathByAuthorityId } from '@/api/casbin'
  import { ref, watch } from 'vue'
  import { ElMessage } from 'element-plus'

  defineOptions({
    name: 'Apis',
  })

  const props = defineProps({
    row: {
      default: function () {
        return {}
      },
      type: Object,
    },
  })

  const apiDefaultProps = ref({
    children: 'children',
    label: 'description',
  })
  const isAllSelected = ref(false)
  // 切换全选状态的方法
  const toggleSelectAll = () => {
    if (apiTree.value) {
      if (!isAllSelected.value) {
        apiTree.value.setCheckedKeys(getAllNodeKeys(apiTreeData.value))
      } else {
        apiTree.value.setCheckedKeys([])
      }
      isAllSelected.value = !isAllSelected.value
    }
  }

  // 获取所有节点的 key
  const getAllNodeKeys = (data) => {
    const keys = []
    const getKeys = (nodes) => {
      nodes.forEach((node) => {
        if (node.children && node.children.length > 0) {
          getKeys(node.children)
        } else {
          keys.push(node.onlyId)
        }
      })
    }
    getKeys(data)
    return keys
  }

  const filterTextName = ref('')
  const filterTextPath = ref('')
  const apiTreeData = ref([])
  const apiTreeIds = ref([])
  const activeUserId = ref('')
  const swaggerData = ref(null)
  const availableRoles = ref([]) // 存储所有可用的角色

  // 从Swagger文档中提取x-role信息
  const extractRolesFromSwagger = (apis, swaggerDoc) => {
    if (!swaggerDoc || !swaggerDoc.paths)
      return apis

    // 用于收集所有角色的集合
    const rolesSet = new Set()

    const processedApis = apis.map((api) => {
      const pathInfo = swaggerDoc.paths[api.path]
      if (pathInfo) {
        const method = api.method.toLowerCase()
        if (pathInfo[method] && pathInfo[method]['x-role']) {
          try {
            // 解析x-role字段，它可能是字符串形式的数组
            let roles = pathInfo[method]['x-role']
            if (typeof roles === 'string') {
              roles = JSON.parse(roles.replace(/'/g, '"'))
            }
            api.roles = roles

            // 将所有角色添加到集合中
            roles.forEach(role => rolesSet.add(role))
          }
          catch (e) {
            console.error('解析x-role失败:', e)
          }
        }
      }
      return api
    })

    // 更新可用角色列表
    availableRoles.value = Array.from(rolesSet)

    return processedApis
  }

  const init = async () => {
    const [swaggerRes, apiRes] = await Promise.all([
      getSwaggerDoc(),
      getAllApis()
    ]);
    swaggerData.value = swaggerRes;
    let apis = apiRes.data.apis;
    if (swaggerData.value) {
      apis = extractRolesFromSwagger(apis, swaggerData.value)
    }

    apiTreeData.value = buildApiTree(apis)
    const res = await getPolicyPathByAuthorityId({
      authorityId: props.row.authorityId,
    })
    activeUserId.value = props.row.authorityId
    apiTreeIds.value = []
    res.data.paths &&
      res.data.paths.forEach((item) => {
        apiTreeIds.value.push('p:' + item.path + 'm:' + item.method)
      })

    // 根据角色自动选中API
    if (props.row.authorityName && swaggerData.value) {
      autoSelectApisByRole(props.row.authorityName)
    }
  }

  // 根据角色名自动选中API
  const autoSelectApisByRole = (roleName) = () => {
    if (!apiTree.value || !swaggerData.value || !swaggerData.value.paths)
      return

    const keysToCheck = []

    // 遍历所有API节点
    const checkNodes = (nodes) => {
      nodes.forEach((node) => {
        if (node.children && node.children.length > 0) {
          checkNodes(node.children)
        }
        else if (node.roles && node.roles.includes(roleName)) {
          keysToCheck.push(node.onlyId)
        }
      })
    }

    checkNodes(apiTreeData.value)

    // 将符合角色的API添加到已选中的列表中
    if (keysToCheck.length > 0) {
      const currentChecked = apiTree.value.getCheckedKeys()
      apiTree.value.setCheckedKeys([...new Set([...currentChecked, ...keysToCheck])])
    }
  }
  // 按角色筛选并选中API
  const selectApisByRole = (roleName) => {
    if (!apiTree.value || !swaggerData.value || !swaggerData.value.paths)
      return

    const keysToCheck = []

    // 遍历所有API节点
    const checkNodes = (nodes) => {
      nodes.forEach((node) => {
        if (node.children && node.children.length > 0) {
          checkNodes(node.children)
        }
        else if (node.roles && node.roles.includes(roleName)) {
          keysToCheck.push(node.onlyId)
        }
      })
    }

    checkNodes(apiTreeData.value)

    // 将符合角色的API添加到已选中的列表中
    if (keysToCheck.length > 0) {
      // 获取当前已选中的keys
      const currentCheckedKeys = apiTree.value.getCheckedKeys()
      // 合并当前选中的和新选中的，使用Set去重
      const mergedKeys = [...new Set([...currentCheckedKeys, ...keysToCheck])]
      // 设置合并后的keys
      apiTree.value.setCheckedKeys(mergedKeys)
      needConfirm.value = true

      ElMessage({
        type: 'success',
        message: `已为角色 "${roleName}" 添加 ${keysToCheck.length} 个API权限`
      })
    } else {
      ElMessage({
        type: 'info',
        message: `没有找到与角色 "${roleName}" 相关的API`,
      })
    }
  }

  init()

  const needConfirm = ref(false)
  // 当树节点变化时，更新全选状态
  const nodeChange = () => {
    needConfirm.value = true
    const checkedKeys = apiTree.value.getCheckedKeys()
    const allKeys = getAllNodeKeys(apiTreeData.value)
    isAllSelected.value = checkedKeys.length === allKeys.length
  }

  // 暴露给外层使用的切换拦截统一方法
  const enterAndNext = () => {
    authApiEnter()
  }

  // 创建api树方法
  const buildApiTree = (apis) => {
    const apiObj = {}
    apis &&
      apis.forEach((item) => {
        item.onlyId = 'p:' + item.path + 'm:' + item.method
        if (Object.prototype.hasOwnProperty.call(apiObj, item.apiGroup)) {
          apiObj[item.apiGroup].push(item)
        } else {
          Object.assign(apiObj, { [item.apiGroup]: [item] })
        }
      })
    const apiTree = []
    for (const key in apiObj) {
      const treeNode = {
        ID: key,
        description: key + '组',
        children: apiObj[key],
      }
      apiTree.push(treeNode)
    }
    return apiTree
  }

  // 关联关系确定
  const apiTree = ref(null)
  const authApiEnter = async () => {
    const checkArr = apiTree.value.getCheckedNodes(true)
    var casbinInfos = []
    checkArr &&
      checkArr.forEach((item) => {
        var casbinInfo = {
          path: item.path,
          method: item.method,
        }
        casbinInfos.push(casbinInfo)
      })
    const res = await UpdateCasbin({
      authorityId: activeUserId.value,
      casbinInfos,
    })
    if (res.code === 0) {
      ElMessage({ type: 'success', message: 'api设置成功' })
    }
  }

  defineExpose({
    needConfirm,
    enterAndNext,
  })

  const filterNode = (value, data) => {
    if (!filterTextName.value && !filterTextPath.value) return true
    let matchesName, matchesPath
    if (!filterTextName.value) {
      matchesName = true
    } else {
      matchesName = data.description && data.description.includes(filterTextName.value)
    }
    if (!filterTextPath.value) {
      matchesPath = true
    } else {
      matchesPath = data.path && data.path.includes(filterTextPath.value)
    }
    return matchesName && matchesPath
  }
  // 当过滤条件变化时，更新全选状态
  watch([filterTextName, filterTextPath], () => {
    apiTree.value.filter('')
    isAllSelected.value = false
  })

  // 选择的角色
  const selectedRole = ref('')
  // 监听角色选择变化
  watch(selectedRole, (newVal) => {
    if (newVal) {
      selectApisByRole(newVal)
    }
  })
</script>
