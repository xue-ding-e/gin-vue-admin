import service from '@/utils/request'

export const getBusinessConfig = () => {
  return service({
    url: '/businessConfig/getBusinessConfig',
    method: 'get',
  })
}

export const setBusinessConfig = (data) => {
  return service({
    url: '/businessConfig/updateBusinessConfig',
    method: 'post',
    data,
  })
}

export const refeshConfigCacheToEtcd = () => {
  return service({
    url: '/businessConfig/refeshConfigCacheToEtcd',
    method: 'get',
  })
}
export const refeshConfigCacheToDB = () => {
  return service({
    url: '/businessConfig/refeshConfigCacheToDB',
    method: 'get',
  })
}
