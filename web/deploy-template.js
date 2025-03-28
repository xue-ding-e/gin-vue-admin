import path from 'node:path'
import { fileURLToPath } from 'node:url'
import Client from 'ssh2-sftp-client'

const __dirname = path.dirname(fileURLToPath(import.meta.url))
const sftp = new Client()

const localDir = path.resolve(__dirname, './dist/')
const remoteDir = '服务器存放前端文件路径'

sftp.connect({
    host: '服务器ip',
    port: '22',
    username: '服务器账号',
    password: '服务器密码', // 或者使用 privateKey: fs.readFileSync('/path/to/key')
})
    .then(() => {
        console.log('连接成功，检测远程目录是否存在...')
        return sftp.exists(remoteDir)
    })
    .then((exists) => {
        if (exists) {
            console.log('远程目录存在，开始删除...')
            // 递归删除远程目录
            return sftp.rmdir(remoteDir, true)
        }
        else {
            console.log('远程目录不存在，跳过删除步骤...')
        }
    })
    .then(() => {
        console.log('开始上传文件...')
        return sftp.uploadDir(localDir, remoteDir)
    })
    .then(() => {
        console.log('上传成功！')
    })
    .catch((err) => {
        console.error('操作失败：', err)
    })
    .finally(() => {
        sftp.end()
    })
