export const flexColumnWidth = (label) => {
  // 获取计算后的样式
  const fontSize = window.getComputedStyle(document.body).fontSize || '14px'
  const fontFamily = window.getComputedStyle(document.body).fontFamily || 'Arial'
  const font = `${fontSize} ${fontFamily}`

  let flexWidth = getTextWidth(label, font)

  // 增加适当的内边距和其他宽度

  if (flexWidth < 80) {
    flexWidth = 80
  }

  return flexWidth + 'px'
}
const getTextWidth = (text, font) => {
  const canvas = document.createElement('canvas')
  const context = canvas.getContext('2d')
  context.font = font
  const metrics = context.measureText(text)
  return metrics.width
}
