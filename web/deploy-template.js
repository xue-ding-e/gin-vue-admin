const Client = require('ssh2-sftp-client');
const path = require('path');

const sftp = new Client();

const localDir = path.resolve(__dirname, './dist/');
const remoteDir = '服务器存放前端文件路径';

sftp.connect({
    host: '服务器ip',
    port: '22',
    username: '服务器账号',
    password: '服务器密码', // 或者使用 privateKey: fs.readFileSync('/path/to/key')
})
    .then(() => {
        console.log('连接成功，开始上传文件...');
        return sftp.uploadDir(localDir, remoteDir);
    })
    .then(() => {
        console.log('上传成功！');
    })
    .catch(err => {
        console.error('上传失败：', err);
    })
    .finally(() => {
        sftp.end();
    });
