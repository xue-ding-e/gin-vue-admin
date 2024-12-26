export const headerWidth = (label) => {
  // 获取计算后的样式
  const fontSize = window.getComputedStyle(document.body).fontSize
  const fontFamily = window.getComputedStyle(document.body).fontFamily
  const font = `${fontSize} ${fontFamily}`

  let flexWidth = getTextWidth(label, font)
  // 在初步测量后的宽度上添加余量
  flexWidth += 25
  if (flexWidth < 80) {
    flexWidth = 80
  }

  return flexWidth + 'px'
}

function getTextWidth(text, font) {
  const canvas = document.createElement('canvas')
  const context = canvas.getContext('2d')
  context.font = font

  // 使用 measureText 获取初步宽度
  const metrics = context.measureText(text)
  let width = metrics.width

  // 检测并处理所有大写字母，根据需要微调补偿值
  const uppercaseLettersCount = (text.match(/[A-Z]/g) || []).length
  const uppercaseOffset = 2 // 每个大写字母额外增加 2 像素，可根据字体情况自由调整
  width += uppercaseLettersCount * uppercaseOffset

  return width
}
