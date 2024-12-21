// 节省性能

// 构建树形数据  ID列和 parentId列
// 例如：
// 数据：[{ID: 1, parentId: 0, name: '节点1'}, {ID: 2, parentId: 1, name: '节点2'}, {ID: 3, parentId: 1, name: '节点3'}]
// 返回：[{ID: 1, parentId: 0, name: '节点1', children: [{ID: 2, parentId: 1, name: '节点2'}, {ID: 3, parentId: 1, name: '节点3'}]}]
export const buildTreeData = (data) => {
  const tree = []
  const record = {}

  data.forEach((item) => {
    record[item.ID] = item
  })

  data.forEach((item) => {
    const parentId = item.parentId || 0
    if (parentId === 0) {
      tree.push(item)
    } else {
      const parent = record[parentId]
      if (parent) {
        if (!parent.children) {
          parent.children = []
        }
        parent.children.push(item)
      } else {
        // 如果找不到父节点，将该项作为根节点
        tree.push(item)
      }
    }
  })

  return tree
}
